package config

import (
	"fmt"
	"log"
	"pretty/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func ConnectToDB() *gorm.DB {
	conn := "user=postgres password=password dbname=test_mekar port=5432 sslmode=disable"
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		fmt.Println("[CONFIG.ConnectDB] error when connect to database")
		log.Fatal(err)
	} else {
		fmt.Println("SUCCES CONNECT TO DATABASE")
	}

	models.Migrate(db)

	return db
}
