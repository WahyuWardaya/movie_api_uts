package routes

import (
	"movie_api/controllers"

	"github.com/gin-gonic/gin"
)

func DirectorRoutes(router *gin.Engine) {
	directors := router.Group("/directors")
	{
		directors.GET("", controllers.GetDirectors)
		directors.POST("", controllers.CreateDirectors)
		directors.GET("/:id", controllers.GetDirectorsByID)
		directors.PUT("/:id", controllers.UpdateDirectors)
		directors.DELETE("/:id", controllers.DeleteDirectors)
	}
}