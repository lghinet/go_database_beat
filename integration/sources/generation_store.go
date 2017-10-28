package sources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type GenerationStore interface {
	GetLastId(key string) int64
	SaveId(generationId int64, key string)
}

type fileGenerationStore struct {
	sync.Mutex
	ids      map[string]int64
	filePath string
}

func NewFileGenerationStore() *fileGenerationStore {
	store := &fileGenerationStore{filePath: "generation_store.json"}
	store.ids = make(map[string]int64, 0)

	file, err := os.Open(store.filePath)
	if err != nil {
		fmt.Println("error:", err)
		return store
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&store.ids)
	if err != nil {
		fmt.Println("error:", err)
	}

	return store
}

func (store *fileGenerationStore) GetLastId(key string) int64 {
	store.Lock()
	value, ok := store.ids[key]
	store.Unlock()

	if !ok {
		store.SaveId(-1, key)
		return -1
	}
	return value
}
func (store *fileGenerationStore) SaveId(value int64, key string) {
	store.Lock()

	store.ids[key] = value
	bytes, _ := json.Marshal(store.ids)
	err := ioutil.WriteFile(store.filePath, bytes, 0644)
	if err != nil {
		fmt.Println("error:", err)
	}
	store.Unlock()
}
