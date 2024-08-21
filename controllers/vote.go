package controllers

import (
	"gin-ranking/cache"
	"gin-ranking/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VoteController struct {
}

func (v VoteController) AddVote(c *gin.Context) {
	//获取用户ID
	userIdstr := c.DefaultPostForm("userId", "0")
	playerIdstr := c.DefaultPostForm("playerId", "0")

	userId, _ := strconv.Atoi(userIdstr)
	playerId, _ := strconv.Atoi(playerIdstr)
	if userId == 0 || playerId == 0 {
		ReturnError(c, 400, "参数错误")
		return
	}

	user, _ := models.GetUserInfo(userId)
	if user.Id == 0 {
		ReturnError(c, 400, "用户不存在")
		return
	}

	player, _ := models.GetPLayerInfo(playerId)
	if player.Id == 0 {
		ReturnError(c, 400, "玩家不存在")
		return
	}

	vote, _ := models.GetVoteInfo(userId, playerId)
	if vote.Id != 0 {
		ReturnError(c, 400, "已经投过票了")
		return
	}

	res, err := models.AddVote(userId, playerId)
	if err == nil {
		models.UpdatePlayer(playerId)

		//更新redis
		var redisKey string
		redisKey = "ranking" + strconv.Itoa(player.Aid)
		cache.Rdb.ZIncrBy(cache.Rctx, redisKey, 1, strconv.Itoa(playerId))
		ReturnSuccess(c, 200, "投票成功", res, 0)
		return
	}

	ReturnError(c, 400, "投票失败")

}
