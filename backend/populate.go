package main

import (
	"backend/models"
)

func PopulateDefault() {
	// Add new 1 year users
	models.AddUser(&models.User{
		ID: "1",
		Email:          "user1A.one@telecomnancy.net",
		Name:           "User 1A One",
		Current_year:   1,
		Promotion_year: 2024,
		Points:         0,
		User_type:      "user",
		Personal_description: "This is a personal description !",
	})
	models.AddUser(&models.User{
		ID: "2",
		Email:          "user1A.two@telecomnancy.net",
		Name:           "User 1A Two",
		Current_year:   1,
		Promotion_year: 2024,
		Points:         0,
		User_type:      "user",
	})
	models.AddUser(&models.User{
		ID: "3",
		Email:          "user1A.three@telecomnancy.net",
		Name:           "User 1A Three",
		Current_year:   1,
		Promotion_year: 2024,
		Points:         0,
		User_type:      "user",
	})
	models.AddUser(&models.User{
		ID: "4",
		Email:          "user1A.four@telecomnancy.net",
		Name:           "User 1A Four",
		Current_year:   1,
		Promotion_year: 2024,
		Points:         0,
		User_type:      "user",
	})

	// Add new 2 years users
	models.AddUser(&models.User{
		ID: "5",
		Email:          "user2A.one@telecomnancy.net",
		Name:           "User 2A One",
		Current_year:   2,
		Promotion_year: 2023,
		Points:         0,
		User_type:      "user",
	})
	models.AddUser(&models.User{
		ID: "6",
		Email:          "user2A.two@telecomnancy.net",
		Name:           "User 2A two",
		Current_year:   2,
		Promotion_year: 2023,
		Points:         0,
		User_type:      "user",
	})

	// Add new 3 years users
	models.AddUser(&models.User{
		ID: "7",
		Email:          "user3A.one@telecomnancy.net",
		Name:           "User 3A One",
		Current_year:   3,
		Promotion_year: 2022,
		Points:         0,
		User_type:      "user",
	})
	models.AddUser(&models.User{
		ID: "8",
		Email:          "user3A.two@telecomnancy.net",
		Name:           "User 3A Two",
		Current_year:   3,
		Promotion_year: 2022,
		Points:         0,
		User_type:      "user",
	})

	// Add 3 admin users
	models.AddUser(&models.User{
		ID: "9",
		Email:          "admin2A.one@telecomnancy.net",
		Name:           "Admin 2A One",
		Current_year:   2,
		Promotion_year: 2023,
		Points:         0,
		User_type:      "admin",
	})
	models.AddUser(&models.User{
		ID: "10",
		Email:          "admin2A.two@telecomnancy.net",
		Name:           "Admin 2A Two",
		Current_year:   2,
		Promotion_year: 2023,
		Points:         0,
		User_type:      "admin",
	})
	models.AddUser(&models.User{
		ID: "11",
		Email:          "admin3A.one@telecomnancy.net",
		Name:           "Admin 3A One",
		Current_year:   3,
		Promotion_year: 2022,
		Points:         0,
		User_type:      "admin",
	})

	// Add stars to the database
	star1 := models.NewStars("5", "1", 0, "Test star 1 bronze")
	star2 := models.NewStars("5", "2", 1, "Test star 2 silver")
	star3 := models.NewStars("6", "3", 2, "Test star 3 gold")
	star4 := models.NewStars("5", "1", 0, "Test star 4 bronze")
	star5 := models.NewStars("5", "2", 1, "Test star 5 silver")
	star6 := models.NewStars("6", "3", 2, "Test star 6 gold")
	star7 := models.NewStars("6", "3", 2, "Test star 7 gold")
	star8 := models.NewStars("6", "3", 2, "Test star 8 gold")
	star9 := models.NewStars("6", "3", 2, "Test star 9 gold")
	star10 := models.NewStars("6", "3", 2, "Test star 10 gold")
	star11 := models.NewStars("7", "5", 0, "Test star 11 not validate")
	star12 := models.NewStars("7", "6", 1, "Test star 12 not validate")

	models.AddStars(star1)
	models.AddStars(star2)
	models.AddStars(star3)
	models.AddStars(star4)
	models.AddStars(star5)
	models.AddStars(star6)
	models.AddStars(star7)
	models.AddStars(star8)
	models.AddStars(star9)
	models.AddStars(star10)
	models.AddStars(star11)
	models.AddStars(star12)

	// Moderate the stars
	models.ModerateStar(1, "9")
	models.ModerateStar(1, "10")

	models.ModerateStar(2, "9")
	models.ModerateStar(2, "10")

	models.ModerateStar(3, "9")
	models.ModerateStar(3, "10")

	models.ModerateStar(4, "9")
	models.ModerateStar(4, "10")

	models.ModerateStar(5, "9")
	models.ModerateStar(5, "10")

	models.ModerateStar(6, "9")
	models.ModerateStar(6, "10")

	models.ModerateStar(7, "9")
	models.ModerateStar(7, "10")

	models.ModerateStar(8, "9")
	models.ModerateStar(8, "10")

	models.ModerateStar(9, "9")
	models.ModerateStar(9, "10")

	models.ModerateStar(10, "9")
	models.ModerateStar(10, "10")

	// Add suggestion
	models.NewSuggestions("Test suggestion 1", "This is a test one", 1).Create()
	models.NewSuggestions("Test suggestion 2", "This is a test two", 1).Create()
	models.NewSuggestions("Test suggestion 3", "This is a test three", 1).Create()

	// Add plannings
	planning1 := models.NewPlaning(
		"planning_pictures_15-07-2022-20-07-2022.jpeg",
		"15/07/2022",
		"20/07/2022")

	planning2 := models.NewPlaning(
		"planning_pictures_21-07-2022-30-07-2022.jpeg",
		"21/07/2022",
		"30/07/2022")

	planning3 := models.NewPlaning(
		"planning_pictures_31-07-2022-05-08-2022.jpeg",
		"31/07/2022",
		"05/08/2022")

	models.AddPlaning(planning1)
	models.AddPlaning(planning2)
	models.AddPlaning(planning3)

	// Add news
	models.NewNewsInte("Il fait chaud hein !!")
	models.NewNewsInte("BigBaz le trou du cul ;)")
	models.NewNewsInte("L'empereur grec")
	models.NewNewsInte("Merci Michèle pour les absences.")
	models.NewNewsInte("Next projet.")
	models.NewNewsInte("Instagrammeur cielllll")
	models.NewNewsInte("Tah le projet de fou").Create()
}
