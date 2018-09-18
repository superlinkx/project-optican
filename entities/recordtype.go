package entities

import "github.com/jinzhu/gorm"

// RecordType Entity used for describing types of records
type RecordType struct {
	gorm.Model
	Name  string `json:"name"`
	Units string `json:"units"`
}
