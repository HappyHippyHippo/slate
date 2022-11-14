package config

import (
	"github.com/spf13/afero"
)

type sourceStrategyDir struct {
	fs             afero.Fs
	decoderFactory IDecoderFactory
}

var _ ISourceStrategy = &sourceStrategyDir{}

func newSourceStrategyDir(fs afero.Fs, decoderFactory IDecoderFactory) (ISourceStrategy, error) {
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}

	return &sourceStrategyDir{
		fs:             fs,
		decoderFactory: decoderFactory,
	}, nil
}

// Accept will check if the source decoderFactory strategy can instantiate a
// new source of the requested type.
func (sourceStrategyDir) Accept(sourceType string) bool {
	return sourceType == SourceTypeDirectory
}

// AcceptFromConfig will check if the source decoderFactory strategy can instantiate
// a source where the data to check comes from a configuration Partial
// instance.
func (s sourceStrategyDir) AcceptFromConfig(cfg IConfig) bool {
	if cfg == nil {
		return false
	}

	if sourceType, e := cfg.String("type"); e == nil {
		return s.Accept(sourceType)
	}

	return false
}

// Create will instantiate the desired file source instance.
func (s sourceStrategyDir) Create(args ...interface{}) (ISource, error) {
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
		return NewSourceDir(path, format, recursive, s.fs, s.decoderFactory)
	}
}

// CreateFromConfig will instantiate the desired file source instance where
// the initialization data comes from a configuration Partial instance.
func (s sourceStrategyDir) CreateFromConfig(cfg IConfig) (ISource, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	if path, e := cfg.String("path"); e != nil {
		return nil, e
	} else if format, e := cfg.String("format", DefaultFileFormat); e != nil {
		return nil, e
	} else if recursive, e := cfg.Bool("recursive", true); e != nil {
		return nil, e
	} else {
		return s.Create(path, format, recursive)
	}
}