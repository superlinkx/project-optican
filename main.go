package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var db *sql.DB

// Record Entity used to describe a health record
type Record struct {
	gorm.Model
	RecordTypeID uint
	RecordType   RecordType `json:"record_type"`
}

// RecordType Entity used for describing types of records
type RecordType struct {
	gorm.Model
	Name  string `json:"name"`
	Units string `json:"units"`
}

func main() {
	gormdb := setDB()
	runMigrations(gormdb)

	var recordType RecordType
	var record []Record
	gormdb.Create(&RecordType{Name: "Name", Units: "em"})
	gormdb.First(&recordType, 1)
	gormdb.Create(&Record{RecordType: recordType})
	gormdb.Preload("RecordType").Find(&record)

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

func setDB() *gorm.DB {
	var err error
	cs := fmt.Sprintf("host=%s user=%s password=%s dbname=healthpack sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	db, err = sql.Open("postgres", cs)
	for err != nil {
		time.Sleep(time.Second)
		glog.Error("Couldn't connect due to error: ", err)
		db, err = sql.Open("postgres", cs)
	}
	gormdb, err := gorm.Open("postgres", db)
	for err != nil {
		time.Sleep(time.Second)
		glog.Error("Couldn't connect due to error: ", err)
		gormdb, err = gorm.Open("postgres", db)
	}
	return gormdb
}

func runMigrations(gormdb *gorm.DB) {
	gormdb.AutoMigrate(&RecordType{})
	gormdb.AutoMigrate(&Record{})
	gormdb.Model(&Record{}).AddForeignKey("record_type_id", "record_types(id)", "RESTRICT", "RESTRICT")
}
