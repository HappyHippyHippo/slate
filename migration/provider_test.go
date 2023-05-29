package migration

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/rdb"
	"gorm.io/gorm"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}

		e := sut.Register(container)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case !container.Has(ID):
			t.Errorf("didn't registered the migrator : %v", sut)
		case !container.Has(DaoID):
			t.Errorf("didn't registered the migrator DAO : %v", sut)
		}
	})

	t.Run("error retrieving db connection factory when retrieving migrator DAO", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (&rdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(rdb.ID, func() (*rdb.ConnectionPool, error) { return nil, expected })

		if _, e := container.Get(DaoID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error retrieving db connection config when retrieving migrator DAO", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(rdb.ConfigID, func() (*gorm.Config, error) { return nil, expected })

		if _, e := container.Get(DaoID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error retrieving connection when retrieving migrator DAO", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		if _, e := container.Get(DaoID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("retrieving migrator DAO", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		rdbCfg := config.Partial{"dialect": "invalid", "host": ":memory:"}
		partial := config.Partial{
			"slate": config.Partial{
				"rdb": config.Partial{
					"connections": config.Partial{
						"primary": rdbCfg}}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		_ = container.Service(config.ID, func() *config.Config { return cfg })
		migrator := NewMockMigrator(ctrl)
		migrator.EXPECT().AutoMigrate(gomock.Any()).Return(nil).Times(1)
		dialector := NewMockDialector(ctrl)
		dialector.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
		dialector.EXPECT().Migrator(gomock.Any()).Return(migrator).Times(1)
		dialectStrategy := NewMockDialectStrategy(ctrl)
		dialectStrategy.EXPECT().Accept(&rdbCfg).Return(true).Times(1)
		dialectStrategy.EXPECT().Create(&rdbCfg).Return(dialector, nil).Times(1)
		_ = container.Service("dialect_strategy", func() rdb.DialectStrategy { return dialectStrategy }, rdb.DialectStrategyTag)

		_ = (&rdb.Provider{}).Boot(container)

		sut, e := container.Get(DaoID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case sut == nil:
			t.Error("didn't returned a reference to the migrator DAO")
		default:
			switch sut.(type) {
			case *Dao:
			default:
				t.Error("didn't returned a migrator DAO reference")
			}
		}
	})

	t.Run("retrieving migrator DAO with db name from env", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := Database
		Database = "secondary"
		defer func() { Database = prev }()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		rdbCfg := config.Partial{"dialect": "invalid", "host": ":memory:"}
		partial := config.Partial{
			"slate": config.Partial{
				"rdb": config.Partial{
					"connections": config.Partial{
						"secondary": rdbCfg}}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		_ = container.Service(config.ID, func() *config.Config { return cfg })
		migrator := NewMockMigrator(ctrl)
		migrator.EXPECT().AutoMigrate(gomock.Any()).Return(nil).Times(1)
		dialector := NewMockDialector(ctrl)
		dialector.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
		dialector.EXPECT().Migrator(gomock.Any()).Return(migrator).Times(1)
		dialectStrategy := NewMockDialectStrategy(ctrl)
		dialectStrategy.EXPECT().Accept(&rdbCfg).Return(true).Times(1)
		dialectStrategy.EXPECT().Create(&rdbCfg).Return(dialector, nil).Times(1)
		_ = container.Service("dialect_strategy", func() rdb.DialectStrategy { return dialectStrategy }, rdb.DialectStrategyTag)

		_ = (&rdb.Provider{}).Boot(container)

		sut, e := container.Get(DaoID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case sut == nil:
			t.Error("didn't returned a reference to the migrator DAO")
		default:
			switch sut.(type) {
			case *Dao:
			default:
				t.Error("didn't returned a migrator DAO reference")
			}
		}
	})

	t.Run("invalid migrator DAO when retrieving the migrator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(DaoID, func() (*Dao, error) { return nil, expected })

		if _, e := container.Get(ID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error retrieving the migrator DAO when retrieving the migrator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(DaoID, func() (*Dao, error) { return nil, expected })

		if _, e := container.Get(ID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("retrieving migrator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		rdbCfg := config.Partial{"dialect": "invalid", "host": ":memory:"}
		partial := config.Partial{
			"slate": config.Partial{
				"rdb": config.Partial{
					"connections": config.Partial{
						"primary": rdbCfg}}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		_ = container.Service(config.ID, func() *config.Config { return cfg })
		migrator := NewMockMigrator(ctrl)
		migrator.EXPECT().AutoMigrate(gomock.Any()).Return(nil).Times(1)
		dialector := NewMockDialector(ctrl)
		dialector.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
		dialector.EXPECT().Migrator(gomock.Any()).Return(migrator).Times(1)
		dialectStrategy := NewMockDialectStrategy(ctrl)
		dialectStrategy.EXPECT().Accept(&rdbCfg).Return(true).Times(1)
		dialectStrategy.EXPECT().Create(&rdbCfg).Return(dialector, nil).Times(1)
		_ = container.Service("dialect_strategy", func() rdb.DialectStrategy { return dialectStrategy }, rdb.DialectStrategyTag)

		_ = (&rdb.Provider{}).Boot(container)

		sut, e := container.Get(ID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case sut == nil:
			t.Error("didn't returned a reference to the migrator")
		default:
			switch sut.(type) {
			case *Migrator:
			default:
				t.Error("didn't returned a migrator reference")
			}
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Boot(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("disable auto migration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		AutoMigrate = false
		defer func() { AutoMigrate = true }()

		container := slate.NewContainer()
		sut := &Provider{}

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the unexpected serr, (%v)", e)
		}
	})

	t.Run("disable migrator auto migration by environment variable", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		AutoMigrate = false
		defer func() { AutoMigrate = true }()

		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the unexpected serr, (%v)", e)
		}
	})

	t.Run("error on retrieving migrator", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ID, func() (*Migrator, error) { return nil, expected })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})
	t.Run("invalid retrieved migrator", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ID, func() (interface{}, error) { return "string", nil })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error on retrieving migrations", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)

		expected := fmt.Errorf("error message")
		rdbCfg := config.Partial{"dialect": "invalid", "host": ":memory:"}
		partial := config.Partial{
			"slate": config.Partial{
				"rdb": config.Partial{
					"connections": config.Partial{
						"primary": rdbCfg}}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		_ = container.Service(config.ID, func() *config.Config { return cfg })
		migrator := NewMockMigrator(ctrl)
		migrator.EXPECT().AutoMigrate(gomock.Any()).Return(nil).Times(1)
		dialector := NewMockDialector(ctrl)
		dialector.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
		dialector.EXPECT().Migrator(gomock.Any()).Return(migrator).Times(1)
		dialectStrategy := NewMockDialectStrategy(ctrl)
		dialectStrategy.EXPECT().Accept(&rdbCfg).Return(true).Times(1)
		dialectStrategy.EXPECT().Create(&rdbCfg).Return(dialector, nil).Times(1)
		_ = container.Service("dialect_strategy", func() rdb.DialectStrategy { return dialectStrategy }, rdb.DialectStrategyTag)
		_ = container.Service("id", func() (interface{}, error) { return nil, expected }, MigrationTag)

		_ = (&rdb.Provider{}).Boot(container)

		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid migration on retrieving migrations", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)

		rdbCfg := config.Partial{"dialect": "invalid", "host": ":memory:"}
		partial := config.Partial{
			"slate": config.Partial{
				"rdb": config.Partial{
					"connections": config.Partial{
						"primary": rdbCfg}}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		_ = container.Service(config.ID, func() *config.Config { return cfg })
		migrator := NewMockMigrator(ctrl)
		migrator.EXPECT().AutoMigrate(gomock.Any()).Return(nil).Times(1)
		dialector := NewMockDialector(ctrl)
		dialector.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
		dialector.EXPECT().Migrator(gomock.Any()).Return(migrator).Times(1)
		dialectStrategy := NewMockDialectStrategy(ctrl)
		dialectStrategy.EXPECT().Accept(&rdbCfg).Return(true).Times(1)
		dialectStrategy.EXPECT().Create(&rdbCfg).Return(dialector, nil).Times(1)
		_ = container.Service("dialect_strategy", func() rdb.DialectStrategy { return dialectStrategy }, rdb.DialectStrategyTag)
		_ = container.Service("id", func() (interface{}, error) { return "string", nil }, MigrationTag)

		_ = (&rdb.Provider{}).Boot(container)

		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("running migrator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)

		rdbCfg := config.Partial{"dialect": "invalid", "host": ":memory:"}
		partial := config.Partial{
			"slate": config.Partial{
				"rdb": config.Partial{
					"connections": config.Partial{
						"primary": rdbCfg}}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		_ = container.Service(config.ID, func() *config.Config { return cfg })
		migrator := NewMockMigrator(ctrl)
		migrator.EXPECT().AutoMigrate(gomock.Any()).Return(nil).Times(1)
		dialector := NewMockDialector(ctrl)
		dialector.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
		dialector.EXPECT().Migrator(gomock.Any()).Return(migrator).Times(1)
		dialectStrategy := NewMockDialectStrategy(ctrl)
		dialectStrategy.EXPECT().Accept(&rdbCfg).Return(true).Times(1)
		dialectStrategy.EXPECT().Create(&rdbCfg).Return(dialector, nil).Times(1)
		_ = container.Service("dialect_strategy", func() rdb.DialectStrategy { return dialectStrategy }, rdb.DialectStrategyTag)
		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(1)).Times(1)
		migration.EXPECT().Up().Times(1)
		_ = container.Service("id", func() (interface{}, error) { return migration, nil }, MigrationTag)

		_ = (&rdb.Provider{}).Boot(container)

		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		}
	})
}
