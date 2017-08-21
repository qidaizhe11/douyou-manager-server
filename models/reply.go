package models

import (
	"time"
)

type Reply struct {
	BaseModel
	UserId        string
	Type          uint // 1：关键词匹配；2：精确匹配
	Answer        string
	IsEmptyAnswer uint
	StartTime     *time.Time
	EndTime       *time.Time
	Priority      uint // 优先级，1/2/3/4/5共5级
	KeywordsCount uint
	Keywords      []ReplyKeywords
}

type ReplyKeywords struct {
	BaseModel
	ReplyId string
	Type    uint // 1：关键词匹配；2：精确匹配
	Text    string
}

type CreateReplyRequest struct {
	UserId        string   `json:"userId" binding:"required"`
	Type          string   `json:"type" binding:"required"`
	Keywords      []string `json:"keywords" binding:"required"`
	IsEmptyAnswer bool     `json:"isEmptyAnswer"`
	Answer        string   `json:"answer"`
}
