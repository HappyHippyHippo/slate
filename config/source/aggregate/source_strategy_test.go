package aggregate

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
)

func Test_SourceStrategy_Accept(t *testing.T) {
	t.Run("don't accept on invalid config pointer", func(t *testing.T) {
		if (&SourceStrategy{
			sources: []config.Source{},
		}).Accept(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		if (&SourceStrategy{
			sources: []config.Source{},
		}).Accept(&config.Partial{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		if (&SourceStrategy{
			sources: []config.Source{},
		}).Accept(&config.Partial{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not env", func(t *testing.T) {
		if (&SourceStrategy{
			sources: []config.Source{},
		}).Accept(&config.Partial{"type": config.UnknownSource}) {
			t.Error("returned true")
		}
	})
}

func Test_SourceStrategy_Create(t *testing.T) {
	t.Run("accept nil config pointer", func(t *testing.T) {
		sut := &SourceStrategy{sources: []config.Source{}}
		src, e := sut.Create(&config.Partial{})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch src.(type) {
			case *Source:
			default:
				t.Error("didn't returned a new env src")
			}
		}
	})

	t.Run("create the source with a single config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		value := config.Partial{"key": "value"}
		source := NewMockSource(ctrl)
		source.EXPECT().Get("", config.Partial{}).Return(value, nil).Times(1)
		sut := &SourceStrategy{sources: []config.Source{source}}

		src, e := sut.Create(&config.Partial{})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *Source:
				if !reflect.DeepEqual(s.Partial, value) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new env src")
			}
		}
	})

	t.Run("create the source with multiple config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		value1 := config.Partial{"key1": "value 1"}
		value2 := config.Partial{"key2": "value 2"}

		expected := config.Partial{"key1": "value 1", "key2": "value 2"}
		source1 := NewMockSource(ctrl)
		source1.EXPECT().Get("", config.Partial{}).Return(value1, nil).Times(1)
		source2 := NewMockSource(ctrl)
		source2.EXPECT().Get("", config.Partial{}).Return(value2, nil).Times(1)
		sut := &SourceStrategy{sources: []config.Source{source1, source2}}

		src, e := sut.Create(&config.Partial{})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *Source:
				if !reflect.DeepEqual(s.Partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new env src")
			}
		}
	})
}
