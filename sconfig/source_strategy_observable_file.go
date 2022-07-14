package sconfig

import (
	"github.com/spf13/afero"
)

type sourceStrategyObservableFile struct {
	sourceStrategyFile
}

var _ ISourceStrategy = &sourceStrategyObservableFile{}

func newSourceStrategyObservableFile(fs afero.Fs, factory IDecoderFactory) (ISourceStrategy, error) {
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	if factory == nil {
		return nil, errNilPointer("dFactory")
	}

	return &sourceStrategyObservableFile{
		sourceStrategyFile: sourceStrategyFile{
			fs:      fs,
			factory: factory,
		},
	}, nil
}

// Accept will check if the source dFactory strategy can instantiate a
// new source of the requested type.
func (sourceStrategyObservableFile) Accept(sourceType string) bool {
	return sourceType == SourceTypeObservableFile
}

// AcceptFromConfig will check if the source dFactory strategy can instantiate
// a source where the data to check comes from a configuration Partial
// instance.
func (s sourceStrategyObservableFile) AcceptFromConfig(cfg IConfig) bool {
	if cfg == nil {
		return false
	}

	if sourceType, err := cfg.String("type"); err == nil {
		return s.Accept(sourceType)
	}

	return false
}

// Create will instantiate the desired observable file source instance.
func (s sourceStrategyObservableFile) Create(args ...interface{}) (ISource, error) {
	if len(args) < 2 {
		return nil, errNilPointer("args")
	}

	if path, ok := args[0].(string); !ok {
		return nil, errConversion(args[0], "string")
	} else if format, ok := args[1].(string); !ok {
		return nil, errConversion(args[1], "string")
	} else {
		return newSourceObservableFile(path, format, s.fs, s.factory)
	}
}

// CreateFromConfig will instantiate the desired file source instance where
// the initialization data comes from a configuration Partial instance.
func (s sourceStrategyObservableFile) CreateFromConfig(cfg IConfig) (ISource, error) {
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
