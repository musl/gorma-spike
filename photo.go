package main

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/musl/hixio/app"
	"github.com/musl/hixio/models"
)

type PhotoController struct {
	*goa.Controller
}

func NewPhotoController(service *goa.Service) *PhotoController {
	return &PhotoController{Controller: service.NewController("PhotoController")}
}

func (c *PhotoController) List(ctx *app.ListPhotoContext) error {
	db := models.NewPhotoDB(ctx.Value("db").(*gorm.DB))
	posts := db.ListPhoto(ctx)
	return ctx.OK(posts)
}

func (c *PhotoController) Show(ctx *app.ShowPhotoContext) error {
	db := models.NewPhotoDB(ctx.Value("db").(*gorm.DB))

	post, err := db.Get(ctx, ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}

	return ctx.OK(post.PhotoToPhoto())
}

func (c *PhotoController) Create(ctx *app.CreatePhotoContext) error {
	db := models.NewPhotoDB(ctx.Value("db").(*gorm.DB))
	post := models.PhotoFromPhotoPayload(ctx.Payload)

	err := db.Add(ctx, post)
	if err != nil {
		return ctx.InternalServerError(fmt.Sprintf("%v", err))
	}

	return ctx.Created(post.PhotoToPhoto())
}

func (c *PhotoController) Delete(ctx *app.DeletePhotoContext) error {
	db := models.NewPhotoDB(ctx.Value("db").(*gorm.DB))

	err := db.Delete(ctx, ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}

	return ctx.NoContent()
}
