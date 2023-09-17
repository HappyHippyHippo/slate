package rest

// EndpointRegister defines an interface to an instance that
// is able to register endpoints to the REST engine/service
type EndpointRegister interface {
	Reg(engine Engine) error
}
