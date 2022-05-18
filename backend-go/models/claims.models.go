package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	User_id int `json:"user_id"`
	jwt.StandardClaims
}
