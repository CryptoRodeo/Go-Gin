package routing

import (
	"example/web-service-gin/album"

	"github.com/gin-gonic/gin"
)

func CreateRoutes(router *gin.Engine) {
	createAlbumRoutes(router)
}

func createAlbumRoutes(r *gin.Engine) {
	r.GET("/albums", album.GetAlbums)
	r.GET("/albums/:id", album.GetAlbumByID)
	r.POST("/albums", album.PostAlbums)
}
