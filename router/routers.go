package router

import (
	"gin-ranking/controllers"
	"gin-ranking/pkg/logger"

	"github.com/gin-contrib/sessions"
	sessions_redis "github.com/gin-contrib/sessions/redis"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	store, _ := sessions_redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	user := r.Group("/user")
	{
		user.POST("/register", controllers.UserController{}.Register)
		user.POST("/login", controllers.UserController{}.Login)
		user.GET("/get/:id", controllers.UserController{}.GetUserInfo)
		user.POST("/userlist", controllers.UserController{}.GetUserList)
		user.POST("/list", controllers.UserController{}.GetList)
		user.POST("/add", controllers.UserController{}.AddUser)
		user.POST("/update", controllers.UserController{}.UpdateUser)
		user.POST("/delete", controllers.UserController{}.DeleteUser)
	}

	order := r.Group("/order")
	{
		order.POST("list", controllers.OrderController{}.GetList)
	}

	player := r.Group("/player")
	{
		player.POST("list", controllers.PlayerController{}.GetList)
		player.POST("ranking", controllers.PlayerController{}.GetRanking)

	}

	vote := r.Group("/vote")
	{
		vote.POST("add", controllers.VoteController{}.AddVote)
	}

	return r
}
