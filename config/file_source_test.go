package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/err"
)

func Test_NewFileSource(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewFileSource("path", FormatYAML, nil, NewMockDecoderFactory(ctrl))
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

		sut, e := NewFileSource("path", FormatYAML, NewMockFs(ctrl), nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("error that may be raised when opening the file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := fmt.Errorf("error message")
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(nil, expected).Times(1)

		sut, e := NewFileSource(path, FormatYAML, fs, NewMockDecoderFactory(ctrl))
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) err when expecting (%v)", e, expected)
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
		decoderFactory.EXPECT().Create(FormatUnknown, file).Return(nil, expected).Times(1)

		sut, e := NewFileSource(path, FormatUnknown, fs, decoderFactory)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) err when expecting (%v)", e, expected)
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
		decoderFactory.EXPECT().Create(FormatYAML, file).Return(decoder, nil).Times(1)

		sut, e := NewFileSource(path, FormatYAML, fs, decoderFactory)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) err when expecting (%v)", e, expected)
		}
	})

	t.Run("creates the config file source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := Config{"field": "value"}
		path := "path"
		file := NewMockFile(ctrl)
		fs := NewMockFs(ctrl)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&partial, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderFactory := NewMockDecoderFactory(ctrl)
		decoderFactory.EXPECT().Create(FormatYAML, file).Return(decoder, nil).Times(1)

		sut, e := NewFileSource(path, FormatYAML, fs, decoderFactory)
		switch {
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch {
			case sut.mutex == nil:
				t.Error("didn't created the access mutex")
			case sut.path != path:
				t.Error("didn't stored the file path")
			case sut.format != FormatYAML:
				t.Error("didn't stored the file content format")
			case sut.fs != fs:
				t.Error("didn't stored the file system adapter reference")
			case sut.decoderFactory != decoderFactory:
				t.Error("didn't stored the decoder factory reference")
			case !reflect.DeepEqual(sut.partial, partial):
				t.Error("didn't stored the decoder information")
			}
		}
	})
}
