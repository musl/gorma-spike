package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("TheDatabase", func() {
	Description("The global storage group for the app")
	Store("postgres", gorma.Postgres, func() {
		Description("This is the postgres relational store")

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
	})
})
