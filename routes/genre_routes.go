package routes

import (
	"movie_api/controllers"

	"github.com/gin-gonic/gin"
)

func GenreRoutes(router *gin.Engine) {
	genres := router.Group("/genres")
	{
		genres.GET("", controllers.GetGenres)
		genres.POST("", controllers.CreateGenres)
		genres.GET("/:id", controllers.GetGenresByID)
		genres.PUT("/:id", controllers.UpdateGenres)
		genres.DELETE("/:id", controllers.DeleteGenres)
	}
}
