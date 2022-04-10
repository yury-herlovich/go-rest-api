package db

import (
	_ "github.com/lib/pq"

	"database/sql"
	"fmt"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "gorestdb"
)

type Database struct {
	*sql.DB
}

var DB *sql.DB

func Init() *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	DB = db

	err = Check()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("DB successfully connected!")
	}

	DB.SetMaxIdleConns(10)

	return DB
}

func Close() {
	DB.Close()
}

func Check() error {
	return DB.Ping()
}
