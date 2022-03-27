package albums

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yury-herlovich/go-rest-api/src/errors"
)

type Album struct {
	ID     int    `json:"id"`
	Title  string `json:"title" binding:"required"`
	Artist string `json:"artist" binding:"required"`
	Year   int    `json:"year" binding:"required,gte=1800,lte=2100"`
}

var albums = []Album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Year: 1958},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Year: 1962},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Year: 1955},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func AddAlbum(c *gin.Context) {
	var newAlbum Album

	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.ParseValidationErrors(err))
		return
	}

	maxId := 1
	for _, album := range albums {
		if album.ID > maxId {
			maxId = album.ID
		}
	}

	newAlbum.ID = maxId + 1

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, errors.ErrorResponse{ErrorMessage: "wrong id"})
		return
	}

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, errors.ErrorResponse{ErrorMessage: "album not found"})
}
