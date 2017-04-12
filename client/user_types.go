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

// Auth Payload is used to auth users.
type authPayload struct {
	// email of a user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// password of the user
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// Validate validates the authPayload type instance.
func (ut *authPayload) Validate() (err error) {
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
	return
}

// Publicize creates AuthPayload from authPayload
func (ut *authPayload) Publicize() *AuthPayload {
	var pub AuthPayload
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.Password != nil {
		pub.Password = *ut.Password
	}
	return &pub
}

// Auth Payload is used to auth users.
type AuthPayload struct {
	// email of a user
	Email string `form:"email" json:"email" xml:"email"`
	// password of the user
	Password string `form:"password" json:"password" xml:"password"`
}

// Validate validates the AuthPayload type instance.
func (ut *AuthPayload) Validate() (err error) {
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "password"))
	}

	if utf8.RuneCountInString(ut.Email) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.email`, ut.Email, utf8.RuneCountInString(ut.Email), 1, true))
	}
	return
}

// Photo Payload is used to create photos.
type photoPayload struct {
	// name of a post
	Alt *string `form:"alt,omitempty" json:"alt,omitempty" xml:"alt,omitempty"`
	// URL to full-size image
	OriginalURL *string `form:"original_url,omitempty" json:"original_url,omitempty" xml:"original_url,omitempty"`
	// is the photo published
	Published *bool `form:"published,omitempty" json:"published,omitempty" xml:"published,omitempty"`
	// URL to thumbnail-size image
	ThumbnailURL *string `form:"thumbnail_url,omitempty" json:"thumbnail_url,omitempty" xml:"thumbnail_url,omitempty"`
}

// Validate validates the photoPayload type instance.
func (ut *photoPayload) Validate() (err error) {
	if ut.Alt == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "alt"))
	}
	if ut.OriginalURL == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "original_url"))
	}
	if ut.ThumbnailURL == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "thumbnail_url"))
	}
	if ut.Published == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "published"))
	}

	if ut.Alt != nil {
		if utf8.RuneCountInString(*ut.Alt) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.alt`, *ut.Alt, utf8.RuneCountInString(*ut.Alt), 1, true))
		}
	}
	if ut.OriginalURL != nil {
		if utf8.RuneCountInString(*ut.OriginalURL) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.original_url`, *ut.OriginalURL, utf8.RuneCountInString(*ut.OriginalURL), 1, true))
		}
	}
	if ut.ThumbnailURL != nil {
		if utf8.RuneCountInString(*ut.ThumbnailURL) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.thumbnail_url`, *ut.ThumbnailURL, utf8.RuneCountInString(*ut.ThumbnailURL), 1, true))
		}
	}
	return
}

// Publicize creates PhotoPayload from photoPayload
func (ut *photoPayload) Publicize() *PhotoPayload {
	var pub PhotoPayload
	if ut.Alt != nil {
		pub.Alt = *ut.Alt
	}
	if ut.OriginalURL != nil {
		pub.OriginalURL = *ut.OriginalURL
	}
	if ut.Published != nil {
		pub.Published = *ut.Published
	}
	if ut.ThumbnailURL != nil {
		pub.ThumbnailURL = *ut.ThumbnailURL
	}
	return &pub
}

// Photo Payload is used to create photos.
type PhotoPayload struct {
	// name of a post
	Alt string `form:"alt" json:"alt" xml:"alt"`
	// URL to full-size image
	OriginalURL string `form:"original_url" json:"original_url" xml:"original_url"`
	// is the photo published
	Published bool `form:"published" json:"published" xml:"published"`
	// URL to thumbnail-size image
	ThumbnailURL string `form:"thumbnail_url" json:"thumbnail_url" xml:"thumbnail_url"`
}

// Validate validates the PhotoPayload type instance.
func (ut *PhotoPayload) Validate() (err error) {
	if ut.Alt == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "alt"))
	}
	if ut.OriginalURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "original_url"))
	}
	if ut.ThumbnailURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "thumbnail_url"))
	}

	if utf8.RuneCountInString(ut.Alt) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.alt`, ut.Alt, utf8.RuneCountInString(ut.Alt), 1, true))
	}
	if utf8.RuneCountInString(ut.OriginalURL) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.original_url`, ut.OriginalURL, utf8.RuneCountInString(ut.OriginalURL), 1, true))
	}
	if utf8.RuneCountInString(ut.ThumbnailURL) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.thumbnail_url`, ut.ThumbnailURL, utf8.RuneCountInString(ut.ThumbnailURL), 1, true))
	}
	return
}

// Post Payload is used to create posts.
type postPayload struct {
	// body of a post
	Body *string `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"`
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
	if utf8.RuneCountInString(ut.Title) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, ut.Title, utf8.RuneCountInString(ut.Title), 1, true))
	}
	return
}

// User Payload is used to create users.
type userPayload struct {
	// email of a user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
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
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Password != nil {
		pub.Password = *ut.Password
	}
	return &pub
}

// User Payload is used to create users.
type UserPayload struct {
	// email of a user
	Email string `form:"email" json:"email" xml:"email"`
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
	if utf8.RuneCountInString(ut.Name) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, ut.Name, utf8.RuneCountInString(ut.Name), 1, true))
	}
	return
}
