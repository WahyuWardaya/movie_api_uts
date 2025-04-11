package routes

import (
	"movie_api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.GET("", controllers.GetUsers)
		users.POST("", controllers.CreateUser)
		users.GET("/:id", controllers.GetUserByID)
		users.PUT("/:id", controllers.UpdateUsers)
		users.DELETE("/:id", controllers.DeleteUsers)
	}
}
