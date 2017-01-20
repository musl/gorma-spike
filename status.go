package main

import (
	"github.com/goadesign/goa"
	"github.com/musl/hixio/app"
)

type StatusController struct {
	*goa.Controller
}

func NewStatusController(service *goa.Service) *StatusController {
	return &StatusController{Controller: service.NewController("StatusController")}
}

func (c *StatusController) Check(ctx *app.CheckStatusContext) error {
	return ctx.OK([]byte("OK"))
}
