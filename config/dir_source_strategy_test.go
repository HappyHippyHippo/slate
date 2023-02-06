package config

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/err"
)

func Test_NewDirSourceStrategy(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewDirSourceStrategy(nil, NewMockDecoderFactory(ctrl))
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

		sut, e := NewDirSourceStrategy(NewMockFs(ctrl), nil)
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

		sut, e := NewDirSourceStrategy(fs, decoderFactory)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case sut.fileSystem != fs:
			t.Error("didn't stored the file system adapter reference")
		case sut.decoderFactory != decoderFactory:
			t.Error("didn't stored the decoder factory reference")
		}
	})
}

func Test_DirSourceStrategy_Accept(t *testing.T) {
	t.Run("don't accept on invalid config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewDirSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.Accept(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewDirSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.Accept(&Config{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewDirSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.Accept(&Config{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if invalid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewDirSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if sut.Accept(&Config{"type": SourceStrategyUnknown}) {
			t.Error("returned true")
		}
	})

	t.Run("accept config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewDirSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		if !sut.Accept(&Config{"type": SourceStrategyDirectory}) {
			t.Error("returned false")
		}
	})
}

func Test_DirSourceStrategy_Create(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewDirSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

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

		sut, _ := NewDirSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&Config{"format": "format", "recursive": true})
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

		sut, _ := NewDirSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&Config{"path": 123, "format": "format", "recursive": true})
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

		sut, _ := NewDirSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&Config{"path": "path", "format": 123, "recursive": true})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("non-bool recursive flag", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewDirSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&Config{"path": "path", "format": 123, "recursive": "true"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("non-bool recursive flag", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewDirSourceStrategy(NewMockFs(ctrl), NewMockDecoderFactory(ctrl))

		src, e := sut.Create(&Config{"path": "path", "format": "format", "recursive": "true"})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("create the dir source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := Config{"field": "value"}
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
		decoder.EXPECT().Decode().Return(&expected, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderFactory := NewMockDecoderFactory(ctrl)
		decoderFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, _ := NewDirSourceStrategy(fs, decoderFactory)

		src, e := sut.Create(&Config{"path": path, "format": DecoderFormatYAML, "recursive": true})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *DirSource:
				if !reflect.DeepEqual(s.config, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new file src")
			}
		}
	})

	t.Run("create the dir source defaulting format if not present in config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := Config{"field": "value"}
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
		decoder.EXPECT().Decode().Return(&expected, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderFactory := NewMockDecoderFactory(ctrl)
		decoderFactory.EXPECT().Create(DecoderFormatYAML, file).Return(decoder, nil).Times(1)

		sut, _ := NewDirSourceStrategy(fs, decoderFactory)

		src, e := sut.Create(&Config{"path": path, "recursive": true})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *DirSource:
				if !reflect.DeepEqual(s.config, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new file src")
			}
		}
	})
}
