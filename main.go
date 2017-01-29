package main

import (
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/musl/hixio/app"
)

func main() {
	// connect to the database
	db, err := gorm.Open("postgres", "host=127.0.0.1 sslmode=disable user=hixio dbname=hixio")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create service
	service := goa.New("hixio")

	// Mount goa middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount app middleware
	service.Use(DBMiddleware(db))

	// Mount auth middleware
	keys, err := LoadJWTPublicKeys("jwt.key.pub")
	if err != nil {
		panic(fmt.Sprintf("Unable to load JWT public key."))
	}
	auth_middleware := jwt.New(jwt.NewSimpleResolver(keys), ValidateJWT(), app.NewJWTSecurity())
	app.UseJWTMiddleware(service, auth_middleware)

	// Mount "auth" controller
	auth_controller := NewAuthController(service, "jwt.key")
	app.MountAuthController(service, auth_controller)

	// Mount "post" controller
	post_controller := NewPostController(service)
	app.MountPostController(service, post_controller)

	// Mount "static" controller
	static_controller := NewStaticController(service)
	app.MountStaticController(service, static_controller)

	// Mount "status" controller
	status_controller := NewStatusController(service)
	app.MountStatusController(service, status_controller)

	// Mount "swagger" controller
	swagger_controller := NewSwaggerController(service)
	app.MountSwaggerController(service, swagger_controller)

	// Mount "user" controller
	user_controller := NewUserController(service)
	app.MountUserController(service, user_controller)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
