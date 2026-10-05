package agodaPrep

type KeyValueStore struct {
	currentStore  map[string]string
	snapshotStore map[int]map[string]string
	nextID        int
}

func (s *KeyValueStore) Set(key string, value string) {
	s.currentStore[key] = value
}

func (s *KeyValueStore) Get(key string) (string, bool) {
	value, exist := s.currentStore[key]
	if exist {
		return value, true
	}

	return "", false
}

func (s *KeyValueStore) Snapshot() int {
	id := s.nextID + 1
	s.snapshotStore[id] = s.currentStore
	return id
}

func (s *KeyValueStore) GetFromSnapshot(snapshotID int, key string) (string, bool) {
	snapshot, exist := s.snapshotStore[snapshotID]
	if !exist {
		return "", false
	}

	value, exist := snapshot[key]
	if !exist {
		return "", false
	}

	return value, true
}
