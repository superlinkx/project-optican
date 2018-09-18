package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/superlinkx/project-healthpack/db"
	"github.com/superlinkx/project-healthpack/entities"
)

func main() {
	gormdb := db.SetDB()
	db.RunMigrationsUp(gormdb)

	var recordType entities.RecordType
	var record []entities.Record
	var tag entities.Tag
	var tags []entities.Tag
	gormdb.Create(&entities.Tag{Name: "Test1"})
	gormdb.Create(&entities.RecordType{Name: "Name", Units: "em"})
	gormdb.First(&tag, 1)
	tags = append(tags, tag)
	gormdb.First(&recordType, 1)
	gormdb.Create(&entities.Record{RecordType: recordType, Tags: tags, Amount: 0.01, Note: "Hello World", RecordTime: time.Now()})
	gormdb.Preload("RecordType").Preload("Tags").Find(&record)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":     "pong",
			"record_type": recordType,
			"record":      record,
		})
	})
	r.Run(":" + os.Getenv("APP_PORT"))
}
