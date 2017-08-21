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

func (user *User) ToUserResponse() UserResponse {

	userResponse := UserResponse{
		Id: user.Id,
		DoubanId: user.DoubanId,
		Nickname: user.Nickname,
		AvatarUrl: user.AvatarUrl,
		Location: user.Location,
		LastLoginTime: user.LastLoginTime,
	}

	switch user.Gender {
	case 1:
		userResponse.Gender = "male"
	case 2:
		userResponse.Gender = "female"
	default:
		userResponse.Gender = ""
	}

	return userResponse
}
