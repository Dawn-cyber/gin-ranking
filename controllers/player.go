package controllers

import (
	"gin-ranking/cache"
	"gin-ranking/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PlayerController struct {
}

func (p PlayerController) GetList(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)
	re, err := models.GetPlayer(aid, "id asc")
	if err != nil {
		ReturnError(c, 500, err.Error())
		return
	}
	ReturnSuccess(c, 200, "success", re, 0)
}

func (p PlayerController) GetRanking(c *gin.Context) {

	idStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(idStr)

	redisKey := "ranking_" + idStr
	rs, err := cache.Rdb.ZRevRange(cache.Rctx, redisKey, 0, -1).Result()
	if err == nil && len(rs) > 0 {
		var players []models.Player
		for _, v := range rs {
			id, _ := strconv.Atoi(v)
			player, _ := models.GetPLayerInfo(id)
			players = append(players, player)
		}
		ReturnSuccess(c, 200, "success", players, 0)
		return
	}

	re, _ := models.GetPlayer(aid, "score desc")
	if err == nil {
		for _, value := range re {
			cache.Rdb.ZAdd(cache.Rctx, redisKey, cache.Zscore(value.Id, value.Score)).Err()
		}
		cache.Rdb.Expire(cache.Rctx, redisKey, 24*time.Hour)
		ReturnSuccess(c, 200, "success", re, 0)
		return
	}
	ReturnError(c, 500, err.Error())
}
