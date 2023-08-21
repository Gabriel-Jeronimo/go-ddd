package aggregate_test

import (
	"testing"

	"github.com/percybolmer/ddd-go/aggregate"
)

func TestCustomer_NewCostumer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		},
		{
			test:        "Valid name",
			name:        "Gabriel Jeronimo",
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(testCase.name)

			if err != testCase.expectedErr {
				t.Errorf("Expected error %v, got %v", testCase.expectedErr, err)
			}
		})
	}
}
