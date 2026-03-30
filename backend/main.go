package main

import (
	"backend/config"
	"backend/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	
	//load env variables
	godotenv.Load()

	// db connection
	db, err := config.ConnectDB()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// repository initialization
	_ = repository.NewPermintaanDarahRepository(db)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		})
	})

	r.Run(":8080")
}
