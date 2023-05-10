package file

import (
	"os"
	"sync"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/config/source"
	"github.com/spf13/afero"
)

// Source defines a config source that read a file content
// and stores its config contents to be used as a config.
type Source struct {
	source.BaseSource
	path           string
	format         string
	fileSystem     afero.Fs
	decoderFactory config.IDecoderFactory
}

var _ config.ISource = &Source{}

// NewSource will instantiate a new configuration source
// that will read a file for configuration info.
func NewSource(
	path,
	format string,
	fileSystem afero.Fs,
	decoderFactory config.IDecoderFactory,
) (*Source, error) {
	// check file system argument reference
	if fileSystem == nil {
		return nil, errNilPointer("fileSystem")
	}
	// check decoder factory argument reference
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}
	// instantiates the config source
	s := &Source{
		BaseSource: source.BaseSource{
			Mutex:  &sync.Mutex{},
			Config: config.Config{},
		},
		path:           path,
		format:         format,
		fileSystem:     fileSystem,
		decoderFactory: decoderFactory,
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
	d, e := s.decoderFactory.Create(s.format, f)
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
	s.Config = *p.(*config.Config)
	s.Mutex.Unlock()
	return nil
}
