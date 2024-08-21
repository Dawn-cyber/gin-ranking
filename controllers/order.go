package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

type Search struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func (o OrderController) GetList(c *gin.Context) {
	//get the params from form
	//id := c.PostForm("id")
	//name := c.DefaultPostForm("name", "default")

	//get params from json by make
	// param := make(map[string]interface{})
	// err := c.BindJSON(&param)
	search := &Search{}
	err := c.BindJSON(&search)
	if err == nil {
		ReturnSuccess(c, 0, search.Name, search, 1)
		return
	}

	ReturnError(c, 0, "question error")

}
