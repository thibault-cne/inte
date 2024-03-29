package models

import "fmt"

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
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ProfileDataResponse struct {
	User    *User   `json:"user"`
	UsersStars  []*Stars `json:"users_stars"`
}

type StarsResponse struct {
	CreatedAt        int64  `json:"created_at"`
	Id                uint    `json:"id"`
	GiverName        string `json:"giver_name"`
	ReceiverName     string `json:"receiver_name"`
	Type              int    `json:"type"`
	Message           string `json:"message"`
	ModerationStatus bool   `json:"moderation_status"`
}

type AllUsersResponse struct {
	Name string `json:"name"`
}

type AllUserWithPointsResponse struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
}

func NewStarsResponse(stars []*Stars) []*StarsResponse {
	out := make([]*StarsResponse, len(stars))

	for i, star := range stars {
		giver, _ := GetUser(star.GiverId)
		receiver, _ := GetUser(star.ReceiverId)

		out[i] = &StarsResponse{
			CreatedAt:        star.CreatedAt.Unix(),
			Id:                star.ID,
			GiverName:        fmt.Sprintf("%s %s", giver.FirstName, giver.LastName),
			ReceiverName:     fmt.Sprintf("%s %s", receiver.FirstName, receiver.LastName),
			Type:              star.Type,
			Message:           star.Message,
			ModerationStatus: star.ModerationStatus,
		}
	}

	return out
}
