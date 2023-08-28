package memory

import (
	"testing"

	"github.com/google/uuid"
	"github.com/percybolmer/ddd-go/aggregate"
	"github.com/percybolmer/ddd-go/domain/customer"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := aggregate.NewCustomer("Drake")

	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:        "No customer by ID",
			id:          uuid.New(),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "Customer by ID",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := repo.Get(testCase.id)

			if err != testCase.expectedErr {
				t.Errorf("Expected error %v, got %v", testCase.expectedErr, err)
			}
		})
	}
}

func TestMemory_AddCustomer(t *testing.T) {
	type testCase struct {
		name        string
		customer    string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add customer",
			customer:    "Drake",
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo := MemoryRepository{
				customers: map[uuid.UUID]aggregate.Customer{},
			}

			cust, err := aggregate.NewCustomer(testCase.customer)

			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(cust)

			if err != testCase.expectedErr {
				t.Errorf("Expected error %v, got %v", testCase.expectedErr, err)
			}

			found, err := repo.Get(cust.GetID())

			if err != nil {
				t.Fatal(err)
			}

			foundId := found.GetID()
			custId := cust.GetID()

			if foundId != custId {
				t.Errorf("Expected id %v, got %v", foundId, custId)
			}
		})
	}
}
