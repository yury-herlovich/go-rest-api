package main

import (
	db "github.com/yury-herlovich/go-rest-api/src/common"

	"github.com/gin-gonic/gin"
	"github.com/yury-herlovich/go-rest-api/src/albums"
)

func main() {
	database := db.Init()
	defer db.Close()

	albumsCtrl := albums.AlbumsController{Database: database}

	r := gin.Default()
	r.GET("/albums", albumsCtrl.GetAlbums)
	r.POST("/albums", albumsCtrl.AddAlbum)
	r.GET("/albums/:id", albumsCtrl.GetAlbum)
	r.DELETE("/albums/:id", albumsCtrl.DeleteAlbum)

	r.Run(":8080")
}
