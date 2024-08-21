package controllers

import (
	"gin-ranking/models"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

type UserApi struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (u UserController) Register(c *gin.Context) {
	//接受用户名和密码
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirmPassword", "")
	if username == "" || password == "" || confirmPassword == "" {
		ReturnError(c, 400, "用户名或密码不能为空")
		return
	}

	if password != confirmPassword {
		ReturnError(c, 400, "两次密码不一致")
		return
	}

	user, err := models.GetUserInfoByUsername(username)
	if user.Id != 0 {
		ReturnError(c, 400, "用户名已存在")
		return
	}

	//添加用户
	id, err := models.AddUserAndPassword(username, EncryMd5(password))
	if err != nil {
		ReturnError(c, 400, err.Error())
		return
	}
	ReturnSuccess(c, 200, "success", id, 0)

}

func (u UserController) Login(c *gin.Context) {
	//接受用户名和密码
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")

	if username == "" || password == "" {
		ReturnError(c, 400, "用户名或密码不能为空")
		return
	}

	user, _ := models.GetUserInfoByUsername(username)
	if user.Id == 0 {
		ReturnError(c, 400, "用户不存在")
		return
	}
	if user.Password != EncryMd5(password) {
		ReturnError(c, 400, "密码错误")
		return
	}
	session := sessions.Default(c)
	session.Set("login-user"+strconv.Itoa(user.Id), user.Id)
	session.Save()
	data := UserApi{Id: user.Id, Name: user.Name}
	ReturnSuccess(c, 200, "success", data, 0)
}

func (u UserController) GetUserInfo(c *gin.Context) {
	//业务逻辑
	idStr := c.Param("id")

	id, _ := strconv.Atoi(idStr)

	user, _ := models.GetUserTest(id)
	ReturnSuccess(c, 0, "success", user, 1)
}

func (u UserController) GetUserList(c *gin.Context) {
	userList, err := models.GetUserList()
	if err != nil {
		ReturnError(c, 404, err.Error())
		return
	} else {
		ReturnSuccess(c, 0, "success", userList, 1)
	}
}

func (u UserController) GetList(c *gin.Context) {
	//logger.Write("日志信息", "user")
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("exception try catch", err)
	// 	}
	// }()

	num1 := 1
	num2 := 0
	num3 := num1 / num2

	ReturnError(c, 404, num3)
}

func (u UserController) AddUser(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	id, err := models.AddUser(name)
	if err != nil {
		ReturnError(c, 404, err)
		return
	} else {
		ReturnSuccess(c, 0, "success", id, 1)
	}
}

func (u UserController) UpdateUser(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)
	models.UpdateUser(id, name)
	ReturnSuccess(c, 0, "update success", nil, 1)

}

func (u UserController) DeleteUser(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)
	err := models.DeleteUser(id)
	if err != nil {
		ReturnError(c, 404, "delete fail")
		return
	}
	ReturnSuccess(c, 0, "delete success", nil, 1)
}
