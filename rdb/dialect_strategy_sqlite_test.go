package rdb

import (
	"errors"
	"strings"
	"testing"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"gorm.io/driver/sqlite"
)

func Test_SqliteDialectStrategy_Accept(t *testing.T) {
	t.Run("refuse if no config", func(t *testing.T) {
		if (&SqliteDialectStrategy{}).Accept(nil) == true {
			t.Error("returned true")
		}
	})

	t.Run("refuse on config parsing", func(t *testing.T) {
		if (&SqliteDialectStrategy{}).Accept(&config.Config{"dialect": 123}) == true {
			t.Error("returned true")
		}
	})

	t.Run("refuse if the dialect name is not mysql", func(t *testing.T) {
		if (&SqliteDialectStrategy{}).Accept(&config.Config{"dialect": "mysql"}) == true {
			t.Error("returned true")
		}
	})

	t.Run("accept if the dialect name is mysql", func(t *testing.T) {
		if (&SqliteDialectStrategy{}).Accept(&config.Config{"dialect": "sQlItE"}) == false {
			t.Error("returned false")
		}
	})
}

func Test_SqliteDialectStrategy_Get(t *testing.T) {
	t.Run("error on nil config", func(t *testing.T) {
		dialect, e := (&SqliteDialectStrategy{}).Get(nil)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("invalid host value on connection configuration", func(t *testing.T) {
		cfg := &config.Config{
			"dialect": "sqlite",
			"host":    123,
		}

		dialect, e := (&SqliteDialectStrategy{}).Get(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("invalid params value on connection configuration", func(t *testing.T) {
		cfg := &config.Config{
			"dialect": "sqlite",
			"host":    "host",
			"params":  123,
		}

		dialect, e := (&SqliteDialectStrategy{}).Get(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("valid connection", func(t *testing.T) {
		expected := "file.db"
		cfg := &config.Config{
			"dialect": "sqlite",
			"host":    "file.db",
		}

		dialect, e := (&SqliteDialectStrategy{}).Get(cfg)
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
		cfg := &config.Config{
			"dialect": "sqlite",
			"host":    expectedPrefix,
			"params": config.Config{
				"param1": "value1",
				"param2": "value2",
			},
		}

		dialect, e := (&SqliteDialectStrategy{}).Get(cfg)
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
