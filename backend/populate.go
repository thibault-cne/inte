package main

import (
	"backend/models"
)

func PopulateDefault() {
	// Add new 1 year users
	user1A := &models.User{
		ID: "1",
		Email:          "user1A.one@telecomnancy.net",
		Name:           "User 1A One",
		CurrentYear:   1,
		PromotionYear: 2024,
		Points:         0,
		UserType:      "user",
		PersonalDescription: "This is a personal description !",
		Studies: "Prepa MP*",
		Hometown: "Paris",
		FacebookId: "User_One",
		SnapchatId: "UserOne",
		GoogleId: "user.one",
		InstagramId: "userTahOne",
	}
	user1A.Create()
	user1A = &models.User{
		ID: "2",
		Email:          "user1A.two@telecomnancy.net",
		Name:           "User 1A Two",
		CurrentYear:   1,
		PromotionYear: 2024,
		Points:         0,
		UserType:      "user",
	}
	user1A.Create()
	user1A = &models.User{
		ID: "3",
		Email:          "user1A.three@telecomnancy.net",
		Name:           "User 1A Three",
		CurrentYear:   1,
		PromotionYear: 2024,
		Points:         0,
		UserType:      "user",
	}
	user1A.Create()
	user1A = &models.User{
		ID: "4",
		Email:          "user1A.four@telecomnancy.net",
		Name:           "User 1A Four",
		CurrentYear:   1,
		PromotionYear: 2024,
		Points:         0,
		UserType:      "user",
	}
	user1A.Create()

	// Add new 2 years users
	user2A := &models.User{
		ID: "5",
		Email:          "user2A.one@telecomnancy.net",
		Name:           "User 2A One",
		CurrentYear:   2,
		PromotionYear: 2023,
		Points:         0,
		UserType:      "user",
	}
	user2A.Create()
	user2A = &models.User{
		ID: "6",
		Email:          "user2A.two@telecomnancy.net",
		Name:           "User 2A two",
		CurrentYear:   2,
		PromotionYear: 2023,
		Points:         0,
		UserType:      "user",
	}
	user2A.Create()

	// Add new 3 years users
	user3A := &models.User{
		ID: "7",
		Email:          "user3A.one@telecomnancy.net",
		Name:           "User 3A One",
		CurrentYear:   3,
		PromotionYear: 2022,
		Points:         0,
		UserType:      "user",
	}
	user3A.Create()
	user3A = &models.User{
		ID: "8",
		Email:          "user3A.two@telecomnancy.net",
		Name:           "User 3A Two",
		CurrentYear:   3,
		PromotionYear: 2022,
		Points:         0,
		UserType:      "user",
	}
	user3A.Create()

	// Add 3 admin users
	u1 := &models.User{
		ID: "9",
		Email:          "admin2A.one@telecomnancy.net",
		Name:           "Admin 2A One",
		CurrentYear:   2,
		PromotionYear: 2023,
		Points:         0,
		UserType:      "admin",
	}
	u1.Create()

	u2 := &models.User{
		ID: "10",
		Email:          "admin2A.two@telecomnancy.net",
		Name:           "Admin 2A Two",
		CurrentYear:   2,
		PromotionYear: 2023,
		Points:         0,
		UserType:      "admin",
	}
	u2.Create()

	u3 := &models.User{
		ID: "11",
		Email:          "admin3A.one@telecomnancy.net",
		Name:           "Admin 3A One",
		CurrentYear:   3,
		PromotionYear: 2022,
		Points:         0,
		UserType:      "admin",
	}
	u3.Create()

	// Create Thibz account so that he already have a description
	t := &models.User{
		ID: "111065978852750383627",
		CurrentYear: 2,
		PromotionYear: 2024,
		PersonalDescription: "Zest parti zest parti les amis",
		Hometown: "Paris",
		Studies: "La chocolaterie de charlie",
	}
	t.Create()

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

	// Stars for Thibault
	star13 := models.NewStars("5", "111065978852750383627", 0, "Test star 13 bronze")
	star14 := models.NewStars("5", "111065978852750383627", 1, "Test star 14 silver")
	star15 := models.NewStars("6", "111065978852750383627", 2, "Test star 15 gold")
	star16 := models.NewStars("5", "111065978852750383627", 0, "Test star 16 bronze")

	// Stars for Tristan
	star17 := models.NewStars("5", "112890394147922856462", 0, "Test star 17 bronze")
	star18 := models.NewStars("5", "112890394147922856462", 1, "Test star 18 silver")
	star19 := models.NewStars("6", "112890394147922856462", 2, "Test star 19 gold")
	star20 := models.NewStars("5", "112890394147922856462", 0, "Test star 20 bronze")

	star1.Create()
	star2.Create()
	star3.Create()
	star4.Create()
	star5.Create()
	star6.Create()
	star7.Create()
	star8.Create()
	star9.Create()
	star10.Create()
	star11.Create()
	star12.Create()
	star13.Create()
	star14.Create()
	star15.Create()
	star16.Create()
	star17.Create()
	star18.Create()
	star19.Create()
	star20.Create()

	// Moderate the stars
	u1.ModerateStar(star1)
	u2.ModerateStar(star1)

	u1.ModerateStar(star2)
	u2.ModerateStar(star2)

	u1.ModerateStar(star3)
	u2.ModerateStar(star3)

	u1.ModerateStar(star4)
	u2.ModerateStar(star4)

	u1.ModerateStar(star5)
	u2.ModerateStar(star5)

	u1.ModerateStar(star6)
	u2.ModerateStar(star6)

	u1.ModerateStar(star7)
	u2.ModerateStar(star7)

	u1.ModerateStar(star8)
	u2.ModerateStar(star8)

	u1.ModerateStar(star9)
	u2.ModerateStar(star9)

	u1.ModerateStar(star10)
	u2.ModerateStar(star10)

	u1.ModerateStar(star13)
	u2.ModerateStar(star13)

	u1.ModerateStar(star14)
	u2.ModerateStar(star14)

	u1.ModerateStar(star15)
	u2.ModerateStar(star15)

	u1.ModerateStar(star16)
	u2.ModerateStar(star16)

	u1.ModerateStar(star17)
	u2.ModerateStar(star17)

	u1.ModerateStar(star18)
	u2.ModerateStar(star18)

	u1.ModerateStar(star19)
	u2.ModerateStar(star19)

	u1.ModerateStar(star20)
	u2.ModerateStar(star20)

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

	planning1.Create()
	planning2.Create()
	planning3.Create()

	// Add news
	models.NewNewsInte("Il fait chaud hein !!").Create()
	models.NewNewsInte("BigBaz le trou du cul").Create()
	models.NewNewsInte("L'empereur grec").Create()
	models.NewNewsInte("Merci Michèle pour les absences.").Create()
	models.NewNewsInte("Next projet.").Create()
	models.NewNewsInte("Instagrammeur cielllll").Create()
	models.NewNewsInte("Tah le projet de fou").Create()
}
