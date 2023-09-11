package memory

import (
	"testing"

	"github.com/percybolmer/ddd-go/domain/product"

	"github.com/google/uuid"
	"github.com/percybolmer/ddd-go/aggregate"
)

func TestMemory_AddProduct(t *testing.T) {
	type TestCase struct {
		name        string
		productName string
		description string
		price       float64
		expectedErr error
	}

	testCases := []TestCase{
		{
			name:        "Add a product",
			productName: "Playstation 4",
			description: "A fine console",
			price:       4000.00,
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo := MemoryRepository{
				products: map[uuid.UUID]aggregate.Product{},
			}

			prod, err := aggregate.NewProduct(testCase.productName, testCase.description, testCase.price)

			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(prod)

			if err != testCase.expectedErr {
				t.Errorf("Expected error %v, got %v", testCase.expectedErr, err)
			}

			found, err := repo.GetByID(prod.GetID())

			if err != nil {
				t.Fatal(err)
			}

			foundId := found.GetID()
			prodId := prod.GetID()

			if foundId != prodId {
				t.Errorf("Expected ID %v, got %v", prodId, foundId)
			}
		})
	}
}
func TestMemory_GetProduct(t *testing.T) {
	type TestCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	prod, err := aggregate.NewProduct("Playstation 4", "A fine console", 4000.00)

	if err != nil {
		t.Fatal(err)
	}

	id := prod.GetID()

	repo := MemoryRepository{
		products: map[uuid.UUID]aggregate.Product{
			id: prod,
		},
	}

	testCases := []TestCase{
		{
			name:        "Customer doesn't exist",
			id:          uuid.New(),
			expectedErr: product.ErrProductNotFound,
		},
		{
			name:        "Customer exists",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := repo.GetByID(testCase.id)

			if err != testCase.expectedErr {
				t.Errorf("Expected error %v, got %v", testCase.expectedErr, err)
			}
		})
	}
}
