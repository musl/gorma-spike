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

// ListPost returns an array of view: default.
func (m *PostDB) ListPost(ctx context.Context) []*app.Post {
	defer goa.MeasureSince([]string{"goa", "db", "post", "listpost"}, time.Now())

	var native []*Post
	var objs []*app.Post
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Post", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.PostToPost())
	}

	return objs
}

// PostToPost loads a Post and builds the default view of media type post.
func (m *Post) PostToPost() *app.Post {
	post := &app.Post{}
	post.Body = m.Body
	post.ID = m.ID
	post.Published = m.Published
	post.Title = m.Title

	return post
}

// OnePost loads a Post and builds the default view of media type post.
func (m *PostDB) OnePost(ctx context.Context, id int) (*app.Post, error) {
	defer goa.MeasureSince([]string{"goa", "db", "post", "onepost"}, time.Now())

	var native Post
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Post", "error", err.Error())
		return nil, err
	}

	view := *native.PostToPost()
	return &view, err
}
