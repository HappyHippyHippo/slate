package file

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
)

func Test_NewObsSource(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		sut, e := NewObsSource("path", "format", nil, config.NewDecoderFactory())
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil decoder factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewObsSource("path", "format", NewMockFs(ctrl), nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error that may be raised when retrieving the file info", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		expected := fmt.Errorf("error message")
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(nil, expected).Times(1)
		decoderFactory := config.NewDecoderFactory()

		sut, e := NewObsSource(path, "format", fs, decoderFactory)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("unexpected (%v) error", e)
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
		decoderFactory := config.NewDecoderFactory()

		sut, e := NewObsSource(path, "format", fs, decoderFactory)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("unexpected (%v) error", e)
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
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept(config.UnknownDecoder).Return(false).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, e := NewObsSource(path, config.UnknownDecoder, fs, decoderFactory)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case errors.Is(e, config.ErrInvalidSource):
			t.Errorf("(%v) when expecting (%v)", e, config.ErrInvalidSource)
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
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("format").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(file).Return(decoder, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, e := NewObsSource(path, "format", fs, decoderFactory)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("(%v) when expecting (%v)", e, expected)
		}
	})

	t.Run("create the config observable file source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		field := "field"
		value := "value"
		expected := config.Partial{field: value}
		file := NewMockFile(ctrl)
		fileInfo := NewMockFileInfo(ctrl)
		fileInfo.EXPECT().ModTime().Return(time.Unix(0, 1)).Times(1)
		fs := NewMockFs(ctrl)
		fs.EXPECT().Stat(path).Return(fileInfo, nil).Times(1)
		fs.EXPECT().OpenFile(path, os.O_RDONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		decoder := NewMockDecoder(ctrl)
		decoder.EXPECT().Decode().Return(&expected, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("format").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(file).Return(decoder, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, e := NewObsSource(path, "format", fs, decoderFactory)
		switch {
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		default:
			switch {
			case sut.path != path:
				t.Error("didn't stored the file path")
			case sut.format != "format":
				t.Error("didn't stored the file content format")
			case sut.fileSystem != fs:
				t.Error("didn't stored the file system adapter reference")
			case sut.decoderCreator != decoderFactory:
				t.Error("didn't stored the decoder factory reference")
			}
		}
	})
}

func Test_ObsSource_Reload(t *testing.T) {
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
		decoder.EXPECT().Decode().Return(&config.Partial{field: value}, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("format").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(file).Return(decoder, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, _ := NewObsSource(path, "format", fs, decoderFactory)

		reloaded, e := sut.Reload()
		switch {
		case reloaded:
			t.Error("flagged that was reloaded")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("unexpected (%v) error", e)
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
		decoder.EXPECT().Decode().Return(&config.Partial{field: value}, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("format").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(file).Return(decoder, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, _ := NewObsSource(path, "format", fs, decoderFactory)

		reloaded, e := sut.Reload()
		switch {
		case reloaded:
			t.Error("flagged that was reloaded")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("unexpected (%v) error", e)
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
		decoder.EXPECT().Decode().Return(&config.Partial{field: value}, nil).Times(1)
		decoder.EXPECT().Close().Return(nil).Times(1)
		decoderStrategy := NewMockDecoderStrategy(ctrl)
		decoderStrategy.EXPECT().Accept("format").Return(true).Times(1)
		decoderStrategy.EXPECT().Create(file).Return(decoder, nil).Times(1)
		decoderFactory := config.NewDecoderFactory()
		_ = decoderFactory.Register(decoderStrategy)

		sut, _ := NewObsSource(path, "format", fs, decoderFactory)

		if reloaded, e := sut.Reload(); reloaded {
			t.Error("flagged that was reloaded")
		} else if e != nil {
			t.Errorf("unexpected (%v) error", e)
		}
	})

	t.Run("should reload a changed source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		field := "field"
		value1 := "value1"
		value2 := "value2"
		expected := config.Partial{field: value2}
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
		decoder1.EXPECT().Decode().Return(&config.Partial{field: value1}, nil).Times(1)
		decoder1.EXPECT().Close().Return(nil).Times(1)
		decoder2 := NewMockDecoder(ctrl)
		decoder2.EXPECT().Decode().Return(&config.Partial{field: value2}, nil).Times(1)
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

		sut, _ := NewObsSource(path, "format", fs, decoderFactory)

		reloaded, e := sut.Reload()
		p, _ := sut.Get("")

		switch {
		case !reloaded:
			t.Error("flagged that was not reloaded")
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !reflect.DeepEqual(expected, p):
			t.Error("didn't stored the check configuration")
		}
	})
}
