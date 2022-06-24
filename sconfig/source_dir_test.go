package sconfig

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serror"
	"io"
	"os"
	"testing"
)

func Test_NewSourceDir(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		src, err := NewSourceDir("path", DecoderFormatYAML, true, nil, &DecoderFactory{})
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

		src, err := NewSourceDir("path", DecoderFormatYAML, true, NewMockFs(ctrl), nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("error that may be raised when opening the dir", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := fmt.Errorf("error message")
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(nil, expected).Times(1)

		src, err := NewSourceDir(path, DecoderFormatYAML, true, fs, &DecoderFactory{})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("error that may be raised when reading the dir", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := fmt.Errorf("error message")
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return(nil, expected).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(dir, nil).Times(1)

		src, err := NewSourceDir(path, DecoderFormatYAML, true, fs, &DecoderFactory{})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("empty dir", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(dir, nil).Times(1)
		factory := DecoderFactory{}
		_ = factory.Register(&decoderStrategyYAML{})

		src, err := NewSourceDir(path, DecoderFormatYAML, true, fs, &factory)
		switch {
		case src == nil:
			t.Error("didn't returned a valid reference")
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		default:
			switch s := src.(type) {
			case *sourceDir:
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
				}
			default:
				t.Error("didn't returned a new file source")
			}
		}
	})

	t.Run("error opening the config file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := fmt.Errorf("error message")
		fileinfoname := "file.yaml"
		fileinfo := NewMockFileInfo(ctrl)
		fileinfo.EXPECT().IsDir().Return(false).Times(1)
		fileinfo.EXPECT().Name().Return(fileinfoname).Times(1)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileinfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(dir, nil).Times(1)
		fs.EXPECT().OpenFile(path+"/"+fileinfoname, os.O_RDONLY, os.FileMode(0o644)).Return(nil, expected).Times(1)

		src, err := NewSourceDir(path, DecoderFormatUnknown, true, fs, &DecoderFactory{})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("error retrieving the proper decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		fileinfoname := "file.yaml"
		fileinfo := NewMockFileInfo(ctrl)
		fileinfo.EXPECT().IsDir().Return(false).Times(1)
		fileinfo.EXPECT().Name().Return(fileinfoname).Times(1)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileinfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		file := NewMockFile(ctrl)
		file.EXPECT().Close().Return(nil).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(dir, nil).Times(1)
		fs.EXPECT().OpenFile(path+"/"+fileinfoname, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)

		src, err := NewSourceDir(path, DecoderFormatUnknown, true, fs, &DecoderFactory{})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrInvalidConfigDecoderFormat):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrInvalidConfigDecoderFormat)
		}
	})

	t.Run("error decoding file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		msg := "error"
		expected := fmt.Errorf("yaml: input error: %s", msg)
		fileinfoname := "file.yaml"
		fileinfo := NewMockFileInfo(ctrl)
		fileinfo.EXPECT().IsDir().Return(false).Times(1)
		fileinfo.EXPECT().Name().Return(fileinfoname).Times(1)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileinfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		file := NewMockFile(ctrl)
		file.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			return 0, fmt.Errorf("%s", msg)
		}).Times(1)
		file.EXPECT().Close().Return(nil).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(dir, nil).Times(1)
		fs.EXPECT().OpenFile(path+"/"+fileinfoname, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		factory := DecoderFactory{}
		_ = factory.Register(&decoderStrategyYAML{})

		src, err := NewSourceDir(path, DecoderFormatYAML, true, fs, &factory)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error", err)
		}
	})

	t.Run("correctly load single file on directory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
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

		src, err := NewSourceDir(path, DecoderFormatYAML, true, fs, &factory)
		switch {
		case src == nil:
			t.Error("didn't returned a valid reference")
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		default:
			switch s := src.(type) {
			case *sourceDir:
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
				}
			default:
				t.Error("didn't returned a new file source")
			}
		}
	})

	t.Run("don't follow sub dirs if not recursive", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		fileinfoname := "file.yaml"
		fileinfo := NewMockFileInfo(ctrl)
		fileinfo.EXPECT().IsDir().Return(false).Times(1)
		fileinfo.EXPECT().Name().Return(fileinfoname).Times(1)
		subdirinfo := NewMockFileInfo(ctrl)
		subdirinfo.EXPECT().IsDir().Return(true).Times(1)
		subdirinfo.EXPECT().Name().Times(0)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileinfo, subdirinfo}, nil).Times(1)
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

		src, err := NewSourceDir(path, DecoderFormatYAML, false, fs, &factory)
		switch {
		case src == nil:
			t.Error("didn't returned a valid reference")
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		default:
			switch s := src.(type) {
			case *sourceDir:
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
				}
			default:
				t.Error("didn't returned a new file source")
			}
		}
	})

	t.Run("error while loading sub dir", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := fmt.Errorf("error message")
		fileinfoname := "file1.yaml"
		fileinfo := NewMockFileInfo(ctrl)
		fileinfo.EXPECT().IsDir().Return(false).Times(1)
		fileinfo.EXPECT().Name().Return(fileinfoname).Times(1)
		subdirinfoname := "subdir"
		subdirinfo := NewMockFileInfo(ctrl)
		subdirinfo.EXPECT().IsDir().Return(true).Times(1)
		subdirinfo.EXPECT().Name().Return(subdirinfoname).Times(1)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileinfo, subdirinfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		file1 := NewMockFile(ctrl)
		file1.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, "field1: value")
			return 13, io.EOF
		}).Times(1)
		file1.EXPECT().Close().Return(nil).Times(1)
		fs := NewMockFs(ctrl)
		gomock.InOrder(
			fs.EXPECT().Open(path).Return(dir, nil).Times(1),
			fs.EXPECT().Open(path+"/"+subdirinfoname).Return(nil, expected).Times(1))
		fs.EXPECT().OpenFile(path+"/"+fileinfoname, os.O_RDONLY, os.FileMode(0o644)).Return(file1, nil).Times(1)
		factory := DecoderFactory{}
		_ = factory.Register(&decoderStrategyYAML{})

		src, err := NewSourceDir(path, DecoderFormatYAML, true, fs, &factory)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error", err)
		}
	})

	t.Run("follow sub dirs if recursive", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		fileinfoname := "file1.yaml"
		fileinfo := NewMockFileInfo(ctrl)
		fileinfo.EXPECT().IsDir().Return(false).Times(1)
		fileinfo.EXPECT().Name().Return(fileinfoname).Times(1)
		subfileinfoname := "file2.yaml"
		subfileinfo := NewMockFileInfo(ctrl)
		subfileinfo.EXPECT().IsDir().Return(false).Times(1)
		subfileinfo.EXPECT().Name().Return(subfileinfoname).Times(1)
		subdirinfoname := "subdir"
		subdirinfo := NewMockFileInfo(ctrl)
		subdirinfo.EXPECT().IsDir().Return(true).Times(1)
		subdirinfo.EXPECT().Name().Return(subdirinfoname).Times(1)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileinfo, subdirinfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		subdir := NewMockFile(ctrl)
		subdir.EXPECT().Readdir(0).Return([]os.FileInfo{subfileinfo}, nil).Times(1)
		subdir.EXPECT().Close().Return(nil).Times(1)
		file1 := NewMockFile(ctrl)
		file1.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, "field1: value")
			return 13, io.EOF
		}).Times(1)
		file1.EXPECT().Close().Return(nil).Times(1)
		file2 := NewMockFile(ctrl)
		file2.EXPECT().Read(gomock.Any()).DoAndReturn(func(buf []byte) (int, error) {
			copy(buf, "field2: value")
			return 13, io.EOF
		}).Times(1)
		file2.EXPECT().Close().Return(nil).Times(1)
		fs := NewMockFs(ctrl)
		gomock.InOrder(
			fs.EXPECT().Open(path).Return(dir, nil).Times(1),
			fs.EXPECT().Open(path+"/"+subdirinfoname).Return(subdir, nil).Times(1))
		gomock.InOrder(
			fs.EXPECT().OpenFile(path+"/"+fileinfoname, os.O_RDONLY, os.FileMode(0o644)).Return(file1, nil).Times(1),
			fs.EXPECT().OpenFile(path+"/"+subdirinfoname+"/"+subfileinfoname, os.O_RDONLY, os.FileMode(0o644)).Return(file2, nil).Times(1))
		factory := DecoderFactory{}
		_ = factory.Register(&decoderStrategyYAML{})

		src, err := NewSourceDir(path, DecoderFormatYAML, true, fs, &factory)
		switch {
		case src == nil:
			t.Error("didn't returned a valid reference")
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		default:
			switch s := src.(type) {
			case *sourceDir:
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
				}
			default:
				t.Error("didn't returned a new file source")
			}
		}
	})
}
