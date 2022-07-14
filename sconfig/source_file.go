package sconfig

import (
	"github.com/spf13/afero"
	"os"
	"sync"
)

type sourceFile struct {
	source
	path    string
	format  string
	fs      afero.Fs
	factory IDecoderFactory
}

var _ ISource = &sourceFile{}

func newSourceFile(path, format string, fs afero.Fs, factory IDecoderFactory) (ISource, error) {
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	if factory == nil {
		return nil, errNilPointer("dFactory")
	}

	s := &sourceFile{
		source: source{
			mutex:   &sync.Mutex{},
			partial: Partial{},
		},
		path:    path,
		format:  format,
		fs:      fs,
		factory: factory,
	}

	if err := s.load(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *sourceFile) load() error {
	f, err := s.fs.OpenFile(s.path, os.O_RDONLY, 0o644)
	if err != nil {
		return err
	}

	d, err := s.factory.Create(s.format, f)
	if err != nil {
		_ = f.Close()
		return err
	}
	defer func() { _ = d.Close() }()

	p, err := d.Decode()
	if err != nil {
		return err
	}

	s.mutex.Lock()
	s.partial = *p.(*Partial)
	s.mutex.Unlock()

	return nil
}
