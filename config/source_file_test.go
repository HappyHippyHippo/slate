package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	serror "github.com/happyhippyhippo/slate/error"
)

func Test_NewSourceFile(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewSourceFile("path", DecoderFormatYAML, nil, NewMockDecoderFactory(ctrl))
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

		sut, e := NewSourceFile("path", DecoderFormatYAML, NewMockFs(ctrl), nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("error that may be raised when opening the file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := fmt.Errorf("error message")
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(nil, expected).Times(1)

		sut, e := NewSourceFile(path, DecoderFormatYAML, fs, NewMockDecoderFactory(ctrl))
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("error that may be raised when creating the decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		path := "path"
		file := NewMockFile(ctrl)
		file.EXPECT().Close().Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoderFactory := NewMockDecoderFactory(ctrl)
		decoderFactory.EXPECT().Create(DecoderFormatUnknown, file).Return(nil, expected).Times(1)

		sut, e := NewSourceFile(path, DecoderFormatUnknown, fs, decoderFactory)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("error that may be raised when running the decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := fmt.Errorf("error message")
		file := NewMockFile(ctrl)
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(nil, expected).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderFactory := NewMockDecoderFactory(ctrl)
		decoderFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, e := NewSourceFile(path, DecoderFormatYAML, fs, decoderFactory)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("creates the config file source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := Partial{"field": "value"}
		path := "path"
		file := NewMockFile(ctrl)
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&partial, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderFactory := NewMockDecoderFactory(ctrl)
		decoderFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, e := NewSourceFile(path, DecoderFormatYAML, fs, decoderFactory)
		switch {
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch s := sut.(type) {
			case *sourceFile:
				switch {
				case s.mutex == nil:
					t.Error("didn't created the access mutex")
				case s.path != path:
					t.Error("didn't stored the file path")
				case s.format != DecoderFormatYAML:
					t.Error("didn't stored the file content format")
				case s.fs != fs:
					t.Error("didn't stored the file system adapter reference")
				case s.decoderFactory != decoderFactory:
					t.Error("didn't stored the decoder decoderFactory reference")
				case !reflect.DeepEqual(s.partial, partial):
					t.Error("didn't stored the decoder information")
				}
			default:
				t.Error("didn't returned a new file source")
			}
		}
	})
}