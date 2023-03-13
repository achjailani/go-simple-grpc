package server

// Contract is an interface
type Contract interface {
	Run(port int) error
}
