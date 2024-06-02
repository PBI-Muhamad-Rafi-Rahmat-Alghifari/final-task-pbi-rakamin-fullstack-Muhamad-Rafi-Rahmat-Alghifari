package router

import (
	"FinalTaskAPI/controllers"
	"FinalTaskAPI/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Halo, selamat datang di Final Task API Rakamin. Gunakan endpoint yang berada di dokumentasi berikut: http://localhost:8080/swagger/index.html",
		})
	})

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.CreateUser)
		userRouter.POST("/login", controllers.LoginUser)
		userRouter.PUT("/:userID", middleware.Authentication(), controllers.EditUser)
		userRouter.DELETE("/:userID", middleware.Authentication(), controllers.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetPhoto)
		photoRouter.PUT("/:photoID", middleware.PhotoAuthorization(), controllers.EditPhoto)
		photoRouter.DELETE("/:photoID", middleware.PhotoAuthorization(), controllers.DeletePhoto)
	}

	return r
}
