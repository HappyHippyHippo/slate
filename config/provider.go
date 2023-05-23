package config

import (
	"github.com/happyhippyhippo/slate"
)

const (
	// ID defines the id to be used as the container
	// registration id of a config instance, as a base id of all other config
	// package instances registered in the application container.
	ID = slate.ID + ".config"

	// DecoderStrategyTag defines the tag to be assigned to all
	// container decoders strategies.
	DecoderStrategyTag = ID + ".decoder.strategy"

	// DecoderFactoryID defines the id to be used as the
	// container registration id of a config decoder factory instance.
	DecoderFactoryID = ID + ".decoder.factory"

	// SourceStrategyTag defines the tag to be assigned to all
	// container source strategies.
	SourceStrategyTag = ID + ".source.strategy"

	// SourceFactoryID defines the id to be used as the
	// container registration id config source factory instance.
	SourceFactoryID = ID + ".source.factory"

	// LoaderID defines the id to be used as the container
	// registration id of a config observer instance.
	LoaderID = ID + ".observer"
)

// Provider defines the slate.config module service provider to be used
// on the application initialization to register the config service.
type Provider struct{}

var _ slate.Provider = &Provider{}

// Register will register the configuration module instances in the
// application container.
func (Provider) Register(
	container *slate.Container,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	// register the services
	_ = container.Service(DecoderFactoryID, NewDecoderFactory)
	_ = container.Service(SourceFactoryID, NewSourceFactory)
	_ = container.Service(ID, NewConfig)
	_ = container.Service(LoaderID, NewLoader)
	return nil
}

// Boot will start the config instance by calling the config
// observer with the defined provider base entry information.
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

	// populate the container decoder factory with all registered decoder strategies
	decoderFactory := p.getDecoderFactory(container)
	for _, s := range p.getDecoderStrategies(container) {
		_ = decoderFactory.Register(s)
	}
	// populate the container source factory with all registered source strategies
	sourceFactory := p.getSourceFactory(container)
	for _, s := range p.getSourceStrategies(container) {
		_ = sourceFactory.Register(s)
	}
	// check if the config observer is active
	if !LoaderActive {
		return nil
	}
	// execute the observer action
	return p.getLoader(container).Load()
}

func (Provider) getDecoderFactory(
	container *slate.Container,
) *DecoderFactory {
	// retrieve the factory entry
	entry, e := container.Get(DecoderFactoryID)
	if e != nil {
		panic(e)
	}
	// validate the retrieved entry type
	instance, ok := entry.(*DecoderFactory)
	if !ok {
		panic(errConversion(entry, "config.*DecoderFactory"))
	}
	return instance
}

func (Provider) getDecoderStrategies(
	container *slate.Container,
) []DecoderStrategy {
	// retrieve the strategies entries
	entries, e := container.Tag(DecoderStrategyTag)
	if e != nil {
		panic(e)
	}
	// type check the retrieved strategies
	var strategies []DecoderStrategy
	for _, entry := range entries {
		s, ok := entry.(DecoderStrategy)
		if !ok {
			panic(errConversion(entry, "config.DecoderStrategy"))
		}
		strategies = append(strategies, s)
	}
	return strategies
}

func (Provider) getSourceFactory(
	container *slate.Container,
) *SourceFactory {
	// retrieve the factory entry
	entry, e := container.Get(SourceFactoryID)
	if e != nil {
		panic(e)
	}
	// validate the retrieved entry type
	instance, ok := entry.(*SourceFactory)
	if !ok {
		panic(errConversion(entry, "config.*SourceFactory"))
	}
	return instance
}

func (Provider) getSourceStrategies(
	container *slate.Container,
) []SourceStrategy {
	// retrieve the strategies entries
	entries, e := container.Tag(SourceStrategyTag)
	if e != nil {
		panic(e)
	}
	// type check the retrieved strategies
	var strategies []SourceStrategy
	for _, entry := range entries {
		s, ok := entry.(SourceStrategy)
		if !ok {
			panic(errConversion(entry, "config.SourceStrategy"))
		}
		strategies = append(strategies, s)
	}
	return strategies
}

func (Provider) getLoader(
	container *slate.Container,
) *Loader {
	// retrieve the observer entry
	entry, e := container.Get(LoaderID)
	if e != nil {
		panic(e)
	}
	// validate the retrieved entry type
	instance, ok := entry.(*Loader)
	if !ok {
		panic(errConversion(entry, "config.loader"))
	}
	return instance
}
