package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/yury-herlovich/go-rest-api/src/albums"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "gorestdb"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("DB successfully connected!")
	}

	r := gin.Default()
	r.GET("/albums", albums.GetAlbums)
	r.POST("/albums", albums.AddAlbum)
	r.GET("/albums/:id", albums.GetAlbum)
	r.DELETE("/albums/:id", albums.DeleteAlbum)

	r.Run(":8080")
}
