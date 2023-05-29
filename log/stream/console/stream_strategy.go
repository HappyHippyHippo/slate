package console

import (
	"strings"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/log"
)

const (
	// Type defines the value to be used to declare a
	// console Log stream type.
	Type = "console"
)

type formatterCreator interface {
	Create(format string, args ...interface{}) (log.Formatter, error)
}

// StreamStrategy defines a console log stream generation strategy.
type StreamStrategy struct {
	formatterCreator formatterCreator
}

var _ log.StreamStrategy = &StreamStrategy{}

// NewStreamStrategy generates a new console log stream
// generation strategy instance.
func NewStreamStrategy(
	formatterCreator *log.FormatterFactory,
) (*StreamStrategy, error) {
	// check formatter factory argument reference
	if formatterCreator == nil {
		return nil, errNilPointer("formatterCreator")
	}
	// instantiates the console stream strategy
	return &StreamStrategy{
		formatterCreator: formatterCreator,
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
	sc := struct {
		Format   string
		Channels []interface{}
		Level    string
	}{}
	if _, e := cfg.Populate("", &sc); e != nil {
		return nil, e
	}
	// validate configuration
	level, e := s.level(sc.Level)
	if e != nil {
		return nil, e
	}
	// create the stream formatter to be given to the console stream
	formatter, e := s.formatterCreator.Create(sc.Format)
	if e != nil {
		return nil, e
	}
	// instantiate the console stream
	return NewStream(formatter, s.channels(sc.Channels), level)
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
		if c, ok := channel.(string); ok {
			result = append(result, c)
		}
	}
	return result
}
