package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("status", func() {
	Description("An API status endpoint")
	BasePath("/status")

	Action("check", func() {
		Description("A basic status-check endpoint")
		Routing(GET(""))
		Response(OK, "text/plain")
	})
})
