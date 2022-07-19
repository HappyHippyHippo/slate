package sconfig

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serr"
	"os"
	"reflect"
	"testing"
)

func Test_NewSourceStrategyDir(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := newSourceStrategyDir(nil, NewMockDecoderFactory(ctrl))
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("nil decoder dFactory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := newSourceStrategyDir(NewMockFs(ctrl), nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("new file source dFactory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		fs := NewMockFs(ctrl)
		dFactory := NewMockDecoderFactory(ctrl)

		sut, e := newSourceStrategyDir(fs, dFactory)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case sut.(*sourceStrategyDir).fs != fs:
			t.Error("didn't stored the file system adapter reference")
		case sut.(*sourceStrategyDir).dFactory != dFactory:
			t.Error("didn't stored the decoder dFactory reference")
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

				sut, _ := newSourceStrategyDir(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))
				if check := sut.Accept(scenario.sourceType); check != scenario.exp {
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

		sut, _ := newSourceStrategyDir(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.AcceptFromConfig(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyDir(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.AcceptFromConfig(&Partial{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyDir(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.AcceptFromConfig(&Partial{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if invalid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyDir(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.AcceptFromConfig(&Partial{"type": SourceTypeUnknown}) {
			t.Error("returned true")
		}
	})

	t.Run("accept config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyDir(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if !sut.AcceptFromConfig(&Partial{"type": SourceTypeDirectory}) {
			t.Error("returned false")
		}
	})
}

func Test_SourceStrategyDir_Create(t *testing.T) {
	t.Run("missing path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyDir(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create()
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("missing format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyDir(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create("path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("missing recursive flag", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyDir(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create("path", "format")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("non-string path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyDir(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create(123, "format", true)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyDir(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create("path", 123, true)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("non-bool recursive flag", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyDir(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create("path", "format", "true")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("create the dir source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := Partial{"field": "value"}
		fileInfoName := "file.yaml"
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().IsDir().Return(false).Times(1)
		fileInfo.EXPECT().Name().Return(fileInfoName).Times(1)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileInfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		file := NewMockFile(ctrl)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(dir, nil).Times(1)
		fs.EXPECT().OpenFile(path+"/"+fileInfoName, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&expected, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, _ := newSourceStrategyDir(fs, dFactory)

		src, e := sut.Create(path, DecoderFormatYAML, true)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
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

		sut, _ := newSourceStrategyDir(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.CreateFromConfig(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("missing path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyDir(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.CreateFromConfig(&Partial{"format": "format", "recursive": true})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrConfigPathNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConfigPathNotFound)
		}
	})

	t.Run("non-string path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyDir(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.CreateFromConfig(&Partial{"path": 123, "format": "format", "recursive": true})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyDir(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.CreateFromConfig(&Partial{"path": "path", "format": 123, "recursive": true})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("non-bool recursive flag", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyDir(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.CreateFromConfig(&Partial{"path": "path", "format": 123, "recursive": "true"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("non-bool recursive flag", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyDir(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.CreateFromConfig(&Partial{"path": "path", "format": "format", "recursive": "true"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("create the dir source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := Partial{"field": "value"}
		fileInfoName := "file.yaml"
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().IsDir().Return(false).Times(1)
		fileInfo.EXPECT().Name().Return(fileInfoName).Times(1)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileInfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		file := NewMockFile(ctrl)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(dir, nil).Times(1)
		fs.EXPECT().OpenFile(path+"/"+fileInfoName, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&expected, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, _ := newSourceStrategyDir(fs, dFactory)

		src, e := sut.CreateFromConfig(&Partial{"path": path, "format": DecoderFormatYAML, "recursive": true})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
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
		fileInfoName := "file.yaml"
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().IsDir().Return(false).Times(1)
		fileInfo.EXPECT().Name().Return(fileInfoName).Times(1)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileInfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		file := NewMockFile(ctrl)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(dir, nil).Times(1)
		fs.EXPECT().OpenFile(path+"/"+fileInfoName, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&expected, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, _ := newSourceStrategyDir(fs, dFactory)

		src, e := sut.CreateFromConfig(&Partial{"path": path, "recursive": true})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
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
