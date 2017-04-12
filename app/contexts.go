//************************************************************************//
// API "hixio": Application Contexts
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
	"strconv"
)

// JWTAuthContext provides the auth jwt action context.
type JWTAuthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *AuthPayload
}

// NewJWTAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller jwt action.
func NewJWTAuthContext(ctx context.Context, service *goa.Service) (*JWTAuthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := JWTAuthContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *JWTAuthContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *JWTAuthContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// CreatePhotoContext provides the photo create action context.
type CreatePhotoContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *PhotoPayload
}

// NewCreatePhotoContext parses the incoming request URL and body, performs validations and creates the
// context used by the photo controller create action.
func NewCreatePhotoContext(ctx context.Context, service *goa.Service) (*CreatePhotoContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreatePhotoContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreatePhotoContext) Created(r *Photo) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.hixio.goa.photo")
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreatePhotoContext) InternalServerError(r string) error {
	ctx.ResponseData.Header().Set("Content-Type", "")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// DeletePhotoContext provides the photo delete action context.
type DeletePhotoContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewDeletePhotoContext parses the incoming request URL and body, performs validations and creates the
// context used by the photo controller delete action.
func NewDeletePhotoContext(ctx context.Context, service *goa.Service) (*DeletePhotoContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := DeletePhotoContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeletePhotoContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeletePhotoContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DeletePhotoContext) InternalServerError(r string) error {
	ctx.ResponseData.Header().Set("Content-Type", "")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ListPhotoContext provides the photo list action context.
type ListPhotoContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListPhotoContext parses the incoming request URL and body, performs validations and creates the
// context used by the photo controller list action.
func NewListPhotoContext(ctx context.Context, service *goa.Service) (*ListPhotoContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListPhotoContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListPhotoContext) OK(r PhotoCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.hixio.goa.photo; type=collection")
	if r == nil {
		r = PhotoCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ShowPhotoContext provides the photo show action context.
type ShowPhotoContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewShowPhotoContext parses the incoming request URL and body, performs validations and creates the
// context used by the photo controller show action.
func NewShowPhotoContext(ctx context.Context, service *goa.Service) (*ShowPhotoContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowPhotoContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowPhotoContext) OK(r *Photo) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.hixio.goa.photo")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowPhotoContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// CreatePostContext provides the post create action context.
type CreatePostContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *PostPayload
}

// NewCreatePostContext parses the incoming request URL and body, performs validations and creates the
// context used by the post controller create action.
func NewCreatePostContext(ctx context.Context, service *goa.Service) (*CreatePostContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreatePostContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreatePostContext) Created(r *Post) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.hixio.goa.post")
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreatePostContext) InternalServerError(r string) error {
	ctx.ResponseData.Header().Set("Content-Type", "")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// DeletePostContext provides the post delete action context.
type DeletePostContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewDeletePostContext parses the incoming request URL and body, performs validations and creates the
// context used by the post controller delete action.
func NewDeletePostContext(ctx context.Context, service *goa.Service) (*DeletePostContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := DeletePostContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeletePostContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeletePostContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DeletePostContext) InternalServerError(r string) error {
	ctx.ResponseData.Header().Set("Content-Type", "")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ListPostContext provides the post list action context.
type ListPostContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListPostContext parses the incoming request URL and body, performs validations and creates the
// context used by the post controller list action.
func NewListPostContext(ctx context.Context, service *goa.Service) (*ListPostContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListPostContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListPostContext) OK(r PostCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.hixio.goa.post; type=collection")
	if r == nil {
		r = PostCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ShowPostContext provides the post show action context.
type ShowPostContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewShowPostContext parses the incoming request URL and body, performs validations and creates the
// context used by the post controller show action.
func NewShowPostContext(ctx context.Context, service *goa.Service) (*ShowPostContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowPostContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowPostContext) OK(r *Post) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.hixio.goa.post")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowPostContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// CheckStatusContext provides the status check action context.
type CheckStatusContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewCheckStatusContext parses the incoming request URL and body, performs validations and creates the
// context used by the status controller check action.
func NewCheckStatusContext(ctx context.Context, service *goa.Service) (*CheckStatusContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CheckStatusContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CheckStatusContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// CreateUserContext provides the user create action context.
type CreateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *UserPayload
}

// NewCreateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller create action.
func NewCreateUserContext(ctx context.Context, service *goa.Service) (*CreateUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateUserContext) Created(r *User) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.hixio.goa.user")
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateUserContext) InternalServerError(r string) error {
	ctx.ResponseData.Header().Set("Content-Type", "")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// DeleteUserContext provides the user delete action context.
type DeleteUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewDeleteUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller delete action.
func NewDeleteUserContext(ctx context.Context, service *goa.Service) (*DeleteUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := DeleteUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteUserContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DeleteUserContext) InternalServerError(r string) error {
	ctx.ResponseData.Header().Set("Content-Type", "")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ListUserContext provides the user list action context.
type ListUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller list action.
func NewListUserContext(ctx context.Context, service *goa.Service) (*ListUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListUserContext) OK(r UserCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.hixio.goa.user; type=collection")
	if r == nil {
		r = UserCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ShowUserContext provides the user show action context.
type ShowUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewShowUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller show action.
func NewShowUserContext(ctx context.Context, service *goa.Service) (*ShowUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OK(r *User) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.hixio.goa.user")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
