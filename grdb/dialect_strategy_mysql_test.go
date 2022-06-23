package grdb

import (
	"errors"
	"github.com/happyhippyhippo/slate/gconfig"
	"github.com/happyhippyhippo/slate/gerror"
	"gorm.io/driver/mysql"
	"strings"
	"testing"
)

func Test_DialectStrategyMySql_Accept(t *testing.T) {
	t.Run("refuse if the dialect name is not mysql", func(t *testing.T) {
		if (&DialectStrategyMySQL{}).Accept("sqlite") == true {
			t.Error("returned true")
		}
	})

	t.Run("accept if the dialect name is mysql", func(t *testing.T) {
		if (&DialectStrategyMySQL{}).Accept("mYsQl") == false {
			t.Error("returned false")
		}
	})
}

func Test_DialectStrategyMySql_Get(t *testing.T) {
	t.Run("invalid username value on connection configuration", func(t *testing.T) {
		cfg := &gconfig.Partial{
			"dialect":  "mysql",
			"username": 123,
		}

		dialect, err := (&DialectStrategyMySQL{}).Get(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("invalid password value on connection configuration", func(t *testing.T) {
		cfg := &gconfig.Partial{
			"dialect":  "mysql",
			"username": "user",
			"password": 123,
		}

		dialect, err := (&DialectStrategyMySQL{}).Get(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("invalid protocol value on connection configuration", func(t *testing.T) {
		cfg := &gconfig.Partial{
			"dialect":  "mysql",
			"username": "user",
			"password": "password",
			"protocol": 123,
		}

		dialect, err := (&DialectStrategyMySQL{}).Get(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("invalid host value on connection configuration", func(t *testing.T) {
		cfg := &gconfig.Partial{
			"dialect":  "mysql",
			"username": "user",
			"password": "password",
			"protocol": "tcp",
			"host":     123,
		}

		dialect, err := (&DialectStrategyMySQL{}).Get(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("invalid port value on connection configuration", func(t *testing.T) {
		cfg := &gconfig.Partial{
			"dialect":  "mysql",
			"username": "user",
			"password": "password",
			"protocol": "tcp",
			"host":     "localhost",
			"port":     "integer",
		}

		dialect, err := (&DialectStrategyMySQL{}).Get(cfg)
		switch {
		case dialect != nil:
			t.Error("return an unexpected valid dialect instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("invalid schema value on connection configuration", func(t *testing.T) {
		cfg := &gconfig.Partial{
			"dialect":  "mysql",
			"username": "user",
			"password": "password",
			"protocol": "tcp",
			"host":     "localhost",
			"port":     3306,
			"schema":   123,
		}

		dialect, err := (&DialectStrategyMySQL{}).Get(cfg)
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
			"dialect":  "mysql",
			"username": "user",
			"password": "password",
			"protocol": "tcp",
			"host":     "localhost",
			"port":     3306,
			"schema":   "mysql",
			"params":   123,
		}

		dialect, err := (&DialectStrategyMySQL{}).Get(cfg)
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
		expected := "user:password@protocol(localhost:123)/rdb"
		cfg := &gconfig.Partial{
			"dialect":  "mysql",
			"username": "user",
			"password": "password",
			"protocol": "protocol",
			"host":     "localhost",
			"port":     123,
			"schema":   "rdb",
		}

		dialect, err := (&DialectStrategyMySQL{}).Get(cfg)
		switch {
		case err != nil:
			t.Errorf("return the unexpected error : (%v)", err)
		case dialect == nil:
			t.Error("didn't return the expected valid dialect instance")
		default:
			if check := dialect.(*mysql.Dialector).Config.DSN; check != expected {
				t.Errorf("dialect composed with the DSN (%v) when expected to be (%v)", check, expected)
			}
		}
	})

	t.Run("valid connection with default protocol and port", func(t *testing.T) {
		expected := "user:password@tcp(localhost:3306)/rdb"
		cfg := &gconfig.Partial{
			"dialect":  "mysql",
			"username": "user",
			"password": "password",
			"host":     "localhost",
			"schema":   "rdb",
		}

		dialect, err := (&DialectStrategyMySQL{}).Get(cfg)
		switch {
		case err != nil:
			t.Errorf("return the unexpected error : (%v)", err)
		case dialect == nil:
			t.Error("didn't return the expected valid dialect instance")
		default:
			switch d := dialect.(type) {
			case *mysql.Dialector:
				if check := d.Config.DSN; check != expected {
					t.Errorf("dialect composed with the DSN (%v) when expected to be (%v)", check, expected)
				}
			default:
				t.Error("didn't return the expected mysql dialect instance")
			}
		}
	})

	t.Run("valid connection with extra params", func(t *testing.T) {
		expectedPrefix := "user:password@tcp(localhost:3306)/rdb?"
		cfg := &gconfig.Partial{
			"dialect":  "mysql",
			"username": "user",
			"password": "password",
			"host":     "localhost",
			"schema":   "rdb",
			"params": gconfig.Partial{
				"param1": "value1",
				"param2": "value2",
			},
		}

		dialect, err := (&DialectStrategyMySQL{}).Get(cfg)
		switch {
		case err != nil:
			t.Errorf("return the unexpected error : (%v)", err)
		case dialect == nil:
			t.Error("didn't return the expected valid dialect instance")
		default:
			switch d := dialect.(type) {
			case *mysql.Dialector:
				dsn := d.Config.DSN
				switch {
				case !strings.HasPrefix(dsn, expectedPrefix):
					t.Errorf("dialect composed with the DSN prefix of (%v) when expected to be (%v)", dsn, expectedPrefix)
				case !strings.Contains(dsn, "&param1=value1"):
					t.Errorf("missing dialect composed with the DSN params (%v)", "&param1=value1")
				case !strings.Contains(dsn, "&param2=value2"):
					t.Errorf("missing dialect composed with the DSN params (%v)", "&param2=value2")
				}
			default:
				t.Error("didn't return the expected mysql dialect instance")
			}
		}
	})
}
