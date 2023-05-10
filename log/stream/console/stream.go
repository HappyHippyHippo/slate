package console

import (
	"os"
	"sort"

	"github.com/happyhippyhippo/slate/log"
	"github.com/happyhippyhippo/slate/log/stream"
)

// Stream defines an instance to a console log output stream.
type Stream struct {
	stream.Stream
}

var _ log.IStream = &Stream{}

// NewStream generate a new console log stream instance.
func NewStream(
	formatter log.IFormatter,
	channels []string,
	level log.Level,
) (*Stream, error) {
	// check formatter argument reference
	if formatter == nil {
		return nil, errNilPointer("formatter")
	}
	// instantiate the console stream with the stdout as the
	// def writer target
	s := &Stream{
		Stream: stream.Stream{
			Formatter: formatter,
			Channels:  channels,
			Level:     level,
			Writer:    os.Stdout,
		},
	}
	// sort the assigned channels list
	sort.Strings(s.Channels)
	return s, nil
}
