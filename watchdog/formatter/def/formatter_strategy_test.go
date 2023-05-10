package def

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/watchdog"
)

func Test_FormatterStrategy_Accept(t *testing.T) {
	t.Run("don't accept if config is a nil pointer", func(t *testing.T) {
		if (&FormatterStrategy{}).Accept(nil) {
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

		if (&FormatterStrategy{}).Accept(config) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept on invalid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		config.EXPECT().Populate("", gomock.AssignableToTypeOf(&struct{ Type string }{})).DoAndReturn(
			func(_ string, data *struct{ Type string }, _ ...bool) (interface{}, error) {
				data.Type = watchdog.UnknownLogFormatterType
				return data, nil
			},
		).Times(1)

		if (&FormatterStrategy{}).Accept(config) {
			t.Error("returned true")
		}
	})

	t.Run("accept on valid type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)
		config.EXPECT().Populate("", gomock.AssignableToTypeOf(&struct{ Type string }{})).DoAndReturn(
			func(_ string, data *struct{ Type string }, _ ...bool) (interface{}, error) {
				data.Type = Type
				return data, nil
			},
		).Times(1)

		if !(&FormatterStrategy{}).Accept(config) {
			t.Error("returned false")
		}
	})
}

func Test_FormatterStrategy_Create(t *testing.T) {
	t.Run("new formatter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config := NewMockConfig(ctrl)

		stream, e := (&FormatterStrategy{}).Create(config)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case stream == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch stream.(type) {
			case *Formatter:
			default:
				t.Error("didn't returned a log formatter")
			}
		}
	})
}
