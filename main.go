package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yury-herlovich/go-rest-api/albums"
)

func main() {
	r := gin.Default()
	r.GET("/albums", albums.GetAlbums)
	r.POST("/albums", albums.AddAlbum)

	r.Run("localhost:8080")
}
