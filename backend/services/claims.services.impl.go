package services

import (
	"backend/models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func NewAccessClaims(user_id int) *models.Claims {
	standard_claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
	}

	return &models.Claims{User_id: user_id, StandardClaims: standard_claims}
}

func NewRefreshClaims(user_id int) *models.Claims {
	standard_claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	return &models.Claims{User_id: user_id, StandardClaims: standard_claims}
}

func NewLoginApiResponse(user_id int) *models.LoginApiResponse {
	access_token, err := create_token(NewAccessClaims(user_id))

	if err != nil {
		return nil
	}

	refresh_token, err := create_token(NewRefreshClaims(user_id))

	if err != nil {
		return nil
	}

	return &models.LoginApiResponse{Access_token: access_token, Refresh_token: refresh_token}
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
