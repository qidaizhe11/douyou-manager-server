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
	Keywords      []ReplyKeywords `gorm:"ForeignKey:Id"`
}

type ReplyKeywords struct {
	BaseModel
	ReplyId string
	Type    uint // 1：关键词匹配；2：精确匹配
	Text    string
}

type ReplyResponse struct {
	UserId string `json:"userId"`
	Type string `json:"type"`
	IsEmptyAnswer bool `json:"isEmptyAnswer"`
	Answer string `json:"answer"`
	Priority int `json:"priority"`
	StartTime time.Time `json:"startTime"`
	EndTime time.Time `json:"endTime"`
	Keywords []ReplyKeywordsResponse `json:"keywords"`
}

type ReplyKeywordsResponse struct {
	ReplyId string `json:"replyId"`
	Type string `json:"type"`
	Text string `json:"text"`
}

type CreateReplyRequest struct {
	UserId        string   `json:"userId" binding:"required"`
	Type          string   `json:"type" binding:"required"`
	Keywords      []string `json:"keywords" binding:"required"`
	IsEmptyAnswer bool     `json:"isEmptyAnswer"`
	Answer        string   `json:"answer"`
}
