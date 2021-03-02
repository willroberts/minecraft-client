package minecraft

import "sync"

type idGenerator struct {
	lastID int32
	lock   sync.Mutex
}

func (gen *idGenerator) GenerateID() int32 {
	gen.lock.Lock()
	defer gen.lock.Unlock()

	id := gen.lastID + 1
	gen.lastID = id

	return id
}
