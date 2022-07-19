package slog

import (
	"github.com/happyhippyhippo/slate/sconfig"
	"github.com/spf13/afero"
)

type streamStrategyRotatingFile struct {
	streamStrategyFile
}

var _ IStreamStrategy = &streamStrategyRotatingFile{}

func newStreamStrategyRotatingFile(fs afero.Fs, fFactory IFormatterFactory) (IStreamStrategy, error) {
	if fs == nil {
		return nil, errNilPointer("sfs")
	}
	if fFactory == nil {
		return nil, errNilPointer("factory")
	}

	return &streamStrategyRotatingFile{
		streamStrategyFile: streamStrategyFile{
			fs:       fs,
			fFactory: fFactory,
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

	if streamType, e := cfg.String("type"); e == nil {
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
		return nil, errConversion(args[3], "slog.Level")
	} else if formatter, e := s.fFactory.Create(format); e != nil {
		return nil, e
	} else if file, e := newStreamRotatingFileWriter(s.fs, path); e != nil {
		return nil, e
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

	if path, e := cfg.String("path"); e != nil {
		return nil, e
	} else if format, e := cfg.String("format"); e != nil {
		return nil, e
	} else if channels, e := s.channels(cfg); e != nil {
		return nil, e
	} else if level, e := s.level(cfg); e != nil {
		return nil, e
	} else {
		return s.Create(path, format, channels, level)
	}
}
