package log

import (
	"fmt"
	"io"
	"os"
	"sort"
)

// ConsoleStream defines an instance to a console log output stream.
type ConsoleStream struct {
	Stream
	writer io.Writer
}

var _ IStream = &ConsoleStream{}

// NewConsoleStream generate a new console log stream instance.
func NewConsoleStream(
	formatter IFormatter,
	channels []string,
	level Level,
) (*ConsoleStream, error) {
	// check formatter argument reference
	if formatter == nil {
		return nil, errNilPointer("formatter")
	}
	// instantiate the console stream with the stdout as the
	// default writer target
	s := &ConsoleStream{
		Stream: Stream{
			formatter,
			channels,
			level,
		},
		writer: os.Stdout,
	}
	// sort the assigned channels list
	sort.Strings(s.channels)
	return s, nil
}

// Signal will process the logging signal request and store the logging request
// into the underlying writer if passing the channel and level filtering.
func (s *ConsoleStream) Signal(
	channel string,
	level Level,
	msg string,
	ctx ...Context,
) error {
	// search if the requested channel is in the stream channel list
	i := sort.SearchStrings(s.channels, channel)
	if i == len(s.channels) || s.channels[i] != channel {
		return nil
	}
	// write the message to the stream
	return s.Broadcast(level, msg, ctx...)
}

// Broadcast will process the logging signal request and store the logging
// request into the underlying writer if passing the level filtering.
func (s *ConsoleStream) Broadcast(
	level Level,
	msg string,
	ctx ...Context,
) error {
	// check if the request level is higher than the associated stream level
	if s.level < level {
		return nil
	}
	// write the message after formatting to the default writer (stdout)
	_, e := fmt.Fprintln(s.writer, s.format(level, msg, ctx...))
	return e
}
