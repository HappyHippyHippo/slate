package glog

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/gerror"
	"io"
	"os"
	"testing"
)

func Test_NewStreamStrategyFile(t *testing.T) {
	t.Run("nil file system adapter", func(t *testing.T) {
		strategy, err := NewStreamStrategyFile(nil, &FormatterFactory{})
		switch {
		case strategy != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("nil formatter factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, err := NewStreamStrategyFile(NewMockFs(ctrl), nil)
		switch {
		case strategy != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("new file stream factory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		if strategy, err := NewStreamStrategyFile(NewMockFs(ctrl), &FormatterFactory{}); strategy == nil {
			t.Errorf("didn't returned a valid reference")
		} else if err != nil {
			t.Errorf("returned the (%v) error", err)
		}
	})
}

func Test_StreamStrategyFile_Accept(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	strategy, _ := NewStreamStrategyFile(NewMockFs(ctrl), &FormatterFactory{})

	t.Run("accept only file type", func(t *testing.T) {
		scenarios := []struct {
			streamType string
			expected   bool
		}{
			{ // _test file type
				streamType: StreamFile,
				expected:   true,
			},
			{ // _test non-file format
				streamType: StreamUnknown,
				expected:   false,
			},
		}

		for _, scenario := range scenarios {
			if check := strategy.Accept(scenario.streamType); check != scenario.expected {
				t.Errorf("returned (%v) for the type (%s)", check, scenario.streamType)
			}
		}
	})
}

func Test_StreamStrategyFile_AcceptFromConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	strategy, _ := NewStreamStrategyFile(NewMockFs(ctrl), &FormatterFactory{})

	t.Run("don't accept if config is a nil pointer", func(t *testing.T) {
		if strategy.AcceptFromConfig(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept on type retrieval error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		config.EXPECT().String("type").Return("", fmt.Errorf("dummy error")).Times(1)

		if strategy.AcceptFromConfig(config) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept on invalid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		config.EXPECT().String("type").Return(StreamUnknown, nil).Times(1)

		if strategy.AcceptFromConfig(config) {
			t.Error("returned true")
		}
	})

	t.Run("accept on valid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		config.EXPECT().String("type").Return(StreamFile, nil).Times(1)

		if !strategy.AcceptFromConfig(config) {
			t.Error("returned false")
		}
	})
}

func Test_StreamStrategyFile_Create(t *testing.T) {
	t.Run("non enough arguments", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := NewStreamStrategyFile(NewMockFs(ctrl), &FormatterFactory{})

		stream, err := strategy.Create(1, 2, 3)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("non-string path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := NewStreamStrategyFile(NewMockFs(ctrl), &FormatterFactory{})

		stream, err := strategy.Create(123, FormatJSON, []string{}, DEBUG)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := NewStreamStrategyFile(NewMockFs(ctrl), &FormatterFactory{})

		stream, err := strategy.Create("path", 123, []string{}, DEBUG)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("non-string list channels", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := NewStreamStrategyFile(NewMockFs(ctrl), &FormatterFactory{})

		stream, err := strategy.Create("path", FormatJSON, "string", DEBUG)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("non-loglevel level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := NewStreamStrategyFile(NewMockFs(ctrl), &FormatterFactory{})

		stream, err := strategy.Create("path", FormatJSON, []string{}, "string")
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("error on creating the formatter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := NewStreamStrategyFile(NewMockFs(ctrl), &FormatterFactory{})

		stream, err := strategy.Create("path", FormatJSON, []string{}, DEBUG)
		switch {
		case stream != nil:
			t.Error("returned a valid stream")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrInvalidLogFormat):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrInvalidLogFormat)
		}
	})

	t.Run("error on opening the file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		fileSystem := NewMockFs(ctrl)
		fileSystem.EXPECT().OpenFile("path", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(nil, expected).Times(1)
		factory := &FormatterFactory{}
		_ = factory.Register(&FormatterStrategyJSON{})
		strategy, _ := NewStreamStrategyFile(fileSystem, factory)

		stream, err := strategy.Create("path", FormatJSON, []string{}, DEBUG)
		switch {
		case stream != nil:
			_ = stream.(io.Closer).Close()
			t.Error("returned a valid stream")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("create the file stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		file := NewMockFile(ctrl)
		fileSystem := NewMockFs(ctrl)
		fileSystem.EXPECT().OpenFile("path", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		factory := &FormatterFactory{}
		_ = factory.Register(&FormatterStrategyJSON{})
		strategy, _ := NewStreamStrategyFile(fileSystem, factory)

		stream, err := strategy.Create("path", FormatJSON, []string{}, DEBUG)
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
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

func Test_StreamStrategyFile_CreateFromConfig(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy, _ := NewStreamStrategyFile(NewMockFs(ctrl), &FormatterFactory{})

		src, err := strategy.CreateFromConfig(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("non-string path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("dummy message")
		config := NewMockConfig(ctrl)
		config.EXPECT().String("path").Return("", expected).Times(1)

		strategy, _ := NewStreamStrategyFile(NewMockFs(ctrl), &FormatterFactory{})

		stream, err := strategy.CreateFromConfig(config)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, expected):
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
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

		strategy, _ := NewStreamStrategyFile(NewMockFs(ctrl), &FormatterFactory{})

		stream, err := strategy.CreateFromConfig(config)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, expected):
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
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

		factory := &FormatterFactory{}
		_ = factory.Register(&FormatterStrategyJSON{})
		strategy, _ := NewStreamStrategyFile(NewMockFs(ctrl), factory)

		stream, err := strategy.CreateFromConfig(config)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, expected):
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
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

		factory := &FormatterFactory{}
		_ = factory.Register(&FormatterStrategyJSON{})
		strategy, _ := NewStreamStrategyFile(NewMockFs(ctrl), factory)

		stream, err := strategy.CreateFromConfig(config)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrConversion)
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

		strategy, _ := NewStreamStrategyFile(NewMockFs(ctrl), &FormatterFactory{})

		stream, err := strategy.CreateFromConfig(config)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, expected):
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("non-log level name level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		gomock.InOrder(
			config.EXPECT().String("path").Return("path", nil),
			config.EXPECT().String("format").Return(FormatJSON, nil),
			config.EXPECT().String("level").Return("invalid", nil),
		)
		config.EXPECT().List("channels").Return([]interface{}{"channel1"}, nil)

		strategy, _ := NewStreamStrategyFile(NewMockFs(ctrl), &FormatterFactory{})

		stream, err := strategy.CreateFromConfig(config)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrInvalidLogLevel):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrInvalidLogLevel)
		}
	})

	t.Run("error on creating the formatter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		gomock.InOrder(
			config.EXPECT().String("path").Return("path", nil),
			config.EXPECT().String("format").Return(FormatUnknown, nil),
			config.EXPECT().String("level").Return(LevelMapName[FATAL], nil),
		)
		config.EXPECT().List("channels").Return([]interface{}{"channel1"}, nil)

		strategy, _ := NewStreamStrategyFile(NewMockFs(ctrl), &FormatterFactory{})

		stream, err := strategy.CreateFromConfig(config)
		switch {
		case stream != nil:
			_ = stream.(io.Closer).Close()
			t.Error("returned a valid stream")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrInvalidLogFormat):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrInvalidLogFormat)
		}
	})

	t.Run("error on opening the file", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		config := NewMockConfig(ctrl)
		gomock.InOrder(
			config.EXPECT().String("path").Return("path", nil),
			config.EXPECT().String("format").Return(FormatJSON, nil),
			config.EXPECT().String("level").Return(LevelMapName[FATAL], nil),
		)
		config.EXPECT().List("channels").Return([]interface{}{"channel1"}, nil)

		fileSystem := NewMockFs(ctrl)
		fileSystem.EXPECT().OpenFile("path", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(nil, expected).Times(1)
		factory := &FormatterFactory{}
		_ = factory.Register(&FormatterStrategyJSON{})
		strategy, _ := NewStreamStrategyFile(fileSystem, factory)

		stream, err := strategy.CreateFromConfig(config)
		switch {
		case stream != nil:
			_ = stream.(io.Closer).Close()
			t.Error("returned a valid stream")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("new stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		gomock.InOrder(
			config.EXPECT().String("path").Return("path", nil),
			config.EXPECT().String("format").Return(FormatJSON, nil),
			config.EXPECT().String("level").Return(LevelMapName[FATAL], nil),
		)
		config.EXPECT().List("channels").Return([]interface{}{"channel1"}, nil)

		file := NewMockFile(ctrl)
		file.EXPECT().Close().Return(nil).Times(1)
		fileSystem := NewMockFs(ctrl)
		fileSystem.EXPECT().OpenFile("path", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).Return(file, nil).Times(1)
		factory := &FormatterFactory{}
		_ = factory.Register(&FormatterStrategyJSON{})
		strategy, _ := NewStreamStrategyFile(fileSystem, factory)

		stream, err := strategy.CreateFromConfig(config)
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
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
