package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Car struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}
