package aggregate

import (
	"reflect"
	"testing"

	"github.com/happyhippyhippo/slate/config"
)

func Test_SourceStrategy_Accept(t *testing.T) {
	t.Run("don't accept on invalid config pointer", func(t *testing.T) {
		if (&SourceStrategy{
			configs: []config.IConfig{},
		}).Accept(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		if (&SourceStrategy{
			configs: []config.IConfig{},
		}).Accept(&config.Config{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		if (&SourceStrategy{
			configs: []config.IConfig{},
		}).Accept(&config.Config{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not env", func(t *testing.T) {
		if (&SourceStrategy{
			configs: []config.IConfig{},
		}).Accept(&config.Config{"type": config.UnknownSourceType}) {
			t.Error("returned true")
		}
	})
}

func Test_SourceStrategy_Create(t *testing.T) {
	t.Run("accept nil config pointer", func(t *testing.T) {
		sut := &SourceStrategy{configs: []config.IConfig{}}
		src, e := sut.Create(&config.Config{})
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
		value := config.Config{"key": "value"}

		expected := value
		sut := &SourceStrategy{configs: []config.IConfig{&value}}

		src, e := sut.Create(&config.Config{})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *Source:
				if !reflect.DeepEqual(s.Config, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new env src")
			}
		}
	})

	t.Run("create the source with multiple config", func(t *testing.T) {
		value1 := config.Config{"key1": "value 1"}
		value2 := config.Config{"key2": "value 2"}

		expected := config.Config{"key1": "value 1", "key2": "value 2"}
		sut := &SourceStrategy{configs: []config.IConfig{&value1, &value2}}

		src, e := sut.Create(&config.Config{})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *Source:
				if !reflect.DeepEqual(s.Config, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new env src")
			}
		}
	})
}
