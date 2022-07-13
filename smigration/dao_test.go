package smigration

import (
	"database/sql/driver"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"reflect"
	"testing"
	"time"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func Test_NewDao(t *testing.T) {
	t.Run("error on auto migrate", func(t *testing.T) {
		db, dbmock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub rdb connection", err)
		}
		defer func() { _ = db.Close() }()

		expected := fmt.Errorf("error message")
		dbmock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("MariaDB"))
		dbmock.ExpectExec("CREATE TABLE `__version`").WillReturnError(expected)

		gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: gormLogger.Discard})

		dao, err := newDao(gdb)
		switch {
		case dao != nil:
			t.Error("return an unexpected valid reference to the dao instance")
		case err == nil:
			t.Error("didn't return the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error insead of the expected (%v)", err, expected)
		}
	})

	t.Run("construct", func(t *testing.T) {
		db, dbmock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub rdb connection", err)
		}
		defer func() { _ = db.Close() }()

		dbmock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("MariaDB"))
		dbmock.ExpectExec("CREATE TABLE `__version`").WillReturnResult(sqlmock.NewResult(0, 0))

		gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: gormLogger.Discard})

		if dao, err := newDao(gdb); err != nil {
			t.Errorf("return the unexpected error : (%v)", err)
		} else if dao == nil {
			t.Error("didn't return the expected parser instance")
		}
	})
}

func Test_Dao_Last(t *testing.T) {
	t.Run("error retrieve last entry", func(t *testing.T) {
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
		if _, err := dao.Last(); err == nil {
			t.Error("didn't returned the expected error")
		} else if check := err.Error(); check != errMsg {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", check, errMsg)
		}
	})

	t.Run("retrieve empty model if no last entry", func(t *testing.T) {
		db, dbmock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub rdb connection", err)
		}
		defer func() { _ = db.Close() }()

		dbmock.
			ExpectQuery("SELECT \\* FROM `__version`").
			WillReturnRows(sqlmock.NewRows(nil))

		gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: gormLogger.Discard})

		dao := &Dao{db: gdb}
		if record, err := dao.Last(); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		} else if record.ID != 0 {
			t.Error("return an unexpected last record")
		}
	})

	t.Run("retrieve last entry", func(t *testing.T) {
		db, dbmock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub rdb connection", err)
		}
		defer func() { _ = db.Close() }()

		id := uint(12)
		version := uint64(23)
		dbmock.
			ExpectQuery("SELECT \\* FROM `__version`").
			WillReturnRows(sqlmock.NewRows([]string{"id", "version", "created_at", "updated_at", "deleted_at"}).AddRow(id, version, nil, nil, nil))
		gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: gormLogger.Discard})

		dao := &Dao{db: gdb}
		if record, err := dao.Last(); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		} else if record.ID != id || record.Version != version {
			t.Error("return an invalid last record")
		}
	})
}

func Test_Dao_Up(t *testing.T) {
	t.Run("error updating last entry", func(t *testing.T) {
		db, dbmock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub rdb connection", err)
		}
		defer func() { _ = db.Close() }()

		version := uint64(12)
		errMsg := "error message"
		dbmock.ExpectBegin()
		dbmock.
			ExpectExec("INSERT INTO `__version`").
			WillReturnError(fmt.Errorf("%s", errMsg))
		dbmock.ExpectRollback()

		gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: gormLogger.Discard})

		dao := &Dao{db: gdb}
		if _, err := dao.Up(version); err == nil {
			t.Error("didn't returned the expected error")
		} else if check := err.Error(); check != errMsg {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", check, errMsg)
		}
	})

	t.Run("updating last entry", func(t *testing.T) {
		db, dbmock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub rdb connection", err)
		}
		defer func() { _ = db.Close() }()

		version := uint64(12)
		dbmock.ExpectBegin()
		dbmock.
			ExpectExec("INSERT INTO `__version`").
			WithArgs(version, AnyTime{}, AnyTime{}, nil).
			WillReturnResult(sqlmock.NewResult(1, 1))
		dbmock.ExpectCommit()

		gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: gormLogger.Discard})

		dao := &Dao{db: gdb}
		if record, err := dao.Up(version); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		} else if reflect.DeepEqual(Record{Version: version}, record) {
			t.Errorf("returned the unexpected model (%v) when expecting : %v", record, Record{Version: version})
		}
	})
}

func Test_Dao_Down(t *testing.T) {
	t.Run("error removing last entry", func(t *testing.T) {
		db, dbmock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub rdb connection", err)
		}
		defer func() { _ = db.Close() }()

		id := uint(12)
		version := uint64(23)
		errMsg := "error message"
		dbmock.ExpectBegin()
		dbmock.
			ExpectExec("DELETE FROM `__version` WHERE `__version`.`id` = ?").
			WithArgs(id).
			WillReturnError(fmt.Errorf("%s", errMsg))
		dbmock.ExpectRollback()

		gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: gormLogger.Discard})

		dao := &Dao{db: gdb}
		if err := dao.Down(Record{ID: id, Version: version}); err == nil {
			t.Error("didn't returned the expected error")
		} else if check := err.Error(); check != errMsg {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", check, errMsg)
		}
	})

	t.Run("removing last entry", func(t *testing.T) {
		db, dbmock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub rdb connection", err)
		}
		defer func() { _ = db.Close() }()

		id := uint(12)
		version := uint64(23)
		dbmock.ExpectBegin()
		dbmock.
			ExpectExec("DELETE FROM `__version` WHERE `__version`.`id` = ?").
			WithArgs(id).
			WillReturnResult(sqlmock.NewResult(1, 1))
		dbmock.ExpectCommit()

		gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: gormLogger.Discard})

		dao := &Dao{db: gdb}
		if err := dao.Down(Record{ID: id, Version: version}); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		}
	})
}
