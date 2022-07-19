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

func newSourceObservableRest(client HTTPClient, uri, format string, dFactory IDecoderFactory, timestampPath, configPath string) (ISourceObservable, error) {
	if client == nil {
		return nil, errNilPointer("client")
	}
	if dFactory == nil {
		return nil, errNilPointer("dFactory")
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
			dFactory:   dFactory,
			configPath: configPath,
		},
		timestampPath: timestampPath,
		timestamp:     time.Unix(0, 0),
	}

	if _, e := s.Reload(); e != nil {
		return nil, e
	}
	return s, nil
}

// Reload will check if the source has been updated, and, if so, reload the
// source configuration Partial content.
func (s *sourceObservableRest) Reload() (bool, error) {
	r, e := s.request()
	if e != nil {
		return false, e
	}

	var t time.Time
	if t, e = s.searchTimestamp(r); e != nil {
		return false, e
	}

	if s.timestamp.Equal(time.Unix(0, 0)) || s.timestamp.Before(t) {
		var p IConfig
		if p, e = s.searchConfig(r); e != nil {
			return false, e
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
	var e error

	var ts interface{}
	if ts, e = body.(*Partial).path(s.timestampPath); e != nil {
		return time.Now(), errConfigRestPathNotFound(s.timestampPath)
	}

	switch ts.(type) {
	case string:
	default:
		return time.Now(), errConversion(ts, "string")
	}

	var t time.Time
	if t, e = time.Parse(time.RFC3339, ts.(string)); e != nil {
		return time.Now(), e
	}

	return t, nil
}
