package config

import (
	"errors"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/err"
)

func Test_NewObservableFileSourceStrategy(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewObservableFileSourceStrategy(nil, NewMockDecoderFactory(ctrl))
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("nil decoder factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewObservableFileSourceStrategy(NewMockFs(ctrl), nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("new file source factory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		fs := NewMockFs(ctrl)
		decoderFactory := NewMockDecoderFactory(ctrl)

		sut, e := NewObservableFileSourceStrategy(fs, decoderFactory)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case sut.fs != fs:
			t.Error("didn't stored the file system adapter reference")
		case sut.decoderFactory != decoderFactory:
			t.Error("didn't stored the decoder factory reference")
		}
	})
}

func Test_ObservableFileSourceStrategy_Accept(t *testing.T) {
	t.Run("don't accept on invalid config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableFileSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.Accept(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableFileSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.Accept(&Config{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableFileSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.Accept(&Config{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if invalid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableFileSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.Accept(&Config{"type": SourceUnknown}) {
			t.Error("returned true")
		}
	})

	t.Run("accept config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableFileSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if !sut.Accept(&Config{"type": SourceObservableFile}) {
			t.Error("returned false")
		}
	})
}

func Test_ObservableFileSourceStrategy_Create(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableFileSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("missing path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableFileSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&Config{"format": "format"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ConfigPathNotFound):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.ConfigPathNotFound)
		}
	})

	t.Run("non-string path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableFileSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&Config{"path": 123, "format": "format"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewObservableFileSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&Config{"path": "path", "format": 123})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("create the observable file source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		field := "field"
		value := "value"
		expected := Config{field: value}
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
		decoderFactory.EXPECT().Create(FormatYAML, file).Return(decoder, nil).Times(1)

		sut, _ := NewObservableFileSourceStrategy(fs, decoderFactory)

		src, e := sut.Create(&Config{"path": path, "format": FormatYAML})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *ObservableFileSource:
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
		expected := Config{field: value}
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
		decoderFactory.EXPECT().Create(FormatYAML, file).Return(decoder, nil).Times(1)

		sut, _ := NewObservableFileSourceStrategy(fs, decoderFactory)

		src, e := sut.Create(&Config{"path": path})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *ObservableFileSource:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new file src")
			}
		}
	})
}
