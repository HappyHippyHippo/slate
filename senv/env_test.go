package senv

import (
	"os"
	"reflect"
	"testing"
)

func Test_Bool(t *testing.T) {
	env := "__ENV_VARIABLE__"

	t.Run("no senv value", func(t *testing.T) {
		value := ""
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		t.Run("should return the default value (true)", func(t *testing.T) {
			if check := Bool(env, false); check != false {
				t.Error("didn't returned the default value of false")
			}
		})
		t.Run("should return the default value (false)", func(t *testing.T) {
			if check := Bool(env, true); check != true {
				t.Error("didn't returned the default value of true")
			}
		})
	})

	t.Run("senv value with the string 'true'", func(t *testing.T) {
		value := "true"
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := Bool(env, false); check != true {
			t.Error("didn't returned the defined environment value of true")
		}
	})

	t.Run("senv value with the string 'TRUE'", func(t *testing.T) {
		value := "TRUE"
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := Bool(env, false); check != true {
			t.Error("didn't returned the defined environment value of true")
		}
	})

	t.Run("senv value with the string '1'", func(t *testing.T) {
		value := "1"
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := Bool(env, false); check != true {
			t.Error("didn't returned the defined environment value of true")
		}
	})
}

func Test_Int(t *testing.T) {
	def := 123
	env := "__ENV_VARIABLE__"

	t.Run("no senv value", func(t *testing.T) {
		value := ""
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := Int(env, def); check != def {
			t.Errorf("returned the (%v) string value, instead of the expected (%v) default value", check, def)
		}
	})

	t.Run("senv value with an invalid string", func(t *testing.T) {
		value := "abc"
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := Int(env, def); check != def {
			t.Errorf("returned the (%v) string value, instead of the expected (%v)", check, def)
		}
	})

	t.Run("senv value with a valid string", func(t *testing.T) {
		value := "456"
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := Int(env, def); check != 456 {
			t.Errorf("returned the (%v) string value, instead of the expected (%v) default value", check, value)
		}
	})
}

func Test_String(t *testing.T) {
	def := "default"
	env := "__ENV_VARIABLE__"

	t.Run("no senv value", func(t *testing.T) {
		value := ""
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := String(env, def); check != def {
			t.Errorf("returned the (%v) string value, instead of the expected (%v) default value", check, def)
		}
	})

	t.Run("senv value with a string", func(t *testing.T) {
		value := "senv string"
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := String(env, def); check != value {
			t.Errorf("returned the (%v) string value, instead of the expected (%v) default value", check, value)
		}
	})
}

func Test_List(t *testing.T) {
	def := []string{"default"}
	env := "__ENV_VARIABLE__"

	t.Run("no senv value", func(t *testing.T) {
		value := ""
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := List(env, def); !reflect.DeepEqual(check, def) {
			t.Errorf("returned the (%v) string value, instead of the expected (%v) default value", check, def)
		}
	})

	t.Run("senv value with a single string", func(t *testing.T) {
		value := "string"
		cur := os.Getenv(env)
		_ = os.Setenv(env, value)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := List(env, def); !reflect.DeepEqual(check, []string{value}) {
			t.Errorf("returned the (%v) string value, instead of the expected (%v) default value", check, value)
		}
	})

	t.Run("senv value with multi strings", func(t *testing.T) {
		value1 := "string1"
		value2 := "string2"
		expected := []string{value1, value2}
		cur := os.Getenv(env)
		_ = os.Setenv(env, value1+","+value2)
		defer func() { _ = os.Setenv(env, cur) }()

		if check := List(env, def); !reflect.DeepEqual(check, expected) {
			t.Errorf("returned the (%v) string value, instead of the expected (%v) default value", check, expected)
		}
	})
}
