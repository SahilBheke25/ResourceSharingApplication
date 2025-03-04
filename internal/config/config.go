package config

import (
	"os"
)

// func LoadEnv() {
// 	//github.com/joho/godotenv
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("No .env file found")
// 	}
// }

func GetEnv(key string) string {
	return os.Getenv(key)
}
