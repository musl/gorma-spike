package main

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/net/context"
	"net/http"
)

/*
* Build a middleware that distributes a thread-safe reference to a
* gorm DB to all requests.
 */
func DBMiddleware(db *gorm.DB) goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			newctx := context.WithValue(ctx, "db", db)
			return h(newctx, rw, req)
		}
	}
}
