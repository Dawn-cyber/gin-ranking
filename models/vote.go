package models

import (
	"gin-ranking/dao"
	"time"
)

type Vote struct {
	Id         int   `json:"id"`
	UserId     int   `json:"user_id"`
	PlayerId   int   `json:"player_id"`
	CreateTime int64 `json:"create_time"`
}

func (Vote) TableName() string {
	return "vote"
}

func GetVoteInfo(userId int, PlayerId int) (Vote, error) {
	var vote Vote
	err := dao.Db.Where("user_id = ? and player_id = ?", userId, PlayerId).First(&vote).Error
	return vote, err
}

func AddVote(userId int, PlayerId int) (int, error) {
	vote := Vote{
		UserId:     userId,
		PlayerId:   PlayerId,
		CreateTime: time.Now().Unix(),
	}
	err := dao.Db.Create(&vote).Error
	return vote.Id, err
}
