package apiresponseservices

import (
	"backend/models"
	stars_services "backend/services/stars.services"
	users_services "backend/services/users.services"
)

func NewProfileDataResponse(user *models.User) *models.ProfileDataResponse {
	stars, err := stars_services.GetStars(user.ID)

	if err != nil {
		panic(err)
	}

	return &models.ProfileDataResponse{
		Username:    user.Name,
		Points:      user.Points,
		Color:       user.Color,
		Users_stars: stars,
	}
}

func NewStarsResponse(stars []models.Stars) []models.StarsResponse {
	stars_response := make([]models.StarsResponse, len(stars))

	for i, star := range stars {
		giver, _ := users_services.GetUser(star.Giver_id)
		receiver, _ := users_services.GetUser(star.Receiver_id)

		stars_response[i] = models.StarsResponse{
			Created_at:        star.CreatedAt.Unix(),
			Id:                star.ID,
			Giver_name:        giver.Name,
			Receiver_name:     receiver.Name,
			Type:              star.Type,
			Message:           star.Message,
			Moderation_status: star.Moderation_status,
		}
	}

	return stars_response
}

func NewAllUsersResponse(users []models.User) []models.AllUsersResponse {
	users_response := make([]models.AllUsersResponse, len(users))

	for i, user := range users {
		users_response[i] = models.AllUsersResponse{
			Name: user.Name,
		}
	}

	return users_response
}

func NewAllUsersWithPointsResponse(users []models.User) []models.AllUserWithPointsResponse {
	users_response := make([]models.AllUserWithPointsResponse, len(users))

	for i, user := range users {
		users_response[i] = models.AllUserWithPointsResponse{
			Name:   user.Name,
			Points: user.Points,
		}
	}

	return users_response
}
