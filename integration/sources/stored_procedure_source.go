package sources

import (
	"charisma-beat/integration"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

type storedProcedureEventSource struct {
	dsn             string
	sql             string
	name            string
	eventMapper     integration.EventMapper
	generationStore GenerationStore

	rows      *sql.Rows
	vals      []interface{}
	cols      []string
	row       map[string]interface{}
	conn      *sql.DB
	unsendRow map[string]interface{}
	done      bool
}

func NewStoredProcedureEventSource(dsn string, sql string, name string,
	eventMapper integration.EventMapper, generationStore GenerationStore) *storedProcedureEventSource {
	return &storedProcedureEventSource{
		dsn:             dsn,
		sql:             sql,
		eventMapper:     eventMapper,
		name:            name,
		generationStore: generationStore,
	}
}

func (source *storedProcedureEventSource) Open() error {
	var err error
	source.conn, err = sql.Open("mssql", source.dsn)
	if err != nil {
		log.Println("Cannot connect: ", err.Error())
		return err
	}

	source.rows, err = source.conn.Query(source.sql, source.generationStore.GetLastId(source.name))
	if err != nil {
		log.Println("Query failed:", err.Error())
		return err
	}

	source.cols, _ = source.rows.Columns()

	source.vals = make([]interface{}, len(source.cols))
	source.row = make(map[string]interface{}, len(source.cols))
	for i := 0; i < len(source.cols); i++ {
		source.vals[i] = new(interface{})
	}

	return nil
}

func (source *storedProcedureEventSource) Close() {
	source.conn.Close()
	source.rows.Close()
}

func (source *storedProcedureEventSource) Next() (integration.Event, bool, error) {
	var i = 0
	if source.unsendRow != nil {
		if source.eventMapper.ShouldAddToCollection(source.row) {
			source.eventMapper.AddToCollection(source.row)
		} else {
			source.eventMapper.Map(source.row)
		}
		i = 1
		source.unsendRow = nil
	}

	for source.rows.Next() {
		err := source.rows.Scan(source.vals...)
		if err != nil {
			return nil, false, err
		}

		for i, colName := range source.cols {
			val := source.vals[i].(*interface{})
			source.row[colName] = *val
		}

		if source.eventMapper.ShouldAddToCollection(source.row) {
			source.eventMapper.AddToCollection(source.row)
		} else {
			if i > 0 {
				source.unsendRow = source.row
				return source.eventMapper.GetEvent(), true, nil
			}
			source.eventMapper.Map(source.row)
		}
		i++
	}

	if i > 0 {
		return source.eventMapper.GetEvent(), true, nil
	} else {
		source.SaveNextGenerationId()
		return nil, false, nil
	}
}

func (source *storedProcedureEventSource) SaveNextGenerationId() {
	if source.rows.NextResultSet() && source.rows.Next() {
		var generationId int64
		source.rows.Scan(&generationId)
		source.generationStore.SaveId(generationId, source.name)
	}
}
