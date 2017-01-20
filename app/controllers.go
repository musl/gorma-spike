//************************************************************************//
// API "hixio": Application Controllers
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/musl/hixio/design
// --out=$(GOPATH)/src/github.com/musl/hixio
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// PostController is the controller interface for the Post actions.
type PostController interface {
	goa.Muxer
	Create(*CreatePostContext) error
	Delete(*DeletePostContext) error
	List(*ListPostContext) error
	Show(*ShowPostContext) error
}

// MountPostController "mounts" a Post resource controller on the given service.
func MountPostController(service *goa.Service, ctrl PostController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreatePostContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*PostPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/api/v1/posts", ctrl.MuxHandler("Create", h, unmarshalCreatePostPayload))
	service.LogInfo("mount", "ctrl", "Post", "action", "Create", "route", "POST /api/v1/posts")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeletePostContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	service.Mux.Handle("DELETE", "/api/v1/posts/:id", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Post", "action", "Delete", "route", "DELETE /api/v1/posts/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListPostContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/posts", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Post", "action", "List", "route", "GET /api/v1/posts")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowPostContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/posts/:id", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Post", "action", "Show", "route", "GET /api/v1/posts/:id")
}

// unmarshalCreatePostPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreatePostPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &postPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// SpaController is the controller interface for the Spa actions.
type SpaController interface {
	goa.Muxer
	goa.FileServer
}

// MountSpaController "mounts" a Spa resource controller on the given service.
func MountSpaController(service *goa.Service, ctrl SpaController) {
	initService(service)
	var h goa.Handler

	h = ctrl.FileHandler("/*filepath", "static/")
	service.Mux.Handle("GET", "/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Spa", "files", "static/", "route", "GET /*filepath")

	h = ctrl.FileHandler("/", "static/index.html")
	service.Mux.Handle("GET", "/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Spa", "files", "static/index.html", "route", "GET /")
}

// StatusController is the controller interface for the Status actions.
type StatusController interface {
	goa.Muxer
	Check(*CheckStatusContext) error
}

// MountStatusController "mounts" a Status resource controller on the given service.
func MountStatusController(service *goa.Service, ctrl StatusController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCheckStatusContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Check(rctx)
	}
	service.Mux.Handle("GET", "/api/v1/status", ctrl.MuxHandler("Check", h, nil))
	service.LogInfo("mount", "ctrl", "Status", "action", "Check", "route", "GET /api/v1/status")
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler

	h = ctrl.FileHandler("/swagger-ui/*filepath", "swagger-ui/dist/")
	service.Mux.Handle("GET", "/swagger-ui/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger-ui/dist/", "route", "GET /swagger-ui/*filepath")

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")

	h = ctrl.FileHandler("/swagger-ui/", "swagger-ui/dist/index.html")
	service.Mux.Handle("GET", "/swagger-ui/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger-ui/dist/index.html", "route", "GET /swagger-ui/")
}
