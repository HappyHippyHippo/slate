package stream

import (
	"fmt"
	"io"
	"sort"

	"github.com/happyhippyhippo/slate/log"
)

// Stream defines the base interaction with a Log stream instance.
type Stream struct {
	Formatter log.IFormatter
	Channels  []string
	Level     log.Level
	Writer    io.Writer
}

// HasChannel will validate if the stream is listening to a specific
// logging channel.
func (s *Stream) HasChannel(
	channel string,
) bool {
	// search the requested string in the already ordered
	// stream channel pool list
	i := sort.SearchStrings(s.Channels, channel)
	return i < len(s.Channels) && s.Channels[i] == channel
}

// ListChannels retrieves the list of Channels that the stream is listening.
func (s *Stream) ListChannels() []string {
	return s.Channels
}

// AddChannel register a channel to the list of Channels that the
// stream is listening.
func (s *Stream) AddChannel(
	channel string,
) {
	// check if the adding channel is not already in the stream
	// channel pool list
	if !s.HasChannel(channel) {
		// add the requested channel and sort the channel pool list
		s.Channels = append(s.Channels, channel)
		sort.Strings(s.Channels)
	}
}

// RemoveChannel removes a channel from the list of Channels that the
// stream is listening.
func (s *Stream) RemoveChannel(
	channel string,
) {
	// search for the channel pool position of the channel to be removed
	i := sort.SearchStrings(s.Channels, channel)
	// check if the channel was not found
	if i == len(s.Channels) || s.Channels[i] != channel {
		return
	}
	// remove the channel from the channel pool list
	s.Channels = append(s.Channels[:i], s.Channels[i+1:]...)
}

// Format will try to format a logging message.
func (s *Stream) Format(
	level log.Level,
	message string,
	ctx ...log.Context,
) string {
	// check if a valid formatter reference is present, if so, return
	// the formatter response of the message content
	if s.Formatter != nil {
		message = s.Formatter.Format(level, message, ctx...)
	}
	// return just the message if no formatter is present
	return message
}

// Signal will process the logging signal request and store the logging request
// into the underlying writer if passing the channel and level filtering.
func (s *Stream) Signal(
	channel string,
	level log.Level,
	msg string,
	ctx ...log.Context,
) error {
	// search if the requested channel is in the stream channel list
	i := sort.SearchStrings(s.Channels, channel)
	if i == len(s.Channels) || s.Channels[i] != channel {
		return nil
	}
	// write the message to the stream
	return s.Broadcast(level, msg, ctx...)
}

// Broadcast will process the logging signal request and store the logging
// request into the underlying writer if passing the level filtering.
func (s *Stream) Broadcast(
	level log.Level,
	msg string,
	ctx ...log.Context,
) error {
	// check if the request level is higher than the associated stream level
	if s.Level < level {
		return nil
	}
	// write the message after formatting to the def writer (stdout)
	_, e := fmt.Fprintln(s.Writer, s.Format(level, msg, ctx...))
	return e
}
