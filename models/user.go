package models

import (
	"gin-ranking/dao"
	"time"
)

type User struct {
	Id         int
	Name       string
	Password   string
	CreateTime int64
	UpdateTime int64
}

func (User) TableName() string {
	return "user"
}

func GetUserInfoByUsername(username string) (User, error) {
	var user User
	err := dao.Db.Where("name = ?", username).First(&user).Error
	return user, err
}

func GetUserInfo(id int) (User, error) {
	var user User
	err := dao.Db.Where("id = ?", id).First(&user).Error
	return user, err
}

func AddUserAndPassword(username string, password string) (int, error) {
	user := User{
		Name:       username,
		Password:   password,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}

func GetUserTest(id int) (User, error) {
	var user User
	err := dao.Db.Where("id = ?", id).First(&user).Error

	return user, err
}

func GetUserList() ([]User, error) {
	var user []User
	err := dao.Db.Where("id < ?", 3).Find(&user).Error
	return user, err

}

func AddUser(name string) (int, error) {
	user := User{Name: name}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}

func UpdateUser(id int, name string) {
	dao.Db.Model(&User{}).Where("id = ?", id).Update("name", name)
}

func DeleteUser(id int) error {
	err := dao.Db.Where("id = ?", id).Delete(&User{}).Error
	return err
}
