package smigration

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serr"
	"github.com/pkg/errors"
	"reflect"
	"testing"
)

func Test_NewMigrator(t *testing.T) {
	t.Run("nil dao", func(t *testing.T) {
		sut, e := NewMigrator(nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("new yaml decoder adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dao := NewMockDao(ctrl)

		sut, e := NewMigrator(dao)
		switch {
		case e != nil:
			t.Errorf("returned the unexpectederror (%v)", e)
		case sut == nil:
			t.Error("didn't return the expected migrator instance")
		case !reflect.DeepEqual(sut.(*migrator).dao, dao):
			t.Error("didn't stored the given dao")
		}
	})
}

func Test_Migrator_AddMigration(t *testing.T) {
	t.Run("error adding nil migration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewMigrator(NewMockDao(ctrl))

		if e := sut.AddMigration(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("adding migration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		migration := NewMockMigration(ctrl)
		sut, _ := NewMigrator(NewMockDao(ctrl))

		if e := sut.AddMigration(migration); e != nil {
			t.Errorf("returned the unexpected (%v) error", e)
		} else if !reflect.DeepEqual(sut.(*migrator).migrations, []IMigration{migration}) {
			t.Error("didn't stored the registering migration")
		}
	})
}

func Test_Migrator_Current(t *testing.T) {
	t.Run("error retrieving the last migration version", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		dao := NewMockDao(ctrl)
		dao.EXPECT().Last().Return(Record{}, expected).Times(1)

		sut, _ := NewMigrator(dao)

		if _, e := sut.Current(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", e, expected)
		}
	})

	t.Run("retrieves the last migration register version", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		version := uint64(23)
		dao := NewMockDao(ctrl)
		dao.EXPECT().Last().Return(Record{Version: version}, nil).Times(1)

		sut, _ := NewMigrator(dao)

		if current, e := sut.Current(); e != nil {
			t.Errorf("returned the unexpected (%v) error", e)
		} else if current != version {
			t.Errorf("unexpectedly returned the %v version when expecting %v", current, version)
		}
	})
}

func Test_Migrator_Migrate(t *testing.T) {
	t.Run("no-op when no migrations", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dao := NewMockDao(ctrl)
		sut, _ := NewMigrator(dao)

		if e := sut.Migrate(); e != nil {
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})

	t.Run("error while retrieving the last version", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		dao := NewMockDao(ctrl)
		dao.EXPECT().Last().Return(Record{}, expected).Times(1)
		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(1)).Times(0)
		sut, _ := NewMigrator(dao)
		_ = sut.AddMigration(migration)

		if e := sut.Migrate(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", e, expected)
		}
	})

	t.Run("no-op if no migrations higher than current last stored", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dao := NewMockDao(ctrl)
		dao.EXPECT().Last().Return(Record{Version: uint64(1)}, nil).Times(1)
		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(1)).Times(1)
		sut, _ := NewMigrator(dao)
		_ = sut.AddMigration(migration)

		if e := sut.Migrate(); e != nil {
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})

	t.Run("error while executing a migration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		dao := NewMockDao(ctrl)
		dao.EXPECT().Last().Return(Record{Version: uint64(0)}, nil).Times(1)
		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(1)).Times(1)
		migration.EXPECT().Up().Return(expected).Times(1)
		sut, _ := NewMigrator(dao)
		_ = sut.AddMigration(migration)

		if e := sut.Migrate(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", e, expected)
		}
	})

	t.Run("error while raising the migration version", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		dao := NewMockDao(ctrl)
		dao.EXPECT().Last().Return(Record{Version: uint64(0)}, nil).Times(1)
		dao.EXPECT().Up(uint64(1)).Return(Record{}, expected).Times(1)
		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(1)).Times(1)
		migration.EXPECT().Up().Return(nil).Times(1)
		sut, _ := NewMigrator(dao)
		_ = sut.AddMigration(migration)

		if e := sut.Migrate(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", e, expected)
		}
	})

	t.Run("execute missing migrations (in order)", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dao := NewMockDao(ctrl)
		dao.EXPECT().Last().Return(Record{Version: uint64(1)}, nil).Times(1)
		gomock.InOrder(
			dao.EXPECT().Up(uint64(2)).Return(Record{Version: uint64(2)}, nil),
			dao.EXPECT().Up(uint64(3)).Return(Record{Version: uint64(3)}, nil),
		)

		migration1 := NewMockMigration(ctrl)
		migration1.EXPECT().Version().Return(uint64(1)).AnyTimes()
		migration2 := NewMockMigration(ctrl)
		migration2.EXPECT().Version().Return(uint64(2)).AnyTimes()
		migration2.EXPECT().Up().Return(nil).Times(1)
		migration3 := NewMockMigration(ctrl)
		migration3.EXPECT().Version().Return(uint64(3)).AnyTimes()
		migration3.EXPECT().Up().Return(nil).Times(1)
		sut, _ := NewMigrator(dao)
		_ = sut.AddMigration(migration3)
		_ = sut.AddMigration(migration1)
		_ = sut.AddMigration(migration2)

		if e := sut.Migrate(); e != nil {
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}

func Test_Migrator_Up(t *testing.T) {
	t.Run("no-op when no migrations", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dao := NewMockDao(ctrl)
		sut, _ := NewMigrator(dao)

		if e := sut.Up(); e != nil {
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})

	t.Run("error while retrieving the last version", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		dao := NewMockDao(ctrl)
		dao.EXPECT().Last().Return(Record{}, expected).Times(1)
		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(1)).Times(0)
		sut, _ := NewMigrator(dao)
		_ = sut.AddMigration(migration)

		if e := sut.Up(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", e, expected)
		}
	})

	t.Run("no-op if no migrations higher than current last stored", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dao := NewMockDao(ctrl)
		dao.EXPECT().Last().Return(Record{Version: uint64(1)}, nil).Times(1)
		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(1)).Times(1)
		sut, _ := NewMigrator(dao)
		_ = sut.AddMigration(migration)

		if e := sut.Up(); e != nil {
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})

	t.Run("error while executing a migration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		dao := NewMockDao(ctrl)
		dao.EXPECT().Last().Return(Record{Version: uint64(0)}, nil).Times(1)
		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(1)).Times(1)
		migration.EXPECT().Up().Return(expected).Times(1)
		sut, _ := NewMigrator(dao)
		_ = sut.AddMigration(migration)

		if e := sut.Up(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", e, expected)
		}
	})

	t.Run("error while raising the migration version", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		dao := NewMockDao(ctrl)
		dao.EXPECT().Last().Return(Record{Version: uint64(1)}, nil).Times(1)
		dao.EXPECT().Up(uint64(2)).Return(Record{}, expected).Times(1)
		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(2)).Times(1)
		migration.EXPECT().Up().Return(nil).Times(1)
		sut, _ := NewMigrator(dao)
		_ = sut.AddMigration(migration)

		if e := sut.Up(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", e, expected)
		}
	})

	t.Run("execute next migration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dao := NewMockDao(ctrl)
		dao.EXPECT().Last().Return(Record{Version: uint64(1)}, nil).Times(1)
		dao.EXPECT().Up(uint64(2)).Return(Record{Version: uint64(2)}, nil).Times(1)
		migration1 := NewMockMigration(ctrl)
		migration1.EXPECT().Version().Return(uint64(1)).AnyTimes()
		migration2 := NewMockMigration(ctrl)
		migration2.EXPECT().Version().Return(uint64(2)).AnyTimes()
		migration2.EXPECT().Up().Return(nil).Times(1)
		migration3 := NewMockMigration(ctrl)
		migration3.EXPECT().Version().Return(uint64(3)).AnyTimes()
		sut, _ := NewMigrator(dao)
		_ = sut.AddMigration(migration3)
		_ = sut.AddMigration(migration1)
		_ = sut.AddMigration(migration2)

		if e := sut.Up(); e != nil {
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}

func Test_Migrator_Down(t *testing.T) {
	t.Run("no-op when no migrations", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dao := NewMockDao(ctrl)
		sut, _ := NewMigrator(dao)

		if e := sut.Down(); e != nil {
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})

	t.Run("error while retrieving the last version", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		dao := NewMockDao(ctrl)
		dao.EXPECT().Last().Return(Record{}, expected).Times(1)
		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(1)).Times(0)
		sut, _ := NewMigrator(dao)
		_ = sut.AddMigration(migration)

		if e := sut.Down(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", e, expected)
		}
	})

	t.Run("error while executing a migration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		version := uint64(34)
		expected := fmt.Errorf("error message")
		dao := NewMockDao(ctrl)
		dao.EXPECT().Last().Return(Record{Version: version}, nil).Times(1)
		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(version).Times(1)
		migration.EXPECT().Down().Return(expected).Times(1)
		sut, _ := NewMigrator(dao)
		_ = sut.AddMigration(migration)

		if e := sut.Down(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", e, expected)
		}
	})

	t.Run("error while removing the migration version", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id := uint(12)
		version := uint64(34)
		expected := fmt.Errorf("error message")
		dao := NewMockDao(ctrl)
		dao.EXPECT().Last().Return(Record{ID: id, Version: version}, nil).Times(1)
		dao.EXPECT().Down(Record{ID: id, Version: version}).Return(expected).Times(1)
		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(version).Times(1)
		migration.EXPECT().Down().Return(nil).Times(1)
		sut, _ := NewMigrator(dao)
		_ = sut.AddMigration(migration)

		if e := sut.Down(); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the unexpected (%v) error instead of the expected (%v)", e, expected)
		}
	})

	t.Run("no-op if the migration was not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dao := NewMockDao(ctrl)
		dao.EXPECT().Last().Return(Record{Version: uint64(2)}, nil).Times(1)
		migration1 := NewMockMigration(ctrl)
		migration1.EXPECT().Version().Return(uint64(1)).AnyTimes()
		sut, _ := NewMigrator(dao)
		_ = sut.AddMigration(migration1)

		if e := sut.Down(); e != nil {
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})

	t.Run("remove migration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dao := NewMockDao(ctrl)
		dao.EXPECT().Last().Return(Record{Version: uint64(2)}, nil).Times(1)
		dao.EXPECT().Down(Record{Version: uint64(2)}).Return(nil).Times(1)
		migration1 := NewMockMigration(ctrl)
		migration1.EXPECT().Version().Return(uint64(1)).AnyTimes()
		migration2 := NewMockMigration(ctrl)
		migration2.EXPECT().Version().Return(uint64(2)).AnyTimes()
		migration2.EXPECT().Down().Return(nil).Times(1)
		migration3 := NewMockMigration(ctrl)
		migration3.EXPECT().Version().Return(uint64(3)).AnyTimes()
		sut, _ := NewMigrator(dao)
		_ = sut.AddMigration(migration3)
		_ = sut.AddMigration(migration1)
		_ = sut.AddMigration(migration2)

		if e := sut.Down(); e != nil {
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}
