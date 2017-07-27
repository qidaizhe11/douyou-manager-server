package main

import (
	//"log"
	//"time"

	"github.com/gin-gonic/gin"

	"./models"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func main() {

	db, err := models.InitDB()
	if err != nil {
		println("err open databases", err)
		return
	}

	defer db.Close()

	models.CreateTableUser(db)

	router := gin.New()

	router.GET("/ping", ping)

	router.Run(":8001")
}
