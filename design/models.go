package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("TheDatabase", func() {
	Description("The global storage group for the app")
	Store("postgres", gorma.Postgres, func() {
		Description("This is the postgres relational store")

		Model("User", func() {
			Description("This is the User model")

			RendersTo(UserMedia)
			BuildsFrom(func() {
				Payload("user", "create")
			})

			Field("id", gorma.Integer, func() {
				Description("Surrogate key for User model")
				PrimaryKey()
			})

			Field("Name", gorma.String, func() {})
			Field("Email", gorma.String, func() {})
			Field("Password", gorma.String, func() {})

			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		Model("Post", func() {
			Description("This is the Post model")

			RendersTo(PostMedia)
			BuildsFrom(func() {
				Payload("post", "create")
			})

			Field("id", gorma.Integer, func() {
				Description("Surrogate key for Post model")
				PrimaryKey()
			})

			Field("title", gorma.String, func() {})
			Field("body", gorma.String, func() {})
			Field("published", gorma.Boolean, func() {})

			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})

		Model("Photo", func() {
			Description("This is the Photo model")

			RendersTo(PhotoMedia)
			BuildsFrom(func() {
				Payload("photo", "create")
			})

			Field("id", gorma.Integer, func() {
				Description("Surrogate key for Photo model")
				PrimaryKey()
			})

			Field("alt", gorma.String, func() {})
			Field("original_url", gorma.String, func() {})
			Field("thumbnail_url", gorma.String, func() {})
			Field("published", gorma.Boolean, func() {})

			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})
	})
})
