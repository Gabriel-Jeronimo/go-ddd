package memory

import (
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/percybolmer/ddd-go/aggregate"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
	ErrUpdateCustomer      = errors.New("failed to update the customer in the repository")
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

func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := mr.customer[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, ErrCustomerNotFound
}

func (mr *MemoryRepository) Add(c aggregate.Customer) error {
	if mr.customer == nil {
		mr.Lock()
		mr.customer = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}

	if _, ok := mr.customer[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", ErrFailedToAddCustomer)
	}

	mr.Lock()
	mr.customer[c.GetID()] = c
	mr.Unlock()
	return nil
}

func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	if _, ok := mr.customer[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", ErrCustomerNotFound)
	}

	mr.Lock()
	mr.customer[c.GetID()] = c
	mr.Unlock()
	return nil
}
