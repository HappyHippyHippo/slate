package config

import (
	"sync"
	"time"

	"github.com/spf13/afero"
)

// ObservableFileSource defines a config source that read a file content
// and stores its config contents to be used as a config.
// The source will also be checked for changes recurrently, so it can
// update the stored configuration information.
type ObservableFileSource struct {
	FileSource
	timestamp time.Time
}

var _ IObservableSource = &ObservableFileSource{}

// NewObservableFileSource will instantiate a new configuration source
// that will read a file for configuration info, opening the
// possibility for on-the-fly update on source content change.
func NewObservableFileSource(
	path,
	format string,
	fs afero.Fs,
	decoderFactory IDecoderFactory,
) (*ObservableFileSource, error) {
	// check file system argument reference
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	// check decoder factory argument reference
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}
	// instantiates the config source
	s := &ObservableFileSource{
		FileSource: FileSource{
			Source: Source{
				mutex:   &sync.RWMutex{},
				partial: Config{},
			},
			path:           path,
			format:         format,
			fs:             fs,
			decoderFactory: decoderFactory,
		},
		timestamp: time.Unix(0, 0),
	}
	// load the file config content
	if _, e := s.Reload(); e != nil {
		return nil, e
	}
	return s, nil
}

// Reload will check if the source has been updated, and, if so, reload the
// source configuration config content.
func (s *ObservableFileSource) Reload() (bool, error) {
	// get the file stats, so we can store the modification time
	fi, e := s.fs.Stat(s.path)
	if e != nil {
		return false, e
	}
	// check if the file modification time is greater than the stored one
	t := fi.ModTime()
	if s.timestamp.Equal(time.Unix(0, 0)) || s.timestamp.Before(t) {
		// load the file content
		if e := s.load(); e != nil {
			return false, e
		}
		// update the stored config content modification time
		s.mutex.Lock()
		s.timestamp = t
		s.mutex.Unlock()
		return true, nil
	}
	return false, nil
}
