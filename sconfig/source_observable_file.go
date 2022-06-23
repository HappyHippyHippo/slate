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

var _ SourceObservable = &sourceObservableFile{}

// NewSourceObservableFile instantiate a new source that treats a file
// as the origin of the configuration content. This file source will be
// periodically checked for changes and loaded if so.
func NewSourceObservableFile(path, format string, fs afero.Fs, factory *DecoderFactory) (SourceObservable, error) {
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	if factory == nil {
		return nil, errNilPointer("factory")
	}

	s := &sourceObservableFile{
		sourceFile: sourceFile{
			source: source{
				mutex:   &sync.RWMutex{},
				partial: Partial{},
			},
			path:    path,
			format:  format,
			fs:      fs,
			factory: factory,
		},
		timestamp: time.Unix(0, 0),
	}

	if _, err := s.Reload(); err != nil {
		return nil, err
	}
	return s, nil
}

// Reload will check if the source has been updated, and, if so, reload the
// source configuration Partial content.
func (s *sourceObservableFile) Reload() (bool, error) {
	fi, err := s.fs.Stat(s.path)
	if err != nil {
		return false, err
	}

	t := fi.ModTime()
	if s.timestamp.Equal(time.Unix(0, 0)) || s.timestamp.Before(t) {
		if err := s.load(); err != nil {
			return false, err
		}
		s.mutex.Lock()
		s.timestamp = t
		s.mutex.Unlock()
		return true, nil
	}
	return false, nil
}
