package parrainageservices

import (
	"backend/db"
	"backend/models"
	usersservices "backend/services/users.services"
)

// Function to retrieve all unadopted users to display them during the parrainge process
func RetrieveUnadoptedUsers() []string {
	var adoptions []*models.Adoption
	var users []*models.User

	unadoptedUsers := make([]string, 0)

	usersResult := db.DB.Where("current_year = ?", 1).Find(&users)

	if usersResult.Error != nil {
		panic(usersResult.Error)
	}

	for _, user := range users {
		adoptions := db.DB.Where("step_son_id = ?", user.ID).Find(&adoptions)

		if adoptions.Error != nil {
			panic(adoptions.Error)
		}

		if adoptions.RowsAffected == 0 {
			unadoptedUsers = append(unadoptedUsers, user.Name)
		}
	}

	return unadoptedUsers
}

func RetrievePendingWishes() []*models.PendingParrainage {
	var pendingWishes []*models.PendingParrainage
	var parrainages []*models.Parrainage

	result := db.DB.Find(&parrainages)

	if result.Error != nil {
		panic(result.Error)
	}

	for _, parr := range parrainages {
		user, _ := usersservices.GetUser(parr.GodFatherId)

		pendingWish := createPendingParrainage(user.Name)

		if parr.FirstWish != "" {
			user, _ = usersservices.GetUser(parr.FirstWish)
			pendingWish.UserWishesNames = append(pendingWish.UserWishesNames, user.Name)
		}

		if parr.SecondWish != "" {
			user, _ = usersservices.GetUser(parr.SecondWish)
			pendingWish.UserWishesNames = append(pendingWish.UserWishesNames, user.Name)
		}

		if parr.ThirdWish != "" {
			user, _ = usersservices.GetUser(parr.ThirdWish)
			pendingWish.UserWishesNames = append(pendingWish.UserWishesNames, user.Name)
		}

		if parr.FourthWish != "" {
			user, _ = usersservices.GetUser(parr.FourthWish)
			pendingWish.UserWishesNames = append(pendingWish.UserWishesNames, user.Name)
		}

		if parr.FifthWish != "" {
			user, _ = usersservices.GetUser(parr.FifthWish)
			pendingWish.UserWishesNames = append(pendingWish.UserWishesNames, user.Name)
		}

		pendingWishes = append(pendingWishes, pendingWish)
	}

	return pendingWishes
}

func GrantWishByNames(godFatherName string, stepSonName string) {
	godFather := usersservices.GetUserByName(godFatherName)
	stepSon := usersservices.GetUserByName(stepSonName)

	parr := RetrieveCurrentParrainage(godFather.ID)

	parr.GrantWhish(stepSon.ID)
}
