package controller

import (
	"github.com/MeiramSh/dostap/internal/controller/middleware"
	"github.com/MeiramSh/dostap/pkg"
	"github.com/gin-gonic/gin"

	"github.com/MeiramSh/dostap/internal/controller/auth"
	"github.com/MeiramSh/dostap/internal/controller/user"
	"github.com/MeiramSh/dostap/internal/controller/friend"
	"github.com/MeiramSh/dostap/internal/repository"
)

func Setup(app pkg.Application, router *gin.Engine) {
	db := app.DB

	loginController := &auth.AuthController{
		UserRepository: repository.NewUserRepository(db),
	}

	userController := &user.UserController{
		UserRepository: repository.NewUserRepository(db),
	}
	friendController := &friend.FriendController{
		FriendRepository: repository.NewFriendRepository(db),
	}
	router.POST("/signup", loginController.Signup)
	router.POST("/signin", loginController.Signin)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello,World"})
	})
	router.Use(middleware.JWTAuth(`access-key`))
	api := router.Group("api/v0")
	{
		userRouter := api.Group("/user")
		{

			userRouter.GET("/profile", userController.GetProfile)
		}
		friendRouter := api.Group("/friend")
		{
			friendRouter.POST("/request", friendController.SendFriendRequest)
		}
	}
}
