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
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.GET("/", controllers.GetAllPhotos)
		photoRouter.GET("/:id", controllers.GetPhotoByID)
		photoRouter.DELETE("/:id", controllers.DeletePhoto)
	}

	socialMediaRouter := r.Group("/socialmedia")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/", controllers.GetAllSocial)
		socialMediaRouter.GET("/:id", controllers.GetSocialById)
		socialMediaRouter.PUT("/:id", middlewares.SocialAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:id", middlewares.SocialAuthorization(), controllers.DeleteSocial)

	}
	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("photo_id=:id/", controllers.CreateComment)
		commentRouter.PUT("/:id", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:id", middlewares.CommentAuthorization(), controllers.DeleteComment)
		commentRouter.GET("/", controllers.GetAllComment)
		commentRouter.GET("/:id", controllers.GetCommentByID)
	}
	return r
}
