package watchdog

type logAdapter interface {
	Start() error
	Error(e error) error
	Done() error
}

// Watchdog defines the instance used to overlook a process execution.
type Watchdog struct {
	logAdapter logAdapter
}

// NewWatchdog generates a new watchdog instance.
func NewWatchdog(
	logAdapter *LogAdapter,
) (*Watchdog, error) {
	// check logAdapter argument reference
	if logAdapter == nil {
		return nil, errNilPointer("logAdapter")
	}
	// return the created watchdog instance
	return &Watchdog{
		logAdapter: logAdapter,
	}, nil
}

// Run will run a process overlooked by the current watchdog instance.
func (w *Watchdog) Run(process Processor) (e error) {
	// create the goroutine signal channels
	closed := make(chan struct{})
	errored := make(chan struct{})
	runner := func() {
		defer func() {
			// get the error instance
			if resp := recover(); resp != nil {
				if typedResp, ok := resp.(error); ok {
					e = typedResp
				}
				// signal error goroutine execution status
				errored <- struct{}{}
			}
		}()
		// run the process method
		e = process.Runner()()
		// signal correct termination of the goroutine
		closed <- struct{}{}
	}
	// log the starting of the watchdog process
	_ = w.logAdapter.Start()
	for {
		// run the method
		go runner()
		// wait for the method result signals
		select {
		case <-errored:
			// log the error
			_ = w.logAdapter.Error(e)
		case <-closed:
			// log the execution termination and
			// terminate the watchdog
			_ = w.logAdapter.Done()
			return e
		}
	}
}
