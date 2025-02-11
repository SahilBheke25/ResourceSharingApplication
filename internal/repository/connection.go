package repository

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	dbUser   = "postgres"
	password = "aim"
	dbname   = "resourcesharing"
)

func InitializeDatabase() *sql.DB {
	var DB *sql.DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, 5432, dbUser, password, dbname)

	var err error

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	return DB
}
