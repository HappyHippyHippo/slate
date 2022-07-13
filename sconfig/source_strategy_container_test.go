package sconfig

import (
	"errors"
	"github.com/happyhippyhippo/slate/serror"
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
			strategy := &sourceStrategyContainer{
				partials: []IConfig{},
			}
			if check := strategy.Accept(scenario.sourceType); check != scenario.expected {
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
		strategy := &sourceStrategyContainer{partials: []IConfig{&value}}

		src, err := strategy.Create()
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
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
		strategy := &sourceStrategyContainer{partials: []IConfig{&value1, &value2}}

		src, err := strategy.Create()
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
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
		src, err := (&sourceStrategyContainer{
			partials: []IConfig{},
		}).CreateFromConfig(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("create the source with a single partial", func(t *testing.T) {
		value := Partial{"key": "value"}

		expected := value
		strategy := &sourceStrategyContainer{partials: []IConfig{&value}}

		src, err := strategy.CreateFromConfig(&Partial{})
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
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
		strategy := &sourceStrategyContainer{partials: []IConfig{&value1, &value2}}

		src, err := strategy.CreateFromConfig(&Partial{})
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
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
