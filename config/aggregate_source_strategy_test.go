package config

import (
	"reflect"
	"testing"
)

func Test_SourceStrategyAggregate_Accept(t *testing.T) {
	t.Run("don't accept on invalid config pointer", func(t *testing.T) {
		if (&AggregateSourceStrategy{
			configs: []IConfig{},
		}).Accept(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		if (&AggregateSourceStrategy{
			configs: []IConfig{},
		}).Accept(&Config{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		if (&AggregateSourceStrategy{
			configs: []IConfig{},
		}).Accept(&Config{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not env", func(t *testing.T) {
		if (&AggregateSourceStrategy{
			configs: []IConfig{},
		}).Accept(&Config{"type": UnknownSourceType}) {
			t.Error("returned true")
		}
	})
}

func Test_AggregateSourceStrategy_Create(t *testing.T) {
	t.Run("accept nil config pointer", func(t *testing.T) {
		sut := &AggregateSourceStrategy{configs: []IConfig{}}
		src, e := sut.Create(&Config{})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch src.(type) {
			case *AggregateSource:
			default:
				t.Error("didn't returned a new env src")
			}
		}
	})

	t.Run("create the source with a single config", func(t *testing.T) {
		value := Config{"key": "value"}

		expected := value
		sut := &AggregateSourceStrategy{configs: []IConfig{&value}}

		src, e := sut.Create(&Config{})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *AggregateSource:
				if !reflect.DeepEqual(s.config, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new env src")
			}
		}
	})

	t.Run("create the source with multiple config", func(t *testing.T) {
		value1 := Config{"key1": "value 1"}
		value2 := Config{"key2": "value 2"}

		expected := Config{"key1": "value 1", "key2": "value 2"}
		sut := &AggregateSourceStrategy{configs: []IConfig{&value1, &value2}}

		src, e := sut.Create(&Config{})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *AggregateSource:
				if !reflect.DeepEqual(s.config, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new env src")
			}
		}
	})
}
