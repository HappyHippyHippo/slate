package sconfig

import (
	"errors"
	"github.com/happyhippyhippo/slate/serror"
	"os"
	"reflect"
	"testing"
)

func Test_SourceStrategyEnv_Accept(t *testing.T) {
	t.Run("accept only env type", func(t *testing.T) {
		scenarios := []struct {
			sourceType string
			expected   bool
		}{
			{ // _test env type
				sourceType: SourceTypeEnv,
				expected:   true,
			},
			{ // _test non-env type
				sourceType: SourceTypeUnknown,
				expected:   false,
			},
		}

		for _, scenario := range scenarios {
			strategy := &sourceStrategyEnv{}
			if check := strategy.Accept(scenario.sourceType); check != scenario.expected {
				t.Errorf("for the type (%s), returned (%v)", scenario.sourceType, check)
			}
		}
	})
}

func Test_SourceStrategyEnv_AcceptFromConfig(t *testing.T) {
	t.Run("don't accept on invalid config pointer", func(t *testing.T) {
		if (&sourceStrategyEnv{}).AcceptFromConfig(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		if (&sourceStrategyEnv{}).AcceptFromConfig(&Partial{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		if (&sourceStrategyEnv{}).AcceptFromConfig(&Partial{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not env", func(t *testing.T) {
		if (&sourceStrategyEnv{}).AcceptFromConfig(&Partial{"type": SourceTypeUnknown}) {
			t.Error("returned true")
		}
	})
}

func Test_SourceStrategyEnv_Create(t *testing.T) {
	t.Run("missing mappings", func(t *testing.T) {
		strategy := &sourceStrategyEnv{}

		src, err := strategy.Create()
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("non-map mappings", func(t *testing.T) {
		strategy := &sourceStrategyEnv{}

		src, err := strategy.Create(123)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("create the source with a map mappings", func(t *testing.T) {
		env := "env"
		value := "value"
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, "") }()

		path := "root"
		expected := Partial{path: value}
		strategy := &sourceStrategyEnv{}

		src, err := strategy.Create(map[string]string{env: path})
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceEnv:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new env src")
			}
		}
	})
}

func Test_SourceStrategyEnv_CreateFromConfig(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		src, err := (&sourceStrategyEnv{}).CreateFromConfig(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("non-map mappings", func(t *testing.T) {
		src, err := (&sourceStrategyEnv{}).CreateFromConfig(&Partial{"mappings": 123})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrInvalidConfigSourcePartial):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrInvalidConfigSourcePartial)
		}
	})

	t.Run("create the source", func(t *testing.T) {
		env := "env"
		value := "value"
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, "") }()

		path := "root"
		expected := Partial{path: value}

		src, err := (&sourceStrategyEnv{}).CreateFromConfig(&Partial{"mappings": Partial{env: path}})
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceEnv:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new env src")
			}
		}
	})

	t.Run("no mappings on config", func(t *testing.T) {
		expected := Partial{}

		src, err := (&sourceStrategyEnv{}).CreateFromConfig(&Partial{})
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceEnv:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new env src")
			}
		}
	})
}
