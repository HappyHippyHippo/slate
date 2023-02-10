package log

import (
	"github.com/happyhippyhippo/slate"
)

const (
	// ID defines the id to be used as the container
	// registration id of a logger instance, as a base id of all other log
	// package instances registered in the application container.
	ID = slate.ID + ".log"

	// FormatterStrategyTag defines the tag to be assigned to all
	// container formatter strategies.
	FormatterStrategyTag = ID + ".formatter.strategy"

	// JSONFormatterStrategyID defines the id to be used as
	// the container registration id of a logger json formatter factory
	// strategy instance.
	JSONFormatterStrategyID = ID + ".formatter.strategy.json"

	// FormatterFactoryID defines the id to be used as the
	// container registration id of a logger formatter factory instance.
	FormatterFactoryID = ID + ".formatter.factory"

	// StreamStrategyTag defines the tag to be assigned to all
	// container stream strategies.
	StreamStrategyTag = ID + ".stream.strategy"

	// ConsoleStreamStrategyID defines the id to be used as the
	// container registration id of a logger console stream factory strategy
	// instance.
	ConsoleStreamStrategyID = ID + ".stream.strategy.console"

	// FileStreamStrategyID defines the id to be used as the
	// container registration id of a logger file stream factory strategy
	// instance.
	FileStreamStrategyID = ID + ".stream.strategy.file"

	// RotatingFileStreamStrategyID defines the id to be used as the
	// container registration id of a logger rotating file stream factory
	// strategy instance.
	RotatingFileStreamStrategyID = ID + ".stream.strategy.rotating_file"

	// StreamFactoryID defines the id to be used as the container
	// registration id of a logger stream factory instance.
	StreamFactoryID = ID + ".stream.factory"

	// LoaderID defines the id to be used as the container
	// registration id of a logger loader instance.
	LoaderID = ID + ".loader"
)

// Provider defines the slate.log module service provider to be used on
// the application initialization to register the logging service.
type Provider struct{}

var _ slate.IProvider = &Provider{}

// Register will register the logger package instances in the
// application container.
func (p Provider) Register(
	container ...slate.IContainer,
) error {
	// check container argument reference
	if len(container) == 0 || container[0] == nil {
		return errNilPointer("container")
	}
	// add formatter strategies and factory
	_ = container[0].Service(JSONFormatterStrategyID, NewJSONFormatterStrategy, FormatterStrategyTag)
	_ = container[0].Service(FormatterFactoryID, NewFormatterFactory)
	// add stream strategies and factory
	_ = container[0].Service(ConsoleStreamStrategyID, NewConsoleStreamStrategy, StreamStrategyTag)
	_ = container[0].Service(FileStreamStrategyID, NewFileStreamStrategy, StreamStrategyTag)
	_ = container[0].Service(RotatingFileStreamStrategyID, NewRotatingFileStreamStrategy, StreamStrategyTag)
	_ = container[0].Service(StreamFactoryID, NewStreamFactory)
	// add log and loader
	_ = container[0].Service(ID, NewLog)
	_ = container[0].Service(LoaderID, NewLoader)
	return nil
}

// Boot will start the logger package config instance by calling the
// logger loader with the defined provider base entry information.
func (p Provider) Boot(
	container ...slate.IContainer,
) error {
	// check container argument reference
	if len(container) == 0 || container[0] == nil {
		return errNilPointer("container")
	}
	// populate the container formatter factory with
	// all registered formatter strategies
	formatterFactory, e := p.getFormatterFactory(container[0])
	if e != nil {
		return e
	}
	formatterStrategies, e := p.getFormatterStrategies(container[0])
	if e != nil {
		return e
	}
	for _, strategy := range formatterStrategies {
		_ = formatterFactory.Register(strategy)
	}
	// populate the container stream factory with all
	// registered stream strategies
	streamFactory, e := p.getStreamFactory(container[0])
	if e != nil {
		return e
	}
	streamStrategies, e := p.getStreamStrategies(container[0])
	if e != nil {
		return e
	}
	for _, strategy := range streamStrategies {
		_ = streamFactory.Register(strategy)
	}
	// check if the log loader is active
	if !LoaderActive {
		return nil
	}
	// get the container registered loader
	loader, e := p.getLoader(container[0])
	if e != nil {
		return e
	}
	// execute the loader action
	return loader.Load()
}

func (Provider) getFormatterFactory(
	container slate.IContainer,
) (IFormatterFactory, error) {
	// retrieve the factory entry
	entry, e := container.Get(FormatterFactoryID)
	if e != nil {
		return nil, e
	}
	// validate the retrieved entry type
	instance, ok := entry.(IFormatterFactory)
	if !ok {
		return nil, errConversion(entry, "log.IFormatterFactory")
	}
	return instance, nil
}

func (Provider) getFormatterStrategies(
	container slate.IContainer,
) ([]IFormatterStrategy, error) {
	// retrieve the strategies entries
	entries, e := container.Tag(FormatterStrategyTag)
	if e != nil {
		return nil, e
	}
	// type check the retrieved strategies
	var strategies []IFormatterStrategy
	for _, entry := range entries {
		if instance, ok := entry.(IFormatterStrategy); ok {
			strategies = append(strategies, instance)
		}
	}
	return strategies, nil
}

func (Provider) getStreamFactory(
	container slate.IContainer,
) (IStreamFactory, error) {
	// retrieve the factory entry
	entry, e := container.Get(StreamFactoryID)
	if e != nil {
		return nil, e
	}
	// validate the retrieved entry type
	instance, ok := entry.(IStreamFactory)
	if !ok {
		return nil, errConversion(entry, "log.IStreamFactory")
	}
	return instance, nil
}

func (Provider) getStreamStrategies(
	container slate.IContainer,
) ([]IStreamStrategy, error) {
	// retrieve the strategies entries
	entries, e := container.Tag(StreamStrategyTag)
	if e != nil {
		return nil, e
	}
	// type check the retrieved strategies
	var strategies []IStreamStrategy
	for _, entry := range entries {
		if instance, ok := entry.(IStreamStrategy); ok {
			strategies = append(strategies, instance)
		}
	}
	return strategies, nil
}

func (Provider) getLoader(
	container slate.IContainer,
) (ILoader, error) {
	// retrieve the loader entry
	entry, e := container.Get(LoaderID)
	if e != nil {
		return nil, e
	}
	// validate the retrieved entry type
	instance, ok := entry.(ILoader)
	if !ok {
		return nil, errConversion(entry, "log.ILoader")
	}
	return instance, nil
}
