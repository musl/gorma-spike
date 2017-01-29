package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateUserPath computes a request path to the create action of user.
func CreateUserPath() string {
	return fmt.Sprintf("/api/v1/users")
}

// creates a user
func (c *Client) CreateUser(ctx context.Context, path string, payload *UserPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateUserRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateUserRequest create the request corresponding to the create action endpoint of the user resource.
func (c *Client) NewCreateUserRequest(ctx context.Context, path string, payload *UserPayload, contentType string) (*http.Request, error) {
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

// DeleteUserPath computes a request path to the delete action of user.
func DeleteUserPath(id int) string {
	return fmt.Sprintf("/api/v1/users/%v", id)
}

// deletes a user
func (c *Client) DeleteUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteUserRequest create the request corresponding to the delete action endpoint of the user resource.
func (c *Client) NewDeleteUserRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListUserPath computes a request path to the list action of user.
func ListUserPath() string {
	return fmt.Sprintf("/api/v1/users")
}

// lists all users
func (c *Client) ListUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListUserRequest create the request corresponding to the list action endpoint of the user resource.
func (c *Client) NewListUserRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ShowUserPath computes a request path to the show action of user.
func ShowUserPath(id int) string {
	return fmt.Sprintf("/api/v1/users/%v", id)
}

// shows a user
func (c *Client) ShowUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowUserRequest create the request corresponding to the show action endpoint of the user resource.
func (c *Client) NewShowUserRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}
