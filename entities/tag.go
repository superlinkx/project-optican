package entities

import "github.com/jinzhu/gorm"

// Tag Entity used to store tags
type Tag struct {
	gorm.Model
	Name string `json:"name"`
}
