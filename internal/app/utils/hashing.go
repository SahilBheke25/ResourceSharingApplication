package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
)

// hashing of the given password
func HashPassword(password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("password cannot be empty")
	}
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// compares a hashed password with a plaintext password
func CheckPasswordHash(password, hash string) bool {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return false
	}
	log.Println(hashedPassword, " ", hash)
	return hashedPassword == hash
}
