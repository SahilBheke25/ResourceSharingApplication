package middleware

import (
	"fmt"
	"time"

	"github.com/SahilBheke25/ResourceSharingApplication/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type Auth interface {
	VerifyToken(tokenString string) (int, error)
	CreateToken(userID int) (string, error)
}

type auth struct {
}

func NewAuthService() Auth {
	return auth{}
}

func (auth auth) VerifyToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}

	userIDFloat, ok := claims["userId"].(float64) // JWT stores numbers as float64
	if !ok {
		return 0, fmt.Errorf("userId missing in token")
	}

	return int(userIDFloat), nil
}

// var secretKey = []byte("secret-key")
var secretKey = []byte(config.GetJwtSecret())

func (auth auth) CreateToken(userID int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userID,
			"exp":    time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
