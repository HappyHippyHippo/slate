package sconfig

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serr"
	"os"
	"reflect"
	"testing"
)

func Test_NewSourceFile(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := newSourceFile("path", DecoderFormatYAML, nil, NewMockDecoderFactory(ctrl))
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

		sut, e := newSourceFile("path", DecoderFormatYAML, NewMockFs(ctrl), nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("error that may be raised when opening the file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := fmt.Errorf("error message")
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(nil, expected).Times(1)

		sut, e := newSourceFile(path, DecoderFormatYAML, fs, NewMockDecoderFactory(ctrl))
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
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatUnknown, file).Return(nil, expected).Times(1)

		sut, e := newSourceFile(path, DecoderFormatUnknown, fs, dFactory)
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
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, e := newSourceFile(path, DecoderFormatYAML, fs, dFactory)
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
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, e := newSourceFile(path, DecoderFormatYAML, fs, dFactory)
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
				case s.dFactory != dFactory:
					t.Error("didn't stored the decoder dFactory reference")
				case !reflect.DeepEqual(s.partial, partial):
					t.Error("didn't stored the decoder information")
				}
			default:
				t.Error("didn't returned a new file source")
			}
		}
	})
}
