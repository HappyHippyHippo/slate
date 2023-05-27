package watchdog

import (
	"sync"
)

type kennelReg struct {
	process  IProcess
	watchdog *Watchdog
}

type factory interface {
	Create(service string) (*Watchdog, error)
}

// Kennel define an instance that will manage a group of watchdog
// instances, and is used to run them in parallel.
type Kennel struct {
	factory factory
	regs    map[string]kennelReg
}

// NewKennel will generate a new kennel instance.
func NewKennel(factory *Factory) (*Kennel, error) {
	// check factory argument reference
	if factory == nil {
		return nil, errNilPointer("factory")
	}
	// return the created factory instance
	return &Kennel{
		factory: factory,
		regs:    map[string]kennelReg{},
	}, nil
}

// Add will create a new watchdog instance that will guard the
// requested process instance.
func (k *Kennel) Add(process IProcess) error {
	// check if there is a watchdog for the requested service
	if _, ok := k.regs[process.Service()]; ok {
		return errDuplicateService(process.Service())
	}
	// create the watchdog for the requested process
	wd, e := k.factory.Create(process.Service())
	if e != nil {
		return e
	}
	// store the process and the created watchdog in the kennel
	k.regs[process.Service()] = kennelReg{
		process:  process,
		watchdog: wd,
	}
	return nil
}

// Run will execute all the registered processes in their
// respective watchdogs.
func (k *Kennel) Run() error {
	// check if there is watchdogs to run
	if len(k.regs) == 0 {
		return nil
	}
	var result error
	// start all the registered watchdogs
	wg := sync.WaitGroup{}
	for _, reg := range k.regs {
		wg.Add(1)
		// run the registered process
		go func(reg kennelReg) {
			// run the process on a created watchdog
			e := reg.watchdog.Run(reg.process)
			if e != nil {
				result = e
			}
			// signal the wait group that the watchdog terminated
			wg.Done()
		}(reg)
	}
	// wait for all started watchdogs processes
	wg.Wait()
	return result
}
