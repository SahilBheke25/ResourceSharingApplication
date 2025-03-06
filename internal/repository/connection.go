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

// var (
// 	host     = config.GetEnv("DB_HOST")
// 	dbUser   = config.GetEnv("DB_USER")
// 	password = config.GetEnv("DB_PASSWORD")
// 	dbname   = config.GetEnv("DB_NAME")
// )

func InitializeDatabase() *sql.DB {

	var port int = 5432
	var DB *sql.DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, dbUser, password, dbname)

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
