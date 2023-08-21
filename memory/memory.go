package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/percybolmer/ddd-go/aggregate"
)

type MemoryRepository struct {
	customer map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customer: make(map[uuid.UUID]aggregate.Customer),
	}
}