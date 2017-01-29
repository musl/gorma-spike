package main

import "github.com/goadesign/goa"

// StaticController implements the spa resource.
type StaticController struct {
	*goa.Controller
}

// NewStaticController creates a spa controller.
func NewStaticController(service *goa.Service) *StaticController {
	return &StaticController{Controller: service.NewController("StaticController")}
}
