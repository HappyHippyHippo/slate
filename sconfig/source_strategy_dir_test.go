package sconfig

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serror"
	"io"
	"os"
	"reflect"
	"testing"
)

func Test_NewSourceStrategyDir(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		strategy, err := newSourceStrategyDir(nil, &(DecoderFactory{}))
		switch {
		case strategy != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("nil decoder factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, err := newSourceStrategyDir(NewMockFs(ctrl), nil)
		switch {
		case strategy != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("new file source factory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		fs := NewMockFs(ctrl)
		factory := &(DecoderFactory{})

		strategy, err := newSourceStrategyDir(fs, factory)
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		case strategy.(*sourceStrategyDir).fs != fs:
			t.Error("didn't stored the file system adapter reference")
		case strategy.(*sourceStrategyDir).factory != factory:
			t.Error("didn't stored the decoder factory reference")
		}
	})
}

func Test_SourceStrategyDir_Accept(t *testing.T) {
	t.Run("accept only file type", func(t *testing.T) {
		scenarios := []struct {
			sourceType string
			exp        bool
		}{
			{ // _test file type
				sourceType: SourceTypeDirectory,
				exp:        true,
			},
			{ // _test non-file type
				sourceType: SourceTypeUnknown,
				exp:        false,
			},
		}

		for _, scenario := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)
				defer func() { ctrl.Finish() }()
				strategy, _ := newSourceStrategyDir(NewMockFs(ctrl), &(DecoderFactory{}))
				if check := strategy.Accept(scenario.sourceType); check != scenario.exp {
					t.Errorf("for the type (%s), returned (%v)", scenario.sourceType, check)
				}
			}
			test()
		}
	})
}

func Test_SourceStrategyDir_AcceptFromConfig(t *testing.T) {
	t.Run("don't accept on invalid config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyDir(NewMockFs(ctrl), &(DecoderFactory{}))

		if strategy.AcceptFromConfig(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyDir(NewMockFs(ctrl), &(DecoderFactory{}))

		if strategy.AcceptFromConfig(&Partial{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyDir(NewMockFs(ctrl), &(DecoderFactory{}))

		if strategy.AcceptFromConfig(&Partial{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if invalid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyDir(NewMockFs(ctrl), &(DecoderFactory{}))

		if strategy.AcceptFromConfig(&Partial{"type": SourceTypeUnknown}) {
			t.Error("returned true")
		}
	})

	t.Run("accept config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyDir(NewMockFs(ctrl), &(DecoderFactory{}))

		if !strategy.AcceptFromConfig(&Partial{"type": SourceTypeDirectory}) {
			t.Error("returned false")
		}
	})
}

func Test_SourceStrategyDir_Create(t *testing.T) {
	t.Run("missing path", func(t *testing.T) {
		strategy := &sourceStrategyDir{}

		src, err := strategy.Create()
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("missing format", func(t *testing.T) {
		strategy := &sourceStrategyDir{}

		src, err := strategy.Create("path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("missing recursive flag", func(t *testing.T) {
		strategy := &sourceStrategyDir{}

		src, err := strategy.Create("path", "format")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("non-string path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyDir(NewMockFs(ctrl), &(DecoderFactory{}))

		src, err := strategy.Create(123, "format", true)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyDir(NewMockFs(ctrl), &(DecoderFactory{}))

		src, err := strategy.Create("path", 123, true)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("non-bool recursive flag", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyDir(NewMockFs(ctrl), &(DecoderFactory{}))

		src, err := strategy.Create("path", "format", "true")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("create the dir source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := Partial{"field": "value"}
		fileinfoname := "file.yaml"
		fileinfo := NewMockFileInfo(ctrl)
		fileinfo.EXPECT().IsDir().Return(false).Times(1)
		fileinfo.EXPECT().Name().Return(fileinfoname).Times(1)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileinfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		file := NewMockFile(ctrl)
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, "field: value")
			return 12, io.EOF
		}).Times(1)
		file.EXPECT().Close().Return(nil).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(dir, nil).Times(1)
		fs.EXPECT().OpenFile(path+"/"+fileinfoname, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		factory := DecoderFactory{}
		_ = factory.Register(&decoderStrategyYAML{})

		strategy, _ := newSourceStrategyDir(fs, &factory)

		src, err := strategy.Create(path, DecoderFormatYAML, true)
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceDir:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new file src")
			}
		}
	})
}

func Test_SourceStrategyDir_CreateFromConfig(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyDir(NewMockFs(ctrl), &(DecoderFactory{}))

		src, err := strategy.CreateFromConfig(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("missing path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyDir(NewMockFs(ctrl), &(DecoderFactory{}))

		src, err := strategy.CreateFromConfig(&Partial{"format": "format", "recursive": true})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConfigPathNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConfigPathNotFound)
		}
	})

	t.Run("non-string path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyDir(NewMockFs(ctrl), &(DecoderFactory{}))

		src, err := strategy.CreateFromConfig(&Partial{"path": 123, "format": "format", "recursive": true})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyDir(NewMockFs(ctrl), &(DecoderFactory{}))

		src, err := strategy.CreateFromConfig(&Partial{"path": "path", "format": 123, "recursive": true})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("non-bool recursive flag", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyDir(NewMockFs(ctrl), &(DecoderFactory{}))

		src, err := strategy.CreateFromConfig(&Partial{"path": "path", "format": 123, "recursive": "true"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("non-bool recursive flag", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := newSourceStrategyDir(NewMockFs(ctrl), &(DecoderFactory{}))

		src, err := strategy.CreateFromConfig(&Partial{"path": "path", "format": "format", "recursive": "true"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("create the dir source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := Partial{"field": "value"}
		fileinfoname := "file.yaml"
		fileinfo := NewMockFileInfo(ctrl)
		fileinfo.EXPECT().IsDir().Return(false).Times(1)
		fileinfo.EXPECT().Name().Return(fileinfoname).Times(1)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileinfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		file := NewMockFile(ctrl)
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, "field: value")
			return 12, io.EOF
		}).Times(1)
		file.EXPECT().Close().Return(nil).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(dir, nil).Times(1)
		fs.EXPECT().OpenFile(path+"/"+fileinfoname, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		factory := DecoderFactory{}
		_ = factory.Register(&decoderStrategyYAML{})

		strategy, _ := newSourceStrategyDir(fs, &factory)

		src, err := strategy.CreateFromConfig(&Partial{"path": path, "format": DecoderFormatYAML, "recursive": true})
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceDir:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new file src")
			}
		}
	})

	t.Run("create the dir source defaulting format if not present in config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := Partial{"field": "value"}
		fileinfoname := "file.yaml"
		fileinfo := NewMockFileInfo(ctrl)
		fileinfo.EXPECT().IsDir().Return(false).Times(1)
		fileinfo.EXPECT().Name().Return(fileinfoname).Times(1)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileinfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		file := NewMockFile(ctrl)
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, "field: value")
			return 12, io.EOF
		}).Times(1)
		file.EXPECT().Close().Return(nil).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(dir, nil).Times(1)
		fs.EXPECT().OpenFile(path+"/"+fileinfoname, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		factory := DecoderFactory{}
		_ = factory.Register(&decoderStrategyYAML{})

		strategy, _ := newSourceStrategyDir(fs, &factory)

		src, err := strategy.CreateFromConfig(&Partial{"path": path, "recursive": true})
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceDir:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new file src")
			}
		}
	})
}
