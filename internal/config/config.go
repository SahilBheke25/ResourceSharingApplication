package config

import (
	"fmt"
	"os"
	"strconv"
)

func GetDbConfig() string {

	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, name)

	return psqlInfo
}

func GetJwtSecret() string {
	return os.Getenv("JWT_SECRET")
}
