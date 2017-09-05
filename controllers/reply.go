package controllers

import (
	//"../models"
	"douyou-manager-server/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	//"time"
)

type ReplyController struct {
	Controller
	db *gorm.DB
}

func NewReplyController(db *gorm.DB) *ReplyController {
	return &ReplyController{db: db}
}

func (ctl ReplyController) GetListByUserId(c *gin.Context) {
	userId := c.Query("userId")

	replies := []models.Reply{}
	err := ctl.db.Where("user_id = ?", userId).Find(&replies).Error

	if err != nil {
		ctl.ErrorResponse(c, http.StatusInternalServerError, "数据库读写错误")
		return
	}

	for i, _ := range replies {
		ctl.db.Model(replies[i]).Related(&replies[i].Keywords)
	}

	//gotReplies := replies

	ctl.SuccessResponse(c, replies)
}

func (ctl ReplyController) Create(c *gin.Context) {
	var requestJson models.CreateReplyRequest
	err := c.BindJSON(&requestJson)
	if err != nil {
		ctl.ErrorResponse(c, http.StatusBadRequest, "请求参数错误")
		return
	}

	reply := models.Reply{
		UserId:        requestJson.UserId,
		IsEmptyAnswer: 0,
		Answer:        requestJson.Answer,
	}

	reply.Id = bson.NewObjectId().Hex()

	if requestJson.IsEmptyAnswer {
		reply.IsEmptyAnswer = 1
	}
	switch requestJson.Type {
	case "keyword":
		reply.Type = 1
	case "fulltext":
		reply.Type = 2
	}

	var replyKeywords []models.ReplyKeywords

	for _, keyword := range requestJson.Keywords {
		replyKeyword := models.ReplyKeywords{
			Type:    reply.Type,
			ReplyId: reply.Id,
			Text:    keyword,
		}
		replyKeyword.Id = bson.NewObjectId().Hex()
		replyKeywords = append(replyKeywords, replyKeyword)
	}

	tx := ctl.db.Begin()

	for _, replyKeyword := range replyKeywords {
		if err := tx.Create(&replyKeyword).Error; err != nil {
			tx.Rollback()
			ctl.ErrorResponse(c, http.StatusInternalServerError, "数据库读写错误")
			return
		}
	}

	reply.KeywordsCount = uint(len(replyKeywords))

	if err := tx.Create(&reply).Error; err != nil {
		tx.Rollback()
		ctl.ErrorResponse(c, http.StatusInternalServerError, "数据库读写错误")
		return
	}

	tx.Commit()

	ctl.SuccessResponse(c, gin.H{})
}
