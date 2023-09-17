package file

import (
	"os"
	"sync"

	"github.com/spf13/afero"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/config/source"
)

type decoderCreator interface {
	Create(format string, args ...interface{}) (config.Decoder, error)
}

// Source defines a config source that read a file content
// and stores its config contents to be used as a config.
type Source struct {
	source.Source
	path           string
	format         string
	fileSystem     afero.Fs
	decoderCreator decoderCreator
}

var _ config.Source = &Source{}

// NewSource will instantiate a new configuration source
// that will read a file for configuration info.
func NewSource(
	path,
	format string,
	fileSystem afero.Fs,
	decoderCreator decoderCreator,
) (*Source, error) {
	// check file system argument reference
	if fileSystem == nil {
		return nil, errNilPointer("fileSystem")
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
		path:           path,
		format:         format,
		fileSystem:     fileSystem,
		decoderCreator: decoderCreator,
	}
	// Load the file config content
	if e := s.load(); e != nil {
		return nil, e
	}
	return s, nil
}

func (s *Source) load() error {
	// open the source target file
	f, e := s.fileSystem.OpenFile(s.path, os.O_RDONLY, 0o644)
	if e != nil {
		return e
	}
	// creates the decoder to parse the file content
	d, e := s.decoderCreator.Create(s.format, f)
	if e != nil {
		_ = f.Close()
		return e
	}
	defer func() { _ = d.Close() }()
	// decode the file content
	p, e := d.Decode()
	if e != nil {
		return e
	}
	// store the parsed content into the source local config
	s.Mutex.Lock()
	s.Partial = *p
	s.Mutex.Unlock()
	return nil
}
