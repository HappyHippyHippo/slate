package smigration

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serror"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"reflect"
	"testing"
)

func Test_NewMigrator(t *testing.T) {
	t.Run("nil dao", func(t *testing.T) {
		check, err := newMigrator(nil)
		switch {
		case check != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("new yaml decoder adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dao := &Dao{}

		check, err := newMigrator(dao)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case check == nil:
			t.Error("didn't return the expected migrator instance")
		case !reflect.DeepEqual(check.(*migrator).dao, dao):
			t.Error("didn't stored the given dao")
		}
	})
}

func Test_Migrator_AddMigration(t *testing.T) {
	t.Run("error adding nil migration", func(t *testing.T) {
		sut, _ := newMigrator(&Dao{})

		if err := sut.AddMigration(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("adding migration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		migration := NewMockMigration(ctrl)
		sut, _ := newMigrator(&Dao{})

		if err := sut.AddMigration(migration); err != nil {
			t.Errorf("returned the unexpected (%v) error", err)
		} else if !reflect.DeepEqual(sut.(*migrator).migrations, []IMigration{migration}) {
			t.Error("didn't stored the registering migration")
		}
	})
}

func Test_Migrator_Current(t *testing.T) {
	t.Run("error retrieving the last migration version", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db, dbmock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub rdb connection", err)
		}
		defer func() { _ = db.Close() }()

		errMsg := "error message"
		dbmock.
			ExpectQuery("SELECT \\* FROM `__version`").
			WillReturnError(fmt.Errorf("%s", errMsg))
		gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: gormLogger.Discard})
		dao := &Dao{db: gdb}

		sut, _ := newMigrator(dao)

		if _, err := sut.Current(); err == nil {
			t.Error("didn't returned the expected error")
		} else if check := err.Error(); check != errMsg {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", check, errMsg)
		}
	})

	t.Run("retrieves the last migration register version", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		version := uint64(23)
		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{Logger: gormLogger.Discard})
		_ = db.AutoMigrate(&Record{})
		_ = db.Create(&Record{Version: version})
		dao := &Dao{db}

		sut, _ := newMigrator(dao)

		if current, err := sut.Current(); err != nil {
			t.Errorf("returned the unexpected (%v) error", err)
		} else if current != version {
			t.Errorf("unexpectedly returned the %v version when expecting %v", current, version)
		}
	})
}

func Test_Migrator_Migrate(t *testing.T) {
	t.Run("no-op when no migrations", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dao := &Dao{}
		sut, _ := newMigrator(dao)

		if err := sut.Migrate(); err != nil {
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})

	t.Run("error while retrieving the last version", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db, dbmock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub rdb connection", err)
		}
		defer func() { _ = db.Close() }()

		errMsg := "error message"
		dbmock.
			ExpectQuery("SELECT \\* FROM `__version`").
			WillReturnError(fmt.Errorf("%s", errMsg))
		gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: gormLogger.Discard})
		dao := &Dao{db: gdb}

		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(1)).Times(0)
		sut, _ := newMigrator(dao)
		_ = sut.AddMigration(migration)

		if err := sut.Migrate(); err == nil {
			t.Error("didn't returned the expected error")
		} else if check := err.Error(); check != errMsg {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", check, errMsg)
		}
	})

	t.Run("no-op if no migrations higher than current last stored", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{Logger: gormLogger.Discard})
		_ = db.AutoMigrate(&Record{})
		_ = db.Create(&Record{Version: 2})
		dao := &Dao{db}

		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(1)).Times(1)
		sut, _ := newMigrator(dao)
		_ = sut.AddMigration(migration)

		if err := sut.Migrate(); err != nil {
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})

	t.Run("error while executing a migration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{Logger: gormLogger.Discard})
		_ = db.AutoMigrate(&Record{})
		dao := &Dao{db}

		expected := fmt.Errorf("error message")
		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(1)).Times(1)
		migration.EXPECT().Up().Return(expected).Times(1)
		sut, _ := newMigrator(dao)
		_ = sut.AddMigration(migration)

		if check := sut.Migrate(); check == nil {
			t.Error("didn't returned the expected error")
		} else if !reflect.DeepEqual(check, expected) {
			t.Errorf("returned the (%v) error insead of the expected (%v)", check, expected)
		}
	})

	t.Run("error while raising the migration version", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db, dbmock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub rdb connection", err)
		}
		defer func() { _ = db.Close() }()

		errMsg := "error message"
		dbmock.
			ExpectQuery("SELECT \\* FROM `__version`").
			WillReturnRows(sqlmock.NewRows([]string{"id", "version", "created_at", "updated_at", "deleted_at"}).AddRow(1, 1, nil, nil, nil))
		dbmock.ExpectBegin()
		dbmock.
			ExpectExec("INSERT INTO `__version`").
			WillReturnError(fmt.Errorf("%s", errMsg))
		dbmock.ExpectRollback()
		gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: gormLogger.Discard})
		dao := &Dao{db: gdb}

		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(2)).Times(1)
		migration.EXPECT().Up().Return(nil).Times(1)
		sut, _ := newMigrator(dao)
		_ = sut.AddMigration(migration)

		if check := sut.Migrate(); check == nil {
			t.Error("didn't returned the expected error")
		} else if check.Error() != errMsg {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", check, errMsg)
		}
	})

	t.Run("execute missing migrations (in order)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{Logger: gormLogger.Discard})
		_ = db.AutoMigrate(&Record{})
		_ = db.Create(&Record{Version: 1})
		dao := &Dao{db}

		migration1 := NewMockMigration(ctrl)
		migration1.EXPECT().Version().Return(uint64(1)).AnyTimes()
		migration2 := NewMockMigration(ctrl)
		migration2.EXPECT().Version().Return(uint64(2)).AnyTimes()
		migration2.EXPECT().Up().Return(nil).Times(1)
		migration3 := NewMockMigration(ctrl)
		migration3.EXPECT().Version().Return(uint64(3)).AnyTimes()
		migration3.EXPECT().Up().Return(nil).Times(1)
		sut, _ := newMigrator(dao)
		_ = sut.AddMigration(migration3)
		_ = sut.AddMigration(migration1)
		_ = sut.AddMigration(migration2)

		if err := sut.Migrate(); err != nil {
			t.Errorf("returned the unexpected (%v) error", err)
		} else if current, _ := sut.Current(); current != uint64(3) {
			t.Errorf("didn't terminated with the expected version of %v", uint64(3))
		}
	})
}

func Test_Migrator_Up(t *testing.T) {
	t.Run("no-op when no migrations", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dao := &Dao{}
		sut, _ := newMigrator(dao)

		if err := sut.Up(); err != nil {
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})

	t.Run("error while retrieving the last version", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db, dbmock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub rdb connection", err)
		}
		defer func() { _ = db.Close() }()

		errMsg := "error message"
		dbmock.
			ExpectQuery("SELECT \\* FROM `__version`").
			WillReturnError(fmt.Errorf("%s", errMsg))
		gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: gormLogger.Discard})
		dao := &Dao{db: gdb}

		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(1)).Times(0)
		sut, _ := newMigrator(dao)
		_ = sut.AddMigration(migration)

		if err := sut.Up(); err == nil {
			t.Error("didn't returned the expected error")
		} else if check := err.Error(); check != errMsg {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", check, errMsg)
		}
	})

	t.Run("no-op if no migrations higher than current last stored", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{Logger: gormLogger.Discard})
		_ = db.AutoMigrate(&Record{})
		_ = db.Create(&Record{Version: 2})
		dao := &Dao{db}

		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(1)).Times(1)
		sut, _ := newMigrator(dao)
		_ = sut.AddMigration(migration)

		if err := sut.Up(); err != nil {
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})

	t.Run("error while executing a migration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{Logger: gormLogger.Discard})
		_ = db.AutoMigrate(&Record{})
		dao := &Dao{db}

		expected := fmt.Errorf("error message")
		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(1)).Times(1)
		migration.EXPECT().Up().Return(expected).Times(1)
		sut, _ := newMigrator(dao)
		_ = sut.AddMigration(migration)

		if check := sut.Up(); check == nil {
			t.Error("didn't returned the expected error")
		} else if !reflect.DeepEqual(check, expected) {
			t.Errorf("returned the (%v) error insead of the expected (%v)", check, expected)
		}
	})

	t.Run("error while raising the migration version", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db, dbmock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub rdb connection", err)
		}
		defer func() { _ = db.Close() }()

		errMsg := "error message"
		dbmock.
			ExpectQuery("SELECT \\* FROM `__version`").
			WillReturnRows(sqlmock.NewRows([]string{"id", "version", "created_at", "updated_at", "deleted_at"}).AddRow(1, 1, nil, nil, nil))
		dbmock.ExpectBegin()
		dbmock.
			ExpectExec("INSERT INTO `__version`").
			WillReturnError(fmt.Errorf("%s", errMsg))
		dbmock.ExpectRollback()
		gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: gormLogger.Discard})
		dao := &Dao{db: gdb}

		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(2)).Times(1)
		migration.EXPECT().Up().Return(nil).Times(1)
		sut, _ := newMigrator(dao)
		_ = sut.AddMigration(migration)

		if check := sut.Up(); check == nil {
			t.Error("didn't returned the expected error")
		} else if check.Error() != errMsg {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", check, errMsg)
		}
	})

	t.Run("execute next migration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{Logger: gormLogger.Discard})
		_ = db.AutoMigrate(&Record{})
		_ = db.Create(&Record{Version: 1})
		dao := &Dao{db}

		migration1 := NewMockMigration(ctrl)
		migration1.EXPECT().Version().Return(uint64(1)).AnyTimes()
		migration2 := NewMockMigration(ctrl)
		migration2.EXPECT().Version().Return(uint64(2)).AnyTimes()
		migration2.EXPECT().Up().Return(nil).Times(1)
		migration3 := NewMockMigration(ctrl)
		migration3.EXPECT().Version().Return(uint64(3)).AnyTimes()
		sut, _ := newMigrator(dao)
		_ = sut.AddMigration(migration3)
		_ = sut.AddMigration(migration1)
		_ = sut.AddMigration(migration2)

		if err := sut.Up(); err != nil {
			t.Errorf("returned the unexpected (%v) error", err)
		} else if current, _ := sut.Current(); current != uint64(2) {
			t.Errorf("didn't terminated with the expected version of %v", uint64(2))
		}
	})
}

func Test_Migrator_Down(t *testing.T) {
	t.Run("no-op when no migrations", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dao := &Dao{}
		sut, _ := newMigrator(dao)

		if err := sut.Down(); err != nil {
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})

	t.Run("error while retrieving the last version", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db, dbmock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub rdb connection", err)
		}
		defer func() { _ = db.Close() }()

		errMsg := "error message"
		dbmock.
			ExpectQuery("SELECT \\* FROM `__version`").
			WillReturnError(fmt.Errorf("%s", errMsg))
		gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: gormLogger.Discard})
		dao := &Dao{db: gdb}

		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(1)).Times(0)
		sut, _ := newMigrator(dao)
		_ = sut.AddMigration(migration)

		if err := sut.Down(); err == nil {
			t.Error("didn't returned the expected error")
		} else if check := err.Error(); check != errMsg {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", check, errMsg)
		}
	})

	t.Run("error while executing a migration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := uint(12)
		version := uint64(34)
		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{Logger: gormLogger.Discard})
		_ = db.AutoMigrate(&Record{})
		_ = db.Create(&Record{ID: id, Version: version})
		dao := &Dao{db}

		expected := fmt.Errorf("error message")
		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(version).Times(1)
		migration.EXPECT().Down().Return(expected).Times(1)
		sut, _ := newMigrator(dao)
		_ = sut.AddMigration(migration)

		if check := sut.Down(); check == nil {
			t.Error("didn't returned the expected error")
		} else if !reflect.DeepEqual(check, expected) {
			t.Errorf("returned the (%v) error insead of the expected (%v)", check, expected)
		}
	})

	t.Run("error while removing the migration version", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db, dbmock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub rdb connection", err)
		}
		defer func() { _ = db.Close() }()

		id := uint(12)
		version := uint64(34)
		errMsg := "error message"
		dbmock.
			ExpectQuery("SELECT \\* FROM `__version`").
			WillReturnRows(sqlmock.NewRows([]string{"id", "version", "created_at", "updated_at", "deleted_at"}).AddRow(id, version, nil, nil, nil))
		dbmock.ExpectBegin()
		dbmock.
			ExpectExec("DELETE FROM `__version` WHERE `__version`.`id` = ?").
			WithArgs(id).
			WillReturnError(fmt.Errorf("%s", errMsg))
		dbmock.ExpectRollback()
		gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: gormLogger.Discard})
		dao := &Dao{db: gdb}

		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(version).Times(1)
		migration.EXPECT().Down().Return(nil).Times(1)
		sut, _ := newMigrator(dao)
		_ = sut.AddMigration(migration)

		if check := sut.Down(); check == nil {
			t.Error("didn't returned the expected error")
		} else if check.Error() != errMsg {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", check, errMsg)
		}
	})

	t.Run("no-op if the migration was not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{Logger: gormLogger.Discard})
		_ = db.AutoMigrate(&Record{})
		_ = db.Create(&Record{Version: 1})
		_ = db.Create(&Record{Version: 2})
		dao := &Dao{db}

		migration1 := NewMockMigration(ctrl)
		migration1.EXPECT().Version().Return(uint64(1)).AnyTimes()
		sut, _ := newMigrator(dao)
		_ = sut.AddMigration(migration1)

		if err := sut.Down(); err != nil {
			t.Errorf("returned the unexpected (%v) error", err)
		} else if current, _ := sut.Current(); current != uint64(2) {
			t.Errorf("didn't terminated with the expected version of %v", uint64(2))
		}
	})

	t.Run("remove migration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{Logger: gormLogger.Discard})
		_ = db.AutoMigrate(&Record{})
		_ = db.Create(&Record{Version: 1})
		_ = db.Create(&Record{Version: 2})
		dao := &Dao{db}

		migration1 := NewMockMigration(ctrl)
		migration1.EXPECT().Version().Return(uint64(1)).AnyTimes()
		migration2 := NewMockMigration(ctrl)
		migration2.EXPECT().Version().Return(uint64(2)).AnyTimes()
		migration2.EXPECT().Down().Return(nil).Times(1)
		migration3 := NewMockMigration(ctrl)
		migration3.EXPECT().Version().Return(uint64(3)).AnyTimes()
		sut, _ := newMigrator(dao)
		_ = sut.AddMigration(migration3)
		_ = sut.AddMigration(migration1)
		_ = sut.AddMigration(migration2)

		if err := sut.Down(); err != nil {
			t.Errorf("returned the unexpected (%v) error", err)
		} else if current, _ := sut.Current(); current != uint64(1) {
			t.Errorf("didn't terminated with the expected version of %v", uint64(1))
		}
	})
}
