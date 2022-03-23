package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type album struct {
	ID     int    `json:"id"`
	Title  string `json:"title" binding:"required"`
	Artist string `json:"artist" binding:"required"`
	Year   int    `json:"year" binding:"required,gte=1800,lte=2100"`
}

var albums = []album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Year: 1958},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Year: 1962},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Year: 1955},
}

type ErrorResponse struct {
	ErrorMessage string     `json:"errorMessage"`
	Errors       []ErrorMsg `json:"errors"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
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
	var newAlbum album

	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ParseValidationErrors(err))
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

func ParseValidationErrors(err error) ErrorResponse {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
		}

		return ErrorResponse{ErrorMessage: "validation error", Errors: out}
	}

	return ErrorResponse{ErrorMessage: "unknown error"}
}
