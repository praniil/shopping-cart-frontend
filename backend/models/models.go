package models

import "gorm.io/gorm"

// used to map the data from go to database

type Products struct {
	gorm.Model
	Title     string `json:"title"`
	Price     int64  `json:"price"`
	Thumbnail string `json:"thumbnail"`
	Image     string `json:"image"`
	Quantity  int64  `json:"class"`
}
