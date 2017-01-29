//************************************************************************//
// API "hixio": Application User Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/musl/hixio/design
// --out=$(GOPATH)/src/github.com/musl/hixio
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// Post Payload is used to create posts.
type postPayload struct {
	// body of a post
	Body *string `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"`
	// surrogate key of a post
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// is the post published
	Published *bool `form:"published,omitempty" json:"published,omitempty" xml:"published,omitempty"`
	// name of a post
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the postPayload type instance.
func (ut *postPayload) Validate() (err error) {
	if ut.Title == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if ut.Body == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body"))
	}
	if ut.Published == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "published"))
	}

	if ut.Body != nil {
		if utf8.RuneCountInString(*ut.Body) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.body`, *ut.Body, utf8.RuneCountInString(*ut.Body), 1, true))
		}
	}
	if ut.ID != nil {
		if *ut.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, *ut.ID, 1, true))
		}
	}
	if ut.Title != nil {
		if utf8.RuneCountInString(*ut.Title) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, *ut.Title, utf8.RuneCountInString(*ut.Title), 1, true))
		}
	}
	return
}

// Publicize creates PostPayload from postPayload
func (ut *postPayload) Publicize() *PostPayload {
	var pub PostPayload
	if ut.Body != nil {
		pub.Body = *ut.Body
	}
	if ut.ID != nil {
		pub.ID = ut.ID
	}
	if ut.Published != nil {
		pub.Published = *ut.Published
	}
	if ut.Title != nil {
		pub.Title = *ut.Title
	}
	return &pub
}

// Post Payload is used to create posts.
type PostPayload struct {
	// body of a post
	Body string `form:"body" json:"body" xml:"body"`
	// surrogate key of a post
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// is the post published
	Published bool `form:"published" json:"published" xml:"published"`
	// name of a post
	Title string `form:"title" json:"title" xml:"title"`
}

// Validate validates the PostPayload type instance.
func (ut *PostPayload) Validate() (err error) {
	if ut.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if ut.Body == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body"))
	}

	if utf8.RuneCountInString(ut.Body) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.body`, ut.Body, utf8.RuneCountInString(ut.Body), 1, true))
	}
	if ut.ID != nil {
		if *ut.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, *ut.ID, 1, true))
		}
	}
	if utf8.RuneCountInString(ut.Title) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, ut.Title, utf8.RuneCountInString(ut.Title), 1, true))
	}
	return
}

// Post Payload is used to create users.
type userPayload struct {
	// email of a user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// surrogate key of a user
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name of a user
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// password of the user
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// Validate validates the userPayload type instance.
func (ut *userPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "password"))
	}

	if ut.Email != nil {
		if utf8.RuneCountInString(*ut.Email) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, *ut.Email, utf8.RuneCountInString(*ut.Email), 1, true))
		}
	}
	if ut.ID != nil {
		if *ut.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, *ut.ID, 1, true))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 1, true))
		}
	}
	return
}

// Publicize creates UserPayload from userPayload
func (ut *userPayload) Publicize() *UserPayload {
	var pub UserPayload
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.ID != nil {
		pub.ID = ut.ID
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Password != nil {
		pub.Password = *ut.Password
	}
	return &pub
}

// Post Payload is used to create users.
type UserPayload struct {
	// email of a user
	Email string `form:"email" json:"email" xml:"email"`
	// surrogate key of a user
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name of a user
	Name string `form:"name" json:"name" xml:"name"`
	// password of the user
	Password string `form:"password" json:"password" xml:"password"`
}

// Validate validates the UserPayload type instance.
func (ut *UserPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "password"))
	}

	if utf8.RuneCountInString(ut.Email) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, ut.Email, utf8.RuneCountInString(ut.Email), 1, true))
	}
	if ut.ID != nil {
		if *ut.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, *ut.ID, 1, true))
		}
	}
	if utf8.RuneCountInString(ut.Name) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, ut.Name, utf8.RuneCountInString(ut.Name), 1, true))
	}
	return
}