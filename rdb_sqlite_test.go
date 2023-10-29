//go:build sqlite

package slate

import (
	"errors"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"gorm.io/driver/sqlite"
)

func Test_RdbSqliteDialectCreator(t *testing.T) {
	t.Run("Accept", func(t *testing.T) {
		t.Run("refuse if no config", func(t *testing.T) {
			if NewRdbSqliteDialectCreator().Accept(nil) == true {
				t.Error("returned true")
			}
		})

		t.Run("refuse on config parsing", func(t *testing.T) {
			if NewRdbSqliteDialectCreator().Accept(&ConfigPartial{"dialect": 123}) == true {
				t.Error("returned true")
			}
		})

		t.Run("refuse if the dialect name is not mysql", func(t *testing.T) {
			if NewRdbSqliteDialectCreator().Accept(&ConfigPartial{"dialect": "mysql"}) == true {
				t.Error("returned true")
			}
		})

		t.Run("accept if the dialect name is mysql", func(t *testing.T) {
			if NewRdbSqliteDialectCreator().Accept(&ConfigPartial{"dialect": "sQlItE"}) == false {
				t.Error("returned false")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("error on nil config", func(t *testing.T) {
			dialect, e := NewRdbSqliteDialectCreator().Create(nil)
			switch {
			case dialect != nil:
				t.Error("return an unexpected valid dialect instance")
			case e == nil:
				t.Error("didn't return the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})

		t.Run("invalid host value on connection configuration", func(t *testing.T) {
			config := ConfigPartial{
				"dialect": "sqlite",
				"host":    123,
			}

			dialect, e := NewRdbSqliteDialectCreator().Create(&config)
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
				"dialect": "sqlite",
				"host":    "host",
				"params":  123,
			}

			dialect, e := NewRdbSqliteDialectCreator().Create(&config)
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
			expected := "file.db"
			config := ConfigPartial{
				"dialect": "sqlite",
				"host":    "file.db",
			}

			dialect, e := NewRdbSqliteDialectCreator().Create(&config)
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
			config := ConfigPartial{
				"dialect": "sqlite",
				"host":    expectedPrefix,
				"params": ConfigPartial{
					"param1": "value1",
					"param2": "value2",
				},
			}

			dialect, e := NewRdbSqliteDialectCreator().Create(&config)
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
	})
}

func Test_RdbSqliteServiceRegister(t *testing.T) {
	t.Run("NewRdbSqliteServiceRegister", func(t *testing.T) {
		t.Run("create", func(t *testing.T) {
			if NewRdbSqliteServiceRegister() == nil {
				t.Error("didn't returned a valid reference")
			}
		})

		t.Run("create with app reference", func(t *testing.T) {
			app := NewApp()
			if sut := NewRdbSqliteServiceRegister(app); sut == nil {
				t.Error("didn't returned a valid reference")
			} else if sut.App != app {
				t.Error("didn't stored the app reference")
			}
		})
	})

	t.Run("Provide", func(t *testing.T) {
		t.Run("nil container", func(t *testing.T) {
			if e := NewRdbSqliteServiceRegister().Provide(nil); e == nil {
				t.Error("didn't return the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})

		t.Run("register components", func(t *testing.T) {
			container := NewServiceContainer()
			sut := NewRdbSqliteServiceRegister()

			e := sut.Provide(container)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case !container.Has(RdbSqliteDialectCreatorContainerID):
				t.Errorf("no sqlite dialect creator : %v", sut)
			}
		})

		t.Run("dialect creator is correctly tagged", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewRdbServiceRegister().Provide(container)
			_ = NewRdbSqliteServiceRegister().Provide(container)

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
						if _, ok := creator.(*RdbSqliteDialectCreator); ok {
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
