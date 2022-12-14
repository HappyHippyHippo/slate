package config

import (
	"os"
	"sync"

	"github.com/spf13/afero"
)

// FileSource defines a config source that read a file content
// and stores its config contents to be used as a config.
type FileSource struct {
	Source
	path           string
	format         string
	fs             afero.Fs
	decoderFactory IDecoderFactory
}

var _ ISource = &FileSource{}

// NewFileSource will instantiate a new configuration source
// that will read a file for configuration info.
func NewFileSource(
	path,
	format string,
	fs afero.Fs,
	decoderFactory IDecoderFactory,
) (*FileSource, error) {
	// check file system argument reference
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	// check decoder factory argument reference
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}
	// instantiates the config source
	s := &FileSource{
		Source: Source{
			mutex:   &sync.Mutex{},
			partial: Config{},
		},
		path:           path,
		format:         format,
		fs:             fs,
		decoderFactory: decoderFactory,
	}
	// load the file config content
	if e := s.load(); e != nil {
		return nil, e
	}
	return s, nil
}

func (s *FileSource) load() error {
	// open the source target file
	f, e := s.fs.OpenFile(s.path, os.O_RDONLY, 0o644)
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
	s.mutex.Lock()
	s.partial = *p.(*Config)
	s.mutex.Unlock()
	return nil
}
