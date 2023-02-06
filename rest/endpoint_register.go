package rest

// IEndpointRegister defines an interface to an instance that
// is able to register endpoints to the REST engine/service
type IEndpointRegister interface {
	Reg(engine Engine) error
}
