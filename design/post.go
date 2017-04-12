package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var PostPayload = Type("PostPayload", func() {
	Description("Post Payload is used to create posts.")
	Attribute("title", String, "name of a post", func() {
		MinLength(1)
	})
	Attribute("body", String, "body of a post", func() {
		MinLength(1)
	})
	Attribute("published", Boolean, "is the post published")
	Required("title", "body", "published")
})

var PostMedia = MediaType("application/vnd.hixio.goa.post", func() {
	TypeName("post")
	Reference(PostPayload)

	Attributes(func() {
		Attribute("id", Integer, "Unique Post ID")
		Attribute("title")
		Attribute("body")
		Attribute("published")
		Attribute("created_at", DateTime)
		Attribute("updated_at", DateTime)
		Required("id", "title", "body", "published")
	})

	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("body")
		Attribute("created_at", DateTime)
		Attribute("updated_at", DateTime)
		Attribute("published")
	})
})

var _ = Resource("post", func() {
	Description("A blog post")
	BasePath("/posts")
	DefaultMedia(PostMedia)

	Security(JWT, func() {
		Scope("api:access")
	})

	Action("list", func() {
		Description("lists all publisged posts")
		Routing(GET(""))
		NoSecurity()
		Response(OK, CollectionOf(PostMedia))
	})

	Action("show", func() {
		Description("shows a post")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer)
		})
		NoSecurity()
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Description("creates a post")
		Routing(POST(""))
		Payload(PostPayload)
		Response(Created, PostMedia)
		Response(InternalServerError, String)
		//Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Description("deletes a post")
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id", Integer)
		})
		Response(NoContent)
		Response(NotFound)
		Response(InternalServerError, String)
		//Response(BadRequest, ErrorMedia)
	})
})
