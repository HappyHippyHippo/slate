package slog

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serr"
	"io"
	"os"
	"testing"
	"time"
)

func Test_NewStreamStrategyFileRotate(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := newStreamStrategyRotatingFile(nil, NewMockFormatterFactory(ctrl))
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("nil formatter factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := newStreamStrategyRotatingFile(NewMockFs(ctrl), nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("new file stream factory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := newStreamStrategyRotatingFile(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))
		switch {
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		}
	})
}

func Test_StreamStrategyFileRotate_Accept(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sut, _ := newStreamStrategyRotatingFile(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

	t.Run("accept only file type", func(t *testing.T) {
		scenarios := []struct {
			streamType string
			expected   bool
		}{
			{ // _test rotating file type
				streamType: StreamRotatingFile,
				expected:   true,
			},
			{ // _test non-rotating-file format
				streamType: StreamUnknown,
				expected:   false,
			},
		}

		for _, scenario := range scenarios {
			if check := sut.Accept(scenario.streamType); check != scenario.expected {
				t.Errorf("returned (%v) for the type (%s)", check, scenario.streamType)
			}
		}
	})
}

func Test_StreamStrategyFileRotate_AcceptFromConfig(t *testing.T) {
	t.Run("don't accept if sconfig is a nil pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newStreamStrategyRotatingFile(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		if sut.AcceptFromConfig(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept on type retrieval error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		config.EXPECT().String("type").Return("", fmt.Errorf("dummy error")).Times(1)

		sut, _ := newStreamStrategyRotatingFile(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		if sut.AcceptFromConfig(config) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept on invalid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		config.EXPECT().String("type").Return(StreamUnknown, nil).Times(1)

		sut, _ := newStreamStrategyRotatingFile(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		if sut.AcceptFromConfig(config) {
			t.Error("returned true")
		}
	})

	t.Run("accept on valid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		config.EXPECT().String("type").Return(StreamRotatingFile, nil).Times(1)

		sut, _ := newStreamStrategyRotatingFile(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		if !sut.AcceptFromConfig(config) {
			t.Error("returned false")
		}
	})
}

func Test_StreamStrategyFileRotate_Create(t *testing.T) {
	t.Run("non enough arguments", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newStreamStrategyRotatingFile(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		stream, e := sut.Create(1, 2, 3)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("non-string path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newStreamStrategyRotatingFile(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		stream, e := sut.Create(123, FormatJSON, []string{}, DEBUG)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newStreamStrategyRotatingFile(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		stream, e := sut.Create("path", 123, []string{}, DEBUG)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("non-string list channels", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newStreamStrategyRotatingFile(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		stream, e := sut.Create("path", FormatJSON, "string", DEBUG)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("non-slog level level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newStreamStrategyRotatingFile(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		stream, e := sut.Create("path", FormatJSON, []string{}, "string")
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("error on creating the formatter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		fFactory := NewMockFormatterFactory(ctrl)
		fFactory.EXPECT().Create(FormatJSON).Return(nil, expected).Times(1)

		sut, _ := newStreamStrategyRotatingFile(NewMockFs(ctrl), fFactory)

		stream, e := sut.Create("path", FormatJSON, []string{}, DEBUG)
		switch {
		case stream != nil:
			t.Error("returned a valid stream")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("error on opening the file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		path := "path-%s"
		expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
		fileSystem := NewMockFs(ctrl)
		fileSystem.EXPECT().OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(nil, expected).Times(1)
		formatter := NewMockFormatter(ctrl)
		fFactory := NewMockFormatterFactory(ctrl)
		fFactory.EXPECT().Create(FormatJSON).Return(formatter, nil).Times(1)

		sut, _ := newStreamStrategyRotatingFile(fileSystem, fFactory)

		stream, e := sut.Create(path, FormatJSON, []string{}, DEBUG)
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

	t.Run("create the file stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path-%s"
		expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
		file := NewMockFile(ctrl)
		fileSystem := NewMockFs(ctrl)
		fileSystem.EXPECT().OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		formatter := NewMockFormatter(ctrl)
		fFactory := NewMockFormatterFactory(ctrl)
		fFactory.EXPECT().Create(FormatJSON).Return(formatter, nil).Times(1)

		sut, _ := newStreamStrategyRotatingFile(fileSystem, fFactory)

		stream, e := sut.Create(path, FormatJSON, []string{}, DEBUG)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case stream == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch stream.(type) {
			case *streamFile:
			default:
				t.Error("didn't returned a new file stream")
			}
		}
	})
}

func Test_StreamStrategyFileRotate_CreateFromConfig(t *testing.T) {
	t.Run("error on nil sconfig pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newStreamStrategyRotatingFile(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		src, e := sut.CreateFromConfig(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("non-string path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("dummy message")
		config := NewMockConfig(ctrl)
		config.EXPECT().String("path").Return("", expected).Times(1)

		sut, _ := newStreamStrategyRotatingFile(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		stream, e := sut.CreateFromConfig(config)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, expected):
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("dummy message")
		config := NewMockConfig(ctrl)
		gomock.InOrder(
			config.EXPECT().String("path").Return("path", nil),
			config.EXPECT().String("format").Return("", expected),
		)

		sut, _ := newStreamStrategyRotatingFile(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		stream, e := sut.CreateFromConfig(config)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, expected):
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("non-list channels", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("dummy message")
		config := NewMockConfig(ctrl)
		gomock.InOrder(
			config.EXPECT().String("path").Return("path", nil),
			config.EXPECT().String("format").Return(FormatJSON, nil),
		)
		config.EXPECT().List("channels").Return(nil, expected)

		sut, _ := newStreamStrategyRotatingFile(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		stream, e := sut.CreateFromConfig(config)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, expected):
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("non-strict string list channels", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		gomock.InOrder(
			config.EXPECT().String("path").Return("path", nil),
			config.EXPECT().String("format").Return(FormatJSON, nil),
		)
		config.EXPECT().List("channels").Return([]interface{}{123}, nil)

		sut, _ := newStreamStrategyRotatingFile(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		stream, e := sut.CreateFromConfig(config)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("non-string level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("dummy message")
		config := NewMockConfig(ctrl)
		gomock.InOrder(
			config.EXPECT().String("path").Return("path", nil),
			config.EXPECT().String("format").Return(FormatJSON, nil),
			config.EXPECT().String("level").Return("", expected),
		)
		config.EXPECT().List("channels").Return([]interface{}{"channel1"}, nil)

		sut, _ := newStreamStrategyRotatingFile(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		stream, e := sut.CreateFromConfig(config)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, expected):
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("non-slog level name level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		gomock.InOrder(
			config.EXPECT().String("path").Return("path", nil),
			config.EXPECT().String("format").Return(FormatJSON, nil),
			config.EXPECT().String("level").Return("invalid", nil),
		)
		config.EXPECT().List("channels").Return([]interface{}{"channel1"}, nil)

		sut, _ := newStreamStrategyRotatingFile(NewMockFs(ctrl), NewMockFormatterFactory(ctrl))

		stream, e := sut.CreateFromConfig(config)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrInvalidLogLevel):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrInvalidLogLevel)
		}
	})

	t.Run("error on creating the formatter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")

		config := NewMockConfig(ctrl)
		gomock.InOrder(
			config.EXPECT().String("path").Return("path", nil),
			config.EXPECT().String("format").Return(FormatUnknown, nil),
			config.EXPECT().String("level").Return(LevelMapName[FATAL], nil),
		)
		config.EXPECT().List("channels").Return([]interface{}{"channel1"}, nil)

		fFactory := NewMockFormatterFactory(ctrl)
		fFactory.EXPECT().Create(FormatUnknown).Return(nil, expected).Times(1)

		sut, _ := newStreamStrategyRotatingFile(NewMockFs(ctrl), fFactory)

		stream, e := sut.CreateFromConfig(config)
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

	t.Run("error on opening the file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path-%s"
		expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
		expected := fmt.Errorf("error message")

		config := NewMockConfig(ctrl)
		gomock.InOrder(
			config.EXPECT().String("path").Return(path, nil),
			config.EXPECT().String("format").Return(FormatJSON, nil),
			config.EXPECT().String("level").Return(LevelMapName[FATAL], nil),
		)
		config.EXPECT().List("channels").Return([]interface{}{"channel1"}, nil)

		fileSystem := NewMockFs(ctrl)
		fileSystem.EXPECT().OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(nil, expected).Times(1)
		formatter := NewMockFormatter(ctrl)
		fFactory := NewMockFormatterFactory(ctrl)
		fFactory.EXPECT().Create(FormatJSON).Return(formatter, nil).Times(1)

		sut, _ := newStreamStrategyRotatingFile(fileSystem, fFactory)

		stream, e := sut.CreateFromConfig(config)
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

		path := "path-%s"
		expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))

		config := NewMockConfig(ctrl)
		gomock.InOrder(
			config.EXPECT().String("path").Return(path, nil),
			config.EXPECT().String("format").Return(FormatJSON, nil),
			config.EXPECT().String("level").Return(LevelMapName[FATAL], nil),
		)
		config.EXPECT().List("channels").Return([]interface{}{"channel1"}, nil)

		file := NewMockFile(ctrl)
		file.EXPECT().Close().Return(nil).Times(1)
		fileSystem := NewMockFs(ctrl)
		fileSystem.EXPECT().OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		formatter := NewMockFormatter(ctrl)
		fFactory := NewMockFormatterFactory(ctrl)
		fFactory.EXPECT().Create(FormatJSON).Return(formatter, nil).Times(1)

		sut, _ := newStreamStrategyRotatingFile(fileSystem, fFactory)

		stream, e := sut.CreateFromConfig(config)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case stream == nil:
			t.Error("didn't returned a valid reference")
		default:
			_ = stream.(io.Closer).Close()
			switch stream.(type) {
			case *streamFile:
			default:
				t.Error("didn't returned a new file stream")
			}
		}
	})
}
