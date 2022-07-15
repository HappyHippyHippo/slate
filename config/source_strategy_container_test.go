package config

import (
	"errors"
	"github.com/happyhippyhippo/slate/err"
	"reflect"
	"testing"
)

func Test_SourceStrategyContainer_Accept(t *testing.T) {
	t.Run("accept only env type", func(t *testing.T) {
		scenarios := []struct {
			sourceType string
			expected   bool
		}{
			{ // _test env type
				sourceType: SourceTypeContainer,
				expected:   true,
			},
			{ // _test non-env type
				sourceType: SourceTypeUnknown,
				expected:   false,
			},
		}

		for _, scenario := range scenarios {
			sut := &sourceStrategyContainer{
				partials: []IConfig{},
			}
			if check := sut.Accept(scenario.sourceType); check != scenario.expected {
				t.Errorf("for the type (%s), returned (%v)", scenario.sourceType, check)
			}
		}
	})
}

func Test_SourceStrategyContainer_AcceptFromConfig(t *testing.T) {
	t.Run("don't accept on invalid config pointer", func(t *testing.T) {
		if (&sourceStrategyContainer{
			partials: []IConfig{},
		}).AcceptFromConfig(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		if (&sourceStrategyContainer{
			partials: []IConfig{},
		}).AcceptFromConfig(&Partial{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		if (&sourceStrategyContainer{
			partials: []IConfig{},
		}).AcceptFromConfig(&Partial{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not env", func(t *testing.T) {
		if (&sourceStrategyContainer{
			partials: []IConfig{},
		}).AcceptFromConfig(&Partial{"type": SourceTypeUnknown}) {
			t.Error("returned true")
		}
	})
}

func Test_SourceStrategyContainer_Create(t *testing.T) {
	t.Run("create the source with a single partial", func(t *testing.T) {
		value := Partial{"key": "value"}

		expected := value
		sut := &sourceStrategyContainer{partials: []IConfig{&value}}

		src, e := sut.Create()
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceContainer:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new env src")
			}
		}
	})

	t.Run("create the source with multiple partials", func(t *testing.T) {
		value1 := Partial{"key1": "value 1"}
		value2 := Partial{"key2": "value 2"}

		expected := Partial{"key1": "value 1", "key2": "value 2"}
		sut := &sourceStrategyContainer{partials: []IConfig{&value1, &value2}}

		src, e := sut.Create()
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceContainer:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new env src")
			}
		}
	})
}

func Test_SourceStrategyContainer_CreateFromConfig(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		src, e := (&sourceStrategyContainer{
			partials: []IConfig{},
		}).CreateFromConfig(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrNilPointer)
		}
	})

	t.Run("create the source with a single partial", func(t *testing.T) {
		value := Partial{"key": "value"}

		expected := value
		sut := &sourceStrategyContainer{partials: []IConfig{&value}}

		src, e := sut.CreateFromConfig(&Partial{})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceContainer:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new env src")
			}
		}
	})

	t.Run("create the source with multiple partial", func(t *testing.T) {
		value1 := Partial{"key1": "value 1"}
		value2 := Partial{"key2": "value 2"}

		expected := Partial{"key1": "value 1", "key2": "value 2"}
		sut := &sourceStrategyContainer{partials: []IConfig{&value1, &value2}}

		src, e := sut.CreateFromConfig(&Partial{})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceContainer:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new env src")
			}
		}
	})
}
