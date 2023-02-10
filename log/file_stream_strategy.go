package log

import (
	"os"

	"github.com/happyhippyhippo/slate/config"
	"github.com/spf13/afero"
)

const (
	// FileStreamType defines the value to be used to declare a
	// file Log stream type.
	FileStreamType = "file"
)

type fileStreamConfig struct {
	Path     string
	Format   string
	Channels []interface{}
	Level    string
}

// FileStreamStrategy defines a file log stream generation strategy.
type FileStreamStrategy struct {
	StreamStrategy
	fs               afero.Fs
	formatterFactory IFormatterFactory
}

var _ IStreamStrategy = &FileStreamStrategy{}

// NewFileStreamStrategy generates a new file log stream
// generation strategy instance.
func NewFileStreamStrategy(
	fs afero.Fs,
	formatterFactory IFormatterFactory,
) (*FileStreamStrategy, error) {
	// check file system argument reference
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	// check formatter factory argument reference
	if formatterFactory == nil {
		return nil, errNilPointer("formatterFactory")
	}
	// instantiate the file stream strategy
	return &FileStreamStrategy{
		fs:               fs,
		formatterFactory: formatterFactory,
	}, nil
}

// Accept will check if the stream factory strategy can instantiate
// a stream where the data to check comes from a configuration partial
// instance.
func (s FileStreamStrategy) Accept(
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
		return sc.Type == FileStreamType
	}
	return false
}

// Create will instantiate the desired stream instance where
// the initialization data comes from a configuration instance.
func (s FileStreamStrategy) Create(
	cfg config.IConfig,
) (IStream, error) {
	// check config argument reference
	if cfg == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sc := fileStreamConfig{}
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
	file, e := s.fs.OpenFile(sc.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if e != nil {
		return nil, e
	}
	// instantiate the console stream
	return NewFileStream(file, formatter, s.channels(sc.Channels), level)
}
