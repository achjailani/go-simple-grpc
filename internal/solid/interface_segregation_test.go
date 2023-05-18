package solid_test

import (
	"github/achjailani/go-simple-grpc/internal/solid"
	"testing"
)

func TestInterfaceSegregation(t *testing.T) {
	simplePrinter := &solid.SimplePrinter{}
	simplePrinter.Print("Sample Document")

	multiFunctionPrinter := &solid.MultiFunctionPrinter{}
	multiFunctionPrinter.Print("Multi-Function Document")
	multiFunctionPrinter.Scan()
}
