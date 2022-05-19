package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	User_id int `json:"user_id"`
	jwt.StandardClaims
}
