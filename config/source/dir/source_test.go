package dir

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
)

func Test_NewSource(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewSource("path", "format", true, nil, config.NewDecoderFactory())
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil decoder factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewSource("path", "format", true, NewMockFs(ctrl), nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error that may be raised when opening the dir", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := fmt.Errorf("error message")
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(nil, expected).Times(1)
		decoderFactory := config.NewDecoderFactory()

		sut, e := NewSource(path, "format", true, fs, decoderFactory)
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
		decoderFactory := config.NewDecoderFactory()

		sut, e := NewSource(path, "format", true, fs, decoderFactory)
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
		decoderFactory := config.NewDecoderFactory()

		sut, e := NewSource(path, "format", true, fs, decoderFactory)
		switch {
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch {
			case sut.Mutex == nil:
				t.Error("didn't created the access mutex")
			case sut.path != path:
				t.Error("didn't stored the file path")
			case sut.format != "format":
				t.Error("didn't stored the file content format")
			case sut.fs != fs:
				t.Error("didn't stored the file system adapter reference")
			case sut.decoderCreator != decoderFactory:
				t.Error("didn't stored the decoder factory reference")
			}
		}
	})

	t.Run("error opening the config file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := fmt.Errorf("error message")
		fileInfoName := "file.yaml"
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().IsDir().Return(false).Times(1)
		fileInfo.EXPECT().Name().Return(fileInfoName).Times(1)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileInfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(dir, nil).Times(1)
		fs.EXPECT().OpenFile(path+"/"+fileInfoName, os.O_RDONLY, os.FileMode(0o644)).Return(nil, expected).Times(1)
		decoderFactory := config.NewDecoderFactory()

		sut, e := NewSource(path, config.UnknownDecoder, true, fs, decoderFactory)
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

		path := "path"
		fileInfoName := "file.yaml"
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().IsDir().Return(false).Times(1)
		fileInfo.EXPECT().Name().Return(fileInfoName).Times(1)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileInfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		file := NewMockFile(ctrl)
		file.EXPECT().Close().Return(nil).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(dir, nil).Times(1)
		fs.EXPECT().OpenFile(path+"/"+fileInfoName, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept(config.UnknownDecoder).Return(false).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, e := NewSource(path, config.UnknownDecoder, true, fs, decoderFactory)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case errors.Is(e, config.ErrInvalidSource):
			t.Errorf("returned the (%v) error when expecting (%v)", e, config.ErrInvalidSource)
		}
	})

	t.Run("error decoding file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		msg := "error message"
		expected := fmt.Errorf("yaml: input error: %s", msg)
		fileInfoName := "file.yaml"
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().IsDir().Return(false).Times(1)
		fileInfo.EXPECT().Name().Return(fileInfoName).Times(1)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileInfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		file := NewMockFile(ctrl)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(dir, nil).Times(1)
		fs.EXPECT().OpenFile(path+"/"+fileInfoName, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(nil, expected).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("format").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(file).Return(decoder, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, e := NewSource(path, "format", true, fs, decoderFactory)
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

		partial := &config.Partial{"field": "value"}
		path := "path"
		fileInfoName := "file.yaml"
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().IsDir().Return(false).Times(1)
		fileInfo.EXPECT().Name().Return(fileInfoName).Times(1)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileInfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		file := NewMockFile(ctrl)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(dir, nil).Times(1)
		fs.EXPECT().OpenFile(path+"/"+fileInfoName, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(partial, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("format").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(file).Return(decoder, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, e := NewSource(path, "format", true, fs, decoderFactory)
		switch {
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch {
			case sut.Mutex == nil:
				t.Error("didn't created the access mutex")
			case sut.path != path:
				t.Error("didn't stored the file path")
			case sut.format != "format":
				t.Error("didn't stored the file content format")
			case sut.fs != fs:
				t.Error("didn't stored the file system adapter reference")
			case sut.decoderCreator != decoderFactory:
				t.Error("didn't stored the decoder factory reference")
			case !reflect.DeepEqual(sut.Partial, *partial):
				t.Errorf("didn't loaded the content correctly having (%v) when expecting (%v)", sut.Partial, *partial)
			}
		}
	})

	t.Run("don't follow sub dirs if not recursive", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := &config.Partial{"field": "value"}
		path := "path"
		fileInfoName := "file.yaml"
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().IsDir().Return(false).Times(1)
		fileInfo.EXPECT().Name().Return(fileInfoName).Times(1)
		subDirInfo := NewMockFileInfo(ctrl)
		subDirInfo.EXPECT().IsDir().Return(true).Times(1)
		subDirInfo.EXPECT().Name().Times(0)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileInfo, subDirInfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		file := NewMockFile(ctrl)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Open(path).Return(dir, nil).Times(1)
		fs.EXPECT().OpenFile(path+"/"+fileInfoName, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(partial, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("format").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(file).Return(decoder, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, e := NewSource(path, "format", false, fs, decoderFactory)
		switch {
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch {
			case sut.Mutex == nil:
				t.Error("didn't created the access mutex")
			case sut.path != path:
				t.Error("didn't stored the file path")
			case sut.format != "format":
				t.Error("didn't stored the file content format")
			case sut.fs != fs:
				t.Error("didn't stored the file system adapter reference")
			case sut.decoderCreator != decoderFactory:
				t.Error("didn't stored the decoder factory reference")
			case !reflect.DeepEqual(sut.Partial, *partial):
				t.Errorf("didn't loaded the content correctly having (%v) when expecting (%v)", sut.Partial, *partial)
			}
		}
	})

	t.Run("error while loading sub dir", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial1 := &config.Partial{"field": "value"}
		path := "path"
		expected := fmt.Errorf("error message")
		fileInfoName := "file1.yaml"
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().IsDir().Return(false).Times(1)
		fileInfo.EXPECT().Name().Return(fileInfoName).Times(1)
		subDirInfoName := "sub_dir"
		subDirInfo := NewMockFileInfo(ctrl)
		subDirInfo.EXPECT().IsDir().Return(true).Times(1)
		subDirInfo.EXPECT().Name().Return(subDirInfoName).Times(1)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileInfo, subDirInfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		file1 := NewMockFile(ctrl)
		fs := NewMockFs(ctrl)
		gomock.InOrder(
			fs.EXPECT().Open(path).Return(dir, nil).Times(1),
			fs.EXPECT().Open(path+"/"+subDirInfoName).Return(nil, expected).Times(1))
		fs.EXPECT().OpenFile(path+"/"+fileInfoName, os.O_RDONLY, os.FileMode(0o644)).Return(file1, nil).Times(1)
		decoder1 := NewMockDecoder(ctrl)
		decoder1.EXPECT().Decode().Return(partial1, nil).Times(1)
		decoder1.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("format").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(file1).Return(decoder1, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, e := NewSource(path, "format", true, fs, decoderFactory)
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

		partial1 := &config.Partial{"field1": "value"}
		partial2 := &config.Partial{"field2": "value"}
		expected := config.Partial{"field1": "value", "field2": "value"}
		path := "path"
		fileInfoName := "file1.yaml"
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().IsDir().Return(false).Times(1)
		fileInfo.EXPECT().Name().Return(fileInfoName).Times(1)
		subFileInfoName := "file2.yaml"
		subFileInfo := NewMockFileInfo(ctrl)
		subFileInfo.EXPECT().IsDir().Return(false).Times(1)
		subFileInfo.EXPECT().Name().Return(subFileInfoName).Times(1)
		subDirInfoName := "sub_dir"
		subDirInfo := NewMockFileInfo(ctrl)
		subDirInfo.EXPECT().IsDir().Return(true).Times(1)
		subDirInfo.EXPECT().Name().Return(subDirInfoName).Times(1)
		dir := NewMockFile(ctrl)
		dir.EXPECT().Readdir(0).Return([]os.FileInfo{fileInfo, subDirInfo}, nil).Times(1)
		dir.EXPECT().Close().Return(nil).Times(1)
		subDir := NewMockFile(ctrl)
		subDir.EXPECT().Readdir(0).Return([]os.FileInfo{subFileInfo}, nil).Times(1)
		subDir.EXPECT().Close().Return(nil).Times(1)
		file1 := NewMockFile(ctrl)
		file2 := NewMockFile(ctrl)
		fs := NewMockFs(ctrl)
		gomock.InOrder(
			fs.EXPECT().Open(path).Return(dir, nil).Times(1),
			fs.EXPECT().Open(path+"/"+subDirInfoName).Return(subDir, nil).Times(1))
		gomock.InOrder(
			fs.EXPECT().OpenFile(path+"/"+fileInfoName, os.O_RDONLY, os.FileMode(0o644)).Return(file1, nil).Times(1),
			fs.EXPECT().OpenFile(path+"/"+subDirInfoName+"/"+subFileInfoName, os.O_RDONLY, os.FileMode(0o644)).Return(file2, nil).Times(1))
		decoder1 := NewMockDecoder(ctrl)
		decoder1.EXPECT().Decode().Return(partial1, nil).Times(1)
		decoder1.EXPECT().Close().Return(nil).Times(1)
		decoder2 := NewMockDecoder(ctrl)
		decoder2.EXPECT().Decode().Return(partial2, nil).Times(1)
		decoder2.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		gomock.InOrder(
			decoderStrategy.EXPECT().Accept("format").Return(true),
			decoderStrategy.EXPECT().Accept("format").Return(true),
		)
		gomock.InOrder(
			decoderStrategy.EXPECT().Create(file1).Return(decoder1, nil),
			decoderStrategy.EXPECT().Create(file2).Return(decoder2, nil),
		)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, e := NewSource(path, "format", true, fs, decoderFactory)
		switch {
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch {
			case sut.Mutex == nil:
				t.Error("didn't created the access mutex")
			case sut.path != path:
				t.Error("didn't stored the file path")
			case sut.format != "format":
				t.Error("didn't stored the file content format")
			case sut.fs != fs:
				t.Error("didn't stored the file system adapter reference")
			case sut.decoderCreator != decoderFactory:
				t.Error("didn't stored the decoder factory reference")
			case !reflect.DeepEqual(sut.Partial, expected):
				t.Errorf("didn't loaded the content correctly having (%v) when expecting (%v)", sut.Partial, expected)
			}
		}
	})
}
