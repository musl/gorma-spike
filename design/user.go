package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var UserPayload = Type("UserPayload", func() {
	Description("Post Payload is used to create users.")
	Attribute("id", Integer, "surrogate key of a user", func() {
		Minimum(1)
	})
	Attribute("name", String, "name of a user", func() {
		MinLength(1)
	})
	Attribute("email", String, "email of a user", func() {
		MinLength(1)
	})
	Attribute("password", String, "password of the user")
	Required("name", "email", "password")
})

var UserMedia = MediaType("application/vnd.hixio.goa.user", func() {
	TypeName("user")
	Reference(UserPayload)

	Attributes(func() {
		Attribute("id", Integer, "Unique Post ID")
		Attribute("name")
		Attribute("email")
		Required("id", "name", "email")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("email")
	})
})

var _ = Resource("user", func() {
	Description("A User")
	BasePath("/users")
	DefaultMedia(UserMedia)

	Security(JWT, func() {
		Scope("api:access")
	})

	Action("list", func() {
		Description("lists all users")
		Routing(GET(""))
		Response(OK, CollectionOf(UserMedia))
	})

	Action("show", func() {
		Description("shows a user")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer)
		})
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Description("creates a user")
		Routing(POST(""))
		Payload(UserPayload)
		Response(Created, UserMedia)
		Response(InternalServerError, String)
	})

	Action("delete", func() {
		Description("deletes a user")
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id", Integer)
		})
		Response(NoContent)
		Response(NotFound)
		Response(InternalServerError, String)
	})
})
