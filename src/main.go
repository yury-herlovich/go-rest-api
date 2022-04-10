package main

import (
	db "github.com/yury-herlovich/go-rest-api/src/common"

	"github.com/gin-gonic/gin"
	"github.com/yury-herlovich/go-rest-api/src/albums"
)

func main() {
	db.Init()
	defer db.Close()

	r := gin.Default()
	r.GET("/albums", albums.GetAlbums)
	r.POST("/albums", albums.AddAlbum)
	r.GET("/albums/:id", albums.GetAlbum)
	r.DELETE("/albums/:id", albums.DeleteAlbum)

	r.Run(":8080")
}
