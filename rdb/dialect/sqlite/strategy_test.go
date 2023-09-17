//go:build sqlite

package sqlite

import (
	"errors"
	"strings"
	"testing"

	"gorm.io/driver/sqlite"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
)

func Test_DialectStrategy_Accept(t *testing.T) {
	t.Run("refuse if no config", func(t *testing.T) {
		if (&DialectStrategy{}).Accept(nil) == true {
			t.Error("returned true")
		}
	})

	t.Run("refuse on config parsing", func(t *testing.T) {
		if (&DialectStrategy{}).Accept(config.Partial{"dialect": 123}) == true {
			t.Error("returned true")
		}
	})

	t.Run("refuse if the dialect name is not mysql", func(t *testing.T) {
		if (&DialectStrategy{}).Accept(config.Partial{"dialect": "mysql"}) == true {
			t.Error("returned true")
		}
	})

	t.Run("accept if the dialect name is mysql", func(t *testing.T) {
		if (&DialectStrategy{}).Accept(config.Partial{"dialect": "sQlItE"}) == false {
			t.Error("returned false")
		}
	})
}

func Test_DialectStrategy_Create(t *testing.T) {
	t.Run("error on nil config", func(t *testing.T) {
		dialect, e := (&DialectStrategy{}).Create(nil)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("invalid host value on connection configuration", func(t *testing.T) {
		cfg := config.Partial{
			"dialect": "sqlite",
			"host":    123,
		}

		dialect, e := (&DialectStrategy{}).Create(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expected (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("invalid params value on connection configuration", func(t *testing.T) {
		cfg := config.Partial{
			"dialect": "sqlite",
			"host":    "host",
			"params":  123,
		}

		dialect, e := (&DialectStrategy{}).Create(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expected (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("valid connection", func(t *testing.T) {
		expected := "file.db"
		cfg := config.Partial{
			"dialect": "sqlite",
			"host":    "file.db",
		}

		dialect, e := (&DialectStrategy{}).Create(cfg)
		switch {
		case e != nil:
			t.Errorf("return the unexpected error : (%v)", e)
		case dialect == nil:
			t.Error("didn't return the expected valid dialect instance")
		default:
			switch d := dialect.(type) {
			case *sqlite.Dialector:
				if check := d.DSN; check != expected {
					t.Errorf("(%v) when expecting (%v)", check, expected)
				}
			default:
				t.Error("didn't return the expected sqlite dialect")
			}
		}
	})

	t.Run("valid connection with extra params", func(t *testing.T) {
		expectedPrefix := ":memory:"
		cfg := config.Partial{
			"dialect": "sqlite",
			"host":    expectedPrefix,
			"params": config.Partial{
				"param1": "value1",
				"param2": "value2",
			},
		}

		dialect, e := (&DialectStrategy{}).Create(cfg)
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
					t.Errorf("(%v) when expecting (%v)", dsn, expectedPrefix)
				case !strings.Contains(dsn, "&param1=value1"):
					t.Errorf("missing params (%v)", "&param1=value1")
				case !strings.Contains(dsn, "&param2=value2"):
					t.Errorf("missing params (%v)", "&param2=value2")
				}
			default:
				t.Error("didn't return the expected sqlite dialect")
			}
		}
	})
}
