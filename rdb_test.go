package slate

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Test_rdb_err(t *testing.T) {
	t.Run("errUnknownRdbDialect", func(t *testing.T) {
		arg := ConfigPartial{"field": "value"}
		context := map[string]interface{}{"field": "value"}
		message := "map[field:value] : unknown database dialect"

		t.Run("creation without context", func(t *testing.T) {
			if e := errUnknownRdbDialect(arg); !errors.Is(e, ErrUnknownRdbDialect) {
				t.Errorf("error not a instance of ErrUnknownRdbDialect")
			} else if e.Error() != message {
				t.Errorf("error message (%v) not same as expected (%v)", e, message)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})

		t.Run("creation with context", func(t *testing.T) {
			if e := errUnknownRdbDialect(arg, context); !errors.Is(e, ErrUnknownRdbDialect) {
				t.Errorf("error not a instance of ErrUnknownRdbDialect")
			} else if e.Error() != message {
				t.Errorf("error message (%v) not same as expected (%v)", e, message)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})
	})
}

func Test_RdbDialectFactory(t *testing.T) {
	t.Run("NewRdbDialectFactory", func(t *testing.T) {
		t.Run("creation with empty creator list", func(t *testing.T) {
			sut := NewRdbDialectFactory(nil)
			if sut == nil {
				t.Error("didn't returned the expected reference")
			}
		})

		t.Run("creation with creator list", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			creator := NewMockRdbDialectCreator(ctrl)

			sut := NewRdbDialectFactory([]RdbDialectCreator{creator})
			if sut == nil {
				t.Error("didn't returned the expected reference")
			} else if (*sut)[0] != creator {
				t.Error("didn't stored the passed creator")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("missing config", func(t *testing.T) {
			sut := NewRdbDialectFactory(nil)

			dialect, e := sut.Create(nil)
			switch {
			case dialect != nil:
				t.Error("return an unexpected valid dialect instance")
			case e == nil:
				t.Error("didn't return the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})

		t.Run("unsupported dialect", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			config := ConfigPartial{"dialect": "unsupported"}
			creator := NewMockRdbDialectCreator(ctrl)
			creator.EXPECT().Accept(&config).Return(false).Times(1)

			sut := NewRdbDialectFactory([]RdbDialectCreator{creator})

			dialect, e := sut.Create(&config)
			switch {
			case dialect != nil:
				t.Error("return an unexpected valid dialect instance")
			case e == nil:
				t.Error("didn't return the expected error")
			case !errors.Is(e, ErrUnknownRdbDialect):
				t.Errorf("(%v) when expected (%v)", e, ErrUnknownRdbDialect)
			}
		})

		t.Run("return creator error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			config := ConfigPartial{"dialect": "unsupported"}
			creator := NewMockRdbDialectCreator(ctrl)
			creator.EXPECT().Accept(&config).Return(true).Times(1)
			creator.EXPECT().Create(&config).Return(nil, expected).Times(1)

			sut := NewRdbDialectFactory([]RdbDialectCreator{creator})

			if _, e := sut.Create(&config); e == nil {
				t.Error("didn't return the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expected (%v)", e, expected)
			}
		})

		t.Run("return creator provided dialect", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			config := ConfigPartial{"dialect": "unsupported"}
			dialect := NewMockGormDialector(ctrl)
			creator := NewMockRdbDialectCreator(ctrl)
			creator.EXPECT().Accept(&config).Return(true).Times(1)
			creator.EXPECT().Create(&config).Return(dialect, nil).Times(1)

			sut := NewRdbDialectFactory([]RdbDialectCreator{creator})

			if check, e := sut.Create(&config); e != nil {
				t.Errorf("return the unexpected error (%v)", e)
			} else if check != dialect {
				t.Error("didn't returned the creator provided dialect")
			}
		})
	})
}

func Test_RdbConnectionFactory(t *testing.T) {
	t.Run("NewRdbConnectionFactory", func(t *testing.T) {
		t.Run("missing dialect factory", func(t *testing.T) {
			sut, e := NewRdbConnectionFactory(nil)
			switch {
			case sut != nil:
				t.Error("return an unexpected valid connection factory")
			case e == nil:
				t.Error("didn't return the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})

		t.Run("valid creation", func(t *testing.T) {
			if sut, e := NewRdbConnectionFactory(NewRdbDialectFactory(nil)); sut == nil {
				t.Error("didn't returned the expected valid connection factory")
			} else if e != nil {
				t.Errorf("return the unexpected error : %v", e)
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("error instantiating dialect", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			config := ConfigPartial{"dialect": "invalid"}

			sut, _ := NewRdbConnectionFactory(NewRdbDialectFactory(nil))
			conn, e := sut.Create(&config, &gorm.Config{})
			switch {
			case conn != nil:
				t.Error("return an unexpected valid connection instance")
			case e == nil:
				t.Error("didn't return the expected error")
			case !errors.Is(e, ErrUnknownRdbDialect):
				t.Errorf("(%v) when expected (%v)", e, ErrUnknownRdbDialect)
			}
		})

		t.Run("error instantiating connector", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			config := ConfigPartial{
				"dialect": "invalid",
				"host":    ":memory:",
			}
			dialect := NewMockGormDialector(ctrl)
			dialect.EXPECT().Initialize(gomock.Any()).Return(expected).Times(1)
			dialectCreator := NewMockRdbDialectCreator(ctrl)
			dialectCreator.EXPECT().Accept(&config).Return(true).Times(1)
			dialectCreator.EXPECT().Create(&config).Return(dialect, nil).Times(1)
			dialectFactory := NewRdbDialectFactory([]RdbDialectCreator{dialectCreator})

			sut, _ := NewRdbConnectionFactory(dialectFactory)

			conn, e := sut.Create(&config, &gorm.Config{Logger: logger.Discard})
			switch {
			case conn != nil:
				t.Error("return an unexpected valid connection instance")
			case e == nil:
				t.Error("didn't return the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expected (%v)", e, expected)
			}
		})

		t.Run("valid connection", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			config := ConfigPartial{
				"dialect": "invalid",
				"host":    ":memory:",
			}
			dialect := NewMockGormDialector(ctrl)
			dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
			dialectCreator := NewMockRdbDialectCreator(ctrl)
			dialectCreator.EXPECT().Accept(&config).Return(true).Times(1)
			dialectCreator.EXPECT().Create(&config).Return(dialect, nil).Times(1)
			dialectFactory := NewRdbDialectFactory([]RdbDialectCreator{dialectCreator})

			sut, _ := NewRdbConnectionFactory(dialectFactory)

			if check, e := sut.Create(&config, &gorm.Config{Logger: logger.Discard}); check == nil {
				t.Error("didn't return the expected connection instance")
			} else if e != nil {
				t.Errorf("return the unexpected error : (%v)", e)
			}
		})
	})
}

func Test_RdbConnectionPool(t *testing.T) {
	t.Run("NewRdbConnectionPool", func(t *testing.T) {
		t.Run("missing configuration", func(t *testing.T) {
			sut, e := NewRdbConnectionPool(nil, &RdbConnectionFactory{})
			switch {
			case sut != nil:
				t.Error("return an unexpected valid connection factory")
			case e == nil:
				t.Error("didn't return the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})

		t.Run("missing connection factory", func(t *testing.T) {
			sut, e := NewRdbConnectionPool(NewConfig(), nil)
			switch {
			case sut != nil:
				t.Error("return an unexpected valid connection factory")
			case e == nil:
				t.Error("didn't return the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})

		t.Run("valid creation", func(t *testing.T) {
			connectionsFactory, _ := NewRdbConnectionFactory(NewRdbDialectFactory(nil))

			if sut, e := NewRdbConnectionPool(NewConfig(), connectionsFactory); sut == nil {
				t.Error("didn't returned the expected valid connection factory ")
			} else if e != nil {
				t.Errorf("return the unexpected error : %v", e)
			}
		})

		t.Run("config change purge all stored connections", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			name := "primary"
			config1 := ConfigPartial{"dialect": "sqlite", "host": ":memory1:"}
			config2 := ConfigPartial{"dialect": "sqlite", "host": ":memory2:"}
			gormCfg := &gorm.Config{Logger: logger.Discard, DisableAutomaticPing: true}
			configs1 := ConfigPartial{}
			_, _ = configs1.Set("slate.rdb.connections", ConfigPartial{name: config1})
			configs2 := ConfigPartial{}
			_, _ = configs2.Set("slate.rdb.connections", ConfigPartial{name + "salt": config2})
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Get("").Return(configs1, nil).MinTimes(1)
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Get("").Return(configs2, nil).MinTimes(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 0, supplier1)
			dialect := NewMockGormDialector(ctrl)
			dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
			dialectCreator := NewMockRdbDialectCreator(ctrl)
			dialectCreator.EXPECT().Accept(&config1).Return(true).Times(1)
			dialectCreator.EXPECT().Create(&config1).Return(dialect, nil).Times(1)
			dialectFactory := NewRdbDialectFactory([]RdbDialectCreator{dialectCreator})
			connectionFactory, _ := NewRdbConnectionFactory(dialectFactory)
			sut, _ := NewRdbConnectionPool(config, connectionFactory)

			_, _ = sut.Get(name, gormCfg)
			if len(sut.connections) != 1 {
				t.Error("didn't store the requested connection instance")
			}

			_ = config.AddSupplier("id2", 10, supplier2)
			if len(sut.connections) != 0 {
				t.Error("didn't removed the stored connection connections")
			}
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("missing requested connection configuration", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			connectionFactory, _ := NewRdbConnectionFactory(NewRdbDialectFactory(nil))
			sut, _ := NewRdbConnectionPool(NewConfig(), connectionFactory)

			conn, e := sut.Get("primary", &gorm.Config{})
			switch {
			case conn != nil:
				t.Error("return an unexpected valid connection instance")
			case e == nil:
				t.Error("didn't return the expected error")
			case !errors.Is(e, ErrConfigPathNotFound):
				t.Errorf("(%v) when expected (%v)", e, ErrConfigPathNotFound)
			}
		})

		t.Run("invalid requested connection configuration", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			configs := ConfigPartial{}
			_, _ = configs.Set("slate.rdb.connections.primary", "string")
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(configs, nil).MinTimes(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 0, supplier)
			dialectCreator := NewMockRdbDialectCreator(ctrl)
			dialectFactory := NewRdbDialectFactory([]RdbDialectCreator{dialectCreator})
			connectionFactory, _ := NewRdbConnectionFactory(dialectFactory)
			sut, _ := NewRdbConnectionPool(config, connectionFactory)

			conn, e := sut.Get("primary", &gorm.Config{})
			switch {
			case conn != nil:
				t.Error("return an unexpected valid connection instance")
			case e == nil:
				t.Error("didn't return the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expected (%v)", e, ErrConversion)
			}
		})

		t.Run("error instantiating connector", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("erorr message")
			partial := ConfigPartial{"data": "string"}
			configs := ConfigPartial{}
			_, _ = configs.Set("slate.rdb.connections.primary", partial)
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(configs, nil).MinTimes(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 0, supplier)
			dialectCreator := NewMockRdbDialectCreator(ctrl)
			dialectCreator.EXPECT().Accept(&partial).Return(true).Times(1)
			dialectCreator.EXPECT().Create(&partial).Return(nil, expected).Times(1)
			dialectFactory := NewRdbDialectFactory([]RdbDialectCreator{dialectCreator})
			connectionFactory, _ := NewRdbConnectionFactory(dialectFactory)
			sut, _ := NewRdbConnectionPool(config, connectionFactory)

			conn, e := sut.Get("primary", &gorm.Config{})
			switch {
			case conn != nil:
				t.Error("return an unexpected valid connection instance")
			case e == nil:
				t.Error("didn't return the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expected (%v)", e, expected)
			}
		})

		t.Run("valid connection", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := ConfigPartial{"data": "string"}
			gormCfg := &gorm.Config{Logger: logger.Discard}
			configs := ConfigPartial{}
			_, _ = configs.Set("slate.rdb.connections.primary", partial)
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(configs, nil).MinTimes(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 0, supplier)
			dialect := NewMockGormDialector(ctrl)
			dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
			dialectCreator := NewMockRdbDialectCreator(ctrl)
			dialectCreator.EXPECT().Accept(&partial).Return(true).Times(1)
			dialectCreator.EXPECT().Create(&partial).Return(dialect, nil).Times(1)
			dialectFactory := NewRdbDialectFactory([]RdbDialectCreator{dialectCreator})
			connectionFactory, _ := NewRdbConnectionFactory(dialectFactory)
			sut, _ := NewRdbConnectionPool(config, connectionFactory)

			if check, e := sut.Get("primary", gormCfg); check == nil {
				t.Error("didn't return the expected connection instance")
			} else if e != nil {
				t.Errorf("return the unexpected error : (%v)", e)
			}
		})

		t.Run("multiple requests only instantiate a single connection", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := ConfigPartial{"data": "string"}
			gormCfg := &gorm.Config{Logger: logger.Discard}
			configs := ConfigPartial{}
			_, _ = configs.Set("slate.rdb.connections.primary", partial)
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(configs, nil).MinTimes(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 0, supplier)
			dialect := NewMockGormDialector(ctrl)
			dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
			dialectCreator := NewMockRdbDialectCreator(ctrl)
			dialectCreator.EXPECT().Accept(&partial).Return(true).Times(1)
			dialectCreator.EXPECT().Create(&partial).Return(dialect, nil).Times(1)
			dialectFactory := NewRdbDialectFactory([]RdbDialectCreator{dialectCreator})
			connectionFactory, _ := NewRdbConnectionFactory(dialectFactory)
			sut, _ := NewRdbConnectionPool(config, connectionFactory)

			conn, _ := sut.Get("primary", gormCfg)
			check, e := sut.Get("primary", gormCfg)
			switch {
			case check == nil:
				t.Error("didn't return the expected connection instance")
			case e != nil:
				t.Errorf("return the unexpected error : (%v)", e)
			case check != conn:
				t.Error("didn't returned the same instance")
			}
		})

		t.Run("reconfigure closes opened connections", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial1 := ConfigPartial{"dialect": "my_dialect", "host": ":memory:"}
			partial2 := ConfigPartial{"data": "string 2"}
			gormCfg := &gorm.Config{Logger: logger.Discard, DisableAutomaticPing: true}
			configs1 := ConfigPartial{}
			_, _ = configs1.Set("slate.rdb.connections.primary", partial1)
			configs2 := ConfigPartial{}
			_, _ = configs2.Set("slate.rdb.connections.primary", partial2)
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Get("").Return(configs1, nil).MinTimes(2)
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Get("").Return(configs2, nil).MinTimes(1)
			config := NewConfig()
			_ = config.AddSupplier("id1", 0, supplier1)
			dialect := NewMockGormDialector(ctrl)
			dialect.
				EXPECT().
				Initialize(gomock.Any()).
				DoAndReturn(func(db *gorm.DB) error {
					db.ConnPool, _ = sql.Open("sqlmock", "name")
					return nil
				}).
				Times(1)
			dialectCreator := NewMockRdbDialectCreator(ctrl)
			dialectCreator.EXPECT().Accept(&partial1).Return(true).Times(1)
			dialectCreator.EXPECT().Create(&partial1).Return(dialect, nil).Times(1)
			dialectFactory := NewRdbDialectFactory([]RdbDialectCreator{dialectCreator})
			connectionFactory, _ := NewRdbConnectionFactory(dialectFactory)
			sut, _ := NewRdbConnectionPool(config, connectionFactory)

			_, _ = sut.Get("primary", gormCfg)

			_ = config.AddSupplier("id2", 0, supplier2)
			if len(sut.connections) != 0 {
				t.Error("didn't closed already opened connections")
			}
		})
	})
}

func Test_RdbServiceRegister(t *testing.T) {
	t.Run("NewRdbServiceRegister", func(t *testing.T) {
		t.Run("create", func(t *testing.T) {
			if NewRdbServiceRegister() == nil {
				t.Error("didn't returned a valid reference")
			}
		})

		t.Run("create with app reference", func(t *testing.T) {
			app := NewApp()
			if sut := NewRdbServiceRegister(app); sut == nil {
				t.Error("didn't returned a valid reference")
			} else if sut.App != app {
				t.Error("didn't stored the app reference")
			}
		})
	})

	t.Run("Provide", func(t *testing.T) {
		t.Run("nil container", func(t *testing.T) {
			if e := NewRdbServiceRegister().Provide(nil); e == nil {
				t.Error("didn't return the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})

		t.Run("register components", func(t *testing.T) {
			container := NewServiceContainer()
			sut := NewRdbServiceRegister()

			e := sut.Provide(container)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case !container.Has(RdbConfigContainerID):
				t.Errorf("no connection configuration : %v", sut)
			case !container.Has(RdbAllDialectCreatorsContainerID):
				t.Errorf("no dialect creator slice : %v", sut)
			case !container.Has(RdbDialectFactoryContainerID):
				t.Errorf("no dialect factory : %v", sut)
			case !container.Has(RdbConnectionFactoryContainerID):
				t.Errorf("no connection factory : %v", sut)
			case !container.Has(RdbContainerID):
				t.Errorf("no connection pool : %v", sut)
			case !container.Has(RdbPrimaryConnectionContainerID):
				t.Errorf("no primary connection handler : %v", sut)
			}
		})

		t.Run("retrieving connection configuration", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewRdbServiceRegister().Provide(container)

			config, e := container.Get(RdbConfigContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case config == nil:
				t.Error("didn't returned a valid reference")
			}
		})

		t.Run("retrieving dialect creators", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewRdbServiceRegister().Provide(container)

			dialectCreator := NewMockRdbDialectCreator(ctrl)
			_ = container.Add("dialect.id", func() RdbDialectCreator {
				return dialectCreator
			}, RdbDialectCreatorTag)

			creators, e := container.Get(RdbAllDialectCreatorsContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case creators == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch s := creators.(type) {
				case []RdbDialectCreator:
					if s[0] != dialectCreator {
						t.Error("didn't return a dialect creator slice populated with the expected creator instance")
					}
				default:
					t.Error("didn't return a dialect creator slice")
				}
			}
		})

		t.Run("retrieving dialect factory", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewRdbServiceRegister().Provide(container)

			factory, e := container.Get(RdbDialectFactoryContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case factory == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch factory.(type) {
				case *RdbDialectFactory:
				default:
					t.Error("didn't return a dialect factory instance")
				}
			}
		})

		t.Run("error retrieving dialect factory when retrieving connection factory", func(t *testing.T) {
			expected := fmt.Errorf("error message")
			container := NewServiceContainer()
			_ = NewConfigServiceRegister().Provide(container)
			_ = NewRdbServiceRegister().Provide(container)
			_ = container.Add(RdbDialectFactoryContainerID, func() (*RdbDialectFactory, error) {
				return nil, expected
			})

			if _, e := container.Get(RdbConnectionFactoryContainerID); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceContainer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("error retrieving configuration when retrieving connection pool", func(t *testing.T) {
			expected := fmt.Errorf("error message")
			container := NewServiceContainer()
			_ = NewRdbServiceRegister().Provide(container)
			_ = container.Add(ConfigContainerID, func() (*Config, error) {
				return nil, expected
			})

			if _, e := container.Get(RdbContainerID); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceContainer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("error retrieving connection factory when retrieving connection pool", func(t *testing.T) {
			expected := fmt.Errorf("error message")
			container := NewServiceContainer()
			_ = NewRdbServiceRegister().Provide(container)
			_ = container.Add(RdbConnectionFactoryContainerID, func() (*RdbConnectionFactory, error) {
				return nil, expected
			})

			if _, e := container.Get(RdbContainerID); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceContainer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("retrieving connection factory", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewConfigServiceRegister().Provide(container)
			_ = NewRdbServiceRegister().Provide(container)

			factory, e := container.Get(RdbConnectionFactoryContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case factory == nil:
				t.Error("didn't return a valid reference")
			default:
				switch factory.(type) {
				case *RdbConnectionFactory:
				default:
					t.Error("didn't return a dialect factory instance")
				}
			}
		})

		t.Run("error retrieving connection factory when retrieving primary connection", func(t *testing.T) {
			expected := fmt.Errorf("error message")
			container := NewServiceContainer()
			_ = NewRdbServiceRegister().Provide(container)
			_ = container.Add(RdbContainerID, func() (*RdbConnectionPool, error) {
				return nil, expected
			})

			if _, e := container.Get(RdbPrimaryConnectionContainerID); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceContainer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("error retrieving connection configuration when retrieving primary connection", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			container := NewServiceContainer()
			_ = NewConfigServiceRegister().Provide(container)
			_ = NewRdbServiceRegister().Provide(container)
			_ = container.Add(RdbConfigContainerID, func() (*gorm.Config, error) {
				return nil, expected
			})

			if _, e := container.Get(RdbPrimaryConnectionContainerID); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceContainer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("valid primary connection", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewConfigServiceRegister().Provide(container)
			_ = NewRdbServiceRegister().Provide(container)

			rdbCfg := ConfigPartial{"dialect": "invalid", "host": ":memory:"}
			partial := ConfigPartial{}
			_, _ = partial.Set("slate.rdb.connections.primary", rdbCfg)
			dialect := NewMockGormDialector(ctrl)
			dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
			dialectCreator := NewMockRdbDialectCreator(ctrl)
			dialectCreator.EXPECT().Accept(&rdbCfg).Return(true).Times(1)
			dialectCreator.EXPECT().Create(&rdbCfg).Return(dialect, nil).Times(1)
			dialectFactory := NewRdbDialectFactory([]RdbDialectCreator{dialectCreator})
			_ = container.Add(RdbDialectFactoryContainerID, func() *RdbDialectFactory {
				return dialectFactory
			})
			source := NewMockConfigSupplier(ctrl)
			source.EXPECT().Get("").Return(partial, nil).Times(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 1, source)
			_ = container.Add(ConfigContainerID, func() *Config {
				return config
			})

			check, e := container.Get(RdbPrimaryConnectionContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case check == nil:
				t.Error("didn't return a valid reference")
			default:
				switch check.(type) {
				case *gorm.DB:
				default:
					t.Error("didn't return a dialect factory instance")
				}
			}
		})

		t.Run("valid primary connection with overridden primary connection", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			primary := "other_primary"
			RdbPrimary = primary
			defer func() { RdbPrimary = "primary" }()

			container := NewServiceContainer()
			_ = NewConfigServiceRegister().Provide(container)
			_ = NewRdbServiceRegister().Provide(container)

			rdbCfg := ConfigPartial{"dialect": "invalid", "host": ":memory:"}
			partial := ConfigPartial{}
			_, _ = partial.Set("slate.rdb.connections.other_primary", rdbCfg)
			dialect := NewMockGormDialector(ctrl)
			dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
			dialectCreator := NewMockRdbDialectCreator(ctrl)
			dialectCreator.EXPECT().Accept(&rdbCfg).Return(true).Times(1)
			dialectCreator.EXPECT().Create(&rdbCfg).Return(dialect, nil).Times(1)
			dialectFactory := NewRdbDialectFactory([]RdbDialectCreator{dialectCreator})
			_ = container.Add(RdbDialectFactoryContainerID, func() *RdbDialectFactory {
				return dialectFactory
			})
			source := NewMockConfigSupplier(ctrl)
			source.EXPECT().Get("").Return(partial, nil).Times(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 1, source)
			_ = container.Add(ConfigContainerID, func() *Config {
				return config
			})

			check, e := container.Get(RdbPrimaryConnectionContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case check == nil:
				t.Error("didn't return a valid reference")
			default:
				switch check.(type) {
				case *gorm.DB:
				default:
					t.Error("didn't return a dialect factory instance")
				}
			}
		})
	})

	t.Run("Boot", func(t *testing.T) {
		t.Run("nil container", func(t *testing.T) {
			if e := NewRdbServiceRegister().Boot(nil); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("run boot", func(t *testing.T) {
			app := NewApp()
			_ = app.Provide(NewRdbServiceRegister())

			if e := app.Boot(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})
	})
}
