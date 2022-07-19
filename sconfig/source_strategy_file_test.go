package sconfig

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serr"
	"os"
	"reflect"
	"testing"
)

func Test_NewSourceStrategyFile(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := newSourceStrategyFile(nil, NewMockDecoderFactory(ctrl))
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

		sut, e := newSourceStrategyFile(NewMockFs(ctrl), nil)
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

		sut, e := newSourceStrategyFile(fs, dFactory)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case sut.(*sourceStrategyFile).fs != fs:
			t.Error("didn't stored the file system adapter reference")
		case sut.(*sourceStrategyFile).dFactory != dFactory:
			t.Error("didn't stored the decoder dFactory reference")
		}
	})
}

func Test_SourceStrategyFile_Accept(t *testing.T) {
	t.Run("accept only file type", func(t *testing.T) {
		scenarios := []struct {
			sourceType string
			exp        bool
		}{
			{ // _test file type
				sourceType: SourceTypeFile,
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

				sut, _ := newSourceStrategyFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))
				if check := sut.Accept(scenario.sourceType); check != scenario.exp {
					t.Errorf("for the type (%s), returned (%v)", scenario.sourceType, check)
				}
			}
			test()
		}
	})
}

func Test_SourceStrategyFile_AcceptFromConfig(t *testing.T) {
	t.Run("don't accept on invalid config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.AcceptFromConfig(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.AcceptFromConfig(&Partial{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.AcceptFromConfig(&Partial{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if invalid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.AcceptFromConfig(&Partial{"type": SourceTypeUnknown}) {
			t.Error("returned true")
		}
	})

	t.Run("accept config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if !sut.AcceptFromConfig(&Partial{"type": SourceTypeFile}) {
			t.Error("returned false")
		}
	})
}

func Test_SourceStrategyFile_Create(t *testing.T) {
	t.Run("missing path", func(t *testing.T) {
		src, e := (&sourceStrategyFile{}).Create()
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
		src, e := (&sourceStrategyFile{}).Create("path")
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

		sut, _ := newSourceStrategyFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create(123, "format")
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

		sut, _ := newSourceStrategyFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create("path", 123)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("create the file source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		field := "field"
		value := "value"
		expected := Partial{field: value}
		file := NewMockFile(ctrl)
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&expected, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, _ := newSourceStrategyFile(fs, dFactory)

		src, e := sut.Create(path, DecoderFormatYAML)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceFile:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new file src")
			}
		}
	})
}

func Test_SourceStrategyFile_CreateFromConfig(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

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

		sut, _ := newSourceStrategyFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.CreateFromConfig(&Partial{"format": "format"})
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

		sut, _ := newSourceStrategyFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.CreateFromConfig(&Partial{"path": 123, "format": "format"})
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

		sut, _ := newSourceStrategyFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.CreateFromConfig(&Partial{"path": "path", "format": 123})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("create the file source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		field := "field"
		value := "value"
		expected := Partial{field: value}
		file := NewMockFile(ctrl)
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&expected, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, _ := newSourceStrategyFile(fs, dFactory)

		src, e := sut.CreateFromConfig(&Partial{"path": path, "format": DecoderFormatYAML})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceFile:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new file src")
			}
		}
	})

	t.Run("create the file source defaulting format if not present in config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		field := "field"
		value := "value"
		expected := Partial{field: value}
		file := NewMockFile(ctrl)
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&expected, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, _ := newSourceStrategyFile(fs, dFactory)

		src, e := sut.CreateFromConfig(&Partial{"path": path})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceFile:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new file src")
			}
		}
	})
}
