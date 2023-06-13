package event

// Listener is a listener interface
type Listener interface {
	Listen(param interface{})
}
