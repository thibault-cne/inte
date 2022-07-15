package db

import (
	planningservices "backend/services/planning.services"
	stars_services "backend/services/stars.services"
	suggestion_services "backend/services/suggestion.services"
)

func populateTestDatabase() {
	// Add 3 stars to user with id 1
	star1 := stars_services.NewStars(1, 1, 1, "Test star 1")
	star2 := stars_services.NewStars(1, 1, 2, "Test star 2")
	star3 := stars_services.NewStars(1, 1, 0, "Test star 3")
	star4 := stars_services.NewStars(1, 1, 0, "Test star 4")
	star5 := stars_services.NewStars(1, 1, 1, "Test star 5")
	star6 := stars_services.NewStars(1, 1, 2, "Test star 6")

	stars_services.AddStars(star1)
	stars_services.AddStars(star2)
	stars_services.AddStars(star3)
	stars_services.AddStars(star4)
	stars_services.AddStars(star5)
	stars_services.AddStars(star6)

	// We moderate the stars to validate them
	stars_services.ModerateStar(1, 1)
	stars_services.ModerateStar(2, 1)
	stars_services.ModerateStar(3, 1)

	stars_services.ModerateStar(1, 2)
	stars_services.ModerateStar(2, 2)
	stars_services.ModerateStar(3, 2)

	// Add suggestion
	suggestion1 := suggestion_services.NewSuggestions("Test suggestion 1", "This is a test one", 1)
	suggestion2 := suggestion_services.NewSuggestions("Test suggestion 2", "This is a test two", 1)
	suggestion3 := suggestion_services.NewSuggestions("Test suggestion 3", "This is a test three", 1)

	suggestion_services.AddSuggestions(suggestion1)
	suggestion_services.AddSuggestions(suggestion2)
	suggestion_services.AddSuggestions(suggestion3)

	// Add plannings
	planning1 := planningservices.NewPlaning(
		"planning_pictures_15-07-2022-20-07-2022.jpeg",
		"15/07/2022",
		"20/07/2022")

	planning2 := planningservices.NewPlaning(
		"planning_pictures_21-07-2022-30-07-2022.jpeg",
		"21/07/2022",
		"30/07/2022")

	planning3 := planningservices.NewPlaning(
		"planning_pictures_31-07-2022-05-08-2022.jpeg",
		"31/07/2022",
		"05/08/2022")

	planningservices.AddPlaning(planning1)
	planningservices.AddPlaning(planning2)
	planningservices.AddPlaning(planning3)
}
