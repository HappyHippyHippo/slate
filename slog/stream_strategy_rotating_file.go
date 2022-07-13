package slog

import (
	"github.com/happyhippyhippo/slate/sconfig"
	"github.com/spf13/afero"
)

type streamStrategyRotatingFile struct {
	streamStrategyFile
}

var _ IStreamStrategy = &streamStrategyRotatingFile{}

func newStreamStrategyRotatingFile(fs afero.Fs, factory IFormatterFactory) (IStreamStrategy, error) {
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	if factory == nil {
		return nil, errNilPointer("factory")
	}

	return &streamStrategyRotatingFile{
		streamStrategyFile: streamStrategyFile{
			fs:      fs,
			factory: factory,
		},
	}, nil
}

// Accept will check if the file stream factory strategy can instantiate a
// stream of the requested type and with the calling parameters.
func (streamStrategyRotatingFile) Accept(streamType string) bool {
	return streamType == StreamRotatingFile
}

// AcceptFromConfig will check if the stream factory strategy can instantiate
// a stream where the data to check comes from a configuration partial
// instance.
func (s streamStrategyRotatingFile) AcceptFromConfig(cfg sconfig.IConfig) bool {
	if cfg == nil {
		return false
	}

	if streamType, err := cfg.String("type"); err == nil {
		return s.Accept(streamType)
	}

	return false
}

// Create will instantiate the desired stream instance.
func (s streamStrategyRotatingFile) Create(args ...interface{}) (IStream, error) {
	if len(args) < 4 {
		return nil, errNilPointer("args")
	}

	if path, ok := args[0].(string); !ok {
		return nil, errConversion(args[0], "string")
	} else if format, ok := args[1].(string); !ok {
		return nil, errConversion(args[1], "string")
	} else if channels, ok := args[2].([]string); !ok {
		return nil, errConversion(args[2], "[]string")
	} else if level, ok := args[3].(Level); !ok {
		return nil, errConversion(args[3], "log.Level")
	} else if formatter, err := s.factory.Create(format); err != nil {
		return nil, err
	} else if file, err := newStreamRotatingFileWriter(s.fs, path); err != nil {
		return nil, err
	} else {
		return newStreamFile(file, formatter, channels, level)
	}
}

// CreateFromConfig will instantiate the desired stream instance where
// the initialization data comes from a configuration instance.
func (s streamStrategyRotatingFile) CreateFromConfig(cfg sconfig.IConfig) (IStream, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	if path, err := cfg.String("path"); err != nil {
		return nil, err
	} else if format, err := cfg.String("format"); err != nil {
		return nil, err
	} else if channels, err := s.channels(cfg); err != nil {
		return nil, err
	} else if level, err := s.level(cfg); err != nil {
		return nil, err
	} else {
		return s.Create(path, format, channels, level)
	}
}
