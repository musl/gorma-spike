package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreatePostPath computes a request path to the create action of post.
func CreatePostPath() string {
	return fmt.Sprintf("/api/v1/posts")
}

// creates a post
func (c *Client) CreatePost(ctx context.Context, path string, payload *PostPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreatePostRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreatePostRequest create the request corresponding to the create action endpoint of the post resource.
func (c *Client) NewCreatePostRequest(ctx context.Context, path string, payload *PostPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DeletePostPath computes a request path to the delete action of post.
func DeletePostPath(id int) string {
	return fmt.Sprintf("/api/v1/posts/%v", id)
}

// deletes a post
func (c *Client) DeletePost(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeletePostRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeletePostRequest create the request corresponding to the delete action endpoint of the post resource.
func (c *Client) NewDeletePostRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListPostPath computes a request path to the list action of post.
func ListPostPath() string {
	return fmt.Sprintf("/api/v1/posts")
}

// lists all posts
func (c *Client) ListPost(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListPostRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListPostRequest create the request corresponding to the list action endpoint of the post resource.
func (c *Client) NewListPostRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowPostPath computes a request path to the show action of post.
func ShowPostPath(id int) string {
	return fmt.Sprintf("/api/v1/posts/%v", id)
}

// shows a post
func (c *Client) ShowPost(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowPostRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowPostRequest create the request corresponding to the show action endpoint of the post resource.
func (c *Client) NewShowPostRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
