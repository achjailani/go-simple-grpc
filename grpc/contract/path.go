package contract

// ProtectedMethods is a function
func ProtectedMethods() map[string]bool {
	return map[string]bool{
		"/foo.Hello/Ping":  false,
		"/foo.Hello/Hello": true,
	}
}
