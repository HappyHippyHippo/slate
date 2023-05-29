package file

import (
	"errors"
	"fmt"
	"github.com/happyhippyhippo/slate/config"
	"io"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/log"
)

func Test_NewStreamStrategyFile(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		sut, e := NewStreamStrategy(nil, log.NewFormatterFactory())
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil formatter factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewStreamStrategy(NewMockFs(ctrl), nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("new file stream factory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		if sut, e := NewStreamStrategy(NewMockFs(ctrl), log.NewFormatterFactory()); sut == nil {
			t.Errorf("didn't returned a valid reference")
		} else if e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}

func Test_StreamStrategyFile_Accept(t *testing.T) {
	t.Run("don't accept if config is a nil pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewStreamStrategy(NewMockFs(ctrl), log.NewFormatterFactory())

		if sut.Accept(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept on type retrieval error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{"type": 123}
		sut, _ := NewStreamStrategy(NewMockFs(ctrl), log.NewFormatterFactory())

		if sut.Accept(&partial) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept on invalid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{"type": "invalid"}
		sut, _ := NewStreamStrategy(NewMockFs(ctrl), log.NewFormatterFactory())

		if sut.Accept(&partial) {
			t.Error("returned true")
		}
	})

	t.Run("accept on valid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{"type": Type}
		sut, _ := NewStreamStrategy(NewMockFs(ctrl), log.NewFormatterFactory())

		if !sut.Accept(&partial) {
			t.Error("returned false")
		}
	})
}

func Test_StreamStrategyFile_Create(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewStreamStrategy(NewMockFs(ctrl), log.NewFormatterFactory())

		src, e := sut.Create(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{"type": Type, "format": 123}
		sut, _ := NewStreamStrategy(NewMockFs(ctrl), log.NewFormatterFactory())

		src, e := sut.Create(&partial)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("non-log level name level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{"type": Type, "format": "format", "level": "invalid"}
		sut, _ := NewStreamStrategy(NewMockFs(ctrl), log.NewFormatterFactory())

		stream, e := sut.Create(&partial)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, log.ErrInvalidLevel):
			t.Errorf("returned the (%v) error when expecting (%v)", e, log.ErrInvalidLevel)
		}
	})

	t.Run("error on creating the formatter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{"type": Type, "format": "format", "level": "fatal"}
		sut, _ := NewStreamStrategy(NewMockFs(ctrl), log.NewFormatterFactory())

		stream, e := sut.Create(&partial)
		switch {
		case stream != nil:
			_ = stream.(io.Closer).Close()
			t.Error("returned a valid stream")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, log.ErrInvalidFormat):
			t.Errorf("returned the (%v) error when expecting (%v)", e, log.ErrInvalidFormat)
		}
	})

	t.Run("error on opening the file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		partial := config.Partial{
			"type":     Type,
			"format":   "format",
			"level":    "fatal",
			"path":     "path",
			"channels": []interface{}{"channel1", "channel2"}}
		fileSystem := NewMockFs(ctrl)
		fileSystem.EXPECT().OpenFile("path", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(nil, expected).Times(1)
		formatter := NewMockFormatter(ctrl)
		formatterStrategy := NewMockFormatterStrategy(ctrl)
		formatterStrategy.EXPECT().Accept("format").Return(true).Times(1)
		formatterStrategy.EXPECT().Create(gomock.Any()).Return(formatter, nil).Times(1)
		formatterFactory := log.NewFormatterFactory()
		_ = formatterFactory.Register(formatterStrategy)

		sut, _ := NewStreamStrategy(fileSystem, formatterFactory)

		stream, e := sut.Create(&partial)
		switch {
		case stream != nil:
			_ = stream.(io.Closer).Close()
			t.Error("returned a valid stream")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("new stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{
			"type":     Type,
			"format":   "format",
			"level":    "fatal",
			"path":     "path",
			"channels": []interface{}{"channel1", "channel2"}}
		file := NewMockFile(ctrl)
		file.EXPECT().Close().Return(nil).Times(1)
		fileSystem := NewMockFs(ctrl)
		fileSystem.EXPECT().OpenFile("path", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		formatter := NewMockFormatter(ctrl)
		formatterStrategy := NewMockFormatterStrategy(ctrl)
		formatterStrategy.EXPECT().Accept("format").Return(true).Times(1)
		formatterStrategy.EXPECT().Create(gomock.Any()).Return(formatter, nil).Times(1)
		formatterFactory := log.NewFormatterFactory()
		_ = formatterFactory.Register(formatterStrategy)

		sut, _ := NewStreamStrategy(fileSystem, formatterFactory)

		stream, e := sut.Create(&partial)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case stream == nil:
			t.Error("didn't returned a valid reference")
		default:
			_ = stream.(io.Closer).Close()
			switch s := stream.(type) {
			case *Stream:
				switch {
				case s.Level != log.FATAL:
					t.Error("didn't created a stream with the correct level")
				case len(s.Channels) != 2:
					t.Error("didn't created a stream with the correct channel list")
				}
			default:
				t.Error("didn't returned a new file stream")
			}
		}
	})
}
