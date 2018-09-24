package db

import (
	"github.com/jinzhu/gorm"
	"github.com/superlinkx/project-optican/entities"
)

// RunMigrationsUp migrates the app to current
func RunMigrationsUp(gormdb *gorm.DB) {
	gormdb.AutoMigrate(&entities.Tag{})
	gormdb.AutoMigrate(&entities.RecordType{})
	gormdb.AutoMigrate(&entities.Record{})
	gormdb.Model(&entities.Record{}).AddForeignKey("record_type_id", "record_types(id)", "RESTRICT", "RESTRICT")
}

// RunMigrationsDown resets the database
func RunMigrationsDown(gormdb *gorm.DB) {
	gormdb.DropTable("record_tags")
	gormdb.DropTable(&entities.Record{})
	gormdb.DropTable(&entities.RecordType{})
	gormdb.DropTable(&entities.Tag{})
}
