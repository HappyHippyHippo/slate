package env

import (
	"os"
	"reflect"
	"testing"
)

func Test_Bool(t *testing.T) {
	env := "__ENV_VARIABLE__"

	t.Run("no environment value", func(t *testing.T) {
		value := ""
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		t.Run("should return the simple value (true)", func(t *testing.T) {
			if check := Bool(env, false); check != false {
				t.Error("returned true")
			}
		})
		t.Run("should return the simple value (false)", func(t *testing.T) {
			if check := Bool(env, true); check != true {
				t.Error("returned false")
			}
		})
	})

	t.Run("environment value with the string 'true'", func(t *testing.T) {
		value := "true"
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := Bool(env, false); check != true {
			t.Error("returned false")
		}
	})

	t.Run("environment value with the string 'TRUE'", func(t *testing.T) {
		value := "TRUE"
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := Bool(env, false); check != true {
			t.Error("returned false")
		}
	})

	t.Run("environment value with the string '1'", func(t *testing.T) {
		value := "1"
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := Bool(env, false); check != true {
			t.Error("returned false")
		}
	})
}

func Test_Int(t *testing.T) {
	def := 123
	env := "__ENV_VARIABLE__"

	t.Run("no environment value", func(t *testing.T) {
		value := ""
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := Int(env, def); check != def {
			t.Errorf("(%v) when expecting (%v)", check, def)
		}
	})

	t.Run("environment value with an invalid string", func(t *testing.T) {
		value := "abc"
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := Int(env, def); check != def {
			t.Errorf("(%v) when expecting (%v)", check, def)
		}
	})

	t.Run("environment value with a valid string", func(t *testing.T) {
		value := "456"
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := Int(env, def); check != 456 {
			t.Errorf("(%v) when expecting (%v)", check, 456)
		}
	})
}

func Test_String(t *testing.T) {
	def := "simple"
	env := "__ENV_VARIABLE__"

	t.Run("no environment value", func(t *testing.T) {
		value := ""
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := String(env, def); check != def {
			t.Errorf("(%v) when expecting (%v)", check, def)
		}
	})

	t.Run("environment value with a string", func(t *testing.T) {
		value := "env string"
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := String(env, def); check != value {
			t.Errorf("(%v) when expecting (%v)", check, value)
		}
	})
}

func Test_List(t *testing.T) {
	def := []string{"simple"}
	env := "__ENV_VARIABLE__"

	t.Run("no environment value", func(t *testing.T) {
		value := ""
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := List(env, def); !reflect.DeepEqual(check, def) {
			t.Errorf("(%v) when expecting (%v)", check, def)
		}
	})

	t.Run("environment value with a single string", func(t *testing.T) {
		value := "string"
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := List(env, def); !reflect.DeepEqual(check, []string{value}) {
			t.Errorf("(%v) when expecting (%v)", check, value)
		}
	})

	t.Run("environment value with multi strings", func(t *testing.T) {
		value1 := "string1"
		value2 := "string2"
		expected := []string{value1, value2}
		cur := os.Getenv(env)
		_ = os.Setenv(env, value1+","+value2)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := List(env, def); !reflect.DeepEqual(check, expected) {
			t.Errorf("(%v) when expecting (%v)", check, expected)
		}
	})
}
