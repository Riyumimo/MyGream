package router

import (
	"MyGream/controllers"
	"MyGream/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {

	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)

	}
	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
	}

	socialMediaRouter := r.Group("/socialmedia")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)

	}
	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("photo_id=:id/", controllers.CreateComment)
	}
	return r
}
