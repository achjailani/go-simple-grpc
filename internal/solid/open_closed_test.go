package solid_test

import (
	"github.com/stretchr/testify/assert"
	"github/achjailani/go-simple-grpc/internal/solid"
	"testing"
)

func TestNewPaymentHandler(t *testing.T) {
	amount := 100.0
	paymentMethod := &solid.CreditCardPayment{}

	paymentHandler := solid.NewPaymentHandler(paymentMethod)
	err := paymentHandler.Pay(amount)
	assert.NoError(t, err)

	gopay := &solid.GoPayPayment{}

	paymentHandler = solid.NewPaymentHandler(gopay)
	err = paymentHandler.Pay(amount)

	assert.NoError(t, err)
}
