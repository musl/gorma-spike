//************************************************************************//
// API "hixio": Models
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/musl/hixio/design
// --out=$(GOPATH)/src/github.com/musl/gorma-spike
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package models

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/musl/gorma-spike/app"
	"golang.org/x/net/context"
	"time"
)

// This is the Post model
type Post struct {
	ID        int `gorm:"primary_key"` // primary key
	Body      string
	Published bool
	Title     string
	CreatedAt time.Time  // timestamp
	DeletedAt *time.Time // nullable timestamp (soft delete)
	UpdatedAt time.Time  // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Post) TableName() string {
	return "posts"

}

// PostDB is the implementation of the storage interface for
// Post.
type PostDB struct {
	Db *gorm.DB
}

// NewPostDB creates a new storage type.
func NewPostDB(db *gorm.DB) *PostDB {
	return &PostDB{Db: db}
}

// DB returns the underlying database.
func (m *PostDB) DB() interface{} {
	return m.Db
}

// PostStorage represents the storage interface.
type PostStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Post, error)
	Get(ctx context.Context, id int) (*Post, error)
	Add(ctx context.Context, post *Post) error
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, id int) error

	ListPost(ctx context.Context) []*app.Post
	OnePost(ctx context.Context, id int) (*app.Post, error)

	UpdateFromPostPayload(ctx context.Context, payload *app.PostPayload, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *PostDB) TableName() string {
	return "posts"

}

// CRUD Functions

// Get returns a single Post as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *PostDB) Get(ctx context.Context, id int) (*Post, error) {
	defer goa.MeasureSince([]string{"goa", "db", "post", "get"}, time.Now())

	var native Post
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Post
func (m *PostDB) List(ctx context.Context) ([]*Post, error) {
	defer goa.MeasureSince([]string{"goa", "db", "post", "list"}, time.Now())

	var objs []*Post
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *PostDB) Add(ctx context.Context, model *Post) error {
	defer goa.MeasureSince([]string{"goa", "db", "post", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Post", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *PostDB) Update(ctx context.Context, model *Post) error {
	defer goa.MeasureSince([]string{"goa", "db", "post", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Post", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *PostDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "post", "delete"}, time.Now())

	var obj Post

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Post", "error", err.Error())
		return err
	}

	return nil
}

// PostFromPostPayload Converts source PostPayload to target Post model
// only copying the non-nil fields from the source.
func PostFromPostPayload(payload *app.PostPayload) *Post {
	post := &Post{}
	post.Body = payload.Body
	post.Published = payload.Published
	post.Title = payload.Title

	return post
}

// UpdateFromPostPayload applies non-nil changes from PostPayload to the model and saves it
func (m *PostDB) UpdateFromPostPayload(ctx context.Context, payload *app.PostPayload, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "post", "updatefrompostPayload"}, time.Now())

	var obj Post
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		goa.LogError(ctx, "error retrieving Post", "error", err.Error())
		return err
	}
	obj.Body = payload.Body
	obj.Published = payload.Published
	obj.Title = payload.Title

	err = m.Db.Save(&obj).Error
	return err
}
