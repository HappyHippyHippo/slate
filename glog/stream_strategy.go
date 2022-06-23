package glog

import (
	"github.com/happyhippyhippo/slate/gconfig"
	"strings"
)

// StreamStrategy interface defines the methods of the stream
// factory strategy that can validate creation requests and instantiation
// of particular type of stream.
type StreamStrategy interface {
	Accept(sourceType string) bool
	AcceptFromConfig(cfg gconfig.Config) bool
	Create(args ...interface{}) (Stream, error)
	CreateFromConfig(cfg gconfig.Config) (Stream, error)
}

type streamStrategy struct{}

func (streamStrategy) level(cfg gconfig.Config) (Level, error) {
	level, err := cfg.String("level")
	if err != nil {
		return FATAL, err
	}

	level = strings.ToLower(level)
	if _, ok := LevelMap[level]; !ok {
		return FATAL, errInvalidLevel(level)
	}
	return LevelMap[level], nil
}

func (streamStrategy) channels(cfg gconfig.Config) ([]string, error) {
	entries, err := cfg.List("channels")
	if err != nil {
		return nil, err
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
