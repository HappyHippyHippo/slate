package file

import (
	"os"
	"strings"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/log"
	"github.com/spf13/afero"
)

const (
	// Type defines the value to be used to declare a
	// file Log stream type.
	Type = "file"
)

type streamConfig struct {
	Path     string
	Format   string
	Channels []interface{}
	Level    string
}

// StreamStrategy defines a file log stream generation strategy.
type StreamStrategy struct {
	fs               afero.Fs
	formatterFactory *log.FormatterFactory
}

var _ log.StreamStrategy = &StreamStrategy{}

// NewStreamStrategy generates a new file log stream
// generation strategy instance.
func NewStreamStrategy(
	fs afero.Fs,
	formatterFactory *log.FormatterFactory,
) (*StreamStrategy, error) {
	// check file system argument reference
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	// check formatter factory argument reference
	if formatterFactory == nil {
		return nil, errNilPointer("formatterFactory")
	}
	// instantiate the file stream strategy
	return &StreamStrategy{
		fs:               fs,
		formatterFactory: formatterFactory,
	}, nil
}

// Accept will check if the stream factory strategy can instantiate
// a stream where the data to check comes from a configuration partial
// instance.
func (s StreamStrategy) Accept(
	cfg *config.Partial,
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
	return sc.Type == Type
}

// Create will instantiate the desired stream instance where
// the initialization data comes from a configuration instance.
func (s StreamStrategy) Create(
	cfg *config.Partial,
) (log.Stream, error) {
	// check config argument reference
	if cfg == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sc := streamConfig{}
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
	return NewStream(file, formatter, s.channels(sc.Channels), level)
}

func (StreamStrategy) level(
	level string,
) (log.Level, error) {
	// check if the retrieved Level string can be mapped to a
	// Level typed value
	level = strings.ToLower(level)
	if _, ok := log.LevelMap[level]; !ok {
		return log.FATAL, errInvalidLevel(level)
	}
	// return the Level typed value of the retrieved Level string
	return log.LevelMap[level], nil
}

func (StreamStrategy) channels(
	list []interface{},
) []string {
	var result []string
	for _, channel := range list {
		c, ok := channel.(string)
		if ok {
			result = append(result, c)
		}
	}
	return result
}
