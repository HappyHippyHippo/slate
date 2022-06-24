package sconfig

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serror"
	"io"
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_NewSourceObservableFile(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		src, err := NewSourceObservableFile("path", DecoderFormatYAML, nil, &DecoderFactory{})
		switch {
		case src != nil:
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

		src, err := NewSourceObservableFile("path", DecoderFormatYAML, NewMockFs(ctrl), nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("error that may be raised when retrieving the file info", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := fmt.Errorf("error message")
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(nil, expected).Times(1)
		factory := DecoderFactory{}

		src, err := NewSourceObservableFile(path, DecoderFormatYAML, fs, &factory)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error", err)
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
		factory := DecoderFactory{}

		src, err := NewSourceObservableFile(path, DecoderFormatYAML, fs, &factory)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error", err)
		}
	})

	t.Run("error that may be raised when creating the decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		file := NewMockFile(ctrl)
		file.EXPECT().Close().Times(1)
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(1)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		factory := DecoderFactory{}

		src, err := NewSourceObservableFile(path, DecoderFormatUnknown, fs, &factory)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrInvalidConfigDecoderFormat):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrInvalidConfigDecoderFormat)
		}
	})

	t.Run("error that may be raised when running the decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		file := NewMockFile(ctrl)
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, fmt.Sprintf("{"))
			return 1, io.EOF
		}).Times(1)
		file.EXPECT().Close().Times(1)
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(1)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		factory := DecoderFactory{}
		_ = factory.Register(&decoderStrategyYAML{})

		src, err := NewSourceObservableFile(path, DecoderFormatYAML, fs, &factory)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != "yaml: line 1: did not find expected node content":
			t.Errorf("returned the (%v) error", err)
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
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, fmt.Sprintf("%s: %s", field, value))
			return 12, io.EOF
		}).Times(1)
		file.EXPECT().Close().Times(1)
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(1)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		factory := DecoderFactory{}
		_ = factory.Register(&decoderStrategyYAML{})

		src, err := NewSourceObservableFile(path, DecoderFormatYAML, fs, &factory)
		switch {
		case src == nil:
			t.Error("didn't returned a valid reference")
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		default:
			switch s := src.(type) {
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
				case s.factory != &factory:
					t.Error("didn't stored the decoder factory reference")
				case !reflect.DeepEqual(s.partial, expected):
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new observable file source")
			}
		}
	})

	t.Run("store the decoded Partial", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		field := "field"
		value := "value"
		expected := Partial{field: value}
		file := NewMockFile(ctrl)
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, fmt.Sprintf("%s: %s", field, value))
			return 12, io.EOF
		}).Times(1)
		file.EXPECT().Close().Times(1)
		fileI := NewMockFileInfo(ctrl)
		fileI.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(fileI, nil).Times(1)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		factory := DecoderFactory{}
		_ = factory.Register(&decoderStrategyYAML{})

		src, err := NewSourceObservableFile(path, DecoderFormatYAML, fs, &factory)
		switch {
		case src == nil:
			t.Error("didn't returned a valid reference")
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		default:
			switch s := src.(type) {
			case *sourceObservableFile:
				if !reflect.DeepEqual(s.partial, expected) {
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
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, fmt.Sprintf("%s: %s", field, value))
			return 12, io.EOF
		}).Times(1)
		file.EXPECT().Close().Times(1)
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
		fs := NewMockFs(ctrl)
		gomock.InOrder(
			fs.EXPECT().Stat(path).Return(fileInfo, nil),
			fs.EXPECT().Stat(path).Return(nil, expected),
		)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		factory := DecoderFactory{}
		_ = factory.Register(&decoderStrategyYAML{})
		src, _ := NewSourceObservableFile(path, DecoderFormatYAML, fs, &factory)

		reloaded, err := src.Reload()
		switch {
		case reloaded:
			t.Error("flagged that was reloaded")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error", err)
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
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, fmt.Sprintf("%s: %s", field, value))
			return 12, io.EOF
		}).Times(1)
		file.EXPECT().Close().Times(1)
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
		factory := DecoderFactory{}
		_ = factory.Register(&decoderStrategyYAML{})
		src, _ := NewSourceObservableFile(path, DecoderFormatYAML, fs, &factory)

		reloaded, err := src.Reload()
		switch {
		case reloaded:
			t.Error("flagged that was reloaded")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error", err)
		}
	})

	t.Run("prevent reload of a unchanged source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		field := "field"
		value := "value"
		file := NewMockFile(ctrl)
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, fmt.Sprintf("%s: %s", field, value))
			return 12, io.EOF
		}).Times(1)
		file.EXPECT().Close().Times(1)
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(2)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(2)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		factory := DecoderFactory{}
		_ = factory.Register(&decoderStrategyYAML{})
		src, _ := NewSourceObservableFile(path, DecoderFormatYAML, fs, &factory)

		if reloaded, err := src.Reload(); reloaded {
			t.Error("flagged that was reloaded")
		} else if err != nil {
			t.Errorf("returned the (%v) error", err)
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
		file1.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, fmt.Sprintf("%src: %src", field, value1))
			return 13, io.EOF
		})
		file1.EXPECT().Close().Times(1)
		file2 := NewMockFile(ctrl)
		file2.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, fmt.Sprintf("%s: %s", field, value2))
			return 13, io.EOF
		})
		file2.EXPECT().Close().Times(1)
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
		factory := DecoderFactory{}
		_ = factory.Register(&decoderStrategyYAML{})
		src, _ := NewSourceObservableFile(path, DecoderFormatYAML, fs, &factory)

		reloaded, err := src.Reload()
		switch {
		case !reloaded:
			t.Error("flagged that was not reloaded")
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case !reflect.DeepEqual(expected, src.(*sourceObservableFile).partial):
			t.Error("didn't stored the check configuration")
		}
	})
}
