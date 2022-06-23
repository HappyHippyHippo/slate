package grdb

import (
	"errors"
	"github.com/happyhippyhippo/slate/gconfig"
	"github.com/happyhippyhippo/slate/gerror"
	"gorm.io/driver/sqlite"
	"strings"
	"testing"
)

func Test_DialectStrategySqlite_Accept(t *testing.T) {
	t.Run("refuse if the dialect name is not mysql", func(t *testing.T) {
		if (&DialectStrategySqlite{}).Accept("mysql") == true {
			t.Error("returned true")
		}
	})

	t.Run("accept if the dialect name is mysql", func(t *testing.T) {
		if (&DialectStrategySqlite{}).Accept("sQlItE") == false {
			t.Error("returned false")
		}
	})
}

func Test_DialectStrategySqlite_Get(t *testing.T) {
	t.Run("invalid host value on connection configuration", func(t *testing.T) {
		cfg := &gconfig.Partial{
			"dialect": "sqlite",
			"host":    123,
		}

		dialect, err := (&DialectStrategySqlite{}).Get(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("invalid params value on connection configuration", func(t *testing.T) {
		cfg := &gconfig.Partial{
			"dialect": "sqlite",
			"host":    "host",
			"params":  123,
		}

		dialect, err := (&DialectStrategySqlite{}).Get(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("valid connection", func(t *testing.T) {
		expected := "file.db"
		cfg := &gconfig.Partial{
			"dialect": "sqlite",
			"host":    "file.db",
		}

		dialect, err := (&DialectStrategySqlite{}).Get(cfg)
		switch {
		case err != nil:
			t.Errorf("return the unexpected error : (%v)", err)
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
		cfg := &gconfig.Partial{
			"dialect": "sqlite",
			"host":    expectedPrefix,
			"params": gconfig.Partial{
				"param1": "value1",
				"param2": "value2",
			},
		}

		dialect, err := (&DialectStrategySqlite{}).Get(cfg)
		switch {
		case err != nil:
			t.Errorf("return the unexpected error : (%v)", err)
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
