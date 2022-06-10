package claimsservices

import (
	"backend/models"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func create_token(claims *models.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	LoadEnv()
	JWT_SECRET := os.Getenv("JWT_SECRET")

	return token.SignedString([]byte(JWT_SECRET))
}

func DecodeToken(token string) (*models.Claims, error) {
	LoadEnv()
	JWT_SECRET := os.Getenv("JWT_SECRET")

	tkn, err := jwt.ParseWithClaims(token, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, err
	}

	claims := tkn.Claims.(*models.Claims)
	return claims, nil
}

func Refresh_token(refresh_token string) (string, error) {
	claims, err := DecodeToken(refresh_token)

	if err != nil {
		return "", err
	}

	return create_token(claims)
}
