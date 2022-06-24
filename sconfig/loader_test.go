package sconfig

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serror"
	"io"
	"os"
	"testing"
	"time"
)

func Test_NewLoader(t *testing.T) {
	cfg := NewConfig(0 * time.Second)
	factory := &SourceFactory{}

	t.Run("nil cfg", func(t *testing.T) {
		load, err := NewLoader(nil, factory)
		switch {
		case load != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("nil source factory", func(t *testing.T) {
		load, err := NewLoader(cfg, nil)
		switch {
		case load != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("new loader", func(t *testing.T) {
		if load, err := NewLoader(cfg, factory); load == nil {
			t.Error("didn't returned a valid reference")
		} else if err != nil {
			t.Errorf("return the (%v) error", err)
		}
	})
}

func Test_Loader_Load(t *testing.T) {
	LoaderSourceID = "base_source_id"
	LoaderSourcePath = "base_source_path"
	LoaderSourceFormat = DecoderFormatYAML
	defer func() {
		LoaderSourceID = "main"
		LoaderSourcePath = "config/config.yaml"
		LoaderSourceFormat = DecoderFormatYAML
	}()

	t.Run("error getting the base source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(LoaderSourcePath, os.O_RDONLY, os.FileMode(0o644)).Return(nil, expected).Times(1)
		decoderFactory := &DecoderFactory{}
		_ = decoderFactory.Register(&decoderStrategyYAML{})
		sourceFactory := &SourceFactory{}
		fileStrategy, _ := NewSourceStrategyFile(fs, decoderFactory)
		_ = sourceFactory.Register(fileStrategy)
		obsFileStrategy, _ := NewSourceStrategyObservableFile(fs, decoderFactory)
		_ = sourceFactory.Register(obsFileStrategy)
		envStrategy := &sourceStrategyEnv{}
		_ = sourceFactory.Register(envStrategy)
		cfg := NewConfig(0 * time.Second)
		load, _ := NewLoader(cfg, sourceFactory)

		if err := load.Load(); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("error storing the base source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := "field: value"
		file := NewMockFile(ctrl)
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, data)
			return len(data), io.EOF
		}).Times(1)
		file.EXPECT().Close().Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(LoaderSourcePath, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoderFactory := &DecoderFactory{}
		_ = decoderFactory.Register(&decoderStrategyYAML{})
		sourceFactory := &SourceFactory{}
		fileStrategy, _ := NewSourceStrategyFile(fs, decoderFactory)
		_ = sourceFactory.Register(fileStrategy)
		obsFileStrategy, _ := NewSourceStrategyObservableFile(fs, decoderFactory)
		_ = sourceFactory.Register(obsFileStrategy)
		envStrategy := &sourceStrategyEnv{}
		_ = sourceFactory.Register(envStrategy)
		src := NewMockSource(ctrl)
		src.EXPECT().Get("").Return(Partial{}, nil)
		cfg := NewConfig(0 * time.Second)
		_ = cfg.AddSource(LoaderSourceID, 0, src)
		load, _ := NewLoader(cfg, sourceFactory)

		if err := load.Load(); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrDuplicateConfigSource) {
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrDuplicateConfigSource)
		}
	})

	t.Run("add base source into the config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := "field: value"
		file := NewMockFile(ctrl)
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, data)
			return len(data), io.EOF
		}).Times(1)
		file.EXPECT().Close().Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(LoaderSourcePath, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoderFactory := &DecoderFactory{}
		_ = decoderFactory.Register(&decoderStrategyYAML{})
		sourceFactory := &SourceFactory{}
		fileStrategy, _ := NewSourceStrategyFile(fs, decoderFactory)
		_ = sourceFactory.Register(fileStrategy)
		obsFileStrategy, _ := NewSourceStrategyObservableFile(fs, decoderFactory)
		_ = sourceFactory.Register(obsFileStrategy)
		envStrategy := &sourceStrategyEnv{}
		_ = sourceFactory.Register(envStrategy)
		cfg := NewConfig(0 * time.Second)
		load, _ := NewLoader(cfg, sourceFactory)

		if err := load.Load(); err != nil {
			t.Errorf("returned the (%v) error", err)
		}
	})

	t.Run("invalid list of sources results in an empty sources list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := "config:\n  sources: 123"
		file := NewMockFile(ctrl)
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, data)
			return len(data), io.EOF
		}).Times(1)
		file.EXPECT().Close().Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(LoaderSourcePath, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoderFactory := &DecoderFactory{}
		_ = decoderFactory.Register(&decoderStrategyYAML{})
		sourceFactory := &SourceFactory{}
		fileStrategy, _ := NewSourceStrategyFile(fs, decoderFactory)
		_ = sourceFactory.Register(fileStrategy)
		obsFileStrategy, _ := NewSourceStrategyObservableFile(fs, decoderFactory)
		_ = sourceFactory.Register(obsFileStrategy)
		envStrategy := &sourceStrategyEnv{}
		_ = sourceFactory.Register(envStrategy)
		cfg := NewConfig(0 * time.Second)
		load, _ := NewLoader(cfg, sourceFactory)

		if err := load.Load(); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		}
	})

	t.Run("error on loaded missing id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := "configs:\n  - priority: 12\n"
		file := NewMockFile(ctrl)
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, data)
			return len(data), io.EOF
		}).Times(1)
		file.EXPECT().Close().Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(LoaderSourcePath, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoderFactory := &DecoderFactory{}
		_ = decoderFactory.Register(&decoderStrategyYAML{})
		sourceFactory := &SourceFactory{}
		fileStrategy, _ := NewSourceStrategyFile(fs, decoderFactory)
		_ = sourceFactory.Register(fileStrategy)
		obsFileStrategy, _ := NewSourceStrategyObservableFile(fs, decoderFactory)
		_ = sourceFactory.Register(obsFileStrategy)
		envStrategy := &sourceStrategyEnv{}
		_ = sourceFactory.Register(envStrategy)
		cfg := NewConfig(0 * time.Second)
		load, _ := NewLoader(cfg, sourceFactory)

		if err := load.Load(); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConfigPathNotFound) {
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrConfigPathNotFound)
		}
	})

	t.Run("error on loaded invalid id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := "configs:\n  - id: 12\n"
		file := NewMockFile(ctrl)
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, data)
			return len(data), io.EOF
		}).Times(1)
		file.EXPECT().Close().Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(LoaderSourcePath, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoderFactory := &DecoderFactory{}
		_ = decoderFactory.Register(&decoderStrategyYAML{})
		sourceFactory := &SourceFactory{}
		fileStrategy, _ := NewSourceStrategyFile(fs, decoderFactory)
		_ = sourceFactory.Register(fileStrategy)
		obsFileStrategy, _ := NewSourceStrategyObservableFile(fs, decoderFactory)
		_ = sourceFactory.Register(obsFileStrategy)
		envStrategy := &sourceStrategyEnv{}
		_ = sourceFactory.Register(envStrategy)
		cfg := NewConfig(0 * time.Second)
		load, _ := NewLoader(cfg, sourceFactory)

		if err := load.Load(); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("error on loaded invalid priority", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := "configs:\n  - id: id\n    priority: string"
		file := NewMockFile(ctrl)
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, data)
			return len(data), io.EOF
		}).Times(1)
		file.EXPECT().Close().Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(LoaderSourcePath, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoderFactory := &DecoderFactory{}
		_ = decoderFactory.Register(&decoderStrategyYAML{})
		sourceFactory := &SourceFactory{}
		fileStrategy, _ := NewSourceStrategyFile(fs, decoderFactory)
		_ = sourceFactory.Register(fileStrategy)
		obsFileStrategy, _ := NewSourceStrategyObservableFile(fs, decoderFactory)
		_ = sourceFactory.Register(obsFileStrategy)
		envStrategy := &sourceStrategyEnv{}
		_ = sourceFactory.Register(envStrategy)
		cfg := NewConfig(0 * time.Second)
		load, _ := NewLoader(cfg, sourceFactory)

		if err := load.Load(); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("error on loaded source factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := "configs:\n  - id: id\n    priority: 0"
		file := NewMockFile(ctrl)
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, data)
			return len(data), io.EOF
		}).Times(1)
		file.EXPECT().Close().Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(LoaderSourcePath, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoderFactory := &DecoderFactory{}
		_ = decoderFactory.Register(&decoderStrategyYAML{})
		sourceFactory := &SourceFactory{}
		fileStrategy, _ := NewSourceStrategyFile(fs, decoderFactory)
		_ = sourceFactory.Register(fileStrategy)
		obsFileStrategy, _ := NewSourceStrategyObservableFile(fs, decoderFactory)
		_ = sourceFactory.Register(obsFileStrategy)
		envStrategy := &sourceStrategyEnv{}
		_ = sourceFactory.Register(envStrategy)
		cfg := NewConfig(0 * time.Second)
		load, _ := NewLoader(cfg, sourceFactory)

		if err := load.Load(); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrInvalidConfigSourcePartial) {
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrInvalidConfigSourcePartial)
		}
	})

	t.Run("error on source registration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := "configs:\n  - id: id\n    priority: 0\n    type: file\n    path: path\n    format: yaml"
		file1 := NewMockFile(ctrl)
		file1.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, data)
			return len(data), io.EOF
		}).Times(1)
		file1.EXPECT().Close().Times(1)
		file2 := NewMockFile(ctrl)
		file2.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, "field: value")
			return 12, io.EOF
		}).Times(1)
		file2.EXPECT().Close().Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(LoaderSourcePath, os.O_RDONLY, os.FileMode(0o644)).Return(file1, nil).Times(1)
		fs.EXPECT().OpenFile("path", os.O_RDONLY, os.FileMode(0o644)).Return(file2, nil).Times(1)
		decoderFactory := &DecoderFactory{}
		_ = decoderFactory.Register(&decoderStrategyYAML{})
		sourceFactory := &SourceFactory{}
		fileStrategy, _ := NewSourceStrategyFile(fs, decoderFactory)
		_ = sourceFactory.Register(fileStrategy)
		obsFileStrategy, _ := NewSourceStrategyObservableFile(fs, decoderFactory)
		_ = sourceFactory.Register(obsFileStrategy)
		envStrategy := &sourceStrategyEnv{}
		_ = sourceFactory.Register(envStrategy)
		src := NewMockSource(ctrl)
		src.EXPECT().Get("").Return(Partial{}, nil).AnyTimes()
		cfg := NewConfig(0 * time.Second)
		_ = cfg.AddSource("id", 0, src)
		load, _ := NewLoader(cfg, sourceFactory)

		if err := load.Load(); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrDuplicateConfigSource) {
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrDuplicateConfigSource)
		}
	})

	t.Run("register the loaded source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := "configs:\n  - id: id\n    priority: 0\n    type: file\n    path: path\n    format: yaml"
		file1 := NewMockFile(ctrl)
		file1.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, data)
			return len(data), io.EOF
		}).Times(1)
		file1.EXPECT().Close().Times(1)
		file2 := NewMockFile(ctrl)
		file2.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, "field: value")
			return 12, io.EOF
		}).Times(1)
		file2.EXPECT().Close().Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(LoaderSourcePath, os.O_RDONLY, os.FileMode(0o644)).Return(file1, nil).Times(1)
		fs.EXPECT().OpenFile("path", os.O_RDONLY, os.FileMode(0o644)).Return(file2, nil).Times(1)
		decoderFactory := &DecoderFactory{}
		_ = decoderFactory.Register(&decoderStrategyYAML{})
		sourceFactory := &SourceFactory{}
		fileStrategy, _ := NewSourceStrategyFile(fs, decoderFactory)
		_ = sourceFactory.Register(fileStrategy)
		obsFileStrategy, _ := NewSourceStrategyObservableFile(fs, decoderFactory)
		_ = sourceFactory.Register(obsFileStrategy)
		envStrategy := &sourceStrategyEnv{}
		_ = sourceFactory.Register(envStrategy)
		cfg := NewConfig(0 * time.Second)
		load, _ := NewLoader(cfg, sourceFactory)

		if err := load.Load(); err != nil {
			t.Errorf("returned the (%v) error", err)
		} else if len(cfg.(*manager).sources) != 2 {
			t.Error("didn't registered the requested config source")
		}
	})

	t.Run("load from defined source path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		LoaderSourceListPath = "cfgs"
		defer func() {
			LoaderSourceListPath = "configs"
		}()

		data := "cfgs:\n  - id: id\n    priority: 0\n    type: file\n    path: path\n    format: yaml"
		file1 := NewMockFile(ctrl)
		file1.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, data)
			return len(data), io.EOF
		}).Times(1)
		file1.EXPECT().Close().Times(1)
		file2 := NewMockFile(ctrl)
		file2.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, "field: value")
			return 12, io.EOF
		}).Times(1)
		file2.EXPECT().Close().Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(LoaderSourcePath, os.O_RDONLY, os.FileMode(0o644)).Return(file1, nil).Times(1)
		fs.EXPECT().OpenFile("path", os.O_RDONLY, os.FileMode(0o644)).Return(file2, nil).Times(1)
		decoderFactory := &DecoderFactory{}
		_ = decoderFactory.Register(&decoderStrategyYAML{})
		sourceFactory := &SourceFactory{}
		fileStrategy, _ := NewSourceStrategyFile(fs, decoderFactory)
		_ = sourceFactory.Register(fileStrategy)
		obsFileStrategy, _ := NewSourceStrategyObservableFile(fs, decoderFactory)
		_ = sourceFactory.Register(obsFileStrategy)
		envStrategy := &sourceStrategyEnv{}
		_ = sourceFactory.Register(envStrategy)
		cfg := NewConfig(0 * time.Second)
		load, _ := NewLoader(cfg, sourceFactory)

		if err := load.Load(); err != nil {
			t.Errorf("returned the (%v) error", err)
		} else if len(cfg.(*manager).sources) != 2 {
			t.Error("didn't registered the requested config source")
		}
	})

	t.Run("register the loaded source with default priority if missing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := "configs:\n  - id: id\n    type: file\n    path: path\n    format: yaml"
		file1 := NewMockFile(ctrl)
		file1.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, data)
			return len(data), io.EOF
		}).Times(1)
		file1.EXPECT().Close().Times(1)
		file2 := NewMockFile(ctrl)
		file2.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, "field: value")
			return 12, io.EOF
		}).Times(1)
		file2.EXPECT().Close().Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(LoaderSourcePath, os.O_RDONLY, os.FileMode(0o644)).Return(file1, nil).Times(1)
		fs.EXPECT().OpenFile("path", os.O_RDONLY, os.FileMode(0o644)).Return(file2, nil).Times(1)
		decoderFactory := &DecoderFactory{}
		_ = decoderFactory.Register(&decoderStrategyYAML{})
		sourceFactory := &SourceFactory{}
		fileStrategy, _ := NewSourceStrategyFile(fs, decoderFactory)
		_ = sourceFactory.Register(fileStrategy)
		obsFileStrategy, _ := NewSourceStrategyObservableFile(fs, decoderFactory)
		_ = sourceFactory.Register(obsFileStrategy)
		envStrategy := &sourceStrategyEnv{}
		_ = sourceFactory.Register(envStrategy)
		cfg := NewConfig(0 * time.Second)
		load, _ := NewLoader(cfg, sourceFactory)

		if err := load.Load(); err != nil {
			t.Errorf("returned the (%v) error", err)
		} else if len(cfg.(*manager).sources) != 2 {
			t.Error("didn't registered the requested config source")
		} else if chk := cfg.(*manager).sources[0].priority; chk != 0 {
			t.Errorf("registered the config source with priority of %d when expecting 0", chk)
		}
	})
}
