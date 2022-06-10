package claimsservices

import (
	"backend/models"
	"time"

	"github.com/golang-jwt/jwt"
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
