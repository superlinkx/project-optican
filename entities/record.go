package entities

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Record Entity used to describe a health record
type Record struct {
	gorm.Model
	Amount       float32    `json:"amount"`
	Note         string     `json:"note"`
	RecordTime   time.Time  `json:"record_time"`
	RecordTypeID uint       `json:"record_type_id"`
	RecordType   RecordType `json:"record_type"`
	Tags         []Tag      `gorm:"many2many:record_tags;" json:"tags"`
}
