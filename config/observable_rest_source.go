package config

import (
	"sync"
	"time"
)

// ObservableRestSource defines a config source that read a REST
// service and store a section of the response as the stored config.
// Also, the REST service will be periodically checked for updates.
type ObservableRestSource struct {
	RestSource
	timestampPath string
	timestamp     time.Time
}

var _ IObservableSource = &ObservableRestSource{}

// NewObservableRestSource will instantiate a new configuration source
// that will read a REST endpoint for configuration info, opening the
// possibility for on-the-fly update on source content change.
func NewObservableRestSource(
	client httpClient,
	uri,
	format string,
	decoderFactory IDecoderFactory,
	timestampPath,
	configPath string,
) (*ObservableRestSource, error) {
	// check client argument reference
	if client == nil {
		return nil, errNilPointer("client")
	}
	// check decoder factory argument reference
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}
	// instantiates the config source
	s := &ObservableRestSource{
		RestSource: RestSource{
			Source: Source{
				mutex:   &sync.Mutex{},
				partial: Config{},
			},
			client:         client,
			uri:            uri,
			format:         format,
			decoderFactory: decoderFactory,
			configPath:     configPath,
		},
		timestampPath: timestampPath,
		timestamp:     time.Unix(0, 0),
	}
	// load the config information from the REST service
	if _, e := s.Reload(); e != nil {
		return nil, e
	}
	return s, nil
}

// Reload will check if the source has been updated, and, if so, reload the
// source configuration content.
func (s *ObservableRestSource) Reload() (bool, error) {
	// get the REST service information
	r, e := s.request()
	if e != nil {
		return false, e
	}
	// search for the response timestamp
	var t time.Time
	if t, e = s.searchTimestamp(r); e != nil {
		return false, e
	}
	// check if the response timestamp is greater than the locally stored
	// config information timestamp
	if s.timestamp.Equal(time.Unix(0, 0)) || s.timestamp.Before(t) {
		// get the response config information
		p, e := s.searchConfig(r)
		if e != nil {
			return false, e
		}
		// store the loaded config information and response timestamp
		s.mutex.Lock()
		s.partial = *p
		s.timestamp = t
		s.mutex.Unlock()
		return true, nil
	}
	return false, nil
}

func (s *ObservableRestSource) searchTimestamp(
	body IConfig,
) (time.Time, error) {
	var e error
	// retrieve the timestamp information from the parsed response data
	var ts interface{}
	if ts, e = body.(*Config).path(s.timestampPath); e != nil {
		return time.Now(), errRestPathNotFound(s.timestampPath)
	}
	// check if the timestamp information is a valid string to be parsed
	switch ts.(type) {
	case string:
	default:
		return time.Now(), errConversion(ts, "string")
	}
	// parse the timestamp string
	var t time.Time
	if t, e = time.Parse(time.RFC3339, ts.(string)); e != nil {
		return time.Now(), e
	}
	return t, nil
}
