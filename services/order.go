package services

import (
	"github.com/percybolmer/ddd-go/domain/customer"
	"github.com/percybolmer/ddd-go/domain/customer/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) error {
	c, err := o.customers.Get(customerId)

	if err != nil {
		return err
	}

	// Put all the products in the customer cart
	return nil
}