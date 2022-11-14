package log

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type streamConsole struct {
	stream
	writer io.Writer
}

var _ IStream = &streamConsole{}

func newStreamConsole(formatter IFormatter, channels []string, level Level) (IStream, error) {
	if formatter == nil {
		return nil, errNilPointer("formatter")
	}

	s := &streamConsole{
		stream: stream{
			formatter,
			channels,
			level,
		},
		writer: os.Stdout,
	}

	sort.Strings(s.channels)

	return s, nil
}

// Signal will process the logging signal request and store the logging request
// into the underlying writer if passing the channel and level filtering.
func (s streamConsole) Signal(channel string, level Level, msg string, ctx map[string]interface{}) error {
	i := sort.SearchStrings(s.channels, channel)
	if i == len(s.channels) || s.channels[i] != channel {
		return nil
	}

	return s.Broadcast(level, msg, ctx)
}

// Broadcast will process the logging signal request and store the logging
// request into the underlying writer if passing the level filtering.
func (s streamConsole) Broadcast(level Level, msg string, ctx map[string]interface{}) error {
	if s.level < level {
		return nil
	}

	_, e := fmt.Fprintln(s.writer, s.format(level, msg, ctx))
	return e
}
