package structs

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Title         string `json:"title"`
	Author        string `json:"author"`
	ISBN          string `json:"isbn"`
	PublishedDate string `json:"published_date"`
}
