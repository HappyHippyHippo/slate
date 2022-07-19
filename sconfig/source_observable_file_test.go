package sconfig

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serr"
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_NewSourceObservableFile(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := newSourceObservableFile("path", DecoderFormatYAML, nil, NewMockDecoderFactory(ctrl))
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

		sut, e := newSourceObservableFile("path", DecoderFormatYAML, NewMockFs(ctrl), nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("error that may be raised when retrieving the file info", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := fmt.Errorf("error message")
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(nil, expected).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)

		sut, e := newSourceObservableFile(path, DecoderFormatYAML, fs, dFactory)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("error that may be raised when opening the file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := fmt.Errorf("error message")
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(1)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(nil, expected).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)

		sut, e := newSourceObservableFile(path, DecoderFormatYAML, fs, dFactory)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("error that may be raised when creating the decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		path := "path"
		file := NewMockFile(ctrl)
		file.EXPECT().Close().Times(1)
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(1)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatUnknown, file).Return(nil, expected).Times(1)

		sut, e := newSourceObservableFile(path, DecoderFormatUnknown, fs, dFactory)
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

		expected := fmt.Errorf("error message")
		path := "path"
		file := NewMockFile(ctrl)
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(1)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(nil, expected).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, e := newSourceObservableFile(path, DecoderFormatYAML, fs, dFactory)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("create the config observable file source", func(t *testing.T) {
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
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, e := newSourceObservableFile(path, DecoderFormatYAML, fs, dFactory)
		switch {
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch s := sut.(type) {
			case *sourceObservableFile:
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
				case !reflect.DeepEqual(s.partial, expected):
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new observable file source")
			}
		}
	})
}

func Test_SourceObservableFile_Reload(t *testing.T) {
	t.Run("error if fail to retrieving the file info", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		field := "field"
		value := "value"
		expected := fmt.Errorf("error message")
		file := NewMockFile(ctrl)
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
		fs := NewMockFs(ctrl)
		gomock.InOrder(
			fs.EXPECT().Stat(path).Return(fileInfo, nil),
			fs.EXPECT().Stat(path).Return(nil, expected),
		)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&Partial{field: value}, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, _ := newSourceObservableFile(path, DecoderFormatYAML, fs, dFactory)

		reloaded, e := sut.Reload()
		switch {
		case reloaded:
			t.Error("flagged that was reloaded")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("error if fails to load the file content", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		field := "field"
		value := "value"
		expected := fmt.Errorf("error message")
		file := NewMockFile(ctrl)
		fileInfo := NewMockFileInfo(ctrl)
		gomock.InOrder(
			fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)),
			fileInfo.EXPECT().ModTime().Return(time.Unix(0, 2)),
		)
		fs := NewMockFs(ctrl)
		gomock.InOrder(
			fs.EXPECT().Stat(path).Return(fileInfo, nil),
			fs.EXPECT().Stat(path).Return(fileInfo, nil),
		)
		gomock.InOrder(
			fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil),
			fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(nil, expected),
		)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&Partial{field: value}, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, _ := newSourceObservableFile(path, DecoderFormatYAML, fs, dFactory)

		reloaded, e := sut.Reload()
		switch {
		case reloaded:
			t.Error("flagged that was reloaded")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("prevent reload of a unchanged source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		field := "field"
		value := "value"
		file := NewMockFile(ctrl)
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(2)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(2)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&Partial{field: value}, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, _ := newSourceObservableFile(path, DecoderFormatYAML, fs, dFactory)

		if reloaded, e := sut.Reload(); reloaded {
			t.Error("flagged that was reloaded")
		} else if e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("should reload a changed source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		field := "field"
		value1 := "value1"
		value2 := "value2"
		expected := Partial{field: value2}
		file1 := NewMockFile(ctrl)
		file2 := NewMockFile(ctrl)
		fileInfo := NewMockFileInfo(ctrl)
		gomock.InOrder(
			fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)),
			fileInfo.EXPECT().ModTime().Return(time.Unix(0, 2)),
		)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(2)
		gomock.InOrder(
			fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file1, nil),
			fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file2, nil),
		)
		decoder1 := NewMockDecoder(ctrl)
		decoder1.EXPECT().Decode().Return(&Partial{field: value1}, nil).Times(1)
		decoder1.EXPECT().Close().Return(nil).Times(1)
		decoder2 := NewMockDecoder(ctrl)
		decoder2.EXPECT().Decode().Return(&Partial{field: value2}, nil).Times(1)
		decoder2.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		gomock.InOrder(
			dFactory.EXPECT().Create(DecoderFormatYAML, file1).Return(decoder1, nil).Times(1),
			dFactory.EXPECT().Create(DecoderFormatYAML, file2).Return(decoder2, nil).Times(1),
		)

		sut, _ := newSourceObservableFile(path, DecoderFormatYAML, fs, dFactory)

		reloaded, e := sut.Reload()
		switch {
		case !reloaded:
			t.Error("flagged that was not reloaded")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case !reflect.DeepEqual(expected, sut.(*sourceObservableFile).partial):
			t.Error("didn't stored the check configuration")
		}
	})
}
