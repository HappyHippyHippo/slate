package migration

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

// Record defines the rdb record that stores a migration.
type Record struct {
	ID uint `json:"id" xml:"id" gorm:"primaryKey"`

	Version uint64 `json:"model" xml:"model"`

	CreatedAt time.Time  `json:"createdAt" xml:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt" xml:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" xml:"deletedAt" sql:"index"`
}

// TableName defines the table name to be used to manage the migrations.
func (Record) TableName() string {
	return "__version"
}

// IDao defines the interface to the migration manager DAO instance.
type IDao interface {
	Last() (Record, error)
	Up(version uint64) (Record, error)
	Down(last Record) error
}

// Dao defines an object to the migration DAO instance responsible
// to manager the installed migrations.
type Dao struct {
	db *gorm.DB
}

var _ IDao = &Dao{}

func newDao(db *gorm.DB) (IDao, error) {
	if e := db.AutoMigrate(&Record{}); e != nil {
		return nil, e
	}

	return &Dao{db: db}, nil
}

// Last will retrieve the last registered migration
func (d Dao) Last() (Record, error) {
	model := Record{}
	result := d.db.
		Order("created_at desc").
		FirstOrInit(&model, Record{Version: 0})
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Record{}, result.Error
		}
	}
	return model, nil
}

// Up will register a new executed migration
func (d Dao) Up(version uint64) (Record, error) {
	model := Record{Version: version}
	if result := d.db.Create(&model); result.Error != nil {
		return Record{}, result.Error
	}
	return model, nil
}

// Down will remove the last migration register
func (d Dao) Down(last Record) error {
	if last.Version != 0 {
		if result := d.db.Unscoped().Delete(&last); result.Error != nil {
			return result.Error
		}
	}

	return nil
}
