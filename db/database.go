package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

// SetDB sets up connection to db
func SetDB() *gorm.DB {
	cs := fmt.Sprintf("host=%s user=%s password=%s dbname=optican sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("postgres", cs)
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
