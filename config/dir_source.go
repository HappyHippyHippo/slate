package config

import (
	"os"
	"sync"

	"github.com/spf13/afero"
)

// DirSource defines a config source that read a directory files,
// recursive or not, and parse each one and store all the read content
// as a config.
type DirSource struct {
	Source
	path           string
	format         string
	recursive      bool
	fs             afero.Fs
	decoderFactory IDecoderFactory
}

var _ ISource = &DirSource{}

// NewDirSource will instantiate a new configuration source
// that will read a directory files for configuration information.
func NewDirSource(
	path,
	format string,
	recursive bool,
	fs afero.Fs,
	decoderFactory IDecoderFactory,
) (*DirSource, error) {
	// check file system argument reference
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	// check decoder factory argument reference
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}
	// instantiates the config source
	s := &DirSource{
		Source: Source{
			mutex:   &sync.Mutex{},
			partial: Config{},
		},
		path:           path,
		format:         format,
		recursive:      recursive,
		fs:             fs,
		decoderFactory: decoderFactory,
	}
	// load the dir files config content
	if e := s.load(); e != nil {
		return nil, e
	}
	return s, nil
}

func (s *DirSource) load() error {
	// load the source directory contents
	p, e := s.loadDir(s.path)
	if e != nil {
		return e
	}
	// store the parsed content into the source local config
	s.mutex.Lock()
	s.partial = *p
	s.mutex.Unlock()
	return nil
}

func (s *DirSource) loadDir(
	path string,
) (*Config, error) {
	// open the directory stream
	dir, e := s.fs.Open(path)
	if e != nil {
		return nil, e
	}
	defer func() { _ = dir.Close() }()
	// get the dir entry list
	files, e := dir.Readdir(0)
	if e != nil {
		return nil, e
	}
	// parse each founded entry
	loaded := &Config{}
	for _, file := range files {
		// check if is an inner directory
		if file.IsDir() {
			// load the founded directory if the source is
			// configured to be recursive
			if s.recursive {
				p, e := s.loadDir(path + "/" + file.Name())
				if e != nil {
					return nil, e
				}
				// merge the loaded content
				loaded.merge(*p)
			}
		} else {
			// load the file content
			p, e := s.loadFile(path + "/" + file.Name())
			if e != nil {
				return nil, e
			}
			// merge the loaded content
			loaded.merge(*p)
		}
	}
	return loaded, nil
}

func (s *DirSource) loadFile(
	path string,
) (*Config, error) {
	// open the file for reading
	f, e := s.fs.OpenFile(path, os.O_RDONLY, 0o644)
	if e != nil {
		return nil, e
	}
	// get a decoder to parse the read file content
	d, e := s.decoderFactory.Create(s.format, f)
	if e != nil {
		_ = f.Close()
		return nil, e
	}
	defer func() { _ = d.Close() }()
	// decode the read file content
	p, e := d.Decode()
	if e != nil {
		return nil, e
	}
	return p.(*Config), nil
}
