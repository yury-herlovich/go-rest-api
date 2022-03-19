package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     int    `json:"id"`
	Title  string `json:"title" binding:"required"`
	Artist string `json:"artist" binding:"required"`
	Year   int    `json:"year" binding:"required"`
}

var albums = []album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Year: 1958},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Year: 1962},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Year: 1955},
}

func main() {
	r := gin.Default()
	r.GET("/albums", getAlbums)
	r.POST("/albums", addAlbum)

	r.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func addAlbum(c *gin.Context) {
	c.String(http.StatusOK, "add new album")
}
