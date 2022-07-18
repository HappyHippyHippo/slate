package rdb

import (
	"errors"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/err"
	"gorm.io/driver/sqlite"
	"strings"
	"testing"
)

func Test_DialectStrategySqlite_Accept(t *testing.T) {
	t.Run("refuse if the dialect name is not mysql", func(t *testing.T) {
		if (&dialectStrategySqlite{}).Accept("mysql") == true {
			t.Error("returned true")
		}
	})

	t.Run("accept if the dialect name is mysql", func(t *testing.T) {
		if (&dialectStrategySqlite{}).Accept("sQlItE") == false {
			t.Error("returned false")
		}
	})
}

func Test_DialectStrategySqlite_Get(t *testing.T) {
	t.Run("invalid host value on connection configuration", func(t *testing.T) {
		cfg := &config.Partial{
			"dialect": "sqlite",
			"host":    123,
		}

		dialect, e := (&dialectStrategySqlite{}).Get(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, err.ErrConversion):
			t.Errorf("returned the (%v) error when expected (%v)", e, err.ErrConversion)
		}
	})

	t.Run("invalid params value on connection configuration", func(t *testing.T) {
		cfg := &config.Partial{
			"dialect": "sqlite",
			"host":    "host",
			"params":  123,
		}

		dialect, e := (&dialectStrategySqlite{}).Get(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, err.ErrConversion):
			t.Errorf("returned the (%v) error when expected (%v)", e, err.ErrConversion)
		}
	})

	t.Run("valid connection", func(t *testing.T) {
		expected := "file.db"
		cfg := &config.Partial{
			"dialect": "sqlite",
			"host":    "file.db",
		}

		dialect, e := (&dialectStrategySqlite{}).Get(cfg)
		switch {
		case e != nil:
			t.Errorf("return the unexpected error : (%v)", e)
		case dialect == nil:
			t.Error("didn't return the expected valid dialect instance")
		default:
			switch d := dialect.(type) {
			case *sqlite.Dialector:
				if check := d.DSN; check != expected {
					t.Errorf("dialect composed with the DSN (%v) when expected to be (%v)", check, expected)
				}
			default:
				t.Error("didn't return the expected sqlite dialect instance")
			}
		}
	})

	t.Run("valid connection with extra params", func(t *testing.T) {
		expectedPrefix := ":memory:"
		cfg := &config.Partial{
			"dialect": "sqlite",
			"host":    expectedPrefix,
			"params": config.Partial{
				"param1": "value1",
				"param2": "value2",
			},
		}

		dialect, e := (&dialectStrategySqlite{}).Get(cfg)
		switch {
		case e != nil:
			t.Errorf("return the unexpected error : (%v)", e)
		case dialect == nil:
			t.Error("didn't return the expected valid dialect instance")
		default:
			switch d := dialect.(type) {
			case *sqlite.Dialector:
				dsn := d.DSN
				switch {
				case !strings.HasPrefix(dsn, expectedPrefix):
					t.Errorf("dialect composed with the DSN prefix of (%v) when expected to be (%v)", dsn, expectedPrefix)
				case !strings.Contains(dsn, "&param1=value1"):
					t.Errorf("missing dialect composed with the DSN params (%v)", "&param1=value1")
				case !strings.Contains(dsn, "&param2=value2"):
					t.Errorf("missing dialect composed with the DSN params (%v)", "&param2=value2")
				}
			default:
				t.Error("didn't return the expected sqlite dialect instance")
			}
		}
	})
}
