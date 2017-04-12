//************************************************************************//
// API "hixio": Model Helpers
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

// MediaType Retrieval Functions

// ListPhoto returns an array of view: default.
func (m *PhotoDB) ListPhoto(ctx context.Context) []*app.Photo {
	defer goa.MeasureSince([]string{"goa", "db", "photo", "listphoto"}, time.Now())

	var native []*Photo
	var objs []*app.Photo
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Photo", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.PhotoToPhoto())
	}

	return objs
}

// PhotoToPhoto loads a Photo and builds the default view of media type photo.
func (m *Photo) PhotoToPhoto() *app.Photo {
	photo := &app.Photo{}
	photo.Alt = m.Alt
	photo.ID = m.ID
	photo.OriginalURL = m.OriginalURL
	photo.Published = m.Published
	photo.ThumbnailURL = m.ThumbnailURL

	return photo
}

// OnePhoto loads a Photo and builds the default view of media type photo.
func (m *PhotoDB) OnePhoto(ctx context.Context, id int) (*app.Photo, error) {
	defer goa.MeasureSince([]string{"goa", "db", "photo", "onephoto"}, time.Now())

	var native Photo
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Photo", "error", err.Error())
		return nil, err
	}

	view := *native.PhotoToPhoto()
	return &view, err
}
