package services

import (
	"backend-go/models"
)

func NewProfileDataResponse(user *models.User) *models.ProfileDataResponse {
	stars, err := GetStars(user.ID)

	if err != nil {
		panic(err)
	}

	return &models.ProfileDataResponse{
		Username:    user.Name,
		Points:      user.Points,
		Users_stars: stars,
	}
}

func NewStarsResponse(stars []models.Stars) []models.StarsResponse {
	stars_response := make([]models.StarsResponse, len(stars))

	for i, star := range stars {
		giver, _ := GetUser(star.Giver_id)
		receiver, _ := GetUser(star.Receiver_id)

		stars_response[i] = models.StarsResponse{
			Giver_name:        giver.Name,
			Receiver_name:     receiver.Name,
			Type:              star.Type,
			Message:           star.Message,
			Moderation_status: star.Moderation_status,
		}
	}

	return stars_response
}
