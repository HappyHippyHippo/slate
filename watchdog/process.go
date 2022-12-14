package watchdog

// IProcess defines an interface to a watchdog process.
type IProcess interface {
	Service() string
	Runner() func() error
}

// Process defines an instance to a watchdog process that will be
// overlooked by the watchdog.
type Process struct {
	service string
	runner  func() error
}

var _ IProcess = &Process{}

// NewProcess generate a new process instance with the given
// service name and runner method.
func NewProcess(
	service string,
	runner func() error,
) (*Process, error) {
	// check runner function argument reference
	if runner == nil {
		return nil, errNilPointer("runner")
	}
	// return the created process instance
	return &Process{
		service: service,
		runner:  runner,
	}, nil
}

// Service will retrieve the service name.
func (p *Process) Service() string {
	return p.service
}

// Runner retrieve the process runner method.
func (p *Process) Runner() func() error {
	return p.runner
}
