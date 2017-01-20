package main

import (
	"fmt"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/musl/hixio/app"
	"github.com/musl/hixio/models"
)

type PostController struct {
	*goa.Controller
}

func NewPostController(service *goa.Service) *PostController {
	return &PostController{Controller: service.NewController("PostController")}
}

func (c *PostController) List(ctx *app.ListPostContext) error {
	db := models.NewPostDB(ctx.Value("db").(*gorm.DB))
	posts := db.ListPost(ctx)
	return ctx.OK(posts)
}

func (c *PostController) Show(ctx *app.ShowPostContext) error {
	db := models.NewPostDB(ctx.Value("db").(*gorm.DB))

	post, err := db.Get(ctx, ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}

	return ctx.OK(post.PostToPost())
}

func (c *PostController) Create(ctx *app.CreatePostContext) error {
	db := models.NewPostDB(ctx.Value("db").(*gorm.DB))
	p := ctx.Payload
	post := models.Post{
		Title:     p.Title,
		Body:      p.Body,
		Published: p.Published,
	}

	err := db.Add(ctx, &post)
	if err != nil {
		return ctx.InternalServerError(fmt.Sprintf("%v", err))
	}

	return ctx.Created(post.PostToPost())
}

func (c *PostController) Delete(ctx *app.DeletePostContext) error {
	db := models.NewPostDB(ctx.Value("db").(*gorm.DB))

	err := db.Delete(ctx, ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}

	return ctx.NoContent()
}
