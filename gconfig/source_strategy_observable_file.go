package gconfig

import (
	"github.com/spf13/afero"
)

// sourceStrategyObservableFile defines an observable config file
// source instantiation strategy to be used by the config sources factory
// instance.
type sourceStrategyObservableFile struct {
	sourceStrategyFile
}

var _ SourceStrategy = &sourceStrategyObservableFile{}

// NewSourceStrategyObservableFile instantiate a new observable
// file source factory strategy that will enable the source factory to
// instantiate a new observable file configuration source.
func NewSourceStrategyObservableFile(fs afero.Fs, factory *DecoderFactory) (SourceStrategy, error) {
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	if factory == nil {
		return nil, errNilPointer("factory")
	}

	return &sourceStrategyObservableFile{
		sourceStrategyFile: sourceStrategyFile{
			fs:      fs,
			factory: factory,
		},
	}, nil
}

// Accept will check if the source factory strategy can instantiate a
// new source of the requested type.
func (sourceStrategyObservableFile) Accept(stype string) bool {
	return stype == SourceTypeObservableFile
}

// AcceptFromConfig will check if the source factory strategy can instantiate
// a source where the data to check comes from a configuration Partial
// instance.
func (s sourceStrategyObservableFile) AcceptFromConfig(cfg Config) bool {
	if cfg == nil {
		return false
	}

	if stype, err := cfg.String("type"); err == nil {
		return s.Accept(stype)
	}

	return false
}

// Create will instantiate the desired observable file source instance.
func (s sourceStrategyObservableFile) Create(args ...interface{}) (Source, error) {
	if len(args) < 2 {
		return nil, errNilPointer("args")
	}

	if path, ok := args[0].(string); !ok {
		return nil, errConversion(args[0], "string")
	} else if format, ok := args[1].(string); !ok {
		return nil, errConversion(args[1], "string")
	} else {
		return NewSourceObservableFile(path, format, s.fs, s.factory)
	}
}

// CreateFromConfig will instantiate the desired file source instance where
// the initialization data comes from a configuration Partial instance.
func (s sourceStrategyObservableFile) CreateFromConfig(cfg Config) (Source, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	if path, err := cfg.String("path"); err != nil {
		return nil, err
	} else if format, err := cfg.String("format", DefaultFileFormat); err != nil {
		return nil, err
	} else {
		return s.Create(path, format)
	}
}
