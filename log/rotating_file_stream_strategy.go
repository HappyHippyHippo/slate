package log

import (
	"github.com/happyhippyhippo/slate/config"
	"github.com/spf13/afero"
)

const (
	// RotatingFileStreamType defines the value to be used to declare a
	// file Log stream type that rotates regarding the current date.
	RotatingFileStreamType = "rotating-file"
)

type rotatingFileStreamConfig struct {
	Path     string
	Format   string
	Channels []interface{}
	Level    string
}

// RotatingFileStreamStrategy define a new rotating file log
// stream generation strategy.
type RotatingFileStreamStrategy struct {
	FileStreamStrategy
}

var _ IStreamStrategy = &RotatingFileStreamStrategy{}

// NewRotatingFileStreamStrategy generate a new rotating file log stream
// generation strategy.
func NewRotatingFileStreamStrategy(
	fs afero.Fs,
	formatterFactory IFormatterFactory,
) (*RotatingFileStreamStrategy, error) {
	// check file system argument reference
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	// check formatter factory argument reference
	if formatterFactory == nil {
		return nil, errNilPointer("formatterFactory")
	}
	// instantiate the rotating file stream strategy
	return &RotatingFileStreamStrategy{
		FileStreamStrategy: FileStreamStrategy{
			fs:               fs,
			formatterFactory: formatterFactory,
		},
	}, nil
}

// Accept will check if the stream factory strategy can instantiate
// a stream where the data to check comes from a configuration partial
// instance.
func (s RotatingFileStreamStrategy) Accept(
	cfg config.IConfig,
) bool {
	// check config argument reference
	if cfg == nil {
		return false
	}
	// retrieve the data from the configuration
	sc := struct{ Type string }{}
	_, e := cfg.Populate("", &sc)
	if e == nil {
		// return acceptance for the read config type
		return sc.Type == RotatingFileStreamType
	}
	return false
}

// Create will instantiate the desired stream instance where
// the initialization data comes from a configuration instance.
func (s RotatingFileStreamStrategy) Create(
	cfg config.IConfig,
) (IStream, error) {
	// check config argument reference
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	// retrieve the data from the configuration
	sc := rotatingFileStreamConfig{}
	_, e := cfg.Populate("", &sc)
	if e != nil {
		return nil, e
	}
	// validate configuration
	level, e := s.level(sc.Level)
	if e != nil {
		return nil, e
	}
	// create the stream formatter to be given to the console stream
	formatter, e := s.formatterFactory.Create(sc.Format)
	if e != nil {
		return nil, e
	}
	// create the stream writer
	file, e := NewRotatingFileWriter(s.fs, sc.Path)
	if e != nil {
		return nil, e
	}
	// instantiate the console stream
	return NewFileStream(file, formatter, s.channels(sc.Channels), level)
}
