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
	factory   *DecoderFactory
}

var _ Source = &sourceDir{}

func newSourceDir(path, format string, recursive bool, fs afero.Fs, factory *DecoderFactory) (Source, error) {
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	if factory == nil {
		return nil, errNilPointer("factory")
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
		factory:   factory,
	}

	if err := s.load(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *sourceDir) load() error {
	p, err := s.loadDir(s.path)
	if err != nil {
		return err
	}

	s.mutex.Lock()
	s.partial = *p
	s.mutex.Unlock()

	return nil
}

func (s *sourceDir) loadDir(path string) (*Partial, error) {
	dir, err := s.fs.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = dir.Close() }()

	files, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	loaded := &Partial{}

	for _, file := range files {
		if file.IsDir() {
			if s.recursive {
				p, err := s.loadDir(path + "/" + file.Name())
				if err != nil {
					return nil, err
				}
				loaded.merge(*p)
			}
		} else {
			p, err := s.loadFile(path + "/" + file.Name())
			if err != nil {
				return nil, err
			}
			loaded.merge(*p)
		}
	}

	return loaded, nil
}

func (s *sourceDir) loadFile(path string) (*Partial, error) {
	f, err := s.fs.OpenFile(path, os.O_RDONLY, 0o644)
	if err != nil {
		return nil, err
	}

	d, err := s.factory.Create(s.format, f)
	if err != nil {
		_ = f.Close()
		return nil, err
	}
	defer func() { _ = d.Close() }()

	p, err := d.Decode()
	if err != nil {
		return nil, err
	}

	return p.(*Partial), nil
}
