package sconfig

import (
	"sync"
	"time"
)

type sourceObservableRest struct {
	sourceRest
	timestampPath string
	timestamp     time.Time
}

var _ ISourceObservable = &sourceObservableRest{}

func newSourceObservableRest(client HTTPClient, uri, format string, factory IDecoderFactory, timestampPath, configPath string) (ISourceObservable, error) {
	if client == nil {
		return nil, errNilPointer("client")
	}
	if factory == nil {
		return nil, errNilPointer("factory")
	}

	s := &sourceObservableRest{
		sourceRest: sourceRest{
			source: source{
				mutex:   &sync.Mutex{},
				partial: Partial{},
			},
			client:     client,
			uri:        uri,
			format:     format,
			factory:    factory,
			configPath: configPath,
		},
		timestampPath: timestampPath,
		timestamp:     time.Unix(0, 0),
	}

	if _, err := s.Reload(); err != nil {
		return nil, err
	}
	return s, nil
}

// Reload will check if the source has been updated, and, if so, reload the
// source configuration Partial content.
func (s *sourceObservableRest) Reload() (bool, error) {
	r, err := s.request()
	if err != nil {
		return false, err
	}

	var t time.Time
	if t, err = s.searchTimestamp(r); err != nil {
		return false, err
	}

	if s.timestamp.Equal(time.Unix(0, 0)) || s.timestamp.Before(t) {
		var p IConfig
		if p, err = s.searchConfig(r); err != nil {
			return false, err
		}

		s.mutex.Lock()
		s.partial = *p.(*Partial)
		s.timestamp = t
		s.mutex.Unlock()

		return true, nil
	}

	return false, nil
}

func (s *sourceObservableRest) searchTimestamp(body IConfig) (time.Time, error) {
	var err error

	var ts interface{}
	if ts, err = body.(*Partial).path(s.timestampPath); err != nil {
		return time.Now(), errConfigRestPathNotFound(s.timestampPath)
	}

	switch ts.(type) {
	case string:
	default:
		return time.Now(), errConversion(ts, "string")
	}

	var t time.Time
	if t, err = time.Parse(time.RFC3339, ts.(string)); err != nil {
		return time.Now(), err
	}

	return t, nil
}
