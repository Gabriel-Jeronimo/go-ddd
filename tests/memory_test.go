package aggregate_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/percybolmer/ddd-go/aggregate"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := aggregate.NewCustomer()
}
