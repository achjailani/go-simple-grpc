package tests

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github/achjailani/go-simple-grpc/config"
	"github/achjailani/go-simple-grpc/domain/service"
	"github/achjailani/go-simple-grpc/infrastructure/dependency"
	"github/achjailani/go-simple-grpc/infrastructure/persistence"
	"github/achjailani/go-simple-grpc/pkg/logger"
	"github/achjailani/go-simple-grpc/tests/database"
	"github/achjailani/go-simple-grpc/utils"
	"log"
	"os"
)

type BoxTest struct {
	Ctx context.Context
	*dependency.Dependency
}

// Init is a function to initialize tests
func Init() *BoxTest {
	errT := os.Setenv("TEST_MODE", "true")
	if errT != nil {
		log.Fatalf("unable to set test mode, err: %s", errT.Error())
	}

	if err := godotenv.Load(fmt.Sprintf("%s/.env", utils.RootDir())); err != nil {
		log.Fatalf("no .env file provided.")
	}

	cfg := config.New()
	ctx := context.Background()

	db, errConn := persistence.NewDBConnection(cfg)
	if errConn != nil {
		log.Fatalf("unable connect to database, %v", errConn)
	}

	drop := database.NewDrop(db)
	errDrop := drop.Reset(ctx)
	if errDrop != nil {
		log.Fatalf("err drop: %v", errDrop)
	}

	repo := service.NewDBService(db)
	loggr := logger.New(logger.NewConfig())

	return &BoxTest{
		Ctx: ctx,
		Dependency: dependency.New(
			dependency.WithConfig(cfg),
			dependency.WithRepository(repo),
			dependency.WithLogger(loggr),
		),
	}
}
