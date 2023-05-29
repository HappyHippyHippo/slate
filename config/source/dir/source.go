package dir

import (
	"os"
	"sync"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/config/source"
	"github.com/spf13/afero"
)

type decoderCreator interface {
	Create(format string, args ...interface{}) (config.Decoder, error)
}

// Source defines a config source that read a directory files,
// recursive or not, and parse each one and store all the read content
// as a config.
type Source struct {
	source.Source
	path           string
	format         string
	recursive      bool
	fs             afero.Fs
	decoderCreator decoderCreator
}

var _ config.Source = &Source{}

// NewSource will instantiate a new configuration source
// that will read a directory files for configuration information.
func NewSource(
	path,
	format string,
	recursive bool,
	fs afero.Fs,
	decoderCreator decoderCreator,
) (*Source, error) {
	// check file system argument reference
	if fs == nil {
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
		recursive:      recursive,
		fs:             fs,
		decoderCreator: decoderCreator,
	}
	// load the dir files config content
	if e := s.load(); e != nil {
		return nil, e
	}
	return s, nil
}

func (s *Source) load() error {
	// load the source directory contents
	p, e := s.loadDir(s.path)
	if e != nil {
		return e
	}
	// store the parsed content into the source local config
	s.Mutex.Lock()
	s.Partial = *p
	s.Mutex.Unlock()
	return nil
}

func (s *Source) loadDir(
	path string,
) (*config.Partial, error) {
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
	loaded := &config.Partial{}
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
				loaded.Merge(*p)
			}
		} else {
			// load the file content
			p, e := s.loadFile(path + "/" + file.Name())
			if e != nil {
				return nil, e
			}
			// merge the loaded content
			loaded.Merge(*p)
		}
	}
	return loaded, nil
}

func (s *Source) loadFile(
	path string,
) (*config.Partial, error) {
	// open the file for reading
	f, e := s.fs.OpenFile(path, os.O_RDONLY, 0o644)
	if e != nil {
		return nil, e
	}
	// get a decoder to parse the read file content
	d, e := s.decoderCreator.Create(s.format, f)
	if e != nil {
		_ = f.Close()
		return nil, e
	}
	defer func() { _ = d.Close() }()
	// decode the read file content
	return d.Decode()
}
