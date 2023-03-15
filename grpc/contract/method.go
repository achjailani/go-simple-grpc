package contract

// ProtectedMethods is a function to hold grpc service methods
// false value indicates that the method is not protected (no authorization needed)
func ProtectedMethods() map[string]bool {
	return map[string]bool{
		"/foo.Hello/Ping":                   false,
		"/foo.Hello/Hello":                  true,
		"/foo.LogService/SaveHttpLog":       true,
		"/foo.LogService/SaveStreamHttpLog": true,
		"/foo.LogService/FindHttpLog":       true,
		"/foo.LogService/GetHttpLog":        true,
	}
}
