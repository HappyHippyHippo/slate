package file

import (
	"sync"
	"time"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/config/source"
	"github.com/spf13/afero"
)

// ObsSource defines a config source that read a file content
// and stores its config contents to be used as a config.
// The source will also be checked for changes recurrently, so it can
// update the stored configuration information.
type ObsSource struct {
	Source
	timestamp time.Time
}

var _ config.Source = &ObsSource{}
var _ config.ObsSource = &ObsSource{}

// NewObsSource will instantiate a new configuration source
// that will read a file for configuration info, opening the
// possibility for on-the-fly update on source content change.
func NewObsSource(
	path,
	format string,
	fileSystem afero.Fs,
	decoderFactory *config.DecoderFactory,
) (*ObsSource, error) {
	// check file system argument reference
	if fileSystem == nil {
		return nil, errNilPointer("fileSystem")
	}
	// check decoder factory argument reference
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}
	// return the requested observable config source instance
	s := &ObsSource{
		Source: Source{
			Source: source.Source{
				Mutex:   &sync.Mutex{},
				Partial: config.Partial{},
			},
			path:           path,
			format:         format,
			fileSystem:     fileSystem,
			decoderFactory: decoderFactory,
		},
		timestamp: time.Unix(0, 0),
	}
	// Load the file config content
	if _, e := s.Reload(); e != nil {
		return nil, e
	}
	return s, nil
}

// Reload will check if the source has been updated, and, if so, reload the
// source configuration config content.
func (s *ObsSource) Reload() (bool, error) {
	// get the file stats, so we can store the modification time
	fi, e := s.fileSystem.Stat(s.path)
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
		s.timestamp = t
		return true, nil
	}
	return false, nil
}
