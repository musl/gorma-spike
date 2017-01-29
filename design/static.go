package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("static", func() {
	Description("The static file endpoint.")
	Files("/*filepath", "static/")
})
