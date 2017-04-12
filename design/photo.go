package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var PhotoPayload = Type("PhotoPayload", func() {
	Description("Photo Payload is used to create photos.")
	Attribute("alt", String, "name of a post", func() {
		MinLength(1)
	})
	Attribute("original_url", String, "URL to full-size image", func() {
		MinLength(1)
	})
	Attribute("thumbnail_url", String, "URL to thumbnail-size image", func() {
		MinLength(1)
	})
	Attribute("published", Boolean, "is the photo published")
	Required("alt", "original_url", "thumbnail_url", "published")
})

var PhotoMedia = MediaType("application/vnd.hixio.goa.photo", func() {
	TypeName("photo")
	Reference(PhotoPayload)

	Attributes(func() {
		Attribute("id", Integer, "Unique Photo ID")
		Attribute("alt")
		Attribute("original_url")
		Attribute("thumbnail_url")
		Attribute("published")
		Attribute("created_at", DateTime)
		Attribute("updated_at", DateTime)
		Required("id", "alt", "original_url", "thumbnail_url", "published")
	})

	View("default", func() {
		Attribute("id")
		Attribute("alt")
		Attribute("original_url")
		Attribute("thumbnail_url")
		Attribute("published")
		Attribute("created_at")
		Attribute("updated_at")
	})
})

var _ = Resource("photo", func() {
	Description("A blog photo")
	BasePath("/photos")
	DefaultMedia(PhotoMedia)

	Security(JWT, func() {
		Scope("api:access")
	})

	Action("list", func() {
		Description("lists all published photos")
		Routing(GET(""))
		NoSecurity()
		Response(OK, CollectionOf(PhotoMedia))
	})

	Action("show", func() {
		Description("shows a photo")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer)
		})
		NoSecurity()
		Response(OK)
		Response(NotFound)
	})

	Action("create", func() {
		Description("creates a photo")
		Routing(POST(""))
		Payload(PhotoPayload)
		Response(Created, PhotoMedia)
		Response(InternalServerError, String)
		//Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Description("deletes a photo")
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
