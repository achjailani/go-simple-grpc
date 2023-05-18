package solid

import "fmt"

// PaymentMethod is an interface
type PaymentMethod interface {
	pay(amount float64) error
}

// CreditCardPayment is a struct
type CreditCardPayment struct {
}

// Pay is a method
func (c *CreditCardPayment) pay(amount float64) error {
	fmt.Println("successfully paid with Credit Card.")

	return nil
}

// PayPalPayment is a struct
type PayPalPayment struct {
}

// Pay is a method
func (p *PayPalPayment) pay(amount float64) error {
	fmt.Println("successfully paid with PayPal.")

	return nil
}

// PaymentHandler is a struct
type PaymentHandler struct {
	PaymentMethod PaymentMethod
}

// Pay is a method
func (p *PaymentHandler) Pay(amount float64) error {
	return p.PaymentMethod.pay(amount)
}

// NewPaymentHandler is a constructor
func NewPaymentHandler(method PaymentMethod) *PaymentHandler {
	return &PaymentHandler{PaymentMethod: method}
}
