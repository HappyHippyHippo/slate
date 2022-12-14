package log

import (
	"fmt"
	"io"
	"sort"
)

// FileStream defines an instance to a file log output stream.
type FileStream struct {
	Stream
	writer io.Writer
}

var _ IStream = &FileStream{}

// NewFileStream generate a new file log stream instance.
func NewFileStream(
	writer io.Writer,
	formatter IFormatter,
	channels []string,
	level Level,
) (IStream, error) {
	// check the formatter argument reference
	if formatter == nil {
		return nil, errNilPointer("formatter")
	}
	// check the writer argument reference
	if writer == nil {
		return nil, errNilPointer("writer")
	}
	// instantiate the file stream
	s := &FileStream{
		Stream: Stream{
			formatter,
			channels,
			level},
		writer: writer}
	// sort the assigned channels list
	sort.Strings(s.channels)
	return s, nil
}

// Close will terminate the stream stored writer instance.
func (s *FileStream) Close() error {
	var e error
	// check if the stored writer implements the closer interface
	// and close it if so
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
func (s FileStream) Signal(
	channel string,
	level Level,
	msg string,
	ctx map[string]interface{},
) error {
	// search if the requested channel is in the stream channel list
	i := sort.SearchStrings(s.channels, channel)
	if i == len(s.channels) || s.channels[i] != channel {
		return nil
	}
	// write the message to the stream
	return s.Broadcast(level, msg, ctx)
}

// Broadcast will process the logging signal request and store the logging
// request into the underlying file if passing the level filtering.
func (s FileStream) Broadcast(
	level Level,
	msg string,
	ctx map[string]interface{},
) error {
	// check if the request level is higher than the associated stream level
	if s.level < level {
		return nil
	}
	// write the message after formatting to the defined stream writer
	_, e := fmt.Fprintln(s.writer, s.format(level, msg, ctx))
	return e
}
