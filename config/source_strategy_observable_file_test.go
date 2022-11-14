package config

import (
	"errors"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	serror "github.com/happyhippyhippo/slate/error"
)

func Test_NewSourceStrategyFileObservable(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := newSourceStrategyObservableFile(nil, NewMockDecoderFactory(ctrl))
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("nil decoder decoderFactory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := newSourceStrategyObservableFile(NewMockFs(ctrl), nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("new file source decoderFactory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		fs := NewMockFs(ctrl)
		decoderFactory := NewMockDecoderFactory(ctrl)

		sut, e := newSourceStrategyObservableFile(fs, decoderFactory)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case sut.(*sourceStrategyObservableFile).fs != fs:
			t.Error("didn't stored the file system adapter reference")
		case sut.(*sourceStrategyObservableFile).decoderFactory != decoderFactory:
			t.Error("didn't stored the decoder decoderFactory reference")
		}
	})
}

func Test_SourceStrategyFileObservable_Accept(t *testing.T) {
	t.Run("accept only file type", func(t *testing.T) {
		scenarios := []struct {
			sourceType string
			exp        bool
		}{
			{ // _test observable file type
				sourceType: SourceTypeObservableFile,
				exp:        true,
			},
			{ // _test non-observable file type
				sourceType: SourceTypeUnknown,
				exp:        false,
			},
		}

		for _, scenario := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)
				defer func() { ctrl.Finish() }()

				sut, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))
				if check := sut.Accept(scenario.sourceType); check != scenario.exp {
					t.Errorf("for the type (%s), returned (%v)", scenario.sourceType, check)
				}
			}
			test()
		}
	})
}

func Test_SourceStrategyFileObservable_AcceptFromConfig(t *testing.T) {
	t.Run("don't accept on invalid config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.AcceptFromConfig(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.AcceptFromConfig(&Partial{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.AcceptFromConfig(&Partial{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if invalid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.AcceptFromConfig(&Partial{"type": SourceTypeUnknown}) {
			t.Error("returned true")
		}
	})

	t.Run("accept config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if !sut.AcceptFromConfig(&Partial{"type": SourceTypeObservableFile}) {
			t.Error("returned false")
		}
	})
}

func Test_SourceStrategyFileObservable_Create(t *testing.T) {
	t.Run("missing path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create()
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("missing format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create("path")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("non-string path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create(123, "format")
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create("path", 123)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("create the observable file source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		field := "field"
		value := "value"
		expected := Partial{field: value}
		file := NewMockFile(ctrl)
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(1)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&expected, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderFactory := NewMockDecoderFactory(ctrl)
		decoderFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, _ := newSourceStrategyObservableFile(fs, decoderFactory)

		src, e := sut.Create(path, DecoderFormatYAML)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceObservableFile:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new file src")
			}
		}
	})
}

func Test_SourceStrategyFileObservable_CreateFromConfig(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.CreateFromConfig(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("missing path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.CreateFromConfig(&Partial{"format": "format"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrConfigPathNotFound):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConfigPathNotFound)
		}
	})

	t.Run("non-string path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.CreateFromConfig(&Partial{"path": 123, "format": "format"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newSourceStrategyObservableFile(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.CreateFromConfig(&Partial{"path": "path", "format": 123})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("create the observable file source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		field := "field"
		value := "value"
		expected := Partial{field: value}
		file := NewMockFile(ctrl)
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(1)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&expected, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderFactory := NewMockDecoderFactory(ctrl)
		decoderFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, _ := newSourceStrategyObservableFile(fs, decoderFactory)

		src, e := sut.CreateFromConfig(&Partial{"path": path, "format": DecoderFormatYAML})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceObservableFile:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new file src")
			}
		}
	})

	t.Run("create the observable file source defaulting format if not present in config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		field := "field"
		value := "value"
		expected := Partial{field: value}
		file := NewMockFile(ctrl)
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(1)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&expected, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderFactory := NewMockDecoderFactory(ctrl)
		decoderFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, _ := newSourceStrategyObservableFile(fs, decoderFactory)

		src, e := sut.CreateFromConfig(&Partial{"path": path})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceObservableFile:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new file src")
			}
		}
	})
}
