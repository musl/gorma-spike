package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var AuthPayload = Type("AuthPayload", func() {
	Description("Auth Payload is used to auth users.")
	Attribute("email", String, "email of a user", func() {
		MinLength(1)
	})
	Attribute("password", String, "password of the user")
	Required("email", "password")
})

var _ = Resource("auth", func() {
	Description("An authorization service")
	BasePath("/auth")

	Action("jwt", func() {
		Description("Creates a valid JWT")
		Routing(POST("/"))
		Payload(AuthPayload)
		Response(NoContent, func() {
			Headers(func() {
				Header("Authorization", String, "Generated JWT")
			})
		})
		Response(Unauthorized)
	})
})
