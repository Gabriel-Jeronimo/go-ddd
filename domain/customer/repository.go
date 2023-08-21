package customer

import (
	"github.com/google/uuid"
	"github.com/percybolmer/ddd-go/aggregate"
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}