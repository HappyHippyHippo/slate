package watchdog

import (
	"github.com/happyhippyhippo/slate"
)

const (
	// ID defines the id to be used as the container
	// registration id of a watchdog creator instance, and as a
	// base id of all other watchdog package instances registered
	// in the application container.
	ID = slate.ID + ".watchdog"

	// LogFormatterStrategyTag defines the simple tag to be used
	// to identify a watchdog log formatter entry in the container.
	LogFormatterStrategyTag = ID + ".formatter.strategy"

	// LogFormatterFactoryID defines the id to be used as the container
	// registration id of the log formatter creator.
	LogFormatterFactoryID = ID + ".formatter.creator"

	// WatchdogFactoryID defines the id to be used as the container
	// registration id of the watchdog creator.
	WatchdogFactoryID = ID + ".creator"

	// ProcessTag defines the simple tag to be used
	// to identify a watchdog process entry in the container.
	ProcessTag = ID + ".process.tag"
)

// Provider defines the slate.watchdog module service provider to be used on
// the application initialization to register the migrations service.
type Provider struct{}

var _ slate.Provider = &Provider{}

// Register will register the migration package instances in the
// application container
func (p Provider) Register(
	container *slate.Container,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	// add log formatter strategies and creator
	_ = container.Service(LogFormatterFactoryID, NewLogFormatterFactory)
	// add the watchdog creator and kennel
	_ = container.Service(WatchdogFactoryID, NewFactory)
	_ = container.Service(ID, NewKennel)
	return nil
}

// Boot will start the migration package
// If the auto migration is defined as true, ether by global variable or
// by environment variable, the migrator will automatically try to migrate
// to the last registered migration
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

	// populate the container log formatter creator with
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
	container *slate.Container,
) *LogFormatterFactory {
	// retrieve the creator entry
	entry, e := container.Get(LogFormatterFactoryID)
	if e != nil {
		panic(e)
	}
	// validate the retrieved entry type
	if instance, ok := entry.(*LogFormatterFactory); ok {
		return instance
	}
	panic(errConversion(entry, "*watchdog.LogFormatterFactory"))
}

func (Provider) getLogFormatterStrategies(
	container *slate.Container,
) []LogFormatterStrategy {
	// retrieve the strategies entries
	entries, e := container.Tag(LogFormatterStrategyTag)
	if e != nil {
		panic(e)
	}
	// type check the retrieved strategies
	var strategies []LogFormatterStrategy
	for _, entry := range entries {
		s, ok := entry.(LogFormatterStrategy)
		if !ok {
			panic(errConversion(entry, "watchdog.LogFormatterStrategy"))
		}
		strategies = append(strategies, s)
	}
	return strategies
}

func (Provider) getKennel(
	container *slate.Container,
) *Kennel {
	// retrieve the kennel instance
	entry, e := container.Get(ID)
	if e != nil {
		panic(e)
	}
	// validate the retrieved entry type
	if instance, ok := entry.(*Kennel); ok {
		return instance
	}
	panic(errConversion(entry, "*watchdog.Kennel"))
}

func (p Provider) getProcesses(
	container *slate.Container,
) []Processor {
	// retrieve the watchdog processes entries
	tags, e := container.Tag(ProcessTag)
	if e != nil {
		panic(e)
	}
	// type check the retrieved processes
	var processes []Processor
	for _, service := range tags {
		p, ok := service.(Processor)
		if !ok {
			panic(errConversion(service, "watchdog.Processor"))
		}
		processes = append(processes, p)
	}
	return processes
}
