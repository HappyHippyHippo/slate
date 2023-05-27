package log

import (
	"github.com/happyhippyhippo/slate"
)

const (
	// ID defines the id to be used as the container
	// registration id of a logger instance, as a base id of all other logger
	// package instances registered in the application container.
	ID = slate.ID + ".logger"

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

// Provider defines the slate.logger module service provider to be used on
// the application initialization to register the logging service.
type Provider struct{}

var _ slate.Provider = &Provider{}

// Register will register the logger package instances in the
// application container.
func (p Provider) Register(
	container *slate.Container,
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
	// check if the logger loader is active
	if !LoaderActive {
		return nil
	}
	// execute the loader action
	return p.getLoader(container).Load()
}

func (Provider) getFormatterFactory(
	container *slate.Container,
) *FormatterFactory {
	// retrieve the factory entry
	entry, e := container.Get(FormatterFactoryID)
	if e != nil {
		panic(e)
	}
	// validate the retrieved entry type
	instance, ok := entry.(*FormatterFactory)
	if !ok {
		panic(errConversion(entry, "*logger.FormatterFactory"))
	}
	return instance
}

func (Provider) getFormatterStrategies(
	container *slate.Container,
) []FormatterStrategy {
	// retrieve the strategies entries
	entries, e := container.Tag(FormatterStrategyTag)
	if e != nil {
		panic(e)
	}
	// type check the retrieved strategies
	var strategies []FormatterStrategy
	for _, entry := range entries {
		s, ok := entry.(FormatterStrategy)
		if !ok {
			panic(errConversion(entry, "logger.FormatterStrategy"))
		}
		strategies = append(strategies, s)
	}
	return strategies
}

func (Provider) getStreamFactory(
	container *slate.Container,
) *StreamFactory {
	// retrieve the factory entry
	entry, e := container.Get(StreamFactoryID)
	if e != nil {
		panic(e)
	}
	// validate the retrieved entry type
	instance, ok := entry.(*StreamFactory)
	if !ok {
		panic(errConversion(entry, "*logger.StreamFactory"))
	}
	return instance
}

func (Provider) getStreamStrategies(
	container *slate.Container,
) []StreamStrategy {
	// retrieve the strategies entries
	entries, e := container.Tag(StreamStrategyTag)
	if e != nil {
		panic(e)
	}
	// type check the retrieved strategies
	var strategies []StreamStrategy
	for _, entry := range entries {
		s, ok := entry.(StreamStrategy)
		if !ok {
			panic(errConversion(entry, "logger.StreamStrategy"))
		}
		strategies = append(strategies, s)
	}
	return strategies
}

func (Provider) getLoader(
	container *slate.Container,
) *Loader {
	// retrieve the loader entry
	entry, e := container.Get(LoaderID)
	if e != nil {
		panic(e)
	}
	// validate the retrieved entry type
	instance, ok := entry.(*Loader)
	if !ok {
		panic(errConversion(entry, "*logger.Loader"))
	}
	return instance
}
