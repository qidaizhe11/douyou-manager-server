package main

import (
	//"log"
	//"time"

	"github.com/gin-gonic/gin"

	//"./models"
	//"./controllers"
	"douyou-manager-server/models"
	"douyou-manager-server/controllers"
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

	//models.CreateTableUser(db)

	router := gin.Default()

	userController := controllers.NewUserController(db)

	router.GET("/users/:userId", userController.Get)
	router.POST("/users", userController.Create)

	router.GET("/ping", ping)

	router.Run(":8001")
}
