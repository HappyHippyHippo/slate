package sconfig

import (
	"github.com/spf13/afero"
)

// sourceStrategyDir defines a config file source instantiation
// strategy to be used by the config sources factory instance.
type sourceStrategyDir struct {
	fs      afero.Fs
	factory *DecoderFactory
}

var _ SourceStrategy = &sourceStrategyDir{}

// NewSourceStrategyDir instantiate a new dir source factory
// strategy that will enable the source factory to instantiate file
// configuration sources from a system directory.
func NewSourceStrategyDir(fs afero.Fs, factory *DecoderFactory) (SourceStrategy, error) {
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	if factory == nil {
		return nil, errNilPointer("factory")
	}

	return &sourceStrategyDir{
		fs:      fs,
		factory: factory,
	}, nil
}

// Accept will check if the source factory strategy can instantiate a
// new source of the requested type.
func (sourceStrategyDir) Accept(stype string) bool {
	return stype == SourceTypeDirectory
}

// AcceptFromConfig will check if the source factory strategy can instantiate
// a source where the data to check comes from a configuration Partial
// instance.
func (s sourceStrategyDir) AcceptFromConfig(cfg Config) bool {
	if cfg == nil {
		return false
	}

	if stype, err := cfg.String("type"); err == nil {
		return s.Accept(stype)
	}

	return false
}

// Create will instantiate the desired file source instance.
func (s sourceStrategyDir) Create(args ...interface{}) (Source, error) {
	if len(args) < 3 {
		return nil, errNilPointer("args")
	}

	if path, ok := args[0].(string); !ok {
		return nil, errConversion(args[0], "string")
	} else if format, ok := args[1].(string); !ok {
		return nil, errConversion(args[1], "string")
	} else if recursive, ok := args[2].(bool); !ok {
		return nil, errConversion(args[2], "bool")
	} else {
		return NewSourceDir(path, format, recursive, s.fs, s.factory)
	}
}

// CreateFromConfig will instantiate the desired file source instance where
// the initialization data comes from a configuration Partial instance.
func (s sourceStrategyDir) CreateFromConfig(cfg Config) (Source, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	if path, err := cfg.String("path"); err != nil {
		return nil, err
	} else if format, err := cfg.String("format", DefaultFileFormat); err != nil {
		return nil, err
	} else if recursive, err := cfg.Bool("recursive", true); err != nil {
		return nil, err
	} else {
		return s.Create(path, format, recursive)
	}
}
