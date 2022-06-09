package db

import (
	"backend/services"
)

func populate_test_database() {
	// Add 3 stars to user with id 1
	star1 := services.NewStars(1, 1, 1, "Test star 1")
	star2 := services.NewStars(1, 1, 2, "Test star 2")
	star3 := services.NewStars(1, 1, 0, "Test star 3")

	services.AddStars(star1)
	services.AddStars(star2)
	services.AddStars(star3)

	// We moderate the stars to validate them
	services.ModerateStar(1, 1)
	services.ModerateStar(2, 1)
	services.ModerateStar(3, 1)

	services.ModerateStar(1, 2)
	services.ModerateStar(2, 2)
	services.ModerateStar(3, 2)

	// Add points to user with id 1
	services.AddPoints(1, 1, 3)
	services.AddPoints(1, 1, 5)
	services.AddPoints(1, 1, 7)

	// Add suggestion
	suggestion1 := services.NewSuggestions("Test suggestion 1", "This is a test one", 1)
	suggestion2 := services.NewSuggestions("Test suggestion 2", "This is a test two", 1)
	suggestion3 := services.NewSuggestions("Test suggestion 3", "This is a test three", 1)

	services.RegisterSuggestions(suggestion1)
	services.RegisterSuggestions(suggestion2)
	services.RegisterSuggestions(suggestion3)
}
