package albums

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yury-herlovich/go-rest-api/src/errors"
)

var table = "albums"

type AlbumsController struct {
	Database *sql.DB
}

type Album struct {
	ID     string `json:"id"`
	Title  string `json:"title" binding:"required"`
	Artist string `json:"artist" binding:"required"`
	Year   int    `json:"year" binding:"required,gte=1800,lte=2100"`
}

func (c *AlbumsController) GetAlbums(ctx *gin.Context) {
	albums := make([]Album, 0)

	rows, err := c.Database.Query("SELECT id, title, artist, year FROM albums")
	defer rows.Close()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "An error occured")
	}

	for rows.Next() {
		var res Album

		rows.Scan(&res.ID, &res.Title, &res.Artist, &res.Year)
		albums = append(albums, res)
	}

	ctx.IndentedJSON(http.StatusOK, albums)
}

func (c *AlbumsController) AddAlbum(ctx *gin.Context) {
	var newAlbum Album

	if err := ctx.ShouldBindJSON(&newAlbum); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.ParseValidationErrors(err))
		return
	}

	err := c.Database.QueryRow(
		"INSERT INTO albums (title, artist, year) values ($1, $2, $3) RETURNING id, title, artist, year",
		newAlbum.Title,
		newAlbum.Artist,
		newAlbum.Year,
	).Scan(
		&newAlbum.ID,
		&newAlbum.Title,
		&newAlbum.Artist,
		&newAlbum.Year,
	)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "An error occured")
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

// func (c *AlbumsController) GetAlbum(ctx *gin.Context) {
// 	id, err := getIdFromParams(ctx)

// 	if err != nil {
// 		ctx.IndentedJSON(http.StatusBadRequest, errors.ErrorResponse{ErrorMessage: "wrong id"})
// 		return
// 	}

// 	for _, a := range albums {
// 		if a.ID == id {
// 			ctx.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}

// 	ctx.IndentedJSON(http.StatusNotFound, errors.ErrorResponse{ErrorMessage: "album not found"})
// }

// func (c *AlbumsController) DeleteAlbum(ctx *gin.Context) {
// 	id, err := getIdFromParams(ctx)

// 	if err != nil {
// 		ctx.IndentedJSON(http.StatusBadRequest, errors.ErrorResponse{ErrorMessage: "wrong id"})
// 		return
// 	}

// 	for ind, a := range albums {
// 		if a.ID != id {
// 			continue
// 		}

// 		albums = append(albums[:ind], albums[ind+1:]...)
// 		ctx.IndentedJSON(http.StatusOK, a)
// 		return
// 	}

// 	ctx.IndentedJSON(http.StatusNotFound, errors.ErrorResponse{ErrorMessage: "album not found"})

// }

// func getIdFromParams(ctx *gin.Context) (int, error) {
// 	return strconv.Atoi(ctx.Param("id"))
// }
