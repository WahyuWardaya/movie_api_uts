package main

import (
	"movie_api/config"
	"movie_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	router := gin.Default()
	routes.UserRoutes(router)
	routes.GenreRoutes(router)
	routes.RoleRoutes(router)
	routes.ActorRoutes(router)
	routes.DirectorRoutes(router)
	routes.MovieRoutes(router)
	


	router.Run(":3000")
}
