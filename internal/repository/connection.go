package repository

import (
	"database/sql"

	"github.com/SahilBheke25/quick-farm-backend/internal/config"
)

func InitializeDatabase() *sql.DB {

	var DB *sql.DB
	psqlInfo := config.GetDbConfig()

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
