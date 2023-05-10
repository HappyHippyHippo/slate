package file

import (
	"io"
	"sort"

	"github.com/happyhippyhippo/slate/log"
	"github.com/happyhippyhippo/slate/log/stream"
)

// Stream defines an instance to a file log output stream.
type Stream struct {
	stream.Stream
}

var _ log.IStream = &Stream{}

// NewStream generate a new file log stream instance.
func NewStream(
	writer io.Writer,
	formatter log.IFormatter,
	channels []string,
	level log.Level,
) (log.IStream, error) {
	// check the formatter argument reference
	if formatter == nil {
		return nil, errNilPointer("formatter")
	}
	// check the writer argument reference
	if writer == nil {
		return nil, errNilPointer("writer")
	}
	// instantiate the file stream
	s := &Stream{
		Stream: stream.Stream{
			Formatter: formatter,
			Channels:  channels,
			Level:     level,
			Writer:    writer,
		},
	}
	// sort the assigned channels list
	sort.Strings(s.Channels)
	return s, nil
}

// Close will terminate the stream stored writer instance.
func (s *Stream) Close() error {
	var e error
	// check if the stored writer implements the closer interface
	// and close it if so
	if s.Writer != nil {
		if w, ok := s.Writer.(io.Closer); ok {
			e = w.Close()
		}
		s.Writer = nil
	}
	return e
}
