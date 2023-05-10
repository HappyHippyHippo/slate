package watchdog

import (
	"github.com/happyhippyhippo/slate"
)

const (
	// ID defines the id to be used as the container
	// registration id of a watchdog factory instance, and as a
	// base id of all other watchdog package instances registered
	// in the application container.
	ID = slate.ID + ".watchdog"

	// LogFormatterStrategyTag defines the def tag to be used
	// to identify a watchdog log formatter entry in the container.
	LogFormatterStrategyTag = ID + ".formatter.strategy"

	// LogFormatterFactoryID defines the id to be used as the container
	// registration id of the log formatter factory.
	LogFormatterFactoryID = ID + ".formatter.factory"

	// FactoryID defines the id to be used as the container
	// registration id of the watchdog factory.
	FactoryID = ID + ".factory"

	// ProcessTag defines the def tag to be used
	// to identify a watchdog process entry in the container.
	ProcessTag = ID + ".process.tag"
)

// Provider defines the slate.watchdog module service provider to be used on
// the application initialization to register the migrations service.
type Provider struct{}

var _ slate.IProvider = &Provider{}

// Register will register the migration package instances in the
// application container
func (p Provider) Register(
	container slate.IContainer,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	// add log formatter strategies and factory
	_ = container.Service(LogFormatterFactoryID, NewLogFormatterFactory)
	// add the watchdog factory and kennel
	_ = container.Service(FactoryID, NewFactory)
	_ = container.Service(ID, NewKennel)
	return nil
}

// Boot will start the migration package
// If the auto migration is defined as true, ether by global variable or
// by environment variable, the migrator will automatically try to migrate
// to the last registered migration
func (p Provider) Boot(
	container slate.IContainer,
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

	// populate the container log formatter factory with
	// all registered log formatter strategies
	formatterFactory := p.getLogFormatterFactory(container)
	for _, strategy := range p.getLogFormatterStrategies(container) {
		_ = formatterFactory.Register(strategy)
	}
	// populate the watchdog kennel with all the processes
	// registered in the container
	kennel := p.getKennel(container)
	for _, proc := range p.getProcesses(container) {
		if e := kennel.Add(proc); e != nil {
			return e
		}
	}
	return nil
}

func (Provider) getLogFormatterFactory(
	container slate.IContainer,
) ILogFormatterFactory {
	// retrieve the factory entry
	entry, e := container.Get(LogFormatterFactoryID)
	if e != nil {
		panic(e)
	}
	// validate the retrieved entry type
	instance, ok := entry.(ILogFormatterFactory)
	if !ok {
		panic(errConversion(entry, "watchdog.ILogFormatterFactory"))
	}
	return instance
}

func (Provider) getLogFormatterStrategies(
	container slate.IContainer,
) []ILogFormatterStrategy {
	// retrieve the strategies entries
	entries, e := container.Tag(LogFormatterStrategyTag)
	if e != nil {
		panic(e)
	}
	// type check the retrieved strategies
	var strategies []ILogFormatterStrategy
	for _, entry := range entries {
		s, ok := entry.(ILogFormatterStrategy)
		if !ok {
			panic(errConversion(entry, "watchdog.ILogFormatterStrategy"))
		}
		strategies = append(strategies, s)
	}
	return strategies
}

func (Provider) getKennel(
	container slate.IContainer,
) IKennel {
	// retrieve the kennel instance
	entry, e := container.Get(ID)
	if e != nil {
		panic(e)
	}
	// validate the retrieved entry type
	instance, ok := entry.(IKennel)
	if !ok {
		panic(errConversion(entry, "watchdog.IKennel"))
	}
	return instance
}

func (p Provider) getProcesses(
	container slate.IContainer,
) []IProcess {
	// retrieve the watchdog processes entries
	tags, e := container.Tag(ProcessTag)
	if e != nil {
		panic(e)
	}
	// type check the retrieved processes
	var processes []IProcess
	for _, service := range tags {
		p, ok := service.(IProcess)
		if !ok {
			panic(errConversion(service, "watchdog.IProcess"))
		}
		processes = append(processes, p)
	}
	return processes
}
