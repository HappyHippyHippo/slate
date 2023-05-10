package rest

import (
	"errors"
	"sync"
	"time"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/config/source"
)

// ObsSource defines a config source that read a REST
// service and store a section of the response as the stored config.
// Also, the REST service will be periodically checked for updates.
type ObsSource struct {
	Source
	timestampPath string
	timestamp     time.Time
}

var _ config.IObsSource = &ObsSource{}

// NewObsSource will instantiate a new configuration source
// that will read a REST endpoint for configuration info, opening the
// possibility for on-the-fly update on source content change.
func NewObsSource(
	client httpClient,
	uri,
	format string,
	decoderFactory config.IDecoderFactory,
	timestampPath,
	configPath string,
) (*ObsSource, error) {
	// check client argument reference
	if client == nil {
		return nil, errNilPointer("client")
	}
	// check decoder factory argument reference
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}
	// instantiates the config source
	s := &ObsSource{
		Source: Source{
			BaseSource: source.BaseSource{
				Mutex:  &sync.Mutex{},
				Config: config.Config{},
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
func (s *ObsSource) Reload() (bool, error) {
	// get the REST service information
	rc, e := s.request()
	if e != nil {
		return false, e
	}
	// search for the response timestamp
	var t time.Time
	if t, e = s.searchTimestamp(rc); e != nil {
		return false, e
	}
	// check if the response timestamp is greater than the locally stored
	// config information timestamp
	if s.timestamp.Equal(time.Unix(0, 0)) || s.timestamp.Before(t) {
		// get the response config information
		c, e := rc.Config(s.configPath)
		if e != nil {
			if errors.Is(e, config.ErrPathNotFound) {
				return false, errConfigNotFound(s.configPath)
			}
			return false, e
		}
		// store the loaded config information and response timestamp
		s.Mutex.Lock()
		s.Config = *c.(*config.Config)
		s.timestamp = t
		s.Mutex.Unlock()
		return true, nil
	}
	return false, nil
}

func (s *ObsSource) searchTimestamp(
	rc config.IConfig,
) (time.Time, error) {
	// retrieve the timestamp information from the parsed response data
	t, e := rc.String(s.timestampPath)
	if e != nil {
		if errors.Is(e, config.ErrPathNotFound) {
			return time.Now(), errTimestampNotFound(s.timestampPath)
		}
		return time.Now(), e
	}
	// parse the timestamp string
	var tt time.Time
	if tt, e = time.Parse(time.RFC3339, t); e != nil {
		return time.Now(), e
	}
	return tt, nil
}
