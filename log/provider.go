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
	// container Formatter strategies.
	FormatterStrategyTag = ID + ".Formatter.strategy"

	// FormatterFactoryID defines the id to be used as the
	// container registration id of a logger Formatter factory instance.
	FormatterFactoryID = ID + ".Formatter.factory"

	// StreamStrategyTag defines the tag to be assigned to all
	// container stream strategies.
	StreamStrategyTag = ID + ".stream.strategy"

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
	container slate.IContainer,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	_ = container.Service(FormatterFactoryID, NewFormatterFactory)
	_ = container.Service(StreamFactoryID, NewStreamFactory)
	_ = container.Service(ID, NewLog)
	_ = container.Service(LoaderID, NewLoader)
	return nil
}

// Boot will start the logger package config instance by calling the
// logger loader with the defined provider base entry information.
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

	// populate the container Formatter factory with
	// all registered Formatter strategies
	formatterFactory := p.getFormatterFactory(container)
	for _, s := range p.getFormatterStrategies(container) {
		_ = formatterFactory.Register(s)
	}
	// populate the container stream factory with all
	// registered stream strategies
	streamFactory := p.getStreamFactory(container)
	for _, s := range p.getStreamStrategies(container) {
		_ = streamFactory.Register(s)
	}
	// check if the log loader is active
	if !LoaderActive {
		return nil
	}
	// execute the loader action
	return p.getLoader(container).Load()
}

func (Provider) getFormatterFactory(
	container slate.IContainer,
) IFormatterFactory {
	// retrieve the factory entry
	entry, e := container.Get(FormatterFactoryID)
	if e != nil {
		panic(e)
	}
	// validate the retrieved entry type
	instance, ok := entry.(IFormatterFactory)
	if !ok {
		panic(errConversion(entry, "log.IFormatterFactory"))
	}
	return instance
}

func (Provider) getFormatterStrategies(
	container slate.IContainer,
) []IFormatterStrategy {
	// retrieve the strategies entries
	entries, e := container.Tag(FormatterStrategyTag)
	if e != nil {
		panic(e)
	}
	// type check the retrieved strategies
	var strategies []IFormatterStrategy
	for _, entry := range entries {
		s, ok := entry.(IFormatterStrategy)
		if !ok {
			panic(errConversion(entry, "log.IFormatterStrategy"))
		}
		strategies = append(strategies, s)
	}
	return strategies
}

func (Provider) getStreamFactory(
	container slate.IContainer,
) IStreamFactory {
	// retrieve the factory entry
	entry, e := container.Get(StreamFactoryID)
	if e != nil {
		panic(e)
	}
	// validate the retrieved entry type
	instance, ok := entry.(IStreamFactory)
	if !ok {
		panic(errConversion(entry, "log.IStreamFactory"))
	}
	return instance
}

func (Provider) getStreamStrategies(
	container slate.IContainer,
) []IStreamStrategy {
	// retrieve the strategies entries
	entries, e := container.Tag(StreamStrategyTag)
	if e != nil {
		panic(e)
	}
	// type check the retrieved strategies
	var strategies []IStreamStrategy
	for _, entry := range entries {
		s, ok := entry.(IStreamStrategy)
		if !ok {
			panic(errConversion(entry, "log.IStreamStrategy"))
		}
		strategies = append(strategies, s)
	}
	return strategies
}

func (Provider) getLoader(
	container slate.IContainer,
) ILoader {
	// retrieve the loader entry
	entry, e := container.Get(LoaderID)
	if e != nil {
		panic(e)
	}
	// validate the retrieved entry type
	instance, ok := entry.(ILoader)
	if !ok {
		panic(errConversion(entry, "log.ILoader"))
	}
	return instance
}
