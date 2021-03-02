package minecraft

import "sync"

// idGenerator is a thread-safe generator for monotonically-increasing unique IDs.
type idGenerator struct {
	lastID int32
	lock   sync.Mutex
}

// GenerateID increments, saves, and returns a new unique ID.
func (gen *idGenerator) GenerateID() int32 {
	gen.lock.Lock()
	defer gen.lock.Unlock()

	id := gen.lastID + 1
	gen.lastID = id

	return id
}
