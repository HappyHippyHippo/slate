package sconfig

import (
	"errors"
	"github.com/happyhippyhippo/slate/serr"
	"os"
	"reflect"
	"testing"
)

func Test_SourceStrategyEnv_Accept(t *testing.T) {
	t.Run("accept only senv type", func(t *testing.T) {
		scenarios := []struct {
			sourceType string
			expected   bool
		}{
			{ // _test senv type
				sourceType: SourceTypeEnv,
				expected:   true,
			},
			{ // _test non-senv type
				sourceType: SourceTypeUnknown,
				expected:   false,
			},
		}

		for _, scenario := range scenarios {
			sut := &sourceStrategyEnv{}
			if check := sut.Accept(scenario.sourceType); check != scenario.expected {
				t.Errorf("for the type (%s), returned (%v)", scenario.sourceType, check)
			}
		}
	})
}

func Test_SourceStrategyEnv_AcceptFromConfig(t *testing.T) {
	t.Run("don't accept on invalid sconfig pointer", func(t *testing.T) {
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

	t.Run("don't accept if type is not senv", func(t *testing.T) {
		if (&sourceStrategyEnv{}).AcceptFromConfig(&Partial{"type": SourceTypeUnknown}) {
			t.Error("returned true")
		}
	})
}

func Test_SourceStrategyEnv_Create(t *testing.T) {
	t.Run("missing mappings", func(t *testing.T) {
		src, e := (&sourceStrategyEnv{}).Create()
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("non-map mappings", func(t *testing.T) {
		src, e := (&sourceStrategyEnv{}).Create(123)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("create the source with a map mappings", func(t *testing.T) {
		env := "senv"
		value := "value"
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, "") }()

		path := "root"
		expected := Partial{path: value}

		src, e := (&sourceStrategyEnv{}).Create(map[string]string{env: path})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceEnv:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new senv src")
			}
		}
	})
}

func Test_SourceStrategyEnv_CreateFromConfig(t *testing.T) {
	t.Run("error on nil sconfig pointer", func(t *testing.T) {
		src, e := (&sourceStrategyEnv{}).CreateFromConfig(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("non-map mappings", func(t *testing.T) {
		src, e := (&sourceStrategyEnv{}).CreateFromConfig(&Partial{"mappings": 123})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrInvalidConfigSourcePartial):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrInvalidConfigSourcePartial)
		}
	})

	t.Run("create the source", func(t *testing.T) {
		env := "senv"
		value := "value"
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, "") }()

		path := "root"
		expected := Partial{path: value}

		src, e := (&sourceStrategyEnv{}).CreateFromConfig(&Partial{"mappings": Partial{env: path}})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceEnv:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new senv src")
			}
		}
	})

	t.Run("no mappings on sconfig", func(t *testing.T) {
		expected := Partial{}

		src, e := (&sourceStrategyEnv{}).CreateFromConfig(&Partial{})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *sourceEnv:
				if !reflect.DeepEqual(s.partial, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new senv src")
			}
		}
	})
}
