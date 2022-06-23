package glog

import (
	"fmt"
	"io"
	"sort"
)

// streamFile defines a file output logger stream.
type streamFile struct {
	stream
	writer io.Writer
}

var _ Stream = &streamConsole{}

// NewStreamFile instantiate a new file stream object that will write logging
// content into a file.
func NewStreamFile(writer io.Writer, formatter Formatter, channels []string, level Level) (Stream, error) {
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
	var err error
	if s.writer != nil {
		if w, ok := s.writer.(io.Closer); ok {
			err = w.Close()
		}
		s.writer = nil
	}
	return err
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

	_, err := fmt.Fprintln(s.writer, s.format(level, msg, ctx))
	return err
}
