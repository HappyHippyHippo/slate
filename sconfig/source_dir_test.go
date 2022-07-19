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

func Test_NewSourceDir(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := newSourceDir("path", DecoderFormatYAML, true, nil, NewMockDecoderFactory(ctrl))
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

		sut, e := newSourceDir("path", DecoderFormatYAML, true, NewMockFs(ctrl), nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("error that may be raised when opening the dir", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := fmt.Errorf("error message")
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(nil, expected).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)

		sut, e := newSourceDir(path, DecoderFormatYAML, true, fs, dFactory)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
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
		dFactory := NewMockDecoderFactory(ctrl)

		sut, e := newSourceDir(path, DecoderFormatYAML, true, fs, dFactory)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
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
		dFactory := NewMockDecoderFactory(ctrl)

		sut, e := newSourceDir(path, DecoderFormatYAML, true, fs, dFactory)
		switch {
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch s := sut.(type) {
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
				case s.dFactory != dFactory:
					t.Error("didn't stored the decoder dFactory reference")
				}
			default:
				t.Error("didn't returned a new file source")
			}
		}
	})

	t.Run("error opening the sconfig file", func(t *testing.T) {
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
		dFactory := NewMockDecoderFactory(ctrl)

		sut, e := newSourceDir(path, DecoderFormatUnknown, true, fs, dFactory)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("error retrieving the proper decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
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
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatUnknown, file).Return(nil, expected).Times(1)

		sut, e := newSourceDir(path, DecoderFormatUnknown, true, fs, dFactory)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("error decoding file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		msg := "serr"
		expected := fmt.Errorf("yaml: input serr: %s", msg)
		fileinfoname := "file.yaml"
		fileinfo := NewMockFileInfo(ctrl)
		fileinfo.EXPECT().IsDir().Return(false).Times(1)
		fileinfo.EXPECT().Name().Return(fileinfoname).Times(1)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileinfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		file := NewMockFile(ctrl)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(dir, nil).Times(1)
		fs.EXPECT().OpenFile(path+"/"+fileinfoname, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(nil, expected).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, e := newSourceDir(path, DecoderFormatYAML, true, fs, dFactory)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("correctly load single file on directory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := &Partial{"field": "value"}
		path := "path"
		fileinfoname := "file.yaml"
		fileinfo := NewMockFileInfo(ctrl)
		fileinfo.EXPECT().IsDir().Return(false).Times(1)
		fileinfo.EXPECT().Name().Return(fileinfoname).Times(1)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileinfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		file := NewMockFile(ctrl)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(dir, nil).Times(1)
		fs.EXPECT().OpenFile(path+"/"+fileinfoname, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(partial, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, e := newSourceDir(path, DecoderFormatYAML, true, fs, dFactory)
		switch {
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch s := sut.(type) {
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
				case s.dFactory != dFactory:
					t.Error("didn't stored the decoder dFactory reference")
				case !reflect.DeepEqual(s.partial, *partial):
					t.Errorf("didn't loaded the content correctly having (%v) when expecting (%v)", s.partial, *partial)
				}
			default:
				t.Error("didn't returned a new file source")
			}
		}
	})

	t.Run("don't follow sub dirs if not recursive", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := &Partial{"field": "value"}
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
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(dir, nil).Times(1)
		fs.EXPECT().OpenFile(path+"/"+fileinfoname, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(partial, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, e := newSourceDir(path, DecoderFormatYAML, false, fs, dFactory)
		switch {
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch s := sut.(type) {
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
				case s.dFactory != dFactory:
					t.Error("didn't stored the decoder dFactory reference")
				case !reflect.DeepEqual(s.partial, *partial):
					t.Errorf("didn't loaded the content correctly having (%v) when expecting (%v)", s.partial, *partial)
				}
			default:
				t.Error("didn't returned a new file source")
			}
		}
	})

	t.Run("error while loading sub dir", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial1 := &Partial{"field": "value"}
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
		fs := NewMockFs(ctrl)
		gomock.InOrder(
			fs.EXPECT().Open(path).Return(dir, nil).Times(1),
			fs.EXPECT().Open(path+"/"+subdirinfoname).Return(nil, expected).Times(1))
		fs.EXPECT().OpenFile(path+"/"+fileinfoname, os.O_RDONLY, os.FileMode(0o644)).Return(file1, nil).Times(1)
		decoder1 := NewMockDecoder(ctrl)
		decoder1.EXPECT().Decode().Return(partial1, nil).Times(1)
		decoder1.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, file1).Return(decoder1, nil).Times(1)

		sut, e := newSourceDir(path, DecoderFormatYAML, true, fs, dFactory)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("follow sub dirs if recursive", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial1 := &Partial{"field1": "value"}
		partial2 := &Partial{"field2": "value"}
		expected := Partial{"field1": "value", "field2": "value"}
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
		file2 := NewMockFile(ctrl)
		fs := NewMockFs(ctrl)
		gomock.InOrder(
			fs.EXPECT().Open(path).Return(dir, nil).Times(1),
			fs.EXPECT().Open(path+"/"+subdirinfoname).Return(subdir, nil).Times(1))
		gomock.InOrder(
			fs.EXPECT().OpenFile(path+"/"+fileinfoname, os.O_RDONLY, os.FileMode(0o644)).Return(file1, nil).Times(1),
			fs.EXPECT().OpenFile(path+"/"+subdirinfoname+"/"+subfileinfoname, os.O_RDONLY, os.FileMode(0o644)).Return(file2, nil).Times(1))
		decoder1 := NewMockDecoder(ctrl)
		decoder1.EXPECT().Decode().Return(partial1, nil).Times(1)
		decoder1.EXPECT().Close().Return(nil).Times(1)
		decoder2 := NewMockDecoder(ctrl)
		decoder2.EXPECT().Decode().Return(partial2, nil).Times(1)
		decoder2.EXPECT().Close().Return(nil).Times(1)
		dFactory := NewMockDecoderFactory(ctrl)
		dFactory.EXPECT().Create(DecoderFormatYAML, file1).Return(decoder1, nil).Times(1)
		dFactory.EXPECT().Create(DecoderFormatYAML, file2).Return(decoder2, nil).Times(1)

		sut, e := newSourceDir(path, DecoderFormatYAML, true, fs, dFactory)
		switch {
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch s := sut.(type) {
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
				case s.dFactory != dFactory:
					t.Error("didn't stored the decoder dFactory reference")
				case !reflect.DeepEqual(s.partial, expected):
					t.Errorf("didn't loaded the content correctly having (%v) when expecting (%v)", s.partial, expected)
				}
			default:
				t.Error("didn't returned a new file source")
			}
		}
	})
}
