package albums

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yury-herlovich/go-rest-api/src/errors"
)

type AlbumsController struct {
	Database *sql.DB
}

type Album struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title" binding:"required"`
	Artist string    `json:"artist" binding:"required"`
	Year   int       `json:"year" binding:"required,numeric,gte=1000,lte=2100"`
}

type UpdateAlbumDTO struct {
	Title  *string `json:"title" binding:"omitempty"`
	Artist *string `json:"artist" binding:"omitempty"`
	Year   *int    `json:"year" binding:"omitempty,numeric,gte=1000,lte=2100"`
}

// GET /albums
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

// POST /albums
func (c *AlbumsController) AddAlbum(ctx *gin.Context) {
	var album Album

	if err := ctx.ShouldBindJSON(&album); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.ParseValidationErrors(err))
		return
	}

	err := c.Database.QueryRow(
		"INSERT INTO albums (title, artist, year) values ($1, $2, $3) RETURNING id, title, artist, year;",
		album.Title,
		album.Artist,
		album.Year,
	).Scan(
		&album.ID, &album.Title, &album.Artist, &album.Year,
	)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "An error occured")
		return
	}

	ctx.IndentedJSON(http.StatusCreated, album)
}

// GET /albums/:id
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

// PATCH /albums/:id
func (c *AlbumsController) UpdateAlbum(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrorResponse{ErrorMessage: "wrong id"})
		return
	}

	var payload UpdateAlbumDTO
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.ParseValidationErrors(err))
		return
	}

	var album Album
	err = c.Database.QueryRow(
		`UPDATE albums SET
			title = COALESCE($2, title),
			artist = COALESCE($3, artist),
			year = COALESCE($4, year)
		WHERE id = $1
		RETURNING id, title, artist, year;`,
		id, payload.Title, payload.Artist, payload.Year,
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

// DELETE /albums/:id
func (c *AlbumsController) DeleteAlbum(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.ErrorResponse{ErrorMessage: "wrong id"})
		return
	}

	res, err := c.Database.Exec("DELETE FROM albums WHERE id = $1;", id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "An error occured")
		return
	}

	count, err := res.RowsAffected()

	if count == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, errors.ErrorResponse{ErrorMessage: "album not found"})
		return
	}

	ctx.IndentedJSON(http.StatusNoContent, nil)

}
