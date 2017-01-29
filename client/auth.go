package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// JWTAuthPath computes a request path to the jwt action of auth.
func JWTAuthPath() string {
	return fmt.Sprintf("/api/v1/auth")
}

// Creates a valid JWT
func (c *Client) JWTAuth(ctx context.Context, path string, payload *AuthPayload, contentType string) (*http.Response, error) {
	req, err := c.NewJWTAuthRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewJWTAuthRequest create the request corresponding to the jwt action endpoint of the auth resource.
func (c *Client) NewJWTAuthRequest(ctx context.Context, path string, payload *AuthPayload, contentType string) (*http.Request, error) {
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
