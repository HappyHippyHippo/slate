package sconfig

import (
	"bytes"
	"io"
	"net/http"
	"sync"
)

// HTTPClient defines the interface of an instance capable to perform the
// rest config obtain action
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type sourceRest struct {
	source
	client     HTTPClient
	uri        string
	format     string
	factory    *DecoderFactory
	configPath string
}

var _ Source = &sourceRest{}

func newSourceRest(client HTTPClient, uri, format string, factory *DecoderFactory, configPath string) (Source, error) {
	if client == nil {
		return nil, errNilPointer("client")
	}
	if factory == nil {
		return nil, errNilPointer("factory")
	}

	s := &sourceRest{
		source: source{
			mutex:   &sync.Mutex{},
			partial: Partial{},
		},
		client:     client,
		uri:        uri,
		format:     format,
		factory:    factory,
		configPath: configPath,
	}

	if err := s.load(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *sourceRest) load() error {
	r, err := s.request()
	if err != nil {
		return err
	}

	r, err = s.searchConfig(r)
	if err != nil {
		return err
	}

	s.mutex.Lock()
	s.partial = *r.(*Partial)
	s.mutex.Unlock()

	return nil
}

func (s *sourceRest) request() (Config, error) {
	var err error

	var req *http.Request
	if req, err = http.NewRequest(http.MethodGet, s.uri, http.NoBody); err != nil {
		return nil, err
	}

	var res *http.Response
	if res, err = s.client.Do(req); err != nil {
		return nil, err
	}

	b, _ := io.ReadAll(res.Body)

	d, err := s.factory.Create(s.format, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	defer func() { _ = d.Close() }()

	return d.Decode()
}

func (s *sourceRest) searchConfig(body Config) (Config, error) {
	var err error

	var cfg interface{}
	if cfg, err = body.(*Partial).path(s.configPath); err != nil {
		return nil, errConfigRestPathNotFound(s.configPath)
	}

	if p, ok := cfg.(Partial); ok {
		return &p, nil
	}

	return nil, errConversion(cfg, "config.Partial")
}
