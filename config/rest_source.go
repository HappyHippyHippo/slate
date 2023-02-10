package config

import (
	"bytes"
	"io"
	"net/http"
	"sync"
)

// httpClient defines the interface of an instance capable to perform the
// rest config obtain action
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// RestSource defines a config source that read a REST service and
// store a section of the response as the stored config.
type RestSource struct {
	Source
	client         httpClient
	uri            string
	format         string
	decoderFactory IDecoderFactory
	configPath     string
}

var _ ISource = &RestSource{}

// NewRestSource will instantiate a new configuration source
// that will read a REST endpoint for configuration info.
func NewRestSource(
	client httpClient,
	uri,
	format string,
	decoderFactory IDecoderFactory,
	configPath string,
) (*RestSource, error) {
	// check client argument reference
	if client == nil {
		return nil, errNilPointer("client")
	}
	// check decoder factory argument reference
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}
	// instantiates the config source
	s := &RestSource{
		Source: Source{
			mutex:  &sync.Mutex{},
			config: Config{},
		},
		client:         client,
		uri:            uri,
		format:         format,
		decoderFactory: decoderFactory,
		configPath:     configPath,
	}
	// load the config information from the REST service
	if e := s.load(); e != nil {
		return nil, e
	}
	return s, nil
}

func (s *RestSource) load() error {
	// get the REST service information
	rc, e := s.request()
	if e != nil {
		return e
	}
	// retrieve the config information from the service response data
	c, e := s.searchConfig(rc)
	if e != nil {
		return e
	}
	// store the retrieved config
	s.mutex.Lock()
	s.config = *c
	s.mutex.Unlock()
	return nil
}

func (s *RestSource) request() (IConfig, error) {
	var e error
	// create the REST service config request
	var req *http.Request
	if req, e = http.NewRequest(http.MethodGet, s.uri, http.NoBody); e != nil {
		return nil, e
	}
	// call the REST service for the configuration information
	var res *http.Response
	if res, e = s.client.Do(req); e != nil {
		return nil, e
	}
	b, _ := io.ReadAll(res.Body)
	// gat a decoder to parse the service data
	d, e := s.decoderFactory.Create(s.format, bytes.NewReader(b))
	if e != nil {
		return nil, e
	}
	defer func() { _ = d.Close() }()
	// decode the data into a config instance
	return d.Decode()
}

func (s *RestSource) searchConfig(
	body IConfig,
) (*Config, error) {
	var cfg interface{}
	var e error
	// get the config from the response
	if cfg, e = body.(*Config).path(s.configPath); e != nil {
		return nil, errRestConfigNotFound(s.configPath)
	}
	// check if the obtained data is a valid config
	if p, ok := cfg.(Config); ok {
		return &p, nil
	}
	return nil, errConversion(cfg, "config.Config")
}
