package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/musl/hixio/app"
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

func main() {
	// Create service
	service := goa.New("hixio")

	// Mount goa middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// connect to the database
	db, err := gorm.Open("postgres", "host=127.0.0.1 sslmode=disable user=hixio dbname=hixio")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Mount db middleware, look ma, no globals!
	service.Use(DBMiddleware(db))

	// Mount "status" controller
	status_controller := NewStatusController(service)
	app.MountStatusController(service, status_controller)

	// Mount "post" controller
	post_controller := NewPostController(service)
	app.MountPostController(service, post_controller)

	// Mount "spa" controller
	spa_controller := NewSpaController(service)
	app.MountSpaController(service, spa_controller)

	// Mount "swagger" controller
	swagger_controller := NewSwaggerController(service)
	app.MountSwaggerController(service, swagger_controller)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
