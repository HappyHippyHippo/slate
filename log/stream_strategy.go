package log

import (
	"strings"

	"github.com/happyhippyhippo/slate/config"
)

const (
	// UnknownStream defines the value to be used to declare an
	// unknown Log stream type.
	UnknownStream = "unknown"
)

// IStreamStrategy interface defines the methods of the stream
// factory strategy that can validate creation requests and instantiation
// of particular type of stream.
type IStreamStrategy interface {
	Accept(cfg config.IConfig) bool
	Create(cfg config.IConfig) (IStream, error)
}

// StreamStrategy defines a new log stream base instance structure.
type StreamStrategy struct{}

func (StreamStrategy) level(
	level string,
) (Level, error) {
	// check if the retrieved level string can be mapped to a
	// level typed value
	level = strings.ToLower(level)
	if _, ok := LevelMap[level]; !ok {
		return FATAL, errInvalidLevel(level)
	}
	// return the level typed value of the retrieved level string
	return LevelMap[level], nil
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
