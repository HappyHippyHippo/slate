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
	t.Run("nil cfg", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, err := newLoader(nil, NewMockSourceFactory(ctrl))
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("nil source dFactory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, err := newLoader(NewMockManager(ctrl), nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("new loader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		if sut, err := newLoader(NewMockManager(ctrl), NewMockSourceFactory(ctrl)); sut == nil {
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
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(nil, expected).Times(1)
		sut, _ := newLoader(NewMockManager(ctrl), sFactory)

		if err := sut.Load(); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("error storing the base source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		src := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(src, nil).Times(1)
		mngr := NewMockManager(ctrl)
		mngr.EXPECT().AddSource(LoaderSourceID, 0, src).Return(expected).Times(1)
		sut, _ := newLoader(mngr, sFactory)

		if err := sut.Load(); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("add base source into the config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		src := NewMockSource(ctrl)
		src.EXPECT().Get("").Return(Partial{}, nil)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(src, nil).Times(1)
		cfg := NewManager(0 * time.Second)
		sut, _ := newLoader(cfg, sFactory)

		if err := sut.Load(); err != nil {
			t.Errorf("returned the (%v) error", err)
		} else if len(cfg.(*manager).sources) != 1 {
			t.Error("didn't stored the requested base source")
		}
	})

	t.Run("invalid list of sources results in an empty sources list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		src := NewMockSource(ctrl)
		src.EXPECT().Get("").Return(Partial{"configs": 123}, nil)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(src, nil).Times(1)
		cfg := NewManager(0 * time.Second)
		sut, _ := newLoader(cfg, sFactory)

		if err := sut.Load(); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		}
	})

	t.Run("error on loaded missing id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := Partial{"configs": []interface{}{Partial{}}}
		src := NewMockSource(ctrl)
		src.EXPECT().Get("").Return(partial, nil)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(src, nil).Times(1)
		cfg := NewManager(0 * time.Second)
		sut, _ := newLoader(cfg, sFactory)

		if err := sut.Load(); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConfigPathNotFound) {
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrConfigPathNotFound)
		}
	})

	t.Run("error on loaded invalid id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := Partial{"configs": []interface{}{Partial{"id": 12}}}
		src := NewMockSource(ctrl)
		src.EXPECT().Get("").Return(partial, nil)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(src, nil).Times(1)
		cfg := NewManager(0 * time.Second)
		sut, _ := newLoader(cfg, sFactory)

		if err := sut.Load(); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("error on loaded invalid priority", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := Partial{"configs": []interface{}{Partial{"id": "id", "priority": "string"}}}
		src := NewMockSource(ctrl)
		src.EXPECT().Get("").Return(partial, nil)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(src, nil).Times(1)
		cfg := NewManager(0 * time.Second)
		sut, _ := newLoader(cfg, sFactory)

		if err := sut.Load(); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("error on loaded source factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		srcPartial := Partial{"id": "id", "priority": 0}
		partial := Partial{"configs": []interface{}{srcPartial}}
		src := NewMockSource(ctrl)
		src.EXPECT().Get("").Return(partial, nil)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(src, nil).Times(1)
		sFactory.EXPECT().CreateFromConfig(&srcPartial).Return(nil, expected).Times(1)
		cfg := NewManager(0 * time.Second)
		sut, _ := newLoader(cfg, sFactory)

		if err := sut.Load(); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("error on source registration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		srcPartial1 := Partial{"id": "id", "priority": 0}
		srcPartial2 := Partial{"id": "id", "priority": 2}
		partial := Partial{"configs": []interface{}{srcPartial1, srcPartial2}}
		src := NewMockSource(ctrl)
		src.EXPECT().Get("").Return(partial, nil)
		src1 := NewMockSource(ctrl)
		src1.EXPECT().Get("").Return(partial, nil)
		src2 := NewMockSource(ctrl)
		sFactory := NewMockSourceFactory(ctrl)
		sFactory.EXPECT().Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat).Return(src, nil).Times(1)
		gomock.InOrder(
			sFactory.EXPECT().CreateFromConfig(&srcPartial1).Return(src1, nil),
			sFactory.EXPECT().CreateFromConfig(&srcPartial2).Return(src2, nil),
		)
		cfg := NewManager(0 * time.Second)
		sut, _ := newLoader(cfg, sFactory)

		if err := sut.Load(); err == nil {
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
		dFactory := &decoderFactory{}
		_ = dFactory.Register(&decoderStrategyYAML{})
		sFactory := &sourceFactory{}
		fileStrategy, _ := newSourceStrategyFile(fs, dFactory)
		_ = sFactory.Register(fileStrategy)
		obsFileStrategy, _ := newSourceStrategyObservableFile(fs, dFactory)
		_ = sFactory.Register(obsFileStrategy)
		envStrategy := &sourceStrategyEnv{}
		_ = sFactory.Register(envStrategy)
		cfg := NewManager(0 * time.Second)
		sut, _ := newLoader(cfg, sFactory)

		if err := sut.Load(); err != nil {
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
		dFactory := &decoderFactory{}
		_ = dFactory.Register(&decoderStrategyYAML{})
		sFactory := &sourceFactory{}
		fileStrategy, _ := newSourceStrategyFile(fs, dFactory)
		_ = sFactory.Register(fileStrategy)
		obsFileStrategy, _ := newSourceStrategyObservableFile(fs, dFactory)
		_ = sFactory.Register(obsFileStrategy)
		envStrategy := &sourceStrategyEnv{}
		_ = sFactory.Register(envStrategy)
		cfg := NewManager(0 * time.Second)
		sut, _ := newLoader(cfg, sFactory)

		if err := sut.Load(); err != nil {
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
		dFactory := &decoderFactory{}
		_ = dFactory.Register(&decoderStrategyYAML{})
		sFactory := &sourceFactory{}
		fileStrategy, _ := newSourceStrategyFile(fs, dFactory)
		_ = sFactory.Register(fileStrategy)
		obsFileStrategy, _ := newSourceStrategyObservableFile(fs, dFactory)
		_ = sFactory.Register(obsFileStrategy)
		envStrategy := &sourceStrategyEnv{}
		_ = sFactory.Register(envStrategy)
		cfg := NewManager(0 * time.Second)
		sut, _ := newLoader(cfg, sFactory)

		if err := sut.Load(); err != nil {
			t.Errorf("returned the (%v) error", err)
		} else if len(cfg.(*manager).sources) != 2 {
			t.Error("didn't registered the requested config source")
		} else if chk := cfg.(*manager).sources[0].priority; chk != 0 {
			t.Errorf("registered the config source with priority of %d when expecting 0", chk)
		}
	})
}
