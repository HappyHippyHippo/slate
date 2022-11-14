package config

import (
	"os"
	"sync"

	"github.com/spf13/afero"
)

type sourceFile struct {
	source
	path           string
	format         string
	fs             afero.Fs
	decoderFactory IDecoderFactory
}

var _ ISource = &sourceFile{}

// NewSourceFile will instantiate a new configuration source
// that will read a file for configuration info.
func NewSourceFile(path, format string, fs afero.Fs, decoderFactory IDecoderFactory) (ISource, error) {
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}

	s := &sourceFile{
		source: source{
			mutex:   &sync.Mutex{},
			partial: Partial{},
		},
		path:           path,
		format:         format,
		fs:             fs,
		decoderFactory: decoderFactory,
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

	d, e := s.decoderFactory.Create(s.format, f)
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
