package db

import (
	"backend/models"
	planningservices "backend/services/planning.services"
	starsservices "backend/services/stars.services"
	suggestionservices "backend/services/suggestion.services"
	usersservices "backend/services/users.services"
)

func populateTestDatabase() {
	// Add new 1 year users
	usersservices.AddUser(&models.User{
		Email:          "user1A.one@telecomnancy.net",
		Name:           "User 1A One",
		Current_year:   1,
		Promotion_year: 2024,
		Points:         0,
		User_type:      "user",
	})
	usersservices.AddUser(&models.User{
		Email:          "user1A.two@telecomnancy.net",
		Name:           "User 1A Two",
		Current_year:   1,
		Promotion_year: 2024,
		Points:         0,
		User_type:      "user",
	})
	usersservices.AddUser(&models.User{
		Email:          "user1A.three@telecomnancy.net",
		Name:           "User 1A Three",
		Current_year:   1,
		Promotion_year: 2024,
		Points:         0,
		User_type:      "user",
	})
	usersservices.AddUser(&models.User{
		Email:          "user1A.four@telecomnancy.net",
		Name:           "User 1A Four",
		Current_year:   1,
		Promotion_year: 2024,
		Points:         0,
		User_type:      "user",
	})

	// Add new 2 years users
	usersservices.AddUser(&models.User{
		Email:          "user2A.one@telecomnancy.net",
		Name:           "User 2A One",
		Current_year:   2,
		Promotion_year: 2023,
		Points:         0,
		User_type:      "user",
	})
	usersservices.AddUser(&models.User{
		Email:          "user2A.two@telecomnancy.net",
		Name:           "User 2A two",
		Current_year:   2,
		Promotion_year: 2023,
		Points:         0,
		User_type:      "user",
	})

	// Add new 3 years users
	usersservices.AddUser(&models.User{
		Email:          "user3A.one@telecomnancy.net",
		Name:           "User 3A One",
		Current_year:   3,
		Promotion_year: 2022,
		Points:         0,
		User_type:      "user",
	})
	usersservices.AddUser(&models.User{
		Email:          "user3A.two@telecomnancy.net",
		Name:           "User 3A Two",
		Current_year:   3,
		Promotion_year: 2022,
		Points:         0,
		User_type:      "user",
	})

	// Add 3 admin users
	usersservices.AddUser(&models.User{
		Email:          "admin2A.one@telecomnancy.net",
		Name:           "Admin 2A One",
		Current_year:   2,
		Promotion_year: 2023,
		Points:         0,
		User_type:      "admin",
	})
	usersservices.AddUser(&models.User{
		Email:          "admin2A.two@telecomnancy.net",
		Name:           "Admin 2A Two",
		Current_year:   2,
		Promotion_year: 2023,
		Points:         0,
		User_type:      "admin",
	})
	usersservices.AddUser(&models.User{
		Email:          "admin3A.one@telecomnancy.net",
		Name:           "Admin 3A One",
		Current_year:   3,
		Promotion_year: 2022,
		Points:         0,
		User_type:      "admin",
	})

	// Add stars to the database
	star1 := starsservices.NewStars(5, 1, 0, "Test star 1 bronze")
	star2 := starsservices.NewStars(5, 2, 1, "Test star 2 silver")
	star3 := starsservices.NewStars(6, 3, 2, "Test star 3 gold")
	star4 := starsservices.NewStars(5, 1, 0, "Test star 4 bronze")
	star5 := starsservices.NewStars(5, 2, 1, "Test star 5 silver")
	star6 := starsservices.NewStars(6, 3, 2, "Test star 6 gold")
	star7 := starsservices.NewStars(6, 3, 2, "Test star 7 gold")
	star8 := starsservices.NewStars(6, 3, 2, "Test star 8 gold")
	star9 := starsservices.NewStars(6, 3, 2, "Test star 9 gold")
	star10 := starsservices.NewStars(6, 3, 2, "Test star 10 gold")
	star11 := starsservices.NewStars(7, 5, 0, "Test star 11 not validate")
	star12 := starsservices.NewStars(7, 6, 1, "Test star 12 not validate")

	starsservices.AddStars(star1)
	starsservices.AddStars(star2)
	starsservices.AddStars(star3)
	starsservices.AddStars(star4)
	starsservices.AddStars(star5)
	starsservices.AddStars(star6)
	starsservices.AddStars(star7)
	starsservices.AddStars(star8)
	starsservices.AddStars(star9)
	starsservices.AddStars(star10)
	starsservices.AddStars(star11)
	starsservices.AddStars(star12)

	// Moderate the stars
	starsservices.ModerateStar(1, 9)
	starsservices.ModerateStar(1, 10)

	starsservices.ModerateStar(2, 9)
	starsservices.ModerateStar(2, 10)

	starsservices.ModerateStar(3, 9)
	starsservices.ModerateStar(3, 10)

	starsservices.ModerateStar(4, 9)
	starsservices.ModerateStar(4, 10)

	starsservices.ModerateStar(5, 9)
	starsservices.ModerateStar(5, 10)

	starsservices.ModerateStar(6, 9)
	starsservices.ModerateStar(6, 10)

	starsservices.ModerateStar(7, 9)
	starsservices.ModerateStar(7, 10)

	starsservices.ModerateStar(8, 9)
	starsservices.ModerateStar(8, 10)

	starsservices.ModerateStar(9, 9)
	starsservices.ModerateStar(9, 10)

	starsservices.ModerateStar(10, 9)
	starsservices.ModerateStar(10, 10)

	// Add suggestion
	suggestion1 := suggestionservices.NewSuggestions("Test suggestion 1", "This is a test one", 1)
	suggestion2 := suggestionservices.NewSuggestions("Test suggestion 2", "This is a test two", 1)
	suggestion3 := suggestionservices.NewSuggestions("Test suggestion 3", "This is a test three", 1)

	suggestionservices.AddSuggestions(suggestion1)
	suggestionservices.AddSuggestions(suggestion2)
	suggestionservices.AddSuggestions(suggestion3)

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
