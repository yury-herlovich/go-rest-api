package albums

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		"INSERT INTO albums (title, artist, year) values ($1, $2, $3) RETURNING id, title, artist, year;",
		newAlbum.Title,
		newAlbum.Artist,
		newAlbum.Year,
	).Scan(
		&newAlbum.ID, &newAlbum.Title, &newAlbum.Artist, &newAlbum.Year,
	)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "An error occured")
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func (c *AlbumsController) GetAlbum(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, errors.ErrorResponse{ErrorMessage: "wrong id"})
		return
	}

	var album Album
	err = c.Database.QueryRow(
		"SELECT id, title, artist, year FROM albums WHERE id = $1;", id,
	).Scan(
		&album.ID, &album.Title, &album.Artist, &album.Year,
	)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			ctx.IndentedJSON(http.StatusNotFound, errors.ErrorResponse{ErrorMessage: "album not found"})
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, "An error occured")
		}

		return
	}

	ctx.IndentedJSON(http.StatusOK, album)
}

func (c *AlbumsController) DeleteAlbum(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, errors.ErrorResponse{ErrorMessage: "wrong id"})
		return
	}

	res, err := c.Database.Exec("DELETE FROM albums WHERE id = $1;", id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "An error occured")
		return
	}

	count, err := res.RowsAffected()

	if count == 0 {
		ctx.IndentedJSON(http.StatusNotFound, errors.ErrorResponse{ErrorMessage: "album not found"})
		return
	}

	ctx.IndentedJSON(http.StatusNoContent, nil)

}
