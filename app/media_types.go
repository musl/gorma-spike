//************************************************************************//
// API "hixio": Application Media Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/musl/hixio/design
// --out=$(GOPATH)/src/github.com/musl/hixio
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// post media type (default view)
//
// Identifier: application/vnd.hixio.goa.post; view=default
type Post struct {
	// body of a post
	Body string `form:"body" json:"body" xml:"body"`
	// Unique Post ID
	ID int `form:"id" json:"id" xml:"id"`
	// is the post published
	Published bool `form:"published" json:"published" xml:"published"`
	// name of a post
	Title string `form:"title" json:"title" xml:"title"`
}

// Validate validates the Post media type instance.
func (mt *Post) Validate() (err error) {
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.Body == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body"))
	}

	if utf8.RuneCountInString(mt.Body) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.body`, mt.Body, utf8.RuneCountInString(mt.Body), 1, true))
	}
	if mt.ID < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, mt.ID, 1, true))
	}
	if utf8.RuneCountInString(mt.Title) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, mt.Title, utf8.RuneCountInString(mt.Title), 1, true))
	}
	return
}

// postCollection is the media type for an array of post (default view)
//
// Identifier: application/vnd.hixio.goa.post; type=collection; view=default
type PostCollection []*Post

// Validate validates the PostCollection media type instance.
func (mt PostCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Title == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "title"))
		}
		if e.Body == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "body"))
		}

		if utf8.RuneCountInString(e.Body) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].body`, e.Body, utf8.RuneCountInString(e.Body), 1, true))
		}
		if e.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response[*].id`, e.ID, 1, true))
		}
		if utf8.RuneCountInString(e.Title) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].title`, e.Title, utf8.RuneCountInString(e.Title), 1, true))
		}
	}
	return
}

// user media type (default view)
//
// Identifier: application/vnd.hixio.goa.user; view=default
type User struct {
	// email of a user
	Email string `form:"email" json:"email" xml:"email"`
	// Unique Post ID
	ID int `form:"id" json:"id" xml:"id"`
	// name of a user
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}

	if utf8.RuneCountInString(mt.Email) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, mt.Email, utf8.RuneCountInString(mt.Email), 1, true))
	}
	if mt.ID < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, mt.ID, 1, true))
	}
	if utf8.RuneCountInString(mt.Name) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 1, true))
	}
	return
}

// userCollection is the media type for an array of user (default view)
//
// Identifier: application/vnd.hixio.goa.user; type=collection; view=default
type UserCollection []*User

// Validate validates the UserCollection media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}
		if e.Email == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "email"))
		}

		if utf8.RuneCountInString(e.Email) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].email`, e.Email, utf8.RuneCountInString(e.Email), 1, true))
		}
		if e.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response[*].id`, e.ID, 1, true))
		}
		if utf8.RuneCountInString(e.Name) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response[*].name`, e.Name, utf8.RuneCountInString(e.Name), 1, true))
		}
	}
	return
}
