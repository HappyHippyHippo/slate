package config

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
	dFactory   IDecoderFactory
	configPath string
}

var _ ISource = &sourceRest{}

func newSourceRest(client HTTPClient, uri, format string, dFactory IDecoderFactory, configPath string) (ISource, error) {
	if client == nil {
		return nil, errNilPointer("client")
	}
	if dFactory == nil {
		return nil, errNilPointer("dFactory")
	}

	s := &sourceRest{
		source: source{
			mutex:   &sync.Mutex{},
			partial: Partial{},
		},
		client:     client,
		uri:        uri,
		format:     format,
		dFactory:   dFactory,
		configPath: configPath,
	}

	if e := s.load(); e != nil {
		return nil, e
	}

	return s, nil
}

func (s *sourceRest) load() error {
	r, e := s.request()
	if e != nil {
		return e
	}

	r, e = s.searchConfig(r)
	if e != nil {
		return e
	}

	s.mutex.Lock()
	s.partial = *r.(*Partial)
	s.mutex.Unlock()

	return nil
}

func (s *sourceRest) request() (IConfig, error) {
	var e error

	var req *http.Request
	if req, e = http.NewRequest(http.MethodGet, s.uri, http.NoBody); e != nil {
		return nil, e
	}

	var res *http.Response
	if res, e = s.client.Do(req); e != nil {
		return nil, e
	}

	b, _ := io.ReadAll(res.Body)

	d, e := s.dFactory.Create(s.format, bytes.NewReader(b))
	if e != nil {
		return nil, e
	}
	defer func() { _ = d.Close() }()

	return d.Decode()
}

func (s *sourceRest) searchConfig(body IConfig) (IConfig, error) {
	var cfg interface{}
	var e error

	if cfg, e = body.(*Partial).path(s.configPath); e != nil {
		return nil, errConfigRestPathNotFound(s.configPath)
	}

	if p, ok := cfg.(Partial); ok {
		return &p, nil
	}

	return nil, errConversion(cfg, "config.Partial")
}
