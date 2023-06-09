package simple

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
)

func Test_FormatterStrategy_Accept(t *testing.T) {
	t.Run("don't accept if config is a nil pointer", func(t *testing.T) {
		if (&FormatterStrategy{}).Accept(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept on type retrieval error", func(t *testing.T) {
		if (&FormatterStrategy{}).Accept(&config.Partial{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept on invalid type", func(t *testing.T) {
		if (&FormatterStrategy{}).Accept(&config.Partial{"type": "invalid"}) {
			t.Error("returned true")
		}
	})

	t.Run("accept on valid type", func(t *testing.T) {
		if !(&FormatterStrategy{}).Accept(&config.Partial{"type": "simple"}) {
			t.Error("returned false")
		}
	})
}

func Test_FormatterStrategy_Create(t *testing.T) {
	t.Run("new formatter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		stream, e := (&FormatterStrategy{}).Create(&config.Partial{})
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
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
