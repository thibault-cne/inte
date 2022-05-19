package models

import "time"

type GoogleApiResponse struct {
	Iss            string `json:"iss"`
	Nbf            string `json:"nbf"`
	Aud            string `json:"aud"`
	Sub            string `json:"sub"`
	Hd             string `json:"hd"`
	Email          string `json:"email"`
	Email_verified bool   `json:"email_verified"`
	Azp            string `json:"azp"`
	Name           string `json:"name"`
	Picture        string `json:"picture"`
	Given_name     string `json:"given_name"`
	Family_name    string `json:"family_name"`
	Iat            string `json:"iat"`
	Exp            string `json:"exp"`
	Jti            string `json:"jti"`
	Alg            string `json:"alg"`
	Kid            string `json:"kid"`
	Typ            string `json:"typ"`
}

type LoginApiResponse struct {
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}

type ProfileDataResponse struct {
	Username    string  `json:"username"`
	Points      int     `json:"points"`
	Users_stars []Stars `json:"users_stars"`
}

type StarsResponse struct {
	Created_at        time.Time `json:"created_at"`
	Id                int       `json:"id"`
	Giver_name        string    `json:"giver_name"`
	Receiver_name     string    `json:"receiver_name"`
	Type              int       `json:"type"`
	Message           string    `json:"message"`
	Moderation_status bool      `json:"moderation_status"`
}

type AllUsersResponse struct {
	Name string `json:"name"`
}

type AllUserWithPointsResponse struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
}
