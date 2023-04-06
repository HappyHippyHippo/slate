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

	// LogFormatterStrategyTag defines the default tag to be used
	// to identify a watchdog log formatter entry in the container.
	LogFormatterStrategyTag = ID + ".formatter.strategy"

	// DefaultLogFormatterStrategyID defines the id to be used as the
	// container registration id of the default log formatter strategy.
	DefaultLogFormatterStrategyID = ID + ".formatter.strategy.default"

	// LogFormatterFactoryID defines the id to be used as the container
	// registration id of the log formatter factory.
	LogFormatterFactoryID = ID + ".formatter.factory"

	// FactoryID defines the id to be used as the container
	// registration id of the watchdog factory.
	FactoryID = ID + ".factory"

	// ProcessTag defines the default tag to be used
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
	_ = container.Service(DefaultLogFormatterStrategyID, NewDefaultLogFormatterStrategy, LogFormatterStrategyTag)
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
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	// populate the container log formatter factory with
	// all registered log formatter strategies
	formatterFactory, e := p.getLogFormatterFactory(container)
	if e != nil {
		return e
	}
	formatterStrategies, e := p.getLogFormatterStrategies(container)
	if e != nil {
		return e
	}
	for _, strategy := range formatterStrategies {
		_ = formatterFactory.Register(strategy)
	}
	// populate the watchdog kennel with all the processes
	// registered in the container
	kennel, e := p.getKennel(container)
	if e != nil {
		return e
	}
	procs, e := p.getProcesses(container)
	if e != nil {
		return e
	}
	for _, proc := range procs {
		if e := kennel.Add(proc); e != nil {
			return e
		}
	}
	return nil
}

func (Provider) getLogFormatterFactory(
	container slate.IContainer,
) (ILogFormatterFactory, error) {
	// retrieve the factory entry
	entry, e := container.Get(LogFormatterFactoryID)
	if e != nil {
		return nil, e
	}
	// validate the retrieved entry type
	instance, ok := entry.(ILogFormatterFactory)
	if !ok {
		return nil, errConversion(entry, "watchdog.ILogFormatterFactory")
	}
	return instance, nil
}

func (Provider) getLogFormatterStrategies(
	container slate.IContainer,
) ([]ILogFormatterStrategy, error) {
	// retrieve the strategies entries
	entries, e := container.Tag(LogFormatterStrategyTag)
	if e != nil {
		return nil, e
	}
	// type check the retrieved strategies
	var strategies []ILogFormatterStrategy
	for _, entry := range entries {
		if instance, ok := entry.(ILogFormatterStrategy); ok {
			strategies = append(strategies, instance)
		}
	}
	return strategies, nil
}

func (Provider) getKennel(
	container slate.IContainer,
) (IKennel, error) {
	// retrieve the kennel instance
	entry, e := container.Get(ID)
	if e != nil {
		return nil, e
	}
	// validate the retrieved entry type
	instance, ok := entry.(IKennel)
	if !ok {
		return nil, errConversion(entry, "watchdog.IKennel")
	}
	return instance, nil
}

func (p Provider) getProcesses(
	container slate.IContainer,
) ([]IProcess, error) {
	// retrieve the watchdog processes entries
	tags, e := container.Tag(ProcessTag)
	if e != nil {
		return nil, e
	}
	// type check the retrieved processes
	var processes []IProcess
	for _, service := range tags {
		if process, ok := service.(IProcess); ok {
			processes = append(processes, process)
		}
	}
	return processes, nil
}
