package slate

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
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

func Test_MigratorDao(t *testing.T) {
	t.Run("NewMigratorDao", func(t *testing.T) {
		t.Run("nil db", func(t *testing.T) {
			sut, e := NewMigratorDao(nil)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("error on auto migrate", func(t *testing.T) {
			db, mockDB, e := sqlmock.New()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			expected := fmt.Errorf("error message")
			mockDB.
				ExpectQuery("SELECT VERSION()").
				WillReturnRows(sqlmock.
					NewRows([]string{"version"}).
					AddRow("MariaDB"))
			mockDB.
				ExpectExec("CREATE TABLE `__migrations`").
				WillReturnError(expected)

			gdb, e := gorm.Open(
				mysql.New(mysql.Config{Conn: db}),
				&gorm.Config{Logger: logger.Discard},
			)

			sut, e := NewMigratorDao(gdb)
			switch {
			case sut != nil:
				t.Error("unexpected valid reference")
			case e == nil:
				t.Error("didn't return the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			default:
				if e := mockDB.ExpectationsWereMet(); e != nil {
					t.Errorf("DB interactions expectations when met : %v", e)
				}
			}
		})

		t.Run("construct", func(t *testing.T) {
			db, mockDB, e := sqlmock.New()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			mockDB.
				ExpectQuery("SELECT VERSION()").
				WillReturnRows(sqlmock.
					NewRows([]string{"version"}).
					AddRow("MariaDB"))
			mockDB.
				ExpectExec("CREATE TABLE `__migrations`").
				WillReturnResult(sqlmock.
					NewResult(0, 0))

			gdb, e := gorm.Open(
				mysql.New(mysql.Config{Conn: db}),
				&gorm.Config{Logger: logger.Discard},
			)

			if sut, e := NewMigratorDao(gdb); e != nil {
				t.Errorf("return the unexpected error : (%v)", e)
			} else if sut == nil {
				t.Error("didn't return the expected parser instance")
			} else if e := mockDB.ExpectationsWereMet(); e != nil {
				t.Errorf("DB interactions expectations when met : %v", e)
			}
		})
	})

	t.Run("Last", func(t *testing.T) {
		t.Run("error retrieve last entry", func(t *testing.T) {
			db, mockDB, e := sqlmock.New()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			expected := fmt.Errorf("error message")
			mockDB.
				ExpectQuery("SELECT VERSION()").
				WillReturnRows(sqlmock.
					NewRows([]string{"version"}).
					AddRow("MariaDB"))
			mockDB.
				ExpectExec("CREATE TABLE `__migrations`").
				WillReturnResult(sqlmock.
					NewResult(0, 0))
			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnError(expected)

			gdb, e := gorm.Open(
				mysql.New(mysql.Config{Conn: db}),
				&gorm.Config{Logger: logger.Discard},
			)

			sut, _ := NewMigratorDao(gdb)
			if _, e := sut.Last(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			} else if e := mockDB.ExpectationsWereMet(); e != nil {
				t.Errorf("DB interactions expectations when met : %v", e)
			}
		})

		t.Run("retrieve empty model if no last entry", func(t *testing.T) {
			db, mockDB, e := sqlmock.New()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			mockDB.
				ExpectQuery("SELECT VERSION()").
				WillReturnRows(sqlmock.
					NewRows([]string{"version"}).
					AddRow("MariaDB"))
			mockDB.
				ExpectExec("CREATE TABLE `__migrations`").
				WillReturnResult(sqlmock.
					NewResult(0, 0))
			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnRows(sqlmock.
					NewRows(nil))

			gdb, e := gorm.Open(
				mysql.New(mysql.Config{Conn: db}),
				&gorm.Config{Logger: logger.Discard},
			)

			sut, _ := NewMigratorDao(gdb)
			if record, e := sut.Last(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if record.ID != 0 {
				t.Error("return an unexpected last record")
			} else if e := mockDB.ExpectationsWereMet(); e != nil {
				t.Errorf("DB interactions expectations when met : %v", e)
			}
		})

		t.Run("retrieve last entry", func(t *testing.T) {
			db, mockDB, e := sqlmock.New()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
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
				ExpectExec("CREATE TABLE `__migrations`").
				WillReturnResult(sqlmock.
					NewResult(0, 0))
			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "version", "created_at", "updated_at", "deleted_at"}).
					AddRow(id, version, nil, nil, nil))

			gdb, e := gorm.Open(
				mysql.New(mysql.Config{Conn: db}),
				&gorm.Config{Logger: logger.Discard},
			)

			sut, _ := NewMigratorDao(gdb)
			if record, e := sut.Last(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if record.ID != id || record.Version != version {
				t.Error("return an invalid last record")
			} else if e := mockDB.ExpectationsWereMet(); e != nil {
				t.Errorf("DB interactions expectations when met : %v", e)
			}
		})
	})

	t.Run("Up", func(t *testing.T) {
		t.Run("error updating last entry", func(t *testing.T) {
			db, mockDB, e := sqlmock.New()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
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
				ExpectExec("CREATE TABLE `__migrations`").
				WillReturnResult(sqlmock.
					NewResult(0, 0))
			mockDB.ExpectBegin()
			mockDB.
				ExpectExec("INSERT INTO `__migrations`").
				WillReturnError(fmt.Errorf("%s", errMsg))
			mockDB.ExpectRollback()

			gdb, e := gorm.Open(
				mysql.New(mysql.Config{Conn: db}),
				&gorm.Config{Logger: logger.Discard},
			)

			sut, _ := NewMigratorDao(gdb)
			if _, e := sut.Up(version); e == nil {
				t.Error("didn't returned the expected error")
			} else if check := e.Error(); check != errMsg {
				t.Errorf("unexpected (%v) error instead of the expected (%v)", check, errMsg)
			} else if e := mockDB.ExpectationsWereMet(); e != nil {
				t.Errorf("DB interactions expectations when met : %v", e)
			}
		})

		t.Run("updating last entry", func(t *testing.T) {
			db, mockDB, e := sqlmock.New()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			version := uint64(12)
			mockDB.
				ExpectQuery("SELECT VERSION()").
				WillReturnRows(sqlmock.
					NewRows([]string{"version"}).
					AddRow("MariaDB"))
			mockDB.
				ExpectExec("CREATE TABLE `__migrations`").
				WillReturnResult(sqlmock.
					NewResult(0, 0))
			mockDB.ExpectBegin()
			mockDB.
				ExpectExec("INSERT INTO `__migrations`").
				WithArgs(version, AnyTime{}, AnyTime{}, nil).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mockDB.ExpectCommit()

			gdb, e := gorm.Open(
				mysql.New(mysql.Config{Conn: db}),
				&gorm.Config{Logger: logger.Discard},
			)

			sut, _ := NewMigratorDao(gdb)
			if record, e := sut.Up(version); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if reflect.DeepEqual(MigratorRecord{Version: version}, record) {
				t.Errorf("(%v) when expecting (%v)", record, MigratorRecord{Version: version})
			} else if e := mockDB.ExpectationsWereMet(); e != nil {
				t.Errorf("DB interactions expectations when met : %v", e)
			}
		})
	})

	t.Run("Down", func(t *testing.T) {
		t.Run("error removing last entry", func(t *testing.T) {
			db, mockDB, e := sqlmock.New()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
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
				ExpectExec("CREATE TABLE `__migrations`").
				WillReturnResult(sqlmock.
					NewResult(0, 0))
			mockDB.ExpectBegin()
			mockDB.
				ExpectExec("DELETE FROM `__migrations` WHERE `__migrations`.`id` = ?").
				WithArgs(id).
				WillReturnError(fmt.Errorf("%s", errMsg))
			mockDB.ExpectRollback()

			gdb, e := gorm.Open(
				mysql.New(mysql.Config{Conn: db}),
				&gorm.Config{Logger: logger.Discard},
			)

			sut, _ := NewMigratorDao(gdb)
			if e := sut.Down(MigratorRecord{ID: id, Version: version}); e == nil {
				t.Error("didn't returned the expected error")
			} else if check := e.Error(); check != errMsg {
				t.Errorf("(%v) when expecting (%v)", check, errMsg)
			} else if e := mockDB.ExpectationsWereMet(); e != nil {
				t.Errorf("DB interactions expectations when met : %v", e)
			}
		})

		t.Run("removing last entry", func(t *testing.T) {
			db, mockDB, e := sqlmock.New()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
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
				ExpectExec("CREATE TABLE `__migrations`").
				WillReturnResult(sqlmock.
					NewResult(0, 0))
			mockDB.ExpectBegin()
			mockDB.
				ExpectExec("DELETE FROM `__migrations` WHERE `__migrations`.`id` = ?").
				WithArgs(id).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mockDB.ExpectCommit()

			gdb, e := gorm.Open(
				mysql.New(mysql.Config{Conn: db}),
				&gorm.Config{Logger: logger.Discard},
			)

			sut, _ := NewMigratorDao(gdb)
			if e := sut.Down(MigratorRecord{ID: id, Version: version}); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if e := mockDB.ExpectationsWereMet(); e != nil {
				t.Errorf("DB interactions expectations when met : %v", e)
			}
		})
	})
}

func Test_Migrator(t *testing.T) {
	createMockedDB := func() (*sql.DB, sqlmock.Sqlmock, *gorm.DB, error) {
		db, mockDB, e := sqlmock.New()
		if e != nil {
			return nil, nil, nil, e
		}

		mockDB.
			ExpectQuery("SELECT VERSION()").
			WillReturnRows(sqlmock.
				NewRows([]string{"version"}).
				AddRow("MariaDB"))
		mockDB.
			ExpectExec("CREATE TABLE `__migrations`").
			WillReturnResult(sqlmock.
				NewResult(0, 0))

		gdb, e := gorm.Open(
			mysql.New(mysql.Config{Conn: db}),
			&gorm.Config{Logger: logger.Discard},
		)
		if e != nil {
			return nil, nil, nil, e
		}

		return db, mockDB, gdb, nil
	}

	t.Run("NewMigrator", func(t *testing.T) {
		t.Run("nil dao", func(t *testing.T) {
			sut, e := NewMigrator(nil, nil)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("new migrator", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			dao := &MigratorDao{}
			migration := NewMockMigration(ctrl)
			migrations := []Migration{migration}

			sut, e := NewMigrator(dao, migrations)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case sut == nil:
				t.Error("didn't return the expected migrator instance")
			case !reflect.DeepEqual(sut.dao, dao):
				t.Error("didn't stored the given dao")
			case len(sut.migrations) != 1:
				t.Error("didn't stored the migration")
			case sut.migrations[0] != migration:
				t.Error("didn't stored the passed migration")
			}
		})

		t.Run("order the given migrations", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			dao := &MigratorDao{}
			migration1 := NewMockMigration(ctrl)
			migration1.EXPECT().Version().Return(uint64(1)).MinTimes(1)
			migration2 := NewMockMigration(ctrl)
			migration2.EXPECT().Version().Return(uint64(2)).MinTimes(1)
			migration3 := NewMockMigration(ctrl)
			migration3.EXPECT().Version().Return(uint64(3)).MinTimes(1)
			migrations := []Migration{migration3, migration1, migration2}

			sut, e := NewMigrator(dao, migrations)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case sut == nil:
				t.Error("didn't return the expected migrator instance")
			case len(sut.migrations) != 3:
				t.Error("didn't stored the migrations")
			case sut.migrations[0] != migration1:
				t.Error("didn't stored the migration 1 in the proper order")
			case sut.migrations[1] != migration2:
				t.Error("didn't stored the migration 2 in the proper order")
			case sut.migrations[2] != migration3:
				t.Error("didn't stored the migration 3 in the proper order")
			}
		})
	})

	t.Run("AddMigration", func(t *testing.T) {
		t.Run("error adding nil migration", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewMigrator(&MigratorDao{}, nil)

			if e := sut.AddMigration(nil); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("adding migration", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			migration := NewMockMigration(ctrl)
			sut, _ := NewMigrator(&MigratorDao{}, nil)

			if e := sut.AddMigration(migration); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if !reflect.DeepEqual(sut.migrations, []Migration{migration}) {
				t.Error("didn't stored the registering migration")
			}
		})

		t.Run("order the given migrations", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			dao := &MigratorDao{}
			migration1 := NewMockMigration(ctrl)
			migration1.EXPECT().Version().Return(uint64(1)).MinTimes(1)
			migration2 := NewMockMigration(ctrl)
			migration2.EXPECT().Version().Return(uint64(2)).MinTimes(1)
			migration3 := NewMockMigration(ctrl)
			migration3.EXPECT().Version().Return(uint64(3)).MinTimes(1)
			sut, _ := NewMigrator(dao, []Migration{})

			_ = sut.AddMigration(migration3)
			_ = sut.AddMigration(migration1)
			_ = sut.AddMigration(migration2)

			switch {
			case len(sut.migrations) != 3:
				t.Error("didn't stored the migrations")
			case sut.migrations[0] != migration1:
				t.Error("didn't stored the migration 1 in the proper order")
			case sut.migrations[1] != migration2:
				t.Error("didn't stored the migration 2 in the proper order")
			case sut.migrations[2] != migration3:
				t.Error("didn't stored the migration 3 in the proper order")
			}
		})
	})

	t.Run("Current", func(t *testing.T) {
		t.Run("error retrieving the last migration version", func(t *testing.T) {
			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			expected := fmt.Errorf("error message")
			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnError(expected)

			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, nil)

			if _, e := sut.Current(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})

		t.Run("retrieves the last migration register version", func(t *testing.T) {
			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			id := uint(12)
			version := uint64(23)
			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "version", "created_at", "updated_at", "deleted_at"}).
					AddRow(id, version, nil, nil, nil))

			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, nil)

			if current, e := sut.Current(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if current != version {
				t.Errorf("(%v) when expecting (%v)", current, version)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})
	})

	t.Run("Migrate", func(t *testing.T) {
		t.Run("no-op when no migrations", func(t *testing.T) {
			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, nil)

			if e := sut.Migrate(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})

		t.Run("error while retrieving the last version", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			expected := fmt.Errorf("error message")
			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnError(expected)

			migration := NewMockMigration(ctrl)
			migration.EXPECT().Version().Return(uint64(1)).Times(0)
			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, []Migration{migration})

			if e := sut.Migrate(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})

		t.Run("no-op if no migrations higher than current last stored", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "version", "created_at", "updated_at", "deleted_at"}).
					AddRow(1, 1, nil, nil, nil))

			migration := NewMockMigration(ctrl)
			migration.EXPECT().Version().Return(uint64(1)).Times(1)
			migration.EXPECT().Up().Return(nil).Times(0)
			migration.EXPECT().Down().Return(nil).Times(0)
			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, []Migration{migration})

			if e := sut.Migrate(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})

		t.Run("error while executing a migration", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			expected := fmt.Errorf("error message")
			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "version", "created_at", "updated_at", "deleted_at"}).
					AddRow(1, 1, nil, nil, nil))

			migration := NewMockMigration(ctrl)
			migration.EXPECT().Version().Return(uint64(2)).Times(1)
			migration.EXPECT().Up().Return(expected).Times(1)
			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, []Migration{migration})

			if e := sut.Migrate(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})

		t.Run("error while raising the migration version", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			expected := fmt.Errorf("error message")
			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "version", "created_at", "updated_at", "deleted_at"}).
					AddRow(1, 1, nil, nil, nil))
			mockDB.ExpectBegin()
			mockDB.
				ExpectExec("INSERT INTO `__migrations`").
				WithArgs(uint64(2), AnyTime{}, AnyTime{}, nil).
				WillReturnError(expected)
			mockDB.ExpectRollback()

			migration := NewMockMigration(ctrl)
			migration.EXPECT().Version().Return(uint64(2)).Times(1)
			migration.EXPECT().Up().Return(nil).Times(1)
			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, []Migration{migration})

			if e := sut.Migrate(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})

		t.Run("execute missing migrations (in order)", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "version", "created_at", "updated_at", "deleted_at"}).
					AddRow(1, 1, nil, nil, nil))
			mockDB.ExpectBegin()
			mockDB.
				ExpectExec("INSERT INTO `__migrations`").
				WithArgs(uint64(2), AnyTime{}, AnyTime{}, nil).
				WillReturnResult(sqlmock.NewResult(2, 1))
			mockDB.ExpectCommit()
			mockDB.ExpectBegin()
			mockDB.
				ExpectExec("INSERT INTO `__migrations`").
				WithArgs(uint64(3), AnyTime{}, AnyTime{}, nil).
				WillReturnResult(sqlmock.NewResult(3, 1))
			mockDB.ExpectCommit()

			migration1 := NewMockMigration(ctrl)
			migration1.EXPECT().Version().Return(uint64(1)).AnyTimes()
			migration1.EXPECT().Up().Return(nil).Times(0)
			migration2 := NewMockMigration(ctrl)
			migration2.EXPECT().Version().Return(uint64(2)).AnyTimes()
			migration2.EXPECT().Up().Return(nil).Times(1)
			migration3 := NewMockMigration(ctrl)
			migration3.EXPECT().Version().Return(uint64(3)).AnyTimes()
			migration3.EXPECT().Up().Return(nil).Times(1)
			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, []Migration{migration3, migration1, migration2})

			if e := sut.Migrate(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})
	})

	t.Run("Up", func(t *testing.T) {
		t.Run("no-op when no migrations", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, nil)

			if e := sut.Up(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})

		t.Run("error while retrieving the last version", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			expected := fmt.Errorf("error message")
			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnError(expected)

			migration := NewMockMigration(ctrl)
			migration.EXPECT().Version().Return(uint64(1)).Times(0)
			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, []Migration{migration})

			if e := sut.Up(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})

		t.Run("no-op if no migrations higher than current last stored", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "version", "created_at", "updated_at", "deleted_at"}).
					AddRow(1, 1, nil, nil, nil))

			migration := NewMockMigration(ctrl)
			migration.EXPECT().Version().Return(uint64(1)).Times(1)
			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, []Migration{migration})

			if e := sut.Up(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})

		t.Run("error while executing a migration", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			expected := fmt.Errorf("error message")
			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "version", "created_at", "updated_at", "deleted_at"}).
					AddRow(1, 1, nil, nil, nil))

			migration := NewMockMigration(ctrl)
			migration.EXPECT().Version().Return(uint64(2)).Times(1)
			migration.EXPECT().Up().Return(expected).Times(1)
			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, []Migration{migration})

			if e := sut.Up(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})

		t.Run("error while raising the migration version", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			expected := fmt.Errorf("error message")
			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "version", "created_at", "updated_at", "deleted_at"}).
					AddRow(1, 1, nil, nil, nil))
			mockDB.ExpectBegin()
			mockDB.
				ExpectExec("INSERT INTO `__migrations`").
				WithArgs(uint64(2), AnyTime{}, AnyTime{}, nil).
				WillReturnError(expected)
			mockDB.ExpectRollback()

			migration := NewMockMigration(ctrl)
			migration.EXPECT().Version().Return(uint64(2)).Times(1)
			migration.EXPECT().Up().Return(nil).Times(1)
			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, []Migration{migration})

			if e := sut.Up(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})

		t.Run("execute next migration", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "version", "created_at", "updated_at", "deleted_at"}).
					AddRow(1, 1, nil, nil, nil))
			mockDB.ExpectBegin()
			mockDB.
				ExpectExec("INSERT INTO `__migrations`").
				WithArgs(uint64(2), AnyTime{}, AnyTime{}, nil).
				WillReturnResult(sqlmock.NewResult(2, 1))
			mockDB.ExpectCommit()

			migration1 := NewMockMigration(ctrl)
			migration1.EXPECT().Version().Return(uint64(1)).AnyTimes()
			migration1.EXPECT().Up().Return(nil).Times(0)
			migration2 := NewMockMigration(ctrl)
			migration2.EXPECT().Version().Return(uint64(2)).AnyTimes()
			migration2.EXPECT().Up().Return(nil).Times(1)
			migration3 := NewMockMigration(ctrl)
			migration3.EXPECT().Version().Return(uint64(3)).AnyTimes()
			migration3.EXPECT().Up().Return(nil).Times(0)
			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, []Migration{migration3, migration1, migration2})

			if e := sut.Up(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})
	})

	t.Run("Down", func(t *testing.T) {
		t.Run("no-op when no migrations", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, nil)

			if e := sut.Down(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})

		t.Run("error while retrieving the last version", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			expected := fmt.Errorf("error message")
			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnError(expected)

			migration := NewMockMigration(ctrl)
			migration.EXPECT().Version().Return(uint64(1)).Times(0)
			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, []Migration{migration})

			if e := sut.Down(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})

		t.Run("error while executing a migration", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			expected := fmt.Errorf("error message")
			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "version", "created_at", "updated_at", "deleted_at"}).
					AddRow(1, 1, nil, nil, nil))

			migration := NewMockMigration(ctrl)
			migration.EXPECT().Version().Return(uint64(1)).Times(1)
			migration.EXPECT().Down().Return(expected).Times(1)
			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, []Migration{migration})

			if e := sut.Down(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})

		t.Run("error while removing the migration version", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			expected := fmt.Errorf("error message")
			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "version", "created_at", "updated_at", "deleted_at"}).
					AddRow(1, 1, nil, nil, nil))
			mockDB.ExpectBegin()
			mockDB.
				ExpectExec("DELETE FROM `__migrations`").
				WithArgs(uint64(1)).
				WillReturnError(expected)
			mockDB.ExpectRollback()

			migration := NewMockMigration(ctrl)
			migration.EXPECT().Version().Return(uint64(1)).Times(1)
			migration.EXPECT().Down().Return(nil).Times(1)
			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, []Migration{migration})

			if e := sut.Down(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})

		t.Run("no-op if the migration was not found", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "version", "created_at", "updated_at", "deleted_at"}).
					AddRow(2, 2, nil, nil, nil))

			migration := NewMockMigration(ctrl)
			migration.EXPECT().Version().Return(uint64(1)).Times(1)
			migration.EXPECT().Down().Return(nil).Times(0)
			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, []Migration{migration})

			if e := sut.Down(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})

		t.Run("remove migration", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			db, mockDB, gdb, e := createMockedDB()
			if e != nil {
				t.Fatalf("unexpected (%s) error when opening connection", e)
			}
			defer func() { _ = db.Close() }()

			mockDB.
				ExpectQuery("SELECT \\* FROM `__migrations`").
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "version", "created_at", "updated_at", "deleted_at"}).
					AddRow(2, 2, nil, nil, nil))
			mockDB.ExpectBegin()
			mockDB.
				ExpectExec("DELETE FROM `__migrations`").
				WithArgs(uint64(2)).
				WillReturnResult(sqlmock.NewResult(3, 1))
			mockDB.ExpectCommit()

			migration1 := NewMockMigration(ctrl)
			migration1.EXPECT().Version().Return(uint64(1)).AnyTimes()
			migration1.EXPECT().Down().Return(nil).Times(0)
			migration2 := NewMockMigration(ctrl)
			migration2.EXPECT().Version().Return(uint64(2)).AnyTimes()
			migration2.EXPECT().Down().Return(nil).Times(1)
			migration3 := NewMockMigration(ctrl)
			migration3.EXPECT().Version().Return(uint64(3)).AnyTimes()
			migration3.EXPECT().Down().Return(nil).Times(0)
			dao, _ := NewMigratorDao(gdb)
			sut, _ := NewMigrator(dao, []Migration{migration3, migration1, migration2})

			if e := sut.Down(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if dbErr := mockDB.ExpectationsWereMet(); dbErr != nil {
				t.Errorf("DB interactions expectations when met : %v", dbErr)
			}
		})
	})
}

func Test_MigratorServiceRegister(t *testing.T) {
	t.Run("NewMigratorServiceRegister", func(t *testing.T) {
		t.Run("create", func(t *testing.T) {
			if NewMigratorServiceRegister(nil) == nil {
				t.Error("didn't returned a valid reference")
			}
		})

		t.Run("create with app reference", func(t *testing.T) {
			app := NewApp()
			if sut := NewMigratorServiceRegister(app); sut == nil {
				t.Error("didn't returned a valid reference")
			} else if sut.App != app {
				t.Error("didn't stored the app reference")
			}
		})
	})

	t.Run("Provide", func(t *testing.T) {
		t.Run("nil container", func(t *testing.T) {
			if e := NewMigratorServiceRegister(nil).Provide(nil); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})
		t.Run("register components", func(t *testing.T) {
			container := NewServiceContainer()
			sut := NewMigratorServiceRegister(nil)

			e := sut.Provide(container)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case !container.Has(MigratorDAOContainerID):
				t.Errorf("no migrator DAO : %v", sut)
			case !container.Has(MigratorAllMigrationsContainerID):
				t.Errorf("no migrations list : %v", sut)
			case !container.Has(MigratorContainerID):
				t.Errorf("no migrator : %v", sut)
			}
		})

		t.Run("error retrieving db connection factory when retrieving migrator DAO", func(t *testing.T) {
			expected := fmt.Errorf("error message")
			container := NewServiceContainer()
			_ = NewRdbServiceRegister(nil).Provide(container)
			_ = NewMigratorServiceRegister(nil).Provide(container)
			_ = container.Add(RdbContainerID, func() (*RdbConnectionPool, error) {
				return nil, expected
			})

			if _, e := container.Get(MigratorDAOContainerID); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceContainer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("error retrieving db connection config when retrieving migrator DAO", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			container := NewServiceContainer()
			_ = NewConfigServiceRegister(nil).Provide(container)
			_ = NewRdbServiceRegister(nil).Provide(container)
			_ = NewMigratorServiceRegister(nil).Provide(container)
			_ = container.Add(RdbConfigContainerID, func() (*gorm.Config, error) {
				return nil, expected
			})

			if _, e := container.Get(MigratorDAOContainerID); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceContainer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("error retrieving connection when retrieving migrator DAO", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewConfigServiceRegister(nil).Provide(container)
			_ = NewRdbServiceRegister(nil).Provide(container)
			_ = NewMigratorServiceRegister(nil).Provide(container)

			if _, e := container.Get(MigratorDAOContainerID); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceContainer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("retrieving migrator DAO", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewConfigServiceRegister(nil).Provide(container)
			_ = NewRdbServiceRegister(nil).Provide(container)
			_ = NewMigratorServiceRegister(nil).Provide(container)

			rdbCfg := ConfigPartial{"dialect": "invalid", "host": ":memory:"}
			partial := ConfigPartial{}
			_, _ = partial.Set("slate.rdb.connections.primary", rdbCfg)
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(partial, nil).Times(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 0, supplier)
			_ = container.Add(ConfigContainerID, func() *Config { return config })
			migrator := NewMockGormMigrator(ctrl)
			migrator.EXPECT().AutoMigrate(gomock.Any()).Return(nil).Times(1)
			dialect := NewMockGormDialector(ctrl)
			dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
			dialect.EXPECT().Migrator(gomock.Any()).Return(migrator).Times(1)
			dialectCreator := NewMockRdbDialectCreator(ctrl)
			dialectCreator.EXPECT().Accept(&rdbCfg).Return(true).Times(1)
			dialectCreator.EXPECT().Create(&rdbCfg).Return(dialect, nil).Times(1)
			_ = container.Add("dialect_strategy", func() RdbDialectCreator {
				return dialectCreator
			}, RdbDialectCreatorTag)

			_ = NewRdbServiceRegister(nil).Boot(container)

			sut, e := container.Get(MigratorDAOContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case sut == nil:
				t.Error("didn't returned a reference to the migrator DAO")
			default:
				switch sut.(type) {
				case *MigratorDao:
				default:
					t.Error("didn't returned a migrator DAO reference")
				}
			}
		})

		t.Run("retrieving migrator DAO with db name from env", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			prev := MigratorDatabase
			MigratorDatabase = "secondary"
			defer func() { MigratorDatabase = prev }()

			container := NewServiceContainer()
			_ = NewConfigServiceRegister(nil).Provide(container)
			_ = NewRdbServiceRegister(nil).Provide(container)
			_ = NewMigratorServiceRegister(nil).Provide(container)

			rdbCfg := ConfigPartial{"dialect": "invalid", "host": ":memory:"}
			partial := ConfigPartial{}
			_, _ = partial.Set("slate.rdb.connections.secondary", rdbCfg)
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(partial, nil).Times(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 0, supplier)
			_ = container.Add(ConfigContainerID, func() *Config { return config })
			migrator := NewMockGormMigrator(ctrl)
			migrator.EXPECT().AutoMigrate(gomock.Any()).Return(nil).Times(1)
			dialect := NewMockGormDialector(ctrl)
			dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
			dialect.EXPECT().Migrator(gomock.Any()).Return(migrator).Times(1)
			dialectCreator := NewMockRdbDialectCreator(ctrl)
			dialectCreator.EXPECT().Accept(&rdbCfg).Return(true).Times(1)
			dialectCreator.EXPECT().Create(&rdbCfg).Return(dialect, nil).Times(1)
			_ = container.Add("dialect_strategy", func() RdbDialectCreator {
				return dialectCreator
			}, RdbDialectCreatorTag)

			_ = NewRdbServiceRegister(nil).Boot(container)

			sut, e := container.Get(MigratorDAOContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case sut == nil:
				t.Error("didn't returned a reference to the migrator DAO")
			default:
				switch sut.(type) {
				case *MigratorDao:
				default:
					t.Error("didn't returned a migrator DAO reference")
				}
			}
		})

		t.Run("retrieving migrations", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewLogServiceRegister().Provide(container)
			_ = NewMigratorServiceRegister().Provide(container)

			migration := NewMockMigration(ctrl)
			_ = container.Add("migration.id", func() Migration {
				return migration
			}, MigratorMigrationTag)

			migrations, e := container.Get(MigratorAllMigrationsContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case migrations == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch c := migrations.(type) {
				case []Migration:
					found := false
					for _, m := range c {
						if m == migration {
							found = true
						}
					}
					if !found {
						t.Error("didn't return a migration slice populated with the expected migration instance")
					}
				default:
					t.Error("didn't return a migration slice")
				}
			}
		})

		t.Run("invalid migrator DAO when retrieving the migrator", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewConfigServiceRegister(nil).Provide(container)
			_ = NewRdbServiceRegister(nil).Provide(container)
			_ = NewMigratorServiceRegister(nil).Provide(container)
			_ = container.Add(MigratorDAOContainerID, func() string {
				return "string"
			})

			if _, e := container.Get(MigratorContainerID); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceContainer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("error retrieving the migrator DAO when retrieving the migrator", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			container := NewServiceContainer()
			_ = NewConfigServiceRegister(nil).Provide(container)
			_ = NewRdbServiceRegister(nil).Provide(container)
			_ = NewMigratorServiceRegister(nil).Provide(container)
			_ = container.Add(MigratorDAOContainerID, func() (*MigratorDao, error) {
				return nil, expected
			})

			if _, e := container.Get(MigratorContainerID); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceContainer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("retrieving migrator", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewConfigServiceRegister(nil).Provide(container)
			_ = NewRdbServiceRegister(nil).Provide(container)
			_ = NewMigratorServiceRegister(nil).Provide(container)

			rdbCfg := ConfigPartial{"dialect": "invalid", "host": ":memory:"}
			partial := ConfigPartial{}
			_, _ = partial.Set("slate.rdb.connections.primary", rdbCfg)
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(partial, nil).Times(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 0, supplier)
			_ = container.Add(ConfigContainerID, func() *Config { return config })
			migrator := NewMockGormMigrator(ctrl)
			migrator.EXPECT().AutoMigrate(gomock.Any()).Return(nil).Times(1)
			dialector := NewMockGormDialector(ctrl)
			dialector.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
			dialector.EXPECT().Migrator(gomock.Any()).Return(migrator).Times(1)
			dialectCreator := NewMockRdbDialectCreator(ctrl)
			dialectCreator.EXPECT().Accept(&rdbCfg).Return(true).Times(1)
			dialectCreator.EXPECT().Create(&rdbCfg).Return(dialector, nil).Times(1)
			_ = container.Add("dialect_strategy", func() RdbDialectCreator {
				return dialectCreator
			}, RdbDialectCreatorTag)

			_ = NewRdbServiceRegister(nil).Boot(container)

			sut, e := container.Get(MigratorContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
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
	})

	t.Run("Boot", func(t *testing.T) {
		t.Run("nil container", func(t *testing.T) {
			if e := NewMigratorServiceRegister(nil).Boot(nil); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})

		t.Run("disable auto migration", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			MigratorAutoMigrate = false
			defer func() { MigratorAutoMigrate = true }()

			container := NewServiceContainer()
			sut := NewMigratorServiceRegister(nil)

			if e := sut.Boot(container); e != nil {
				t.Errorf("unexpected serr, (%v)", e)
			}
		})

		t.Run("disable migrator auto migration by environment variable", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			MigratorAutoMigrate = false
			defer func() { MigratorAutoMigrate = true }()

			container := NewServiceContainer()
			sut := NewMigratorServiceRegister(nil)
			_ = sut.Provide(container)

			if e := sut.Boot(container); e != nil {
				t.Errorf("unexpected serr, (%v)", e)
			}
		})

		t.Run("error on retrieving migrator", func(t *testing.T) {
			expected := fmt.Errorf("error message")
			container := NewServiceContainer()
			_ = NewConfigServiceRegister(nil).Provide(container)
			_ = NewRdbServiceRegister(nil).Provide(container)

			sut := NewMigratorServiceRegister(nil)
			_ = sut.Provide(container)
			_ = container.Add(MigratorContainerID, func() (*Migrator, error) {
				return nil, expected
			})

			if e := sut.Boot(container); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceContainer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("invalid retrieved migrator", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewConfigServiceRegister(nil).Provide(container)
			_ = NewRdbServiceRegister(nil).Provide(container)

			sut := NewMigratorServiceRegister(nil)
			_ = sut.Provide(container)
			_ = container.Add(MigratorContainerID, func() interface{} {
				return "string"
			})

			if e := sut.Boot(container); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrConversion) {
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("error on retrieving migrations", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewConfigServiceRegister(nil).Provide(container)
			_ = NewRdbServiceRegister(nil).Provide(container)

			expected := fmt.Errorf("error message")
			_ = container.Add("id", func() (interface{}, error) {
				return nil, expected
			}, MigratorMigrationTag)
			sut := NewMigratorServiceRegister(nil)
			_ = sut.Provide(container)

			if e := sut.Boot(container); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceContainer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("running migrator", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewConfigServiceRegister(nil).Provide(container)
			_ = NewRdbServiceRegister(nil).Provide(container)

			rdbCfg := ConfigPartial{"dialect": "sqlite", "host": ":memory:"}
			partial := ConfigPartial{}
			_, _ = partial.Set("slate.rdb.connections.primary", rdbCfg)
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(partial, nil).Times(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 0, supplier)
			_ = container.Add(ConfigContainerID, func() *Config { return config })
			migrator := NewMockGormMigrator(ctrl)
			migrator.EXPECT().AutoMigrate(gomock.Any()).Return(nil).Times(1)
			dialect := NewMockGormDialector(ctrl)
			dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
			dialect.EXPECT().Migrator(gomock.Any()).Return(migrator).Times(1)
			dialectCreator := NewMockRdbDialectCreator(ctrl)
			dialectCreator.EXPECT().Accept(&rdbCfg).Return(true).Times(1)
			dialectCreator.EXPECT().Create(&rdbCfg).Return(dialect, nil).Times(1)
			_ = container.Add("dialect_strategy", func() RdbDialectCreator {
				return dialectCreator
			}, RdbDialectCreatorTag)
			migration := NewMockMigration(ctrl)
			migration.EXPECT().Version().Return(uint64(1)).Times(1)
			migration.EXPECT().Up().Times(1)
			_ = container.Add("id", func() interface{} {
				return migration
			}, MigratorMigrationTag)

			_ = NewRdbServiceRegister(nil).Boot(container)

			sut := NewMigratorServiceRegister(nil)
			_ = sut.Provide(container)

			if e := sut.Boot(container); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})
	})
}
