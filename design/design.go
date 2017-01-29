package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("hixio", func() {
	Description("Mike Hix's Blog Service")
	Host("127.0.0.1:8080")
	BasePath("/api/v1")
})
