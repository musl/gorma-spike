package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CheckStatusPath computes a request path to the check action of status.
func CheckStatusPath() string {
	return fmt.Sprintf("/api/v1/status")
}

// A basic status-check endpoint
func (c *Client) CheckStatus(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewCheckStatusRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCheckStatusRequest create the request corresponding to the check action endpoint of the status resource.
func (c *Client) NewCheckStatusRequest(ctx context.Context, path string) (*http.Request, error) {
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
