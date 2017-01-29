package main

import (
	"fmt"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/musl/hixio/app"
	"github.com/musl/hixio/models"
)

type UserController struct {
	*goa.Controller
}

func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

func (c *UserController) List(ctx *app.ListUserContext) error {
	db := models.NewUserDB(ctx.Value("db").(*gorm.DB))
	users := db.ListUser(ctx)
	return ctx.OK(users)
}

func (c *UserController) Show(ctx *app.ShowUserContext) error {
	db := models.NewUserDB(ctx.Value("db").(*gorm.DB))

	user, err := db.Get(ctx, ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}

	return ctx.OK(user.UserToUser())
}

func (c *UserController) Create(ctx *app.CreateUserContext) error {
	db := models.NewUserDB(ctx.Value("db").(*gorm.DB))
	user := models.UserFromUserPayload(ctx.Payload)

	err := db.Add(ctx, user)
	if err != nil {
		return ctx.InternalServerError(fmt.Sprintf("%v", err))
	}

	return ctx.Created(user.UserToUser())
}

func (c *UserController) Delete(ctx *app.DeleteUserContext) error {
	db := models.NewUserDB(ctx.Value("db").(*gorm.DB))

	err := db.Delete(ctx, ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}

	return ctx.NoContent()
}
