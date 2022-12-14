package migration

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func Test_NewDao(t *testing.T) {
	t.Run("error on auto migrate", func(t *testing.T) {
		db, mockDB, e := sqlmock.New()
		if e != nil {
			t.Fatalf("anerror '%s' was not expected when opening a stub rdb connection", e)
		}
		defer func() { _ = db.Close() }()

		expected := fmt.Errorf("error message")
		mockDB.
			ExpectQuery("SELECT VERSION()").
			WillReturnRows(sqlmock.
				NewRows([]string{"version"}).
				AddRow("MariaDB"))
		mockDB.
			ExpectExec("CREATE TABLE `__version`").
			WillReturnError(expected)

		gdb, e := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: logger.Discard})

		sut, e := NewDao(gdb)
		switch {
		case sut != nil:
			t.Error("return an unexpected valid reference to the dao instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error insead of the expected (%v)", e, expected)
		}
	})

	t.Run("construct", func(t *testing.T) {
		db, mockDB, e := sqlmock.New()
		if e != nil {
			t.Fatalf("anerror '%s' was not expected when opening a stub rdb connection", e)
		}
		defer func() { _ = db.Close() }()

		mockDB.
			ExpectQuery("SELECT VERSION()").
			WillReturnRows(sqlmock.
				NewRows([]string{"version"}).
				AddRow("MariaDB"))
		mockDB.
			ExpectExec("CREATE TABLE `__version`").
			WillReturnResult(sqlmock.
				NewResult(0, 0))

		gdb, e := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: logger.Discard})

		if sut, e := NewDao(gdb); e != nil {
			t.Errorf("return the unexpected error : (%v)", e)
		} else if sut == nil {
			t.Error("didn't return the expected parser instance")
		}
	})
}

func Test_Dao_Last(t *testing.T) {
	t.Run("error retrieve last entry", func(t *testing.T) {
		db, mockDB, e := sqlmock.New()
		if e != nil {
			t.Fatalf("anerror '%s' was not expected when opening a stub rdb connection", e)
		}
		defer func() { _ = db.Close() }()

		expected := fmt.Errorf("error message")
		mockDB.
			ExpectQuery("SELECT VERSION()").
			WillReturnRows(sqlmock.
				NewRows([]string{"version"}).
				AddRow("MariaDB"))
		mockDB.
			ExpectExec("CREATE TABLE `__version`").
			WillReturnResult(sqlmock.
				NewResult(0, 0))
		mockDB.
			ExpectQuery("SELECT \\* FROM `__version`").
			WillReturnError(expected)

		gdb, e := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: logger.Discard})

		sut, _ := NewDao(gdb)
		if _, e := sut.Last(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", e, expected)
		}
	})

	t.Run("retrieve empty model if no last entry", func(t *testing.T) {
		db, mockDB, e := sqlmock.New()
		if e != nil {
			t.Fatalf("anerror '%s' was not expected when opening a stub rdb connection", e)
		}
		defer func() { _ = db.Close() }()

		mockDB.
			ExpectQuery("SELECT VERSION()").
			WillReturnRows(sqlmock.
				NewRows([]string{"version"}).
				AddRow("MariaDB"))
		mockDB.
			ExpectExec("CREATE TABLE `__version`").
			WillReturnResult(sqlmock.
				NewResult(0, 0))
		mockDB.
			ExpectQuery("SELECT \\* FROM `__version`").
			WillReturnRows(sqlmock.
				NewRows(nil))

		gdb, e := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: logger.Discard})

		sut, _ := NewDao(gdb)
		if record, e := sut.Last(); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		} else if record.ID != 0 {
			t.Error("return an unexpected last record")
		}
	})

	t.Run("retrieve last entry", func(t *testing.T) {
		db, mockDB, e := sqlmock.New()
		if e != nil {
			t.Fatalf("anerror '%s' was not expected when opening a stub rdb connection", e)
		}
		defer func() { _ = db.Close() }()

		id := uint(12)
		version := uint64(23)
		mockDB.
			ExpectQuery("SELECT VERSION()").
			WillReturnRows(sqlmock.
				NewRows([]string{"version"}).
				AddRow("MariaDB"))
		mockDB.
			ExpectExec("CREATE TABLE `__version`").
			WillReturnResult(sqlmock.
				NewResult(0, 0))
		mockDB.
			ExpectQuery("SELECT \\* FROM `__version`").
			WillReturnRows(sqlmock.
				NewRows([]string{"id", "version", "created_at", "updated_at", "deleted_at"}).
				AddRow(id, version, nil, nil, nil))
		gdb, e := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: logger.Discard})

		sut, _ := NewDao(gdb)
		if record, e := sut.Last(); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		} else if record.ID != id || record.Version != version {
			t.Error("return an invalid last record")
		}
	})
}

func Test_Dao_Up(t *testing.T) {
	t.Run("error updating last entry", func(t *testing.T) {
		db, mockDB, e := sqlmock.New()
		if e != nil {
			t.Fatalf("anerror '%s' was not expected when opening a stub rdb connection", e)
		}
		defer func() { _ = db.Close() }()

		version := uint64(12)
		errMsg := "error message"
		mockDB.
			ExpectQuery("SELECT VERSION()").
			WillReturnRows(sqlmock.
				NewRows([]string{"version"}).
				AddRow("MariaDB"))
		mockDB.
			ExpectExec("CREATE TABLE `__version`").
			WillReturnResult(sqlmock.
				NewResult(0, 0))
		mockDB.ExpectBegin()
		mockDB.
			ExpectExec("INSERT INTO `__version`").
			WillReturnError(fmt.Errorf("%s", errMsg))
		mockDB.ExpectRollback()

		gdb, e := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: logger.Discard})

		sut, _ := NewDao(gdb)
		if _, e := sut.Up(version); e == nil {
			t.Error("didn't returned the expected error")
		} else if check := e.Error(); check != errMsg {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", check, errMsg)
		}
	})

	t.Run("updating last entry", func(t *testing.T) {
		db, mockDB, e := sqlmock.New()
		if e != nil {
			t.Fatalf("anerror '%s' was not expected when opening a stub rdb connection", e)
		}
		defer func() { _ = db.Close() }()

		version := uint64(12)
		mockDB.
			ExpectQuery("SELECT VERSION()").
			WillReturnRows(sqlmock.
				NewRows([]string{"version"}).
				AddRow("MariaDB"))
		mockDB.
			ExpectExec("CREATE TABLE `__version`").
			WillReturnResult(sqlmock.
				NewResult(0, 0))
		mockDB.ExpectBegin()
		mockDB.
			ExpectExec("INSERT INTO `__version`").
			WithArgs(version, AnyTime{}, AnyTime{}, nil).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mockDB.ExpectCommit()

		gdb, e := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: logger.Discard})

		sut, _ := NewDao(gdb)
		if record, e := sut.Up(version); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		} else if reflect.DeepEqual(Record{Version: version}, record) {
			t.Errorf("returned the unexpected model (%v) when expecting : %v", record, Record{Version: version})
		}
	})
}

func Test_Dao_Down(t *testing.T) {
	t.Run("error removing last entry", func(t *testing.T) {
		db, mockDB, e := sqlmock.New()
		if e != nil {
			t.Fatalf("anerror '%s' was not expected when opening a stub rdb connection", e)
		}
		defer func() { _ = db.Close() }()

		id := uint(12)
		version := uint64(23)
		errMsg := "error message"
		mockDB.
			ExpectQuery("SELECT VERSION()").
			WillReturnRows(sqlmock.
				NewRows([]string{"version"}).
				AddRow("MariaDB"))
		mockDB.
			ExpectExec("CREATE TABLE `__version`").
			WillReturnResult(sqlmock.
				NewResult(0, 0))
		mockDB.ExpectBegin()
		mockDB.
			ExpectExec("DELETE FROM `__version` WHERE `__version`.`id` = ?").
			WithArgs(id).
			WillReturnError(fmt.Errorf("%s", errMsg))
		mockDB.ExpectRollback()

		gdb, e := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: logger.Discard})

		sut, _ := NewDao(gdb)
		if e := sut.Down(Record{ID: id, Version: version}); e == nil {
			t.Error("didn't returned the expected error")
		} else if check := e.Error(); check != errMsg {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", check, errMsg)
		}
	})

	t.Run("removing last entry", func(t *testing.T) {
		db, mockDB, e := sqlmock.New()
		if e != nil {
			t.Fatalf("anerror '%s' was not expected when opening a stub rdb connection", e)
		}
		defer func() { _ = db.Close() }()

		id := uint(12)
		version := uint64(23)
		mockDB.
			ExpectQuery("SELECT VERSION()").
			WillReturnRows(sqlmock.
				NewRows([]string{"version"}).
				AddRow("MariaDB"))
		mockDB.
			ExpectExec("CREATE TABLE `__version`").
			WillReturnResult(sqlmock.
				NewResult(0, 0))
		mockDB.ExpectBegin()
		mockDB.
			ExpectExec("DELETE FROM `__version` WHERE `__version`.`id` = ?").
			WithArgs(id).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mockDB.ExpectCommit()

		gdb, e := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{Logger: logger.Discard})

		sut, _ := NewDao(gdb)
		if e := sut.Down(Record{ID: id, Version: version}); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		}
	})
}
