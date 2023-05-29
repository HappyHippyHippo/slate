package aggregate

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/config/source"
)

const (
	// ID defines the application container registration string for the
	// aggregate source creation strategy.
	ID = source.ID + ".aggregate"

	// SourceTag defines the tag to be assigned
	// to all container defined config sources to be aggregated.
	SourceTag = ID + ".tag"
)

// Provider defines the slate.config module service provider to be used
// on the application initialization to register the config service.
type Provider struct{}

var _ slate.Provider = &Provider{}

// Register will register the configuration section instances in the
// application container.
func (Provider) Register(
	container *slate.Container,
) error {
	if container == nil {
		return errNilPointer("container")
	}
	_ = container.Service(ID, func() *SourceStrategy {
		// get all the registered config partials
		tagged, _ := container.Tag(SourceTag)
		var sources []config.Source
		for _, t := range tagged {
			if p, ok := t.(config.Source); ok {
				sources = append(sources, p)
			}
		}
		// allocate the new source strategy with all retrieved partials
		return &SourceStrategy{sources: sources}
	}, config.SourceStrategyTag)
	return nil
}

// Boot will start the configuration config instance by calling the
// configuration loader with the defined provider base entry information.
func (p Provider) Boot(
	container *slate.Container,
) error {
	if container == nil {
		return errNilPointer("container")
	}
	return nil
}
