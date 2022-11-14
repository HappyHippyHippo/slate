package log

import (
	"os"

	sconfig "github.com/happyhippyhippo/slate/config"
	"github.com/spf13/afero"
)

type streamStrategyFile struct {
	streamStrategy
	fs               afero.Fs
	formatterFactory IFormatterFactory
}

var _ IStreamStrategy = &streamStrategyFile{}

func newStreamStrategyFile(fs afero.Fs, formatterFactory IFormatterFactory) (IStreamStrategy, error) {
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	if formatterFactory == nil {
		return nil, errNilPointer("formatterFactory")
	}

	return &streamStrategyFile{
		fs:               fs,
		formatterFactory: formatterFactory,
	}, nil
}

// Accept will check if the file stream factory strategy can instantiate a
// stream of the requested type and with the calling parameters.
func (streamStrategyFile) Accept(streamType string) bool {
	return streamType == StreamFile
}

// AcceptFromConfig will check if the stream factory strategy can instantiate
// a stream where the data to check comes from a configuration partial
// instance.
func (s streamStrategyFile) AcceptFromConfig(cfg sconfig.IConfig) bool {
	if cfg == nil {
		return false
	}

	if streamType, e := cfg.String("type"); e == nil {
		return s.Accept(streamType)
	}

	return false
}

// Create will instantiate the desired stream instance.
func (s streamStrategyFile) Create(args ...interface{}) (IStream, error) {
	if len(args) < 4 {
		return nil, errNilPointer("args")
	}

	if path, ok := args[0].(string); !ok {
		return nil, errConversion(args[1], "string")
	} else if format, ok := args[1].(string); !ok {
		return nil, errConversion(args[1], "string")
	} else if channels, ok := args[2].([]string); !ok {
		return nil, errConversion(args[2], "[]string")
	} else if level, ok := args[3].(Level); !ok {
		return nil, errConversion(args[3], "log.Level")
	} else if formatter, e := s.formatterFactory.Create(format); e != nil {
		return nil, e
	} else if file, e := s.fs.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644); e != nil {
		return nil, e
	} else {
		return newStreamFile(file, formatter, channels, level)
	}
}

// CreateFromConfig will instantiate the desired stream instance where
// the initialization data comes from a configuration instance.
func (s streamStrategyFile) CreateFromConfig(cfg sconfig.IConfig) (IStream, error) {
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
