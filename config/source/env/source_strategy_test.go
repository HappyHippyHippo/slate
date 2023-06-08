package env

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
)

func Test_NewSourceStrategy(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		if NewSourceStrategy() == nil {
			t.Error("didn't returned the expected reference")
		}
	})
}

func Test_SourceStrategy_Accept(t *testing.T) {
	t.Run("don't accept on invalid config pointer", func(t *testing.T) {
		if (&SourceStrategy{}).Accept(nil) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is missing", func(t *testing.T) {
		if (&SourceStrategy{}).Accept(config.Partial{}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not a string", func(t *testing.T) {
		if (&SourceStrategy{}).Accept(config.Partial{"type": 123}) {
			t.Error("returned true")
		}
	})

	t.Run("don't accept if type is not env", func(t *testing.T) {
		if (&SourceStrategy{}).Accept(config.Partial{"type": config.UnknownSource}) {
			t.Error("returned true")
		}
	})
}

func Test_SourceStrategy_Create(t *testing.T) {
	t.Run("error on nil config pointer", func(t *testing.T) {
		src, e := (&SourceStrategy{}).Create(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("non-map mappings", func(t *testing.T) {
		src, e := (&SourceStrategy{}).Create(config.Partial{"mappings": 123})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("non-string key map mappings", func(t *testing.T) {
		src, e := (&SourceStrategy{}).Create(config.Partial{"mappings": config.Partial{1: "value"}})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("non-string value map mappings", func(t *testing.T) {
		src, e := (&SourceStrategy{}).Create(config.Partial{"mappings": config.Partial{"key": 1}})
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("create the source", func(t *testing.T) {
		env := "env"
		value := "value"
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, "") }()

		path := "root"
		expected := config.Partial{path: value}

		src, e := (&SourceStrategy{}).Create(config.Partial{"mappings": config.Partial{env: path}})
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

	t.Run("no mappings on config", func(t *testing.T) {
		expected := config.Partial{}

		src, e := (&SourceStrategy{}).Create(config.Partial{})
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
