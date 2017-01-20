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

package client

import (
	"github.com/goadesign/goa"
	"net/http"
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

// DecodePost decodes the Post instance encoded in resp body.
func (c *Client) DecodePost(resp *http.Response) (*Post, error) {
	var decoded Post
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
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

// DecodePostCollection decodes the PostCollection instance encoded in resp body.
func (c *Client) DecodePostCollection(resp *http.Response) (PostCollection, error) {
	var decoded PostCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}
