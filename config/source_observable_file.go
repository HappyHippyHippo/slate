package config

import (
	"sync"
	"time"

	"github.com/spf13/afero"
)

type sourceObservableFile struct {
	sourceFile
	timestamp time.Time
}

var _ ISourceObservable = &sourceObservableFile{}

// NewSourceObservableFile will instantiate a new configuration source
// that will read a file for configuration info, opening the
// possibility for on-the-fly update on source content change.
func NewSourceObservableFile(path, format string, fs afero.Fs, decoderFactory IDecoderFactory) (ISourceObservable, error) {
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}

	s := &sourceObservableFile{
		sourceFile: sourceFile{
			source: source{
				mutex:   &sync.RWMutex{},
				partial: Partial{},
			},
			path:           path,
			format:         format,
			fs:             fs,
			decoderFactory: decoderFactory,
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
