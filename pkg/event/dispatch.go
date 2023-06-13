package event

// Dispatch is a function to set new event
func Dispatch(name Name, param interface{}) error {
	return dispatcher.dispatch(name, param)
}
