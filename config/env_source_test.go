package config

import (
	"os"
	"reflect"
	"testing"
)

func Test_NewEnvSource(t *testing.T) {
	t.Run("with empty mappings", func(t *testing.T) {
		sut, e := NewEnvSource(map[string]string{})
		switch {
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch {
			case sut.mutex == nil:
				t.Error("didn't created the access mutex")
			case !reflect.DeepEqual(sut.config, Config{}):
				t.Error("didn't loaded the content correctly")
			}
		}
	})

	t.Run("with root mappings", func(t *testing.T) {
		env := "env"
		value := "value"
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, "") }()

		expected := Config{"id": value}

		sut, e := NewEnvSource(map[string]string{env: "id"})
		switch {
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch {
			case sut.mutex == nil:
				t.Error("didn't created the access mutex")
			case !reflect.DeepEqual(sut.config, expected):
				t.Errorf("didn't loaded the content correctly having (%v) when expecting (%v)", sut.config, expected)
			}
		}
	})

	t.Run("with multi-level mappings", func(t *testing.T) {
		env := "env"
		value := "value"
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, "") }()

		expected := Config{"root": Config{"node": value}}

		sut, e := NewEnvSource(map[string]string{env: "root.node"})
		switch {
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch {
			case sut.mutex == nil:
				t.Error("didn't created the access mutex")
			case !reflect.DeepEqual(sut.config, expected):
				t.Errorf("didn't loaded the content correctly having (%v) when expecting (%v)", sut.config, expected)
			}
		}
	})

	t.Run("with multi-level mapping", func(t *testing.T) {
		_ = os.Setenv("env1", "value")
		defer func() {
			_ = os.Setenv("env1", "")
		}()

		expected := Config{"root": Config{"node": "value"}}

		sut, e := NewEnvSource(map[string]string{"env1": "root.node"})
		switch {
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch {
			case sut.mutex == nil:
				t.Error("didn't created the access mutex")
			case !reflect.DeepEqual(sut.config, expected):
				t.Errorf("didn't loaded the content correctly having (%v) when expecting (%v)", sut.config, expected)
			}
		}
	})
}
