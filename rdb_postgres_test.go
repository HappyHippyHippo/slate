//go:build postgres

package slate

import (
	"errors"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"gorm.io/driver/postgres"
)

func Test_RdbPostgresDialectCreator(t *testing.T) {
	t.Run("Accept", func(t *testing.T) {
		t.Run("refuse if no config", func(t *testing.T) {
			if NewRdbPostgresDialectCreator().Accept(nil) == true {
				t.Error("returned true")
			}
		})

		t.Run("refuse on config parsing", func(t *testing.T) {
			if NewRdbPostgresDialectCreator().Accept(&ConfigPartial{"dialect": 123}) == true {
				t.Error("returned true")
			}
		})

		t.Run("refuse if the dialect name is not postgres", func(t *testing.T) {
			if NewRdbPostgresDialectCreator().Accept(&ConfigPartial{"dialect": "sqlite"}) == true {
				t.Error("returned true")
			}
		})

		t.Run("accept if the dialect name is postgres", func(t *testing.T) {
			if NewRdbPostgresDialectCreator().Accept(&ConfigPartial{"dialect": "postgres"}) == false {
				t.Error("returned false")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("error on nil config", func(t *testing.T) {
			dialect, e := NewRdbPostgresDialectCreator().Create(nil)
			switch {
			case dialect != nil:
				t.Error("return an unexpected valid dialect instance")
			case e == nil:
				t.Error("didn't return the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})

		t.Run("invalid username value on connection configuration", func(t *testing.T) {
			config := ConfigPartial{
				"dialect":  "postgres",
				"username": 123,
			}

			dialect, e := NewRdbPostgresDialectCreator().Create(&config)
			switch {
			case dialect != nil:
				t.Error("return an unexpected valid dialect instance")
			case e == nil:
				t.Error("didn't return the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expected (%v)", e, ErrConversion)
			}
		})

		t.Run("invalid password value on connection configuration", func(t *testing.T) {
			config := ConfigPartial{
				"dialect":  "postgres",
				"username": "user",
				"password": 123,
			}

			dialect, e := NewRdbPostgresDialectCreator().Create(&config)
			switch {
			case dialect != nil:
				t.Error("return an unexpected valid dialect instance")
			case e == nil:
				t.Error("didn't return the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expected (%v)", e, ErrConversion)
			}
		})

		t.Run("invalid host value on connection configuration", func(t *testing.T) {
			config := ConfigPartial{
				"dialect":  "postgres",
				"username": "user",
				"password": "password",
				"host":     123,
			}

			dialect, e := NewRdbPostgresDialectCreator().Create(&config)
			switch {
			case dialect != nil:
				t.Error("return an unexpected valid dialect instance")
			case e == nil:
				t.Error("didn't return the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expected (%v)", e, ErrConversion)
			}
		})

		t.Run("invalid port value on connection configuration", func(t *testing.T) {
			config := ConfigPartial{
				"dialect":  "postgres",
				"username": "user",
				"password": "password",
				"host":     "localhost",
				"port":     "integer",
			}

			dialect, e := NewRdbPostgresDialectCreator().Create(&config)
			switch {
			case dialect != nil:
				t.Error("return an unexpected valid dialect instance")
			case e == nil:
				t.Error("didn't return the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expected (%v)", e, ErrConversion)
			}
		})

		t.Run("invalid schema value on connection configuration", func(t *testing.T) {
			config := ConfigPartial{
				"dialect":  "postgres",
				"username": "user",
				"password": "password",
				"host":     "localhost",
				"port":     3306,
				"schema":   123,
			}

			dialect, e := NewRdbPostgresDialectCreator().Create(&config)
			switch {
			case dialect != nil:
				t.Error("return an unexpected valid dialect instance")
			case e == nil:
				t.Error("didn't return the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expected (%v)", e, ErrConversion)
			}
		})

		t.Run("invalid params value on connection configuration", func(t *testing.T) {
			config := ConfigPartial{
				"dialect":  "postgres",
				"username": "user",
				"password": "password",
				"host":     "localhost",
				"port":     3306,
				"schema":   "mysql",
				"params":   123,
			}

			dialect, e := NewRdbPostgresDialectCreator().Create(&config)
			switch {
			case dialect != nil:
				t.Error("return an unexpected valid dialect instance")
			case e == nil:
				t.Error("didn't return the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expected (%v)", e, ErrConversion)
			}
		})

		t.Run("valid connection", func(t *testing.T) {
			expected := "user=user password=password host=localhost port=123 dbname=rdb"
			config := ConfigPartial{
				"dialect":  "postgres",
				"username": "user",
				"password": "password",
				"host":     "localhost",
				"port":     123,
				"schema":   "rdb",
			}

			dialect, e := NewRdbPostgresDialectCreator().Create(&config)
			switch {
			case e != nil:
				t.Errorf("return the unexpected error : (%v)", e)
			case dialect == nil:
				t.Error("didn't return the expected valid dialect instance")
			default:
				if check := dialect.(*postgres.Dialector).Config.DSN; check != expected {
					t.Errorf("(%v) when expecting (%v)", check, expected)
				}
			}
		})

		t.Run("valid connection with simple port", func(t *testing.T) {
			expected := "user=user password=password host=localhost port=3306 dbname=rdb"
			config := ConfigPartial{
				"dialect":  "postgres",
				"username": "user",
				"password": "password",
				"host":     "localhost",
				"schema":   "rdb",
			}

			dialect, e := NewRdbPostgresDialectCreator().Create(&config)
			switch {
			case e != nil:
				t.Errorf("return the unexpected error : (%v)", e)
			case dialect == nil:
				t.Error("didn't return the expected valid dialect instance")
			default:
				switch d := dialect.(type) {
				case *postgres.Dialector:
					if check := d.Config.DSN; check != expected {
						t.Errorf("(%v) when expecting (%v)", check, expected)
					}
				default:
					t.Error("didn't return the expected mysql dialect")
				}
			}
		})

		t.Run("valid connection with extra params", func(t *testing.T) {
			expectedPrefix := "user=user password=password host=localhost port=3306 dbname=rdb"
			config := ConfigPartial{
				"dialect":  "postgres",
				"username": "user",
				"password": "password",
				"host":     "localhost",
				"schema":   "rdb",
				"params": ConfigPartial{
					"param1": "value1",
					"param2": "value2",
				},
			}

			dialect, e := NewRdbPostgresDialectCreator().Create(&config)
			switch {
			case e != nil:
				t.Errorf("return the unexpected error : (%v)", e)
			case dialect == nil:
				t.Error("didn't return the expected valid dialect instance")
			default:
				switch d := dialect.(type) {
				case *postgres.Dialector:
					dsn := d.Config.DSN
					switch {
					case !strings.HasPrefix(dsn, expectedPrefix):
						t.Errorf("(%v) when expecting (%v)", dsn, expectedPrefix)
					case !strings.Contains(dsn, " param1=value1"):
						t.Errorf("missing params (%v)", "&param1=value1")
					case !strings.Contains(dsn, " param2=value2"):
						t.Errorf("missing params (%v)", "&param2=value2")
					}
				default:
					t.Error("didn't return the expected mysql dialect")
				}
			}
		})
	})
}

func Test_RdbPostgresServiceRegister(t *testing.T) {
	t.Run("NewRdbPostgresServiceRegister", func(t *testing.T) {
		t.Run("create", func(t *testing.T) {
			if NewRdbPostgresServiceRegister() == nil {
				t.Error("didn't returned a valid reference")
			}
		})

		t.Run("create with app reference", func(t *testing.T) {
			app := NewApp()
			if sut := NewRdbPostgresServiceRegister(app); sut == nil {
				t.Error("didn't returned a valid reference")
			} else if sut.App != app {
				t.Error("didn't stored the app reference")
			}
		})
	})

	t.Run("Provide", func(t *testing.T) {
		t.Run("nil container", func(t *testing.T) {
			if e := NewRdbPostgresServiceRegister().Provide(nil); e == nil {
				t.Error("didn't return the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})

		t.Run("register components", func(t *testing.T) {
			container := NewServiceContainer()
			sut := NewRdbPostgresServiceRegister()

			e := sut.Provide(container)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case !container.Has(RdbPostgresDialectCreatorContainerID):
				t.Errorf("no postgres dialect creator : %v", sut)
			}
		})

		t.Run("dialect creator is correctly tagged", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewRdbServiceRegister().Provide(container)
			_ = NewRdbPostgresServiceRegister().Provide(container)

			creators, e := container.Get(RdbAllDialectCreatorsContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case creators == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch c := creators.(type) {
				case []RdbDialectCreator:
					found := false
					for _, creator := range c {
						if _, ok := creator.(*RdbPostgresDialectCreator); ok {
							found = true
						}
					}
					if !found {
						t.Error("didn't return a dialect creator slice populated with the expected creator instance")
					}
				default:
					t.Error("didn't return a dialect creator slice")
				}
			}
		})
	})
}
