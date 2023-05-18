package solid

import "fmt"

// SQLConnectionInterface is an interface
// This is a low-level module
type SQLConnectionInterface interface {
	Connect()
}

// MySQLConnection is a struct
// This is a high-level module
type MySQLConnection struct {
}

// Connect is a method
func (m *MySQLConnection) Connect() {
	fmt.Println("Connect to mysql db")
}

// PostgresConnection is a type
// This is a high-level module
type PostgresConnection struct {
}

// Connect is a method
func (p *PostgresConnection) Connect() {
	fmt.Println("Connect to mysql db")
}

// UserRepository is a struct
type UserRepository struct {
	DB SQLConnectionInterface
}

// NewUserRepository is a constructor
func NewUserRepository(db SQLConnectionInterface) *UserRepository {
	return &UserRepository{DB: db}
}

// Save is a method
func (ur *UserRepository) Save(data interface{}) error {
	fmt.Println("successfully saved")
	return nil
}
