package log

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
)

func Test_NewConsoleStreamStrategy(t *testing.T) {
	t.Run("nil formatter factory", func(t *testing.T) {
		sut, e := NewConsoleStreamStrategy(nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("new console stream factory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		if sut, e := NewConsoleStreamStrategy(NewMockFormatterFactory(ctrl)); sut == nil {
			t.Errorf("didn't returned a valid reference")
		} else if e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}

func Test_ConsoleStreamStrategy_Accept(t *testing.T) {
	t.Run("don't accept if config is a nil pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewConsoleStreamStrategy(NewMockFormatterFactory(ctrl))

		if sut.Accept(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept on type retrieval error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		config.EXPECT().Populate("", gomock.AssignableToTypeOf(&struct{ Type string }{})).DoAndReturn(
			func(_ string, data *struct{ Type string }, _ ...bool) (interface{}, error) {
				return nil, fmt.Errorf("dummy error")
			},
		).Times(1)

		sut, _ := NewConsoleStreamStrategy(NewMockFormatterFactory(ctrl))

		if sut.Accept(config) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept on invalid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		config.EXPECT().Populate("", gomock.AssignableToTypeOf(&struct{ Type string }{})).DoAndReturn(
			func(_ string, data *struct{ Type string }, _ ...bool) (interface{}, error) {
				data.Type = UnknownStream
				return data, nil
			},
		).Times(1)

		sut, _ := NewConsoleStreamStrategy(NewMockFormatterFactory(ctrl))

		if sut.Accept(config) {
			t.Error("returned true")
		}
	})

	t.Run("accept on valid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		config.EXPECT().Populate("", gomock.AssignableToTypeOf(&struct{ Type string }{})).DoAndReturn(
			func(_ string, data *struct{ Type string }, _ ...bool) (interface{}, error) {
				data.Type = ConsoleStreamType
				return data, nil
			},
		).Times(1)

		sut, _ := NewConsoleStreamStrategy(NewMockFormatterFactory(ctrl))

		if !sut.Accept(config) {
			t.Error("returned false")
		}
	})
}

func Test_ConsoleStreamStrategy_Create(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewConsoleStreamStrategy(NewMockFormatterFactory(ctrl))

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

		expected := fmt.Errorf("dummy message")
		config := NewMockConfig(ctrl)
		config.EXPECT().Populate("", gomock.AssignableToTypeOf(&consoleStreamConfig{})).DoAndReturn(
			func(_ string, data *consoleStreamConfig, _ ...bool) (interface{}, error) {
				return nil, expected
			},
		).Times(1)

		sut, _ := NewConsoleStreamStrategy(NewMockFormatterFactory(ctrl))

		stream, e := sut.Create(config)
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
		config.EXPECT().Populate("", gomock.AssignableToTypeOf(&consoleStreamConfig{})).DoAndReturn(
			func(_ string, data *consoleStreamConfig, _ ...bool) (interface{}, error) {
				data.Format = JSONFormatterFormat
				data.Level = "invalid"
				return data, nil
			},
		).Times(1)

		sut, _ := NewConsoleStreamStrategy(NewMockFormatterFactory(ctrl))

		stream, e := sut.Create(config)
		switch {
		case stream != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, ErrInvalidLevel):
			t.Errorf("returned the (%v) error when expecting (%v)", e, ErrInvalidLevel)
		}
	})

	t.Run("error creating the formatter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("dummy message")
		config := NewMockConfig(ctrl)
		config.EXPECT().Populate("", gomock.AssignableToTypeOf(&consoleStreamConfig{})).DoAndReturn(
			func(_ string, data *consoleStreamConfig, _ ...bool) (interface{}, error) {
				data.Format = JSONFormatterFormat
				data.Level = "fatal"
				return data, nil
			},
		).Times(1)
		formatterFactory := NewMockFormatterFactory(ctrl)
		formatterFactory.EXPECT().Create(JSONFormatterFormat).Return(nil, expected).Times(1)

		sut, _ := NewConsoleStreamStrategy(formatterFactory)

		stream, e := sut.Create(config)
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
		config.EXPECT().Populate("", gomock.AssignableToTypeOf(&consoleStreamConfig{})).DoAndReturn(
			func(_ string, data *consoleStreamConfig, _ ...bool) (interface{}, error) {
				data.Format = JSONFormatterFormat
				data.Level = "fatal"
				data.Channels = []interface{}{"channel1", "channel2"}
				return data, nil
			},
		).Times(1)
		formatter := NewMockFormatter(ctrl)
		formatterFactory := NewMockFormatterFactory(ctrl)
		formatterFactory.EXPECT().Create(JSONFormatterFormat).Return(formatter, nil).Times(1)

		sut, _ := NewConsoleStreamStrategy(formatterFactory)

		stream, e := sut.Create(config)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case stream == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := stream.(type) {
			case *ConsoleStream:
				switch {
				case s.level != FATAL:
					t.Error("didn't created a stream with the correct level")
				case len(s.channels) != 2:
					t.Error("didn't created a stream with the correct channel list")
				}
			default:
				t.Error("didn't returned a new console stream")
			}
		}
	})
}
