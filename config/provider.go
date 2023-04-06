package config

import (
	"time"

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

	// YAMLDecoderStrategyID defines the id to be used as the
	// container registration id of a yaml config decoder factory strategy
	// instance.
	YAMLDecoderStrategyID = ID + ".decoder.strategy.yaml"

	// JSONDecoderStrategyID defines the id to be used as the
	// container registration id of a json config decoder factory strategy
	// instance.
	JSONDecoderStrategyID = ID + ".decoder.strategy.json"

	// DecoderFactoryID defines the id to be used as the
	// container registration id of a config decoder factory instance.
	DecoderFactoryID = ID + ".decoder.factory"

	// SourceStrategyTag defines the tag to be assigned to all
	// container source strategies.
	SourceStrategyTag = ID + ".source.strategy"

	// EnvSourceStrategyID defines the id to be used as
	// the container registration id of a config environment source
	// factory strategy instance.
	EnvSourceStrategyID = ID + ".source.strategy.env"

	// FileSourceStrategyID defines the id to be used as the
	// container registration id of a config file source factory strategy
	// instance.
	FileSourceStrategyID = ID + ".source.strategy.file"

	// ObservableFileSourceStrategyID defines the id to be used
	// as the container registration id of a config observable file source
	// factory strategy instance.
	ObservableFileSourceStrategyID = ID + ".source.strategy.observable_file"

	// DirSourceStrategyID defines the id to be used as the
	// container registration id of a config dir source factory strategy
	// instance.
	DirSourceStrategyID = ID + ".source.strategy.dir"

	// RestSourceStrategyID defines the id to be used as the
	// container registration id of a config rest source factory strategy
	// instance.
	RestSourceStrategyID = ID + ".source.strategy.rest"

	// ObservableRestSourceStrategyID defines the id to be used
	// as the container registration id of a config observable rest source
	// factory strategy instance.
	ObservableRestSourceStrategyID = ID + ".source.strategy.observable_rest"

	// AggregateSourceStrategyID defines the id to be used as
	// the container registration id of a container loading config source
	// factory strategy instance.
	AggregateSourceStrategyID = ID + ".source.strategy.aggregate"

	// AggregateSourceTag defines the tag to be assigned
	// to all container defined config partials.
	AggregateSourceTag = ID + ".source.aggregate.tag"

	// SourceFactoryID defines the id to be used as the
	// container registration id config source factory instance.
	SourceFactoryID = ID + ".source.factory"

	// LoaderID defines the id to be used as the container
	// registration id of a config loader instance.
	LoaderID = ID + ".loader"
)

// Provider defines the slate.config module service provider to be used
// on the application initialization to register the config service.
type Provider struct{}

var _ slate.IProvider = &Provider{}

// Register will register the configuration section instances in the
// application container.
func (Provider) Register(
	container slate.IContainer,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	// add decoder strategies abd factory
	_ = container.Service(YAMLDecoderStrategyID, NewYAMLDecoderStrategy, DecoderStrategyTag)
	_ = container.Service(JSONDecoderStrategyID, NewJSONDecoderStrategy, DecoderStrategyTag)
	_ = container.Service(DecoderFactoryID, NewDecoderFactory)
	// add source strategies and factory
	_ = container.Service(EnvSourceStrategyID, NewEnvSourceStrategy, SourceStrategyTag)
	_ = container.Service(FileSourceStrategyID, NewFileSourceStrategy, SourceStrategyTag)
	_ = container.Service(ObservableFileSourceStrategyID, NewObservableFileSourceStrategy, SourceStrategyTag)
	_ = container.Service(DirSourceStrategyID, NewDirSourceStrategy, SourceStrategyTag)
	_ = container.Service(RestSourceStrategyID, NewRestSourceStrategy, SourceStrategyTag)
	_ = container.Service(ObservableRestSourceStrategyID, NewObservableRestSourceStrategy, SourceStrategyTag)
	_ = container.Service(AggregateSourceStrategyID, func() *AggregateSourceStrategy {
		// get all the registered config partials
		tagged, _ := container.Tag(AggregateSourceTag)
		var configs []IConfig
		for _, t := range tagged {
			if p, ok := t.(IConfig); ok {
				configs = append(configs, p)
			}
		}
		// allocate the new source strategy with all retrieved partials
		return &AggregateSourceStrategy{configs: configs}
	}, SourceStrategyTag)
	_ = container.Service(SourceFactoryID, NewSourceFactory)
	// add manager and loader
	_ = container.Service(ID, func() IManager {
		return NewManager(time.Duration(ObserveFrequency) * time.Second)
	})
	_ = container.Service(LoaderID, NewLoader)
	return nil
}

// Boot will start the configuration config instance by calling the
// configuration loader with the defined provider base entry information.
func (p Provider) Boot(
	container slate.IContainer,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	// populate the container decoder factory with all registered decoder strategies
	decoderFactory, e := p.getDecoderFactory(container)
	if e != nil {
		return e
	}
	decoderStrategies, e := p.getDecoderStrategies(container)
	if e != nil {
		return e
	}
	for _, strategy := range decoderStrategies {
		_ = decoderFactory.Register(strategy)
	}
	// populate the container source factory with all registered source strategies
	sourceFactory, e := p.getSourceFactory(container)
	if e != nil {
		return e
	}
	sourceStrategies, e := p.getSourceStrategies(container)
	if e != nil {
		return e
	}
	for _, strategy := range sourceStrategies {
		_ = sourceFactory.Register(strategy)
	}
	// check if the config loader is active
	if !LoaderActive {
		return nil
	}
	// get the container registered loader
	loader, e := p.getLoader(container)
	if e != nil {
		return e
	}
	// execute the loader action
	return loader.Load()
}

func (Provider) getDecoderFactory(
	container slate.IContainer,
) (IDecoderFactory, error) {
	// retrieve the factory entry
	entry, e := container.Get(DecoderFactoryID)
	if e != nil {
		return nil, e
	}
	// validate the retrieved entry type
	instance, ok := entry.(IDecoderFactory)
	if !ok {
		return nil, errConversion(entry, "config.IDecoderFactory")
	}
	return instance, nil
}

func (Provider) getDecoderStrategies(
	container slate.IContainer,
) ([]IDecoderStrategy, error) {
	// retrieve the strategies entries
	entries, e := container.Tag(DecoderStrategyTag)
	if e != nil {
		return nil, e
	}
	// type check the retrieved strategies
	var strategies []IDecoderStrategy
	for _, entry := range entries {
		if instance, ok := entry.(IDecoderStrategy); ok {
			strategies = append(strategies, instance)
		}
	}
	return strategies, nil
}

func (Provider) getSourceFactory(
	container slate.IContainer,
) (ISourceFactory, error) {
	// retrieve the factory entry
	entry, e := container.Get(SourceFactoryID)
	if e != nil {
		return nil, e
	}
	// validate the retrieved entry type
	instance, ok := entry.(ISourceFactory)
	if !ok {
		return nil, errConversion(entry, "config.ISourceFactory")
	}
	return instance, nil
}

func (Provider) getSourceStrategies(
	container slate.IContainer,
) ([]ISourceStrategy, error) {
	// retrieve the strategies entries
	entries, e := container.Tag(SourceStrategyTag)
	if e != nil {
		return nil, e
	}
	// type check the retrieved strategies
	var strategies []ISourceStrategy
	for _, entry := range entries {
		if instance, ok := entry.(ISourceStrategy); ok {
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
		return nil, errConversion(entry, "config.ILoader")
	}
	return instance, nil
}
