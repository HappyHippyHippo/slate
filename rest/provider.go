package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/watchdog"
)

const (
	// ID defines a base id of all other rest
	// package instances registered in the application container.
	ID = slate.ID + ".rest"

	// EngineID defines the id to be used as the
	// container registration id of the rest engine instance.
	EngineID = ID + ".engine"

	// ProcessID defines the id to be used as the
	// container registration id of the rest watchdog process.
	ProcessID = ID + ".process"

	// EndpointRegisterTag defines the tag to be used as the
	// identification of a controller's registration instance.
	EndpointRegisterTag = ID + ".register"
)

// Provider defines the REST services provider instance.
type Provider struct{}

var _ slate.IProvider = &Provider{}

// Register will register the REST section instances in the
// application container.
func (p Provider) Register(
	container ...slate.IContainer,
) error {
	// check container argument reference
	if len(container) == 0 || container[0] == nil {
		return errNilPointer("container")
	}
	// add REST engine
	_ = container[0].Service(EngineID, func() Engine {
		return gin.New()
	})
	// add REST watchdog process instance
	_ = container[0].Service(ProcessID, NewProcess, watchdog.ProcessTag)
	return nil
}

// Boot will start the REST engine with the defined controllers.
func (p Provider) Boot(
	container ...slate.IContainer,
) error {
	// check container argument reference
	if len(container) == 0 || container[0] == nil {
		return errNilPointer("container")
	}
	// retrieve the REST engine
	engine, e := p.getEngine(container[0])
	if e != nil {
		return e
	}
	// retrieve the controller's registration instances
	registers, e := p.getRegisters(container[0])
	if e != nil {
		return e
	}
	// run the registration process of all retrieved registers
	for _, reg := range registers {
		if e := reg.Reg(engine); e != nil {
			return e
		}
	}
	return nil
}

func (Provider) getEngine(
	container slate.IContainer,
) (Engine, error) {
	// retrieve the loader entry
	entry, e := container.Get(EngineID)
	if e != nil {
		return nil, e
	}
	// validate the retrieved entry type
	instance, ok := entry.(Engine)
	if !ok {
		return nil, errConversion(entry, "rest.Engine")
	}
	return instance, nil
}

func (Provider) getRegisters(
	container slate.IContainer,
) ([]IEndpointRegister, error) {
	// retrieve the strategies entries
	entries, e := container.Tag(EndpointRegisterTag)
	if e != nil {
		return nil, e
	}
	// type check the retrieved strategies
	var registers []IEndpointRegister
	for _, entry := range entries {
		if instance, ok := entry.(IEndpointRegister); ok {
			registers = append(registers, instance)
		}
	}
	return registers, nil
}
