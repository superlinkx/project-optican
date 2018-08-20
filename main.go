package main

import (
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

func main() {

	cs := fmt.Sprintf("host=%s user=%s password=%s dbname=healthpack sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	r := gin.Default()

	db, err := gorm.Open("postgres", cs)
	for err != nil {
		time.Sleep(time.Second)
		glog.Error("Couldn't connect due to error: ", err)
		db, err = gorm.Open("postgres", cs)
	}
	defer db.Close()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":        "pong",
			"error":          err,
			"connection_str": cs,
		})
	})
	r.Run(":" + os.Getenv("APP_PORT"))
}
