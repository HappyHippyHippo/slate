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

var _ slate.Provider = &Provider{}

// Register will register the REST section instances in the
// application container.
func (p Provider) Register(
	container *slate.Container,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	_ = container.Service(EngineID, func() Engine { return gin.New() })
	_ = container.Service(ProcessID, NewProcess, watchdog.ProcessTag)
	return nil
}

// Boot will start the REST engine with the defined controllers.
func (p Provider) Boot(
	container *slate.Container,
) (e error) {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}

	defer func() {
		if r := recover(); r != nil {
			e = r.(error)
		}
	}()

	// run the registration process of all retrieved registers
	engine := p.getEngine(container)
	for _, r := range p.getRegisters(container) {
		if e := r.Reg(engine); e != nil {
			return e
		}
	}
	return nil
}

func (Provider) getEngine(
	container *slate.Container,
) Engine {
	// retrieve the loader entry
	entry, e := container.Get(EngineID)
	if e != nil {
		panic(e)
	}
	// validate the retrieved entry type
	if instance, ok := entry.(Engine); ok {
		return instance
	}
	panic(errConversion(entry, "rest.Engine"))
}

func (Provider) getRegisters(
	container *slate.Container,
) []EndpointRegister {
	// retrieve the strategies entries
	entries, e := container.Tag(EndpointRegisterTag)
	if e != nil {
		panic(e)
	}
	// type check the retrieved strategies
	var registers []EndpointRegister
	for _, entry := range entries {
		if instance, ok := entry.(EndpointRegister); ok {
			registers = append(registers, instance)
		}
	}
	return registers
}
