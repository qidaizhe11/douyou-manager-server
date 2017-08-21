package controllers

import (
	//"../models"
	"douyou-manager-server/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

type UserController struct {
	Controller
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db: db}
}

func (ctl UserController) Get(c *gin.Context) {
	id := c.Param("userId")

	user := models.User{}
	//err := ctl.db.First(&user, id).Error
	err := ctl.db.Where("id = ?", id).First(&user).Error

	if err != nil {
		ctl.ErrorResponse(c, http.StatusInternalServerError, "数据库读写错误")
		return
	}

	userResponse := user.ToUserResponse()

	ctl.SuccessResponse(c, userResponse)
}

type CreateUserRequest struct {
	DoubanId  string `json:"doubanId" binding:"required"`
	Nickname  string `json:"nickname" binding:"required"`
	AvatarUrl string `json:"avatarUrl"`
	Gender    string `json:"gender"`
	Location  string `json:"location"`
}

func (ctl UserController) Create(c *gin.Context) {
	var json CreateUserRequest
	err := c.BindJSON(&json)
	if err != nil {
		ctl.ErrorResponse(c, http.StatusBadRequest, "请求参数错误")
		//println("controller, user, create, bindjson error:", err.Error())
		return
	}

	user := models.User{
		DoubanId:      json.DoubanId,
		Nickname:      json.Nickname,
		LastLoginTime: time.Now(),
	}

	user.Id = bson.NewObjectId().Hex()

	if json.AvatarUrl != "" {
		user.AvatarUrl = json.AvatarUrl
	}
	if json.Gender != "" {
		switch json.Gender {
		case "male":
			user.Gender = 1
		case "female":
			user.Gender = 2
		}
	}
	if json.Location != "" {
		user.Location = json.Location
	}

	ctl.db.Create(&user)

	ctl.SuccessResponse(c, gin.H{})
}
