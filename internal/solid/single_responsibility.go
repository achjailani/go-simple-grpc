package solid

import (
	"fmt"
	"time"
)

// Customer is a struct
// it should have only one responsibility or task which is to define
// entity or struct object
type Customer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ============== PERSISTENCE LAYER ==============

// CustomerRepositoryInterface is an interface
type CustomerRepositoryInterface interface {
	Save(consumer *Customer) error
	Find(id int) (*Customer, error)
}

// CustomerRepository is a struct
type CustomerRepository struct {
}

// NewCustomerRepository is a constructor
func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}

// Save is a method
func (c *CustomerRepository) Save(consumer *Customer) error {
	fmt.Println("Successfully saved to database.")

	return nil
}

// Find is a method
func (c *CustomerRepository) Find(id int) (*Customer, error) {
	fmt.Printf("Successfully get customer by id: %d\n", id)

	return &Customer{}, nil
}

var _ CustomerRepositoryInterface = &CustomerRepository{}

// =============== NOTIFICATION LAYER =================

// CustomerNotifierInterface is an interface
type CustomerNotifierInterface interface {
	Send(customer *Customer, message string) error
}

// CustomerNotifier is a struct
type CustomerNotifier struct {
}

// NewCustomerNotifier is a constructor
func NewCustomerNotifier() *CustomerNotifier {
	return &CustomerNotifier{}
}

// Send is a method
func (c CustomerNotifier) Send(customer *Customer, message string) error {
	fmt.Println(message)

	return nil
}

var _ CustomerNotifierInterface = &CustomerNotifier{}
