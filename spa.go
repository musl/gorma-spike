package main

import "github.com/goadesign/goa"

// SpaController implements the spa resource.
type SpaController struct {
	*goa.Controller
}

// NewSpaController creates a spa controller.
func NewSpaController(service *goa.Service) *SpaController {
	return &SpaController{Controller: service.NewController("SpaController")}
}
