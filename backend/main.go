package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	
	//load env variables
	godotenv.Load()
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		panic("DB_DSN is not set")
	}

	// db connection
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		})
	})

	r.Run(":8080")
}
