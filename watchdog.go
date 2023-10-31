package slate

import (
	"fmt"
	"sync"
)

// ----------------------------------------------------------------------------
// defs
// ----------------------------------------------------------------------------

const (
	// WatchdogContainerID defines the id to be used as the Provider
	// registration id of a watchdog creator instance, and as a
	// base id of all other watchdog package connections registered
	// in the application Provider.
	WatchdogContainerID = ContainerID + ".watchdog"

	// WatchdogLogFormatterContainerID defines the simple tag to be used
	// to identify a watchdog log formatter entry in the Provider.
	WatchdogLogFormatterContainerID = WatchdogContainerID + ".formatter"

	// WatchdogLogFormatterCreatorTag defines the simple tag to be used
	// to identify a watchdog log formatter entry in the Provider.
	WatchdogLogFormatterCreatorTag = WatchdogLogFormatterContainerID + ".creator"

	// WatchdogDefaultLogFormatterCreatorContainerID defines the simple tag to be used
	// to identify a watchdog log formatter entry in the Provider.
	WatchdogDefaultLogFormatterCreatorContainerID = WatchdogLogFormatterCreatorTag + ".default"

	// WatchdogAllLogFormatterCreatorsContainerID defines the simple tag to be used
	// to identify a watchdog log formatter entry in the Provider.
	WatchdogAllLogFormatterCreatorsContainerID = WatchdogLogFormatterCreatorTag + ".all"

	// WatchdogLogFormatterFactoryContainerID defines the id to be used as the Provider
	// registration id of the log formatter creator.
	WatchdogLogFormatterFactoryContainerID = WatchdogLogFormatterContainerID + ".factory"

	// WatchdogFactoryContainerID defines the id to be used as the Provider
	// registration id of the watchdog creator.
	WatchdogFactoryContainerID = WatchdogContainerID + ".factory"

	// WatchdogProcessTag defines the simple tag to be used
	// to identify a watchdog process entry in the Provider.
	WatchdogProcessTag = WatchdogContainerID + ".process"

	// WatchdogAllProcessesContainerID defines the simple tag to be used
	// to identify a watchdog log formatter entry in the Provider.
	WatchdogAllProcessesContainerID = WatchdogProcessTag + ".all"

	// WatchdogEnvID defines the watchdog package base environment variable name.
	WatchdogEnvID = EnvID + "_WATCHDOG"

	// WatchdogLogFormatterTypeDefault defines the default log formatter type id
	WatchdogLogFormatterTypeDefault = "default"
)

var (
	// WatchdogConfigPathPrefix defines the configuration path of the watchdog
	// entries of the application.
	WatchdogConfigPathPrefix = EnvString(WatchdogEnvID+"_CONFIG_PATH", "slate.watchdog.services")

	// WatchdogDefaultFormatter defines the default logging formatter instance id.
	WatchdogDefaultFormatter = EnvString(WatchdogEnvID+"_DEFAULT_FORMATTER", WatchdogLogFormatterTypeDefault)

	// WatchdogLogChannel defines the logging signal channel of the watchdogs.
	WatchdogLogChannel = EnvString(WatchdogEnvID+"_LOG_CHANNEL", "watchdog")

	// WatchdogLogStartLevel defines the watchdog starting logging signal
	// message level.
	WatchdogLogStartLevel = EnvString(WatchdogEnvID+"_LOG_START_LEVEL", "notice")

	// WatchdogLogStartMessage defines the watchdog starting logging signal message.
	WatchdogLogStartMessage = EnvString(WatchdogEnvID+"_LOG_START_MESSAGE", "[watchdog:%s] start execution")

	// WatchdogLogErrorLevel defines the watchdog error logging signal
	// message level.
	WatchdogLogErrorLevel = EnvString(WatchdogEnvID+"_LOG_ERROR_LEVEL", "error")

	// WatchdogLogErrorMessage defines the watchdog error logging signal message.
	WatchdogLogErrorMessage = EnvString(WatchdogEnvID+"_LOG_ERROR_MESSAGE", "[watchdog:%s] execution error (%v)")

	// WatchdogLogDoneLevel defines the watchdog termination logging
	// signal message level.
	WatchdogLogDoneLevel = EnvString(WatchdogEnvID+"_LOG_DONE_LEVEL", "notice")

	// WatchdogLogDoneMessage defines the watchdog termination logging
	// signal message.
	WatchdogLogDoneMessage = EnvString(WatchdogEnvID+"_LOG_DONE_MESSAGE", "[watchdog:%s] execution terminated")
)

// ----------------------------------------------------------------------------
// errors
// ----------------------------------------------------------------------------

var (
	// ErrInvalidWatchdogLogWriter defines an error that signal that the
	// given watchdog configuration has defined an unrecognizable log
	// writer type.
	ErrInvalidWatchdogLogWriter = fmt.Errorf("invalid watchdog log writer")

	// ErrDuplicateWatchdog defines an error that signal that the
	// given watchdog service is already registered.
	ErrDuplicateWatchdog = fmt.Errorf("duplicate watchdog")
)

func errInvalidWatchdogLogWriter(
	config *ConfigPartial,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrInvalidWatchdogLogWriter, fmt.Sprintf("%v", config), ctx...)
}

func errDuplicateWatchdog(
	id string,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrDuplicateWatchdog, id, ctx...)
}

// ----------------------------------------------------------------------------
// watchdog log formatter
// ----------------------------------------------------------------------------

// WatchdogLogFormatter defines an interface to a watchdog
// logging message formatter.
type WatchdogLogFormatter interface {
	Start(service string) string
	Error(service string, e error) string
	Done(service string) string
}

// ----------------------------------------------------------------------------
// watchdog log formatter creator
// ----------------------------------------------------------------------------

// WatchdogLogFormatterCreator defines a watchdog logging message
// formatter creator service.
type WatchdogLogFormatterCreator interface {
	Accept(config *ConfigPartial) bool
	Create(config *ConfigPartial) (WatchdogLogFormatter, error)
}

// ----------------------------------------------------------------------------
// watchdog log formatter factory
// ----------------------------------------------------------------------------

// WatchdogLogFormatterFactory defines an object responsible
// to instantiate a new watchdog log formatter.
type WatchdogLogFormatterFactory []WatchdogLogFormatterCreator

// NewWatchdogLogFormatterFactory will instantiate a new logging formatter
// creator instance.
func NewWatchdogLogFormatterFactory(
	creators []WatchdogLogFormatterCreator,
) *WatchdogLogFormatterFactory {
	factory := &WatchdogLogFormatterFactory{}
	for _, creator := range creators {
		*factory = append(*factory, creator)
	}
	return factory
}

// Create will instantiate and return a new watchdog log formatter
// defined by the requested configuration data.
func (f *WatchdogLogFormatterFactory) Create(
	config *ConfigPartial,
) (WatchdogLogFormatter, error) {
	// check the config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// find a creator that accepts the requested log formatter type
	for _, creator := range *f {
		if creator.Accept(config) {
			// create the requested config log formatter
			return creator.Create(config)
		}
	}
	return nil, errInvalidWatchdogLogWriter(config)
}

// ----------------------------------------------------------------------------
// watchdog default log formatter
// ----------------------------------------------------------------------------

// WatchdogDefaultLogFormatter defines an instance to a watchdog
// logging message formatter.
type WatchdogDefaultLogFormatter struct{}

var _ WatchdogLogFormatter = &WatchdogDefaultLogFormatter{}

// NewWatchdogDefaultLogFormatter will instantiate a new default
// watchdog logging message formatter
func NewWatchdogDefaultLogFormatter() *WatchdogDefaultLogFormatter {
	return &WatchdogDefaultLogFormatter{}
}

// Start format a watchdog starting signal message.
func (WatchdogDefaultLogFormatter) Start(
	service string,
) string {
	return fmt.Sprintf(WatchdogLogStartMessage, service)
}

// Error format a watchdog error signal message.
func (WatchdogDefaultLogFormatter) Error(
	service string,
	e error,
) string {
	return fmt.Sprintf(WatchdogLogErrorMessage, service, e)
}

// Done format a watchdog termination signal message.
func (WatchdogDefaultLogFormatter) Done(
	service string,
) string {
	return fmt.Sprintf(WatchdogLogDoneMessage, service)
}

// ----------------------------------------------------------------------------
// watchdog default log formatter creator
// ----------------------------------------------------------------------------

// WatchdogDefaultLogFormatterCreator defines the default log formatter
// creator service.
type WatchdogDefaultLogFormatterCreator struct{}

var _ WatchdogLogFormatterCreator = &WatchdogDefaultLogFormatterCreator{}

// NewWatchdogDefaultLogFormatterCreator will instantiate a new default logging
// formatter creator service.
func NewWatchdogDefaultLogFormatterCreator() *WatchdogDefaultLogFormatterCreator {
	return &WatchdogDefaultLogFormatterCreator{}
}

// Accept will check if the creator is able to create the requested
// log formatter based on the given configuration.
func (s WatchdogDefaultLogFormatterCreator) Accept(
	config *ConfigPartial,
) bool {
	// check the config argument reference
	if config == nil {
		return false
	}
	// retrieve the data from the configuration
	sConfig := struct{ Type string }{
		Type: WatchdogLogFormatterTypeDefault,
	}
	if _, e := config.Populate("", &sConfig); e != nil {
		return false
	}
	// return acceptance for the read config type
	return sConfig.Type == WatchdogLogFormatterTypeDefault
}

// Create will try to generate a log formatter based on the
// passed configuration.
func (s WatchdogDefaultLogFormatterCreator) Create(
	_ *ConfigPartial,
) (WatchdogLogFormatter, error) {
	// return the default log formatter instance
	return &WatchdogDefaultLogFormatter{}, nil
}

// ----------------------------------------------------------------------------
// watchdog log adapter
// ----------------------------------------------------------------------------

type watchdogLogger interface {
	Signal(channel string, level LogLevel, msg string, ctx ...LogContext) error
}

// WatchdogLogAdapter define an instance a watchdog logging adapter.
type WatchdogLogAdapter struct {
	name       string
	channel    string
	startLevel LogLevel
	errorLevel LogLevel
	doneLevel  LogLevel
	logger     watchdogLogger
	formatter  WatchdogLogFormatter
}

// NewWatchdogLogAdapter will create a new watchdog logging adapter.
func NewWatchdogLogAdapter(
	name,
	channel string,
	startLevel,
	errorLevel,
	doneLevel LogLevel,
	logger *Log,
	formatter WatchdogLogFormatter,
) (*WatchdogLogAdapter, error) {
	// check log argument instance
	if logger == nil {
		return nil, errNilPointer("logger")
	}
	// check log formatter argument instance
	if formatter == nil {
		return nil, errNilPointer("formatter")
	}
	// return the created log adapter instance
	return &WatchdogLogAdapter{
		name:       name,
		channel:    channel,
		startLevel: startLevel,
		errorLevel: errorLevel,
		doneLevel:  doneLevel,
		logger:     logger,
		formatter:  formatter,
	}, nil
}

// Start will format and redirect the start logging message to
// the application logger.
func (a *WatchdogLogAdapter) Start() error {
	// propagate the logging signal to the adapter stored log instance
	return a.logger.Signal(a.channel, a.startLevel, a.formatter.Start(a.name))
}

// Error will format and redirect the error logging message to
// the application logger.
func (a *WatchdogLogAdapter) Error(
	e error,
) error {
	// propagate the logging signal to the adapter stored log instance
	return a.logger.Signal(a.channel, a.errorLevel, a.formatter.Error(a.name, e))
}

// Done will format and redirect the termination logging message to
// the application logger.
func (a *WatchdogLogAdapter) Done() error {
	// propagate the logging signal to the adapter stored log instance
	return a.logger.Signal(a.channel, a.doneLevel, a.formatter.Done(a.name))
}

// ----------------------------------------------------------------------------
// watchdog processor
// ----------------------------------------------------------------------------

// WatchdogProcessor defines an interface to a watchdog process.
type WatchdogProcessor interface {
	Service() string
	Runner() func() error
}

// ----------------------------------------------------------------------------
// watchdog
// ----------------------------------------------------------------------------

// Watchdog defines the instance used to overlook a process execution.
type Watchdog struct {
	logAdapter *WatchdogLogAdapter
}

// NewWatchdog generates a new watchdog instance.
func NewWatchdog(
	logAdapter *WatchdogLogAdapter,
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
func (w *Watchdog) Run(
	process WatchdogProcessor,
) (e error) {
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

// ----------------------------------------------------------------------------
// watchdog factory
// ----------------------------------------------------------------------------

// WatchdogFactory defines an instance of a watchdog creator, used
// to create watchdogs related to a configuration entry.
type WatchdogFactory struct {
	config              *Config
	log                 *Log
	logFormatterFactory *WatchdogLogFormatterFactory
}

// NewWatchdogFactory will generate a new watchdog creator instance.
func NewWatchdogFactory(
	config *Config,
	logger *Log,
	logFormatterFactory *WatchdogLogFormatterFactory,
) (*WatchdogFactory, error) {
	// check config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// check log argument reference
	if logger == nil {
		return nil, errNilPointer("logger")
	}
	// check formatter creator reference
	if logFormatterFactory == nil {
		return nil, errNilPointer("logFormatterFactory")
	}
	// return the created watchdog creator instance
	return &WatchdogFactory{
		config:              config,
		log:                 logger,
		logFormatterFactory: logFormatterFactory,
	}, nil
}

// Create will create a new watchdog instance for the required
// service with the name passed as argument.
func (f *WatchdogFactory) Create(
	service string,
) (*Watchdog, error) {
	// get service watchdog configuration
	config, e := f.config.Partial(
		fmt.Sprintf("%s.%s", WatchdogConfigPathPrefix, service),
		ConfigPartial{},
	)
	if e != nil {
		return nil, e
	}
	// parse the retrieved configuration
	wc := struct {
		Name    string
		Channel string
		Level   struct {
			Start string
			Error string
			Done  string
		}
		Formatter string
	}{
		Name:    service,
		Channel: WatchdogLogChannel,
		Level: struct {
			Start string
			Error string
			Done  string
		}{
			Start: WatchdogLogStartLevel,
			Error: WatchdogLogErrorLevel,
			Done:  WatchdogLogDoneLevel,
		},
		Formatter: WatchdogDefaultFormatter,
	}
	if _, e = config.Populate("", &wc); e != nil {
		return nil, e
	}
	// validate the logging levels read from config
	startLevel, ok := LogLevelMap[wc.Level.Start]
	if !ok {
		return nil, errConversion(wc.Level.Start, "log.Level")
	}
	errorLevel, ok := LogLevelMap[wc.Level.Error]
	if !ok {
		return nil, errConversion(wc.Level.Error, "log.Level")
	}
	doneLevel, ok := LogLevelMap[wc.Level.Done]
	if !ok {
		return nil, errConversion(wc.Level.Done, "log.Level")
	}
	// obtain the formatter for the watchdog log adapter
	formatter, e := f.logFormatterFactory.Create(&ConfigPartial{"type": wc.Formatter})
	if e != nil {
		return nil, e
	}
	// generate the watchdog log adapter
	la, _ := NewWatchdogLogAdapter(wc.Name, wc.Channel, startLevel, errorLevel, doneLevel, f.log, formatter)
	// return the generated watchdog
	return NewWatchdog(la)
}

// ----------------------------------------------------------------------------
// watchdog process
// ----------------------------------------------------------------------------

// WatchdogProcess defines an instance to a watchdog process that will be
// overlooked by the watchdog.
type WatchdogProcess struct {
	service string
	runner  func() error
}

var _ WatchdogProcessor = &WatchdogProcess{}

// NewWatchdogProcess generate a new process instance with the given
// service name and runner method.
func NewWatchdogProcess(
	service string,
	runner func() error,
) (*WatchdogProcess, error) {
	// check runner function argument reference
	if runner == nil {
		return nil, errNilPointer("runner")
	}
	// return the created process instance
	return &WatchdogProcess{
		service: service,
		runner:  runner,
	}, nil
}

// Service will retrieve the service name.
func (p *WatchdogProcess) Service() string {
	return p.service
}

// Runner retrieve the process runner method.
func (p *WatchdogProcess) Runner() func() error {
	return p.runner
}

// ----------------------------------------------------------------------------
// watchdog kennel
// ----------------------------------------------------------------------------

type watchdogKennelReg struct {
	process  WatchdogProcessor
	watchdog *Watchdog
}

// WatchdogKennel define an instance that will manage a group of watchdog
// connections, and is used to run them in parallel.
type WatchdogKennel struct {
	watchdogFactory *WatchdogFactory
	regs            map[string]watchdogKennelReg
}

// NewWatchdogKennel will generate a new kennel instance.
func NewWatchdogKennel(
	watchdogFactory *WatchdogFactory,
	processes []WatchdogProcessor,
) (*WatchdogKennel, error) {
	// check creator argument reference
	if watchdogFactory == nil {
		return nil, errNilPointer("watchdogFactory")
	}
	// return the created creator instance
	kennel := &WatchdogKennel{
		watchdogFactory: watchdogFactory,
		regs:            map[string]watchdogKennelReg{},
	}
	// register all the processes
	for _, process := range processes {
		if e := kennel.add(process); e != nil {
			return nil, e
		}
	}
	return kennel, nil
}

// Run will execute all the registered processes in their
// respective watchdogs.
func (k *WatchdogKennel) Run() error {
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
		go func(reg watchdogKennelReg) {
			// run the process on a created watchdog
			if e := reg.watchdog.Run(reg.process); e != nil {
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

// Add will create a new watchdog instance that will guard the
// requested process instance.
func (k *WatchdogKennel) add(
	process WatchdogProcessor,
) error {
	// check if there is a watchdog for the requested service
	if _, ok := k.regs[process.Service()]; ok {
		return errDuplicateWatchdog(process.Service())
	}
	// create the watchdog for the requested process
	wd, e := k.watchdogFactory.Create(process.Service())
	if e != nil {
		return e
	}
	// store the process and the created watchdog in the kennel
	k.regs[process.Service()] = watchdogKennelReg{
		process:  process,
		watchdog: wd,
	}
	return nil
}

// ----------------------------------------------------------------------------
// watchdog service register
// ----------------------------------------------------------------------------

// WatchdogServiceRegister defines the service provider to be used on
// the application initialization to register the relational
// database services.
type WatchdogServiceRegister struct {
	ServiceRegister
}

var _ ServiceProvider = &WatchdogServiceRegister{}

// NewWatchdogServiceRegister will generate a new service registry instance
func NewWatchdogServiceRegister(
	app ...*App,
) *WatchdogServiceRegister {
	return &WatchdogServiceRegister{
		ServiceRegister: *NewServiceRegister(app...),
	}
}

// Provide will register the relational database module services in the
// application Provider.
func (sr WatchdogServiceRegister) Provide(
	container *ServiceContainer,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	// register the services
	_ = container.Add(WatchdogDefaultLogFormatterCreatorContainerID, NewWatchdogDefaultLogFormatterCreator, WatchdogLogFormatterCreatorTag)
	_ = container.Add(WatchdogAllLogFormatterCreatorsContainerID, sr.getLogFormattersCreators(container))
	_ = container.Add(WatchdogLogFormatterFactoryContainerID, NewWatchdogLogFormatterFactory)
	_ = container.Add(WatchdogFactoryContainerID, NewWatchdogFactory)
	_ = container.Add(WatchdogAllProcessesContainerID, sr.getAllProcesses(container))
	_ = container.Add(WatchdogContainerID, NewWatchdogKennel)
	return nil
}

func (WatchdogServiceRegister) getLogFormattersCreators(
	container *ServiceContainer,
) func() []WatchdogLogFormatterCreator {
	return func() []WatchdogLogFormatterCreator {
		// retrieve all the log formatters creators from the Provider
		var creators []WatchdogLogFormatterCreator
		entries, _ := container.Tag(WatchdogLogFormatterCreatorTag)
		for _, entry := range entries {
			// type check the retrieved service
			s, ok := entry.(WatchdogLogFormatterCreator)
			if ok {
				creators = append(creators, s)
			}
		}
		return creators
	}
}

func (WatchdogServiceRegister) getAllProcesses(
	container *ServiceContainer,
) func() []WatchdogProcessor {
	return func() []WatchdogProcessor {
		// retrieve all the watchdog processes from the Provider
		var processes []WatchdogProcessor
		entries, _ := container.Tag(WatchdogProcessTag)
		for _, entry := range entries {
			// type check the retrieved service
			s, ok := entry.(WatchdogProcessor)
			if ok {
				processes = append(processes, s)
			}
		}
		return processes
	}
}
