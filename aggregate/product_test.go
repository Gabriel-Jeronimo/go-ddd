package aggregate_test

import (
	"testing"

	"github.com/percybolmer/ddd-go/aggregate"
)

func TestProduct_NewProduct(t *testing.T) {
	type TestCase struct {
		test        string
		name        string
		description string
		price       float64
		expectedErr error
	}

	testCases := []TestCase{
		{
			test:        "Empty name validation",
			name:        "",
			description: "A nice console",
			price:       4000.00,
			expectedErr: aggregate.ErrMissingValues,
		},
		{
			test:        "Empty description validation",
			name:        "Playstation 4",
			description: "",
			price:       4000.00,
			expectedErr: aggregate.ErrMissingValues,
		},
		{
			test:        "Valid name and description",
			name:        "Playstation 4",
			description: "A nice console",
			price:       4000.00,
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.test, func(t *testing.T) {
			_, err := aggregate.NewProduct(testCase.name, testCase.description, testCase.price)

			if err != testCase.expectedErr {
				t.Errorf("Expected error %v, got %v", testCase.expectedErr, err)
			}
		})
	}
}
