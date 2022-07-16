package parrainageservices

import (
	"backend/models"
	usersservices "backend/services/users.services"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Function to retrieve all unadopted users to display them during the parrainge process
func RetrieveUnadoptedUsers() []string {
	var adoptions []Adoption
	var users []models.User

	unadoptedUsers := make([]string, 0)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	usersResult := db.Where("current_year = ?", 1).Find(&users)

	if usersResult.Error != nil {
		panic(usersResult.Error)
	}

	for _, user := range users {
		adoptions := db.Where("step_son_id = ?", user.ID).Find(&adoptions)

		if adoptions.Error != nil {
			panic(adoptions.Error)
		}

		if adoptions.RowsAffected == 0 {
			unadoptedUsers = append(unadoptedUsers, user.Name)
		}
	}

	return unadoptedUsers
}

func RetrievePendingWishes() []PendingParrainage {
	pendingWishes := make([]PendingParrainage, 0)

	var parrainages []Parrainage

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	result := db.Find(&parrainages)

	if result.Error != nil {
		panic(result.Error)
	}

	for _, parr := range parrainages {
		user, _ := usersservices.GetUser(parr.GodFatherId)

		pendingWish := createPendingParrainage(user.Name)

		if parr.FirstWish != 0 {
			user, _ = usersservices.GetUser(parr.FirstWish)
			pendingWish.UserWishesNames = append(pendingWish.UserWishesNames, user.Name)
		}

		if parr.SecondWish != 0 {
			user, _ = usersservices.GetUser(parr.SecondWish)
			pendingWish.UserWishesNames = append(pendingWish.UserWishesNames, user.Name)
		}

		if parr.ThirdWish != 0 {
			user, _ = usersservices.GetUser(parr.ThirdWish)
			pendingWish.UserWishesNames = append(pendingWish.UserWishesNames, user.Name)
		}

		if parr.FourthWish != 0 {
			user, _ = usersservices.GetUser(parr.FourthWish)
			pendingWish.UserWishesNames = append(pendingWish.UserWishesNames, user.Name)
		}

		if parr.FifthWish != 0 {
			user, _ = usersservices.GetUser(parr.FifthWish)
			pendingWish.UserWishesNames = append(pendingWish.UserWishesNames, user.Name)
		}

		pendingWishes = append(pendingWishes, *pendingWish)
	}

	return pendingWishes
}
