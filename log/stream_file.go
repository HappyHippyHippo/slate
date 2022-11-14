package log

import (
	"fmt"
	"io"
	"sort"
)

type streamFile struct {
	stream
	writer io.Writer
}

var _ IStream = &streamConsole{}

func newStreamFile(writer io.Writer, formatter IFormatter, channels []string, level Level) (IStream, error) {
	if formatter == nil {
		return nil, errNilPointer("formatter")
	}
	if writer == nil {
		return nil, errNilPointer("writer")
	}

	s := &streamFile{
		stream: stream{
			formatter,
			channels,
			level},
		writer: writer}

	sort.Strings(s.channels)

	return s, nil
}

// Close will terminate the stream stored writer instance.
func (s *streamFile) Close() error {
	var e error
	if s.writer != nil {
		if w, ok := s.writer.(io.Closer); ok {
			e = w.Close()
		}
		s.writer = nil
	}
	return e
}

// Signal will process the logging signal request and store the logging request
// into the underlying file if passing the channel and level filtering.
func (s streamFile) Signal(channel string, level Level, msg string, ctx map[string]interface{}) error {
	i := sort.SearchStrings(s.channels, channel)
	if i == len(s.channels) || s.channels[i] != channel {
		return nil
	}

	return s.Broadcast(level, msg, ctx)
}

// Broadcast will process the logging signal request and store the logging
// request into the underlying file if passing the level filtering.
func (s streamFile) Broadcast(level Level, msg string, ctx map[string]interface{}) error {
	if s.level < level {
		return nil
	}

	_, e := fmt.Fprintln(s.writer, s.format(level, msg, ctx))
	return e
}
