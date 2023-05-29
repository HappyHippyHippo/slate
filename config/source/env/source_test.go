package env

import (
	"os"
	"reflect"
	"testing"

	"github.com/happyhippyhippo/slate/config"
)

func Test_NewSource(t *testing.T) {
	t.Run("with empty mappings", func(t *testing.T) {
		sut, e := NewSource(map[string]string{})
		switch {
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch {
			case sut.Mutex == nil:
				t.Error("didn't created the access mutex")
			case !reflect.DeepEqual(sut.Partial, config.Partial{}):
				t.Error("didn't loaded the content correctly")
			}
		}
	})

	t.Run("with empty environment", func(t *testing.T) {
		env := "env"

		expected := config.Partial{}

		sut, e := NewSource(map[string]string{env: "id"})
		switch {
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch {
			case sut.Mutex == nil:
				t.Error("didn't created the access mutex")
			case !reflect.DeepEqual(sut.Partial, expected):
				t.Errorf("didn't loaded the content correctly having (%v) when expecting (%v)", sut.Partial, expected)
			}
		}
	})

	t.Run("with root mappings", func(t *testing.T) {
		env := "env"
		value := "value"
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, "") }()

		expected := config.Partial{"id": value}

		sut, e := NewSource(map[string]string{env: "id"})
		switch {
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch {
			case sut.Mutex == nil:
				t.Error("didn't created the access mutex")
			case !reflect.DeepEqual(sut.Partial, expected):
				t.Errorf("didn't loaded the content correctly having (%v) when expecting (%v)", sut.Partial, expected)
			}
		}
	})

	t.Run("with multi-level mappings", func(t *testing.T) {
		env := "env"
		value := "value"
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, "") }()

		expected := config.Partial{"root": config.Partial{"node": value}}

		sut, e := NewSource(map[string]string{env: "root.node"})
		switch {
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch {
			case sut.Mutex == nil:
				t.Error("didn't created the access mutex")
			case !reflect.DeepEqual(sut.Partial, expected):
				t.Errorf("didn't loaded the content correctly having (%v) when expecting (%v)", sut.Partial, expected)
			}
		}
	})

	t.Run("with multi-level mapping", func(t *testing.T) {
		_ = os.Setenv("env1", "value")
		defer func() {
			_ = os.Setenv("env1", "")
		}()

		expected := config.Partial{"root": config.Partial{"node": "value"}}

		sut, e := NewSource(map[string]string{"env1": "root.node"})
		switch {
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch {
			case sut.Mutex == nil:
				t.Error("didn't created the access mutex")
			case !reflect.DeepEqual(sut.Partial, expected):
				t.Errorf("didn't loaded the content correctly having (%v) when expecting (%v)", sut.Partial, expected)
			}
		}
	})

	t.Run("with multi-level mapping (deeper)", func(t *testing.T) {
		_ = os.Setenv("env1", "value")
		defer func() {
			_ = os.Setenv("env1", "")
		}()

		expected := config.Partial{"root": config.Partial{"node1": config.Partial{"node2": "value"}}}

		sut, e := NewSource(map[string]string{"env1": "root.node1.node2"})
		switch {
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		default:
			switch {
			case sut.Mutex == nil:
				t.Error("didn't created the access mutex")
			case !reflect.DeepEqual(sut.Partial, expected):
				t.Errorf("didn't loaded the content correctly having (%v) when expecting (%v)", sut.Partial, expected)
			}
		}
	})
}
