package models

import (
	"time"
)

type User struct {
	BaseModel
	DoubanId      string
	Nickname      string
	AvatarUrl     string
	Gender        uint // 1: male; 2: female
	Location      string
	LastLoginTime time.Time
}

type UserResponse struct {
	Id            string `json:"id"`
	DoubanId      string `json:"doubanId"`
	Nickname      string `json:"nickname"`
	AvatarUrl     string `json:"avatarUrl"`
	Gender        string `json:"gender"`
	Location      string `json:"location"`
	LastLoginTime time.Time `json:"lastLoginTime"`
}
