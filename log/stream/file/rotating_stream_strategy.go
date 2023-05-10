package file

import (
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/log"
	"github.com/spf13/afero"
)

const (
	// RotatingType defines the value to be used to declare a
	// file Log stream type that rotates regarding the current date.
	RotatingType = "rotating-file"
)

type rotatingStreamConfig struct {
	Path     string
	Format   string
	Channels []interface{}
	Level    string
}

// RotatingStreamStrategy define a new rotating file log
// stream generation strategy.
type RotatingStreamStrategy struct {
	StreamStrategy
}

var _ log.IStreamStrategy = &RotatingStreamStrategy{}

// NewRotatingStreamStrategy generate a new rotating file log stream
// generation strategy.
func NewRotatingStreamStrategy(
	fs afero.Fs,
	formatterFactory log.IFormatterFactory,
) (*RotatingStreamStrategy, error) {
	// check file system argument reference
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	// check formatter factory argument reference
	if formatterFactory == nil {
		return nil, errNilPointer("formatterFactory")
	}
	// instantiate the rotating file stream strategy
	return &RotatingStreamStrategy{
		StreamStrategy: StreamStrategy{
			fs:               fs,
			formatterFactory: formatterFactory,
		},
	}, nil
}

// Accept will check if the stream factory strategy can instantiate
// a stream where the data to check comes from a configuration partial
// instance.
func (s RotatingStreamStrategy) Accept(
	cfg config.IConfig,
) bool {
	// check config argument reference
	if cfg == nil {
		return false
	}
	// retrieve the data from the configuration
	sc := struct{ Type string }{}
	if _, e := cfg.Populate("", &sc); e != nil {
		return false
	}
	// return acceptance for the read config type
	return sc.Type == RotatingType
}

// Create will instantiate the desired stream instance where
// the initialization data comes from a configuration instance.
func (s RotatingStreamStrategy) Create(
	cfg config.IConfig,
) (log.IStream, error) {
	// check config argument reference
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	// retrieve the data from the configuration
	sc := rotatingStreamConfig{}
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
	file, e := NewRotatingWriter(s.fs, sc.Path)
	if e != nil {
		return nil, e
	}
	// instantiate the console stream
	return NewStream(file, formatter, s.channels(sc.Channels), level)
}
