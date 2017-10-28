package sources

type FakeGenerationStore struct {
}

func (FakeGenerationStore) GetLastId(key string) int64 {
	return -1
}
func (FakeGenerationStore) SaveId(value int64, key string) {

}
