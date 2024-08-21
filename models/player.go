package models

import (
	"gin-ranking/dao"

	"github.com/jinzhu/gorm"
)

type Player struct {
	Id          int
	Aid         int
	Ref         string
	Nickname    string
	Declaration string
	Avatar      string
	Score       int
}

func (Player) TableName() string {
	return "player"
}

func GetPlayerList() ([]Player, error) {
	var players []Player
	err := dao.Db.Find(&players).Error
	return players, err
}

func GetPlayer(aid int, sort string) ([]Player, error) {
	var player []Player
	err := dao.Db.Where("aid = ?", aid).Order(sort).Find(&player).Error
	return player, err
}

func GetPLayerInfo(id int) (Player, error) {
	var player Player
	err := dao.Db.Where("id = ?", id).First(&player).Error
	return player, err
}

func UpdatePlayer(id int) {
	dao.Db.Model(&Player{}).Where("id = ?", id).Update("score", gorm.Expr("score + ?", 1))
}
