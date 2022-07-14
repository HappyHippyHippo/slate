package slog

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serror"
	"testing"
)

func Test_NewStreamStrategyConsole(t *testing.T) {
	t.Run("nil formatter factory", func(t *testing.T) {
		strategy, err := newStreamStrategyConsole(nil)
		switch {
		case strategy != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("new console stream factory strategy", func(t *testing.T) {
		if strategy, err := newStreamStrategyConsole(&FormatterFactory{}); strategy == nil {
			t.Errorf("didn't returned a valid reference")
		} else if err != nil {
			t.Errorf("returned the (%v) error", err)
		}
	})
}

func Test_StreamStrategyConsole_Accept(t *testing.T) {
	strategy, _ := newStreamStrategyConsole(&FormatterFactory{})

	t.Run("accept only console type", func(t *testing.T) {
		scenarios := []struct {
			streamType string
			expected   bool
		}{
			{ // _test console type
				streamType: StreamConsole,
				expected:   true,
			},
			{ // _test non-console format
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

func Test_StreamStrategyConsole_AcceptFromConfig(t *testing.T) {
	strategy, _ := newStreamStrategyConsole(&FormatterFactory{})

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
		config.EXPECT().String("type").Return(StreamConsole, nil).Times(1)

		if !strategy.AcceptFromConfig(config) {
			t.Error("returned false")
		}
	})
}

func Test_StreamStrategyConsole_Create(t *testing.T) {
	t.Run("non enough arguments", func(t *testing.T) {
		strategy, _ := newStreamStrategyConsole(&FormatterFactory{})

		stream, err := strategy.Create(1, 2)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		strategy, _ := newStreamStrategyConsole(&FormatterFactory{})

		stream, err := strategy.Create(123, []string{}, DEBUG)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("non-string list channels", func(t *testing.T) {
		strategy, _ := newStreamStrategyConsole(&FormatterFactory{})

		stream, err := strategy.Create(FormatJSON, "string", DEBUG)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("non-loglevel level", func(t *testing.T) {
		strategy, _ := newStreamStrategyConsole(&FormatterFactory{})

		stream, err := strategy.Create(FormatJSON, []string{}, "string")
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("error on creating the formatter", func(t *testing.T) {
		strategy, _ := newStreamStrategyConsole(&FormatterFactory{})

		stream, err := strategy.Create(FormatJSON, []string{}, DEBUG)
		switch {
		case stream != nil:
			t.Error("returned a valid stream")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrInvalidLogFormat):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrInvalidLogFormat)
		}
	})

	t.Run("create the console stream", func(t *testing.T) {
		fFactory := &FormatterFactory{}
		_ = fFactory.Register(&formatterStrategyJSON{})
		strategy, _ := newStreamStrategyConsole(fFactory)

		stream, err := strategy.Create(FormatJSON, []string{}, DEBUG)
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case stream == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch stream.(type) {
			case *streamConsole:
			default:
				t.Error("didn't returned a new console stream")
			}
		}
	})
}

func Test_StreamStrategyConsole_CreateFromConfig(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		strategy, _ := newStreamStrategyConsole(&FormatterFactory{})

		src, err := strategy.CreateFromConfig(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("dummy message")
		config := NewMockConfig(ctrl)
		config.EXPECT().String("format").Return("", expected).Times(1)

		strategy, _ := newStreamStrategyConsole(&FormatterFactory{})

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
		config.EXPECT().String("format").Return(FormatJSON, nil)
		config.EXPECT().List("channels").Return(nil, expected)

		fFactory := &FormatterFactory{}
		_ = fFactory.Register(&formatterStrategyJSON{})
		strategy, _ := newStreamStrategyConsole(fFactory)

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
		config.EXPECT().String("format").Return(FormatJSON, nil)
		config.EXPECT().List("channels").Return([]interface{}{123}, nil)

		fFactory := &FormatterFactory{}
		_ = fFactory.Register(&formatterStrategyJSON{})
		strategy, _ := newStreamStrategyConsole(fFactory)

		stream, err := strategy.CreateFromConfig(config)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("non-string level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("dummy message")
		config := NewMockConfig(ctrl)
		gomock.InOrder(
			config.EXPECT().String("format").Return(FormatJSON, nil),
			config.EXPECT().String("level").Return("", expected),
		)
		config.EXPECT().List("channels").Return([]interface{}{"channel1"}, nil)

		strategy, _ := newStreamStrategyConsole(&FormatterFactory{})

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
			config.EXPECT().String("format").Return(FormatJSON, nil),
			config.EXPECT().String("level").Return("invalid", nil),
		)
		config.EXPECT().List("channels").Return([]interface{}{"channel1"}, nil)

		strategy, _ := newStreamStrategyConsole(&FormatterFactory{})

		stream, err := strategy.CreateFromConfig(config)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrInvalidLogLevel):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrInvalidLogLevel)
		}
	})

	t.Run("new stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		gomock.InOrder(
			config.EXPECT().String("format").Return(FormatJSON, nil),
			config.EXPECT().String("level").Return(LevelMapName[FATAL], nil),
		)
		config.EXPECT().List("channels").Return([]interface{}{"channel1"}, nil)

		fFactory := &FormatterFactory{}
		_ = fFactory.Register(&formatterStrategyJSON{})
		strategy, _ := newStreamStrategyConsole(fFactory)

		stream, err := strategy.CreateFromConfig(config)
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case stream == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch stream.(type) {
			case *streamConsole:
			default:
				t.Error("didn't returned a new console stream")
			}
		}
	})
}
