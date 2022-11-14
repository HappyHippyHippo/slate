package log

import (
	"strings"

	sconfig "github.com/happyhippyhippo/slate/config"
)

// IStreamStrategy interface defines the methods of the stream
// factory strategy that can validate creation requests and instantiation
// of particular type of stream.
type IStreamStrategy interface {
	Accept(sourceType string) bool
	AcceptFromConfig(cfg sconfig.IConfig) bool
	Create(args ...interface{}) (IStream, error)
	CreateFromConfig(cfg sconfig.IConfig) (IStream, error)
}

type streamStrategy struct{}

func (streamStrategy) level(cfg sconfig.IConfig) (Level, error) {
	level, e := cfg.String("level")
	if e != nil {
		return FATAL, e
	}

	level = strings.ToLower(level)
	if _, ok := LevelMap[level]; !ok {
		return FATAL, errInvalidLevel(level)
	}
	return LevelMap[level], nil
}

func (streamStrategy) channels(cfg sconfig.IConfig) ([]string, error) {
	entries, e := cfg.List("channels")
	if e != nil {
		return nil, e
	}

	var channels []string
	for _, entry := range entries {
		channel, ok := entry.(string)
		if !ok {
			return nil, errConversion(entry, "string")
		}
		channels = append(channels, channel)
	}
	return channels, nil
}
