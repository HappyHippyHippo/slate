package log

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	serror "github.com/happyhippyhippo/slate/error"
)

func Test_NewStreamStrategyConsole(t *testing.T) {
	t.Run("nil formatter factory", func(t *testing.T) {
		sut, e := newStreamStrategyConsole(nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("new console stream factory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		if sut, e := newStreamStrategyConsole(NewMockFormatterFactory(ctrl)); sut == nil {
			t.Errorf("didn't returned a valid reference")
		} else if e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}

func Test_StreamStrategyConsole_Accept(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sut, _ := newStreamStrategyConsole(NewMockFormatterFactory(ctrl))

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
			if check := sut.Accept(scenario.streamType); check != scenario.expected {
				t.Errorf("returned (%v) for the type (%s)", check, scenario.streamType)
			}
		}
	})
}

func Test_StreamStrategyConsole_AcceptFromConfig(t *testing.T) {
	t.Run("don't accept if config is a nil pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newStreamStrategyConsole(NewMockFormatterFactory(ctrl))

		if sut.AcceptFromConfig(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept on type retrieval error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		config.EXPECT().String("type").Return("", fmt.Errorf("dummy error")).Times(1)

		sut, _ := newStreamStrategyConsole(NewMockFormatterFactory(ctrl))

		if sut.AcceptFromConfig(config) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept on invalid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		config.EXPECT().String("type").Return(StreamUnknown, nil).Times(1)

		sut, _ := newStreamStrategyConsole(NewMockFormatterFactory(ctrl))

		if sut.AcceptFromConfig(config) {
			t.Error("returned true")
		}
	})

	t.Run("accept on valid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		config.EXPECT().String("type").Return(StreamConsole, nil).Times(1)

		sut, _ := newStreamStrategyConsole(NewMockFormatterFactory(ctrl))

		if !sut.AcceptFromConfig(config) {
			t.Error("returned false")
		}
	})
}

func Test_StreamStrategyConsole_Create(t *testing.T) {
	t.Run("non enough arguments", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newStreamStrategyConsole(NewMockFormatterFactory(ctrl))

		stream, e := sut.Create(1, 2)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newStreamStrategyConsole(NewMockFormatterFactory(ctrl))

		stream, e := sut.Create(123, []string{}, DEBUG)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("non-string list channels", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newStreamStrategyConsole(NewMockFormatterFactory(ctrl))

		stream, e := sut.Create(FormatJSON, "string", DEBUG)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("non-loglevel level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newStreamStrategyConsole(NewMockFormatterFactory(ctrl))

		stream, e := sut.Create(FormatJSON, []string{}, "string")
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
		}
	})

	t.Run("error on creating the formatter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		formatterFactory := NewMockFormatterFactory(ctrl)
		formatterFactory.EXPECT().Create(FormatJSON).Return(nil, expected).Times(1)
		sut, _ := newStreamStrategyConsole(formatterFactory)

		stream, e := sut.Create(FormatJSON, []string{}, DEBUG)
		switch {
		case stream != nil:
			t.Error("returned a valid stream")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("create the console stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		formatter := NewMockFormatter(ctrl)
		formatterFactory := NewMockFormatterFactory(ctrl)
		formatterFactory.EXPECT().Create(FormatJSON).Return(formatter, nil).Times(1)
		sut, _ := newStreamStrategyConsole(formatterFactory)

		stream, e := sut.Create(FormatJSON, []string{}, DEBUG)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
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
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := newStreamStrategyConsole(NewMockFormatterFactory(ctrl))

		src, e := sut.CreateFromConfig(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("non-string format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("dummy message")
		config := NewMockConfig(ctrl)
		config.EXPECT().String("format").Return("", expected).Times(1)

		sut, _ := newStreamStrategyConsole(NewMockFormatterFactory(ctrl))

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
		config.EXPECT().String("format").Return(FormatJSON, nil)
		config.EXPECT().List("channels").Return(nil, expected)

		sut, _ := newStreamStrategyConsole(NewMockFormatterFactory(ctrl))

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
		config.EXPECT().String("format").Return(FormatJSON, nil)
		config.EXPECT().List("channels").Return([]interface{}{123}, nil)

		sut, _ := newStreamStrategyConsole(NewMockFormatterFactory(ctrl))

		stream, e := sut.CreateFromConfig(config)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrConversion)
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

		sut, _ := newStreamStrategyConsole(NewMockFormatterFactory(ctrl))

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

	t.Run("non-log level name level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		gomock.InOrder(
			config.EXPECT().String("format").Return(FormatJSON, nil),
			config.EXPECT().String("level").Return("invalid", nil),
		)
		config.EXPECT().List("channels").Return([]interface{}{"channel1"}, nil)

		sut, _ := newStreamStrategyConsole(NewMockFormatterFactory(ctrl))

		stream, e := sut.CreateFromConfig(config)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrInvalidLogLevel):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrInvalidLogLevel)
		}
	})

	t.Run("error creating the formatter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("dummy message")
		config := NewMockConfig(ctrl)
		gomock.InOrder(
			config.EXPECT().String("format").Return(FormatJSON, nil),
			config.EXPECT().String("level").Return(LevelMapName[FATAL], nil),
		)
		config.EXPECT().List("channels").Return([]interface{}{"channel1"}, nil)
		formatterFactory := NewMockFormatterFactory(ctrl)
		formatterFactory.EXPECT().Create(FormatJSON).Return(nil, expected).Times(1)

		sut, _ := newStreamStrategyConsole(formatterFactory)

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

	t.Run("new stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		gomock.InOrder(
			config.EXPECT().String("format").Return(FormatJSON, nil),
			config.EXPECT().String("level").Return(LevelMapName[FATAL], nil),
		)
		config.EXPECT().List("channels").Return([]interface{}{"channel1"}, nil)
		formatter := NewMockFormatter(ctrl)
		formatterFactory := NewMockFormatterFactory(ctrl)
		formatterFactory.EXPECT().Create(FormatJSON).Return(formatter, nil).Times(1)

		sut, _ := newStreamStrategyConsole(formatterFactory)

		stream, e := sut.CreateFromConfig(config)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
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
