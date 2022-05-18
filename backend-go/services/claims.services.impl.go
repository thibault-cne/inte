package services

import (
	"backend-go/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

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

	return token.SignedString([]byte("secret"))
}

func decode_token(token string) (*models.Claims, error) {
	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, err
	}

	return claims, nil
}

func refresh_token(refresh_token string) (string, error) {
	claims, err := decode_token(refresh_token)

	if err != nil {
		return "", err
	}

	return create_token(claims)
}
