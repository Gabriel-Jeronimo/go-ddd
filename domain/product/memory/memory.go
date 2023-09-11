package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/percybolmer/ddd-go/aggregate"
	"github.com/percybolmer/ddd-go/domain/product"
)

type MemoryRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mr *MemoryRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mr.products[id]; ok {
		return product, nil
	}

	return aggregate.Product{}, product.ErrProductNotFound
}

func (mr *MemoryRepository) Add(p aggregate.Product) error {
	if mr.products == nil {
		mr.Lock()
		mr.products = make(map[uuid.UUID]aggregate.Product)
		mr.Unlock()
	}

	if _, ok := mr.products[p.GetID()]; ok {
		return fmt.Errorf("product already exists: %w", product.ErrProductAlreadyExist)
	}

	mr.Lock()
	mr.products[p.GetID()] = p
	mr.Unlock()

	return nil
}

func (mr *MemoryRepository) Update(p aggregate.Product) error {
	if _, ok := mr.products[p.GetID()]; ok {
		return fmt.Errorf("product doesn't exists: %w", product.ErrProductNotFound)
	}

	mr.Lock()
	mr.products[p.GetID()] = p
	mr.Unlock()

	return nil
}

func (mr *MemoryRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product

	for _, v := range mr.products {
		products = append(products, v)
	}

	return products, nil
}

func (mr *MemoryRepository) Delete(id uuid.UUID) error {
	if _, ok := mr.products[id]; ok {
		return fmt.Errorf("product doesn't exist: %w", product.ErrProductNotFound)
	}

	mr.Lock()
	delete(mr.products, id)
	mr.Unlock()

	return nil
}
