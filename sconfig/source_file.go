package sconfig

import (
	"github.com/spf13/afero"
	"os"
	"sync"
)

type sourceFile struct {
	source
	path     string
	format   string
	fs       afero.Fs
	dFactory IDecoderFactory
}

var _ ISource = &sourceFile{}

func newSourceFile(path, format string, fs afero.Fs, dFactory IDecoderFactory) (ISource, error) {
	if fs == nil {
		return nil, errNilPointer("sfs")
	}
	if dFactory == nil {
		return nil, errNilPointer("dFactory")
	}

	s := &sourceFile{
		source: source{
			mutex:   &sync.Mutex{},
			partial: Partial{},
		},
		path:     path,
		format:   format,
		fs:       fs,
		dFactory: dFactory,
	}

	if e := s.load(); e != nil {
		return nil, e
	}

	return s, nil
}

func (s *sourceFile) load() error {
	f, e := s.fs.OpenFile(s.path, os.O_RDONLY, 0o644)
	if e != nil {
		return e
	}

	d, e := s.dFactory.Create(s.format, f)
	if e != nil {
		_ = f.Close()
		return e
	}
	defer func() { _ = d.Close() }()

	p, e := d.Decode()
	if e != nil {
		return e
	}

	s.mutex.Lock()
	s.partial = *p.(*Partial)
	s.mutex.Unlock()

	return nil
}
