package main

import (
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/musl/hixio/app"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"path/filepath"
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

// LoadJWTPublicKeys loads PEM encoded RSA public keys used to validata and decrypt the JWT.
/*
* Generate with:
* openssl genrsa -out jwt.key 4096
* openssl rsa -in jwt.key -pubout > jwt.key.pub
 */
func LoadJWTPublicKeys(path string) ([]jwt.Key, error) {
	keyFiles, err := filepath.Glob(path)
	if err != nil {
		return nil, err
	}
	keys := make([]jwt.Key, len(keyFiles))
	for i, keyFile := range keyFiles {
		pem, err := ioutil.ReadFile(keyFile)
		if err != nil {
			return nil, err
		}
		key, err := jwtgo.ParseRSAPublicKeyFromPEM([]byte(pem))
		if err != nil {
			return nil, fmt.Errorf("failed to load key %s: %s", keyFile, err)
		}
		keys[i] = key
	}
	if len(keys) == 0 {
		return nil, fmt.Errorf("couldn't load public keys for JWT security")
	}

	return keys, nil
}

/*
*
 */
func ValidateJWT() goa.Middleware {
	validator := func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// TODO validate the JWT claim
			return h(ctx, rw, req)
		}
	}

	m, _ := goa.NewMiddleware(validator)
	return m
}

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
