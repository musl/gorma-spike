package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreatePhotoPath computes a request path to the create action of photo.
func CreatePhotoPath() string {
	return fmt.Sprintf("/api/v1/photos")
}

// creates a photo
func (c *Client) CreatePhoto(ctx context.Context, path string, payload *PhotoPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreatePhotoRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreatePhotoRequest create the request corresponding to the create action endpoint of the photo resource.
func (c *Client) NewCreatePhotoRequest(ctx context.Context, path string, payload *PhotoPayload, contentType string) (*http.Request, error) {
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
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// DeletePhotoPath computes a request path to the delete action of photo.
func DeletePhotoPath(id int) string {
	return fmt.Sprintf("/api/v1/photos/%v", id)
}

// deletes a photo
func (c *Client) DeletePhoto(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeletePhotoRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeletePhotoRequest create the request corresponding to the delete action endpoint of the photo resource.
func (c *Client) NewDeletePhotoRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ListPhotoPath computes a request path to the list action of photo.
func ListPhotoPath() string {
	return fmt.Sprintf("/api/v1/photos")
}

// lists all published photos
func (c *Client) ListPhoto(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListPhotoRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListPhotoRequest create the request corresponding to the list action endpoint of the photo resource.
func (c *Client) NewListPhotoRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowPhotoPath computes a request path to the show action of photo.
func ShowPhotoPath(id int) string {
	return fmt.Sprintf("/api/v1/photos/%v", id)
}

// shows a photo
func (c *Client) ShowPhoto(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowPhotoRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowPhotoRequest create the request corresponding to the show action endpoint of the photo resource.
func (c *Client) NewShowPhotoRequest(ctx context.Context, path string) (*http.Request, error) {
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
