package models

import "time"

type Reply struct {
	BaseModel
	Type uint
	Answer string
	IsEmptyAnswer uint
	startTime time.Time
	endTime time.Time
	priority uint
}

func (reply *Reply) Insert() error {
	db := GetDB()

	return db.Create(reply).Error
}

func (reply *Reply) Update() error {
	db := GetDB()

	return db.Save(reply).Error
}

func ListRepliesByUserId(userId string) ([]*Reply, error) {
	db := GetDB()

	var replies []*Reply
	err := db.Where("user_id = ?", userId).Find(&replies).Error
	return replies, err
}
