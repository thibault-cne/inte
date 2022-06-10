package claimsservices

import (
	"backend/models"
	"strings"
)

func RetrieveUserClaims(reqToken string) (*models.Claims, error) {
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	// Validate token
	claims, err := DecodeToken(reqToken)

	return claims, err
}
