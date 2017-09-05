package main

import (
	//"log"
	//"time"

	"github.com/gin-gonic/gin"

	//"./models"
	//"./controllers"
	"douyou-manager-server/controllers"
	"douyou-manager-server/models"
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

	db.LogMode(true)

	defer db.Close()

	//models.CreateTableUser(db)

	router := gin.Default()

	userController := controllers.NewUserController(db)
	replyController := controllers.NewReplyController(db)

	router.GET("/users/:userId", userController.Get)
	router.POST("/users", userController.Create)
	router.POST("/replies", replyController.Create)
	router.GET("/replies", replyController.GetListByUserId)

	router.GET("/ping", ping)

	router.Run(":8001")
}
