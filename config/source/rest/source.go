package rest

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"sync"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/config/source"
)

type requester interface {
	Do(req *http.Request) (*http.Response, error)
}

type decoderCreator interface {
	Create(format string, args ...interface{}) (config.Decoder, error)
}

// Source defines a config source that read a REST service and
// store a section of the response as the stored config.
type Source struct {
	source.Source
	client         requester
	uri            string
	format         string
	decoderCreator decoderCreator
	configPath     string
}

var _ config.Source = &Source{}

// NewSource will instantiate a new configuration source
// that will read a REST endpoint for configuration info.
func NewSource(
	client requester,
	uri,
	format string,
	decoderCreator decoderCreator,
	configPath string,
) (*Source, error) {
	// check client argument reference
	if client == nil {
		return nil, errNilPointer("client")
	}
	// check decoder factory argument reference
	if decoderCreator == nil {
		return nil, errNilPointer("decoderCreator")
	}
	// instantiates the config source
	s := &Source{
		Source: source.Source{
			Mutex:   &sync.Mutex{},
			Partial: config.Partial{},
		},
		client:         client,
		uri:            uri,
		format:         format,
		decoderCreator: decoderCreator,
		configPath:     configPath,
	}
	// load the config information from the REST service
	if e := s.load(); e != nil {
		return nil, e
	}
	return s, nil
}

func (s *Source) load() error {
	// get the REST service information
	cfg, e := s.request()
	if e != nil {
		return e
	}
	// retrieve the config information from the service response data
	c, e := cfg.Partial(s.configPath)
	if e != nil {
		if errors.Is(e, config.ErrPathNotFound) {
			return errConfigNotFound(s.configPath)
		}
		return e
	}
	// store the retrieved config
	s.Mutex.Lock()
	s.Partial = *c
	s.Mutex.Unlock()
	return nil
}

func (s *Source) request() (*config.Partial, error) {
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
	d, e := s.decoderCreator.Create(s.format, bytes.NewReader(b))
	if e != nil {
		return nil, e
	}
	defer func() { _ = d.Close() }()
	// decode the data into a config instance
	return d.Decode()
}
