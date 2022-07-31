package sconfig

import (
	"github.com/spf13/afero"
	"os"
	"sync"
)

type sourceDir struct {
	source
	path      string
	format    string
	recursive bool
	fs        afero.Fs
	dFactory  IDecoderFactory
}

var _ ISource = &sourceDir{}

// NewSourceDir will instantiate a new configuration source
// that will read a directory files for configuration information.
func NewSourceDir(path, format string, recursive bool, fs afero.Fs, dFactory IDecoderFactory) (ISource, error) {
	if fs == nil {
		return nil, errNilPointer("sfs")
	}
	if dFactory == nil {
		return nil, errNilPointer("dFactory")
	}

	s := &sourceDir{
		source: source{
			mutex:   &sync.Mutex{},
			partial: Partial{},
		},
		path:      path,
		format:    format,
		recursive: recursive,
		fs:        fs,
		dFactory:  dFactory,
	}

	if e := s.load(); e != nil {
		return nil, e
	}

	return s, nil
}

func (s *sourceDir) load() error {
	p, e := s.loadDir(s.path)
	if e != nil {
		return e
	}

	s.mutex.Lock()
	s.partial = *p
	s.mutex.Unlock()

	return nil
}

func (s *sourceDir) loadDir(path string) (*Partial, error) {
	dir, e := s.fs.Open(path)
	if e != nil {
		return nil, e
	}
	defer func() { _ = dir.Close() }()

	files, e := dir.Readdir(0)
	if e != nil {
		return nil, e
	}

	loaded := &Partial{}

	for _, file := range files {
		if file.IsDir() {
			if s.recursive {
				p, e := s.loadDir(path + "/" + file.Name())
				if e != nil {
					return nil, e
				}
				loaded.merge(*p)
			}
		} else {
			p, e := s.loadFile(path + "/" + file.Name())
			if e != nil {
				return nil, e
			}
			loaded.merge(*p)
		}
	}

	return loaded, nil
}

func (s *sourceDir) loadFile(path string) (*Partial, error) {
	f, e := s.fs.OpenFile(path, os.O_RDONLY, 0o644)
	if e != nil {
		return nil, e
	}

	d, e := s.dFactory.Create(s.format, f)
	if e != nil {
		_ = f.Close()
		return nil, e
	}
	defer func() { _ = d.Close() }()

	p, e := d.Decode()
	if e != nil {
		return nil, e
	}

	return p.(*Partial), nil
}
