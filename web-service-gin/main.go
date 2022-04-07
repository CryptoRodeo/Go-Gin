package main

import (
	"example/web-service-gin/album"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	createRoutes(router)
	router.Run("localhost:8080")
}

func createRoutes(router *gin.Engine) {
	router.GET("/albums", album.GetAlbums)
	router.GET("/albums/:id", album.GetAlbumByID)

	router.POST("/albums", album.PostAlbums)
}
