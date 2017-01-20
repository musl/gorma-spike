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

var PostPayload = Type("PostPayload", func() {
	Description("Post Payload is used to create posts.")
	Attribute("id", Integer, "surrogate key of a post", func() {
		Minimum(1)
	})
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
		Required("id", "title", "body", "published")
	})

	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("body")
		Attribute("published")
	})
})

var _ = Resource("status", func() {
	Description("An API status endpoint")
	BasePath("/status")

	Action("check", func() {
		Description("A basic status-check endpoint")
		Routing(GET(""))
		Response(OK, "text/plain")
	})
})

var _ = Resource("post", func() {
	Description("A blog post")
	BasePath("/posts")
	DefaultMedia(PostMedia)

	Action("list", func() {
		Description("lists all posts")
		Routing(GET(""))
		Response(OK, CollectionOf(PostMedia))
	})

	Action("show", func() {
		Description("shows a post")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer)
		})
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

var _ = Resource("spa", func() {
	Description("The static file endpoint.")
	Files("/*filepath", "static/")
})

var _ = Resource("swagger", func() {
	Description("The API Specification by Swagger")
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger-ui/*filepath", "swagger-ui/dist/")
})
