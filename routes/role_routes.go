package routes

import (
	"movie_api/controllers"

	"github.com/gin-gonic/gin"
)

func RoleRoutes(router *gin.Engine) {
	roles := router.Group("/roles")
	{
		roles.GET("", controllers.GetRoles)
		roles.POST("", controllers.CreateRoles)
		roles.GET("/:id", controllers.GetRolesByID)
		roles.PUT("/:id", controllers.UpdateRoles)
		roles.DELETE("/:id", controllers.DeleteRoles)
	}
}
