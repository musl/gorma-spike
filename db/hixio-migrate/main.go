package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/musl/hixio/models"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=hixio dbname=hixio sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to the database: %v", err))
	}
	defer db.Close()

	db.AutoMigrate(&models.User{}, &models.Post{})
}
