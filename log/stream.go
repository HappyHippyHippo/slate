package log

import (
	"sort"
)

// IStream interface defines the interaction methods with a logging stream.
type IStream interface {
	Level() Level

	Signal(channel string, level Level, message string, ctx ...Context) error
	Broadcast(level Level, message string, ctx ...Context) error

	HasChannel(channel string) bool
	ListChannels() []string
	AddChannel(channel string)
	RemoveChannel(channel string)
}

// Stream defines the base interaction with a Log stream instance.
type Stream struct {
	formatter IFormatter
	channels  []string
	level     Level
}

// Level retrieves the logging level filter value of the stream.
func (s *Stream) Level() Level {
	return s.level
}

// HasChannel will validate if the stream is listening to a specific
// logging channel.
func (s *Stream) HasChannel(
	channel string,
) bool {
	// search the requested string in the already ordered
	// stream channel pool list
	i := sort.SearchStrings(s.channels, channel)
	return i < len(s.channels) && s.channels[i] == channel
}

// ListChannels retrieves the list of channels that the stream is listening.
func (s Stream) ListChannels() []string {
	return s.channels
}

// AddChannel register a channel to the list of channels that the
// stream is listening.
func (s *Stream) AddChannel(
	channel string,
) {
	// check if the adding channel is not already in the stream
	// channel pool list
	if !s.HasChannel(channel) {
		// add the requested channel and sort the channel pool list
		s.channels = append(s.channels, channel)
		sort.Strings(s.channels)
	}
}

// RemoveChannel removes a channel from the list of channels that the
// stream is listening.
func (s *Stream) RemoveChannel(
	channel string,
) {
	// search for the channel pool position of the channel to be removed
	i := sort.SearchStrings(s.channels, channel)
	// check if the channel was not found
	if i == len(s.channels) || s.channels[i] != channel {
		return
	}
	// remove the channel from the channel pool list
	s.channels = append(s.channels[:i], s.channels[i+1:]...)
}

func (s *Stream) format(
	level Level,
	message string,
	ctx ...Context,
) string {
	// check if a valid formatter reference is present, if so, return
	// the formatter response of the message content
	if s.formatter != nil {
		message = s.formatter.Format(level, message, ctx...)
	}
	// return just the message if no formatter is present
	return message
}
