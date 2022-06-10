package db

import (
	stars_services "backend/services/stars.services"
	suggestion_services "backend/services/suggestion.services"
	users_services "backend/services/users.services"
)

func populate_test_database() {
	// Add 3 stars to user with id 1
	star1 := stars_services.NewStars(1, 1, 1, "Test star 1")
	star2 := stars_services.NewStars(1, 1, 2, "Test star 2")
	star3 := stars_services.NewStars(1, 1, 0, "Test star 3")

	stars_services.AddStars(star1)
	stars_services.AddStars(star2)
	stars_services.AddStars(star3)

	// We moderate the stars to validate them
	stars_services.ModerateStar(1, 1)
	stars_services.ModerateStar(2, 1)
	stars_services.ModerateStar(3, 1)

	stars_services.ModerateStar(1, 2)
	stars_services.ModerateStar(2, 2)
	stars_services.ModerateStar(3, 2)

	// Add points to user with id 1
	users_services.AddPoints(1, 1, 3)
	users_services.AddPoints(1, 1, 5)
	users_services.AddPoints(1, 1, 7)

	// Add suggestion
	suggestion1 := suggestion_services.NewSuggestions("Test suggestion 1", "This is a test one", 1)
	suggestion2 := suggestion_services.NewSuggestions("Test suggestion 2", "This is a test two", 1)
	suggestion3 := suggestion_services.NewSuggestions("Test suggestion 3", "This is a test three", 1)

	suggestion_services.AddSuggestions(suggestion1)
	suggestion_services.AddSuggestions(suggestion2)
	suggestion_services.AddSuggestions(suggestion3)
}
