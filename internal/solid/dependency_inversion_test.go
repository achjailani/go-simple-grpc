package solid_test

import (
	"github/achjailani/go-simple-grpc/internal/solid"
	"testing"
)

func TestDependencyInversion(t *testing.T) {
	mysqlDB := &solid.MySQLConnection{}

	userRepo := solid.NewUserRepository(mysqlDB)
	userRepo.Save(struct {
		ID   int
		Name string
	}{
		ID:   1,
		Name: "Hello",
	})

	postgresDB := &solid.PostgresConnection{}
	userRepo = solid.NewUserRepository(postgresDB)
	userRepo.Save(struct {
		ID   int
		Name string
	}{
		ID:   1,
		Name: "Hi",
	})
}
