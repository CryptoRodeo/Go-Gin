package main

import (
	"example/web-service-gin/routing"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routing.CreateRoutes(router)
	// another routing example
	// album.CreateRoutes(router)
	router.Run("localhost:8080")
}
