package solid_test

import (
	"github.com/stretchr/testify/assert"
	"github/achjailani/go-simple-grpc/internal/solid"
	"testing"
)

func TestCustomer(t *testing.T) {
	customer := solid.Customer{
		ID:    1,
		Name:  "Orlando",
		Email: "orlando@example.com",
	}

	var err error

	err = solid.NewCustomerRepository().Save(&customer)
	assert.NoError(t, err)

	err = solid.NewCustomerNotifier().Send(&customer, "welcome!")
	assert.NoError(t, err)

}
