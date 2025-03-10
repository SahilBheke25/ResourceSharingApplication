package repository

import (
	"database/sql"
)

func InitializeDatabase(psqlInfo string) *sql.DB {

	// var port string = "5432"
	var DB *sql.DB
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, dbUser, password, dbname)

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
