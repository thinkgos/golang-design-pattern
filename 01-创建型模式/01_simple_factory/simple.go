package simplefactory

import "fmt"

// API is interface
type API interface {
	Say(name string) string
}

const (
	hi = iota
	hello
)

// NewAPI return Api instance by type
func NewAPI(t int) API {
	switch t {
	case hi:
		return new(hiAPI)
	case hello:
		return new(helloAPI)
	default:
		panic("invalid api type")
	}
}

//hiAPI is one of API implement
type hiAPI struct{}

//Say hi to name
func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

//HelloAPI is another API implement
type helloAPI struct{}

//Say hello to name
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
