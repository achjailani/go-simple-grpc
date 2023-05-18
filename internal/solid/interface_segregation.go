package solid

import "fmt"

// Printer is a struct
type Printer interface {
	Print(document string)
}

// Scanner is a struct
type Scanner interface {
	Scan()
}

// SimplePrinter is a struct
type SimplePrinter struct {
}

// Print is a method
func (p *SimplePrinter) Print(document string) {
	fmt.Printf("%s successfully printed!\n", document)
}

// MultiFunctionPrinter is a struct
type MultiFunctionPrinter struct {
	// Multi-function printer implementation}
}

// Print is a method
func (m *MultiFunctionPrinter) Print(document string) {
	fmt.Printf("%s successfully printed!\n", document)
}

// Scan is a method
func (m *MultiFunctionPrinter) Scan() {
	fmt.Println("Document successfully scanned!")
}
