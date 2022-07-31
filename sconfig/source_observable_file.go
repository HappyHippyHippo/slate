package sconfig

import (
	"github.com/spf13/afero"
	"sync"
	"time"
)

type sourceObservableFile struct {
	sourceFile
	timestamp time.Time
}

var _ ISourceObservable = &sourceObservableFile{}

// NewSourceObservableFile will instantiate a new configuration source
// that will read a file for configuration info, opening the
// possibility for on-the-fly update on source content change.
func NewSourceObservableFile(path, format string, fs afero.Fs, dFactory IDecoderFactory) (ISourceObservable, error) {
	if fs == nil {
		return nil, errNilPointer("sfs")
	}
	if dFactory == nil {
		return nil, errNilPointer("dFactory")
	}

	s := &sourceObservableFile{
		sourceFile: sourceFile{
			source: source{
				mutex:   &sync.RWMutex{},
				partial: Partial{},
			},
			path:     path,
			format:   format,
			fs:       fs,
			dFactory: dFactory,
		},
		timestamp: time.Unix(0, 0),
	}

	if _, e := s.Reload(); e != nil {
		return nil, e
	}
	return s, nil
}

// Reload will check if the source has been updated, and, if so, reload the
// source configuration Partial content.
func (s *sourceObservableFile) Reload() (bool, error) {
	fi, e := s.fs.Stat(s.path)
	if e != nil {
		return false, e
	}

	t := fi.ModTime()
	if s.timestamp.Equal(time.Unix(0, 0)) || s.timestamp.Before(t) {
		if e := s.load(); e != nil {
			return false, e
		}
		s.mutex.Lock()
		s.timestamp = t
		s.mutex.Unlock()
		return true, nil
	}
	return false, nil
}
