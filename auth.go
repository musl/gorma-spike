package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"golang.org/x/net/context"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/jinzhu/gorm"
	"github.com/musl/hixio/app"
	"github.com/musl/hixio/models"
	"github.com/satori/go.uuid"

	jwtgo "github.com/dgrijalva/jwt-go"
)

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

type AuthController struct {
	*goa.Controller
	privateKey *rsa.PrivateKey
}

func NewAuthController(service *goa.Service, path string) *AuthController {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Unable to read private key file: %s, %v", path, err))
	}
	privKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		panic(fmt.Sprintf("Unable to load private key file: %s, %v", path, err))
	}
	return &AuthController{
		Controller: service.NewController("AuthController"),
		privateKey: privKey,
	}
}

func (c *AuthController) JWT(ctx *app.JWTAuthContext) error {
	db := ctx.Value("db").(*gorm.DB)
	user := models.User{}

	err := db.Where(&models.User{
		Email:    ctx.Payload.Email,
		Password: ctx.Payload.Password,
	}).First(&user).Error

	if err != nil ||
		ctx.Payload.Email != user.Email ||
		ctx.Payload.Password != user.Password {
		return ctx.Unauthorized()
	}

	token := jwtgo.New(jwtgo.SigningMethodRS512)
	in10m := time.Now().Add(time.Duration(10) * time.Minute).Unix()
	token.Claims = jwtgo.MapClaims{
		"iss":    "Issuer",              // who creates the token and signs it
		"aud":    "Audience",            // to whom the token is intended to be sent
		"exp":    in10m,                 // time when the token will expire (10 minutes from now)
		"jti":    uuid.NewV4().String(), // a unique identifier for the token
		"iat":    time.Now().Unix(),     // when the token was issued/created (now)
		"nbf":    2,                     // time before which the token is not yet valid (2 minutes ago)
		"sub":    "subject",             // the subject/principal is whom the token is about
		"scopes": "api:access",          // token scope - not a standard claim
	}

	signedToken, err := token.SignedString(c.privateKey)
	if err != nil {
		return fmt.Errorf("failed to sign token: %s", err) // internal error
	}

	ctx.ResponseData.Header().Set("Authorization", "Bearer "+signedToken)
	return ctx.NoContent()
}
