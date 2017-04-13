//************************************************************************//
// API "hixio": Models
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/musl/hixio/design
// --out=$(GOPATH)/src/github.com/musl/hixio
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package models

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/musl/hixio/app"
	"golang.org/x/net/context"
	"time"
)

// This is the Photo model
type Photo struct {
	ID           int `gorm:"primary_key"` // primary key
	Alt          string
	Original     string
	OriginalURL  string `gorm:"column:original_u_r_l"`
	Published    bool
	Thumbnail    string
	ThumbnailURL string     `gorm:"column:thumbnail_u_r_l"`
	CreatedAt    time.Time  // timestamp
	DeletedAt    *time.Time // nullable timestamp (soft delete)
	UpdatedAt    time.Time  // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Photo) TableName() string {
	return "photos"

}

// PhotoDB is the implementation of the storage interface for
// Photo.
type PhotoDB struct {
	Db *gorm.DB
}

// NewPhotoDB creates a new storage type.
func NewPhotoDB(db *gorm.DB) *PhotoDB {
	return &PhotoDB{Db: db}
}

// DB returns the underlying database.
func (m *PhotoDB) DB() interface{} {
	return m.Db
}

// PhotoStorage represents the storage interface.
type PhotoStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Photo, error)
	Get(ctx context.Context, id int) (*Photo, error)
	Add(ctx context.Context, photo *Photo) error
	Update(ctx context.Context, photo *Photo) error
	Delete(ctx context.Context, id int) error

	ListPhoto(ctx context.Context) []*app.Photo
	OnePhoto(ctx context.Context, id int) (*app.Photo, error)

	UpdateFromPhotoPayload(ctx context.Context, payload *app.PhotoPayload, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *PhotoDB) TableName() string {
	return "photos"

}

// CRUD Functions

// Get returns a single Photo as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *PhotoDB) Get(ctx context.Context, id int) (*Photo, error) {
	defer goa.MeasureSince([]string{"goa", "db", "photo", "get"}, time.Now())

	var native Photo
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Photo
func (m *PhotoDB) List(ctx context.Context) ([]*Photo, error) {
	defer goa.MeasureSince([]string{"goa", "db", "photo", "list"}, time.Now())

	var objs []*Photo
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *PhotoDB) Add(ctx context.Context, model *Photo) error {
	defer goa.MeasureSince([]string{"goa", "db", "photo", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Photo", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *PhotoDB) Update(ctx context.Context, model *Photo) error {
	defer goa.MeasureSince([]string{"goa", "db", "photo", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Photo", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *PhotoDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "photo", "delete"}, time.Now())

	var obj Photo

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Photo", "error", err.Error())
		return err
	}

	return nil
}

// PhotoFromPhotoPayload Converts source PhotoPayload to target Photo model
// only copying the non-nil fields from the source.
func PhotoFromPhotoPayload(payload *app.PhotoPayload) *Photo {
	photo := &Photo{}
	photo.Alt = payload.Alt
	photo.Original = payload.Original
	photo.Published = payload.Published
	photo.Thumbnail = payload.Thumbnail

	return photo
}

// UpdateFromPhotoPayload applies non-nil changes from PhotoPayload to the model and saves it
func (m *PhotoDB) UpdateFromPhotoPayload(ctx context.Context, payload *app.PhotoPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "photo", "updatefromphotoPayload"}, time.Now())

	var obj Photo
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Photo", "error", err.Error())
		return err
	}
	obj.Alt = payload.Alt
	obj.Original = payload.Original
	obj.Published = payload.Published
	obj.Thumbnail = payload.Thumbnail

	err = m.Db.Save(&obj).Error
	return err
}
