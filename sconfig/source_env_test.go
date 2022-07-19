package sconfig

import (
	"os"
	"reflect"
	"testing"
)

func Test_NewSourceEnv(t *testing.T) {
	t.Run("with empty mappings", func(t *testing.T) {
		sut, e := newSourceEnv(map[string]string{})
		switch {
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch s := sut.(type) {
			case *sourceEnv:
				switch {
				case s.mutex == nil:
					t.Error("didn't created the access mutex")
				case !reflect.DeepEqual(s.partial, Partial{}):
					t.Error("didn't loaded the content correctly")
				}
			default:
				t.Error("didn't returned a new environment source")
			}
		}
	})

	t.Run("with root mappings", func(t *testing.T) {
		env := "senv"
		value := "value"
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, "") }()

		expected := Partial{"id": value}

		sut, e := newSourceEnv(map[string]string{env: "id"})
		switch {
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch s := sut.(type) {
			case *sourceEnv:
				switch {
				case s.mutex == nil:
					t.Error("didn't created the access mutex")
				case !reflect.DeepEqual(s.partial, expected):
					t.Errorf("didn't loaded the content correctly having (%v) when expecting (%v)", s.partial, expected)
				}
			default:
				t.Error("didn't returned a new environment source")
			}
		}
	})

	t.Run("with multi-level mappings", func(t *testing.T) {
		env := "senv"
		value := "value"
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, "") }()

		expected := Partial{"root": Partial{"node": value}}

		sut, e := newSourceEnv(map[string]string{env: "root.node"})
		switch {
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch s := sut.(type) {
			case *sourceEnv:
				switch {
				case s.mutex == nil:
					t.Error("didn't created the access mutex")
				case !reflect.DeepEqual(s.partial, expected):
					t.Errorf("didn't loaded the content correctly having (%v) when expecting (%v)", s.partial, expected)
				}
			default:
				t.Error("didn't returned a new environment source")
			}
		}
	})

	t.Run("with multi-level mapping", func(t *testing.T) {
		_ = os.Setenv("env1", "value")
		defer func() {
			_ = os.Setenv("env1", "")
		}()

		expected := Partial{"root": Partial{"node": "value"}}

		sut, e := newSourceEnv(map[string]string{"env1": "root.node"})
		switch {
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch s := sut.(type) {
			case *sourceEnv:
				switch {
				case s.mutex == nil:
					t.Error("didn't created the access mutex")
				case !reflect.DeepEqual(s.partial, expected):
					t.Errorf("didn't loaded the content correctly having (%v) when expecting (%v)", s.partial, expected)
				}
			default:
				t.Error("didn't returned a new environment source")
			}
		}
	})
}
