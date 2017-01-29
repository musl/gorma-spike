package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("swagger", func() {
	Description("The API Specification by Swagger")
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger-ui/*filepath", "swagger-ui/dist/")
})
