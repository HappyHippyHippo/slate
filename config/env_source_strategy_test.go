package config

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/happyhippyhippo/slate/err"
)

func Test_EnvSourceStrategy_Accept(t *testing.T) {
	t.Run("don't accept on invalid config pointer", func(t *testing.T) {
		if (&EnvSourceStrategy{}).Accept(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		if (&EnvSourceStrategy{}).Accept(&Config{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		if (&EnvSourceStrategy{}).Accept(&Config{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not env", func(t *testing.T) {
		if (&EnvSourceStrategy{}).Accept(&Config{"type": UnknownSourceType}) {
			t.Error("returned true")
		}
	})
}

func Test_EnvSourceStrategy_Create(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		src, e := (&EnvSourceStrategy{}).Create(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("non-map mappings", func(t *testing.T) {
		src, e := (&EnvSourceStrategy{}).Create(&Config{"mappings": 123})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("non-string key map mappings", func(t *testing.T) {
		src, e := (&EnvSourceStrategy{}).Create(&Config{"mappings": Config{1: "value"}})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("non-string value map mappings", func(t *testing.T) {
		src, e := (&EnvSourceStrategy{}).Create(&Config{"mappings": Config{"key": 1}})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("create the source", func(t *testing.T) {
		env := "env"
		value := "value"
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, "") }()

		path := "root"
		expected := Config{path: value}

		src, e := (&EnvSourceStrategy{}).Create(&Config{"mappings": Config{env: path}})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *EnvSource:
				if !reflect.DeepEqual(s.config, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new env src")
			}
		}
	})

	t.Run("no mappings on config", func(t *testing.T) {
		expected := Config{}

		src, e := (&EnvSourceStrategy{}).Create(&Config{})
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case src == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch s := src.(type) {
			case *EnvSource:
				if !reflect.DeepEqual(s.config, expected) {
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new env src")
			}
		}
	})
}
