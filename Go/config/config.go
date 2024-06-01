package config

import (
	"github.com/jinzhu/gorm"
	"github.com/yasir16/simple-api-rbac/structs"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "docker:docker@tcp(db:3306)/godb?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(structs.Book{})
	return db
}
