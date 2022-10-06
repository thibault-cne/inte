package parrainageservices

import (
	"backend/db"
	"backend/models"
	usersservices "backend/services/users.services"
)

func EndParrainageRound() {
	var parrainages []*models.Parrainage

	loopState := true

	for loopState {
		loopState = false

		result := db.DB.Find(&parrainages)

		if result.Error != nil {
			panic(result.Error)
		}

		for _, parr := range parrainages {
			if parr.IsGranted {
				continue
			}

			for i := 1; i < 6; i++ {
				parr.CalculateWishes(i)

				if parr.IsGranted {
					loopState = true
					break
				}
			}
		}
	}

}

func CleanAllParrainage() {
	var parrainages []*models.Parrainage

	result := db.DB.Find(&parrainages)

	if result.Error != nil {
		panic(result.Error)
	}

	for _, parr := range parrainages {
		parr.CleanParrainageRound()
	}
}

func EndParrainageProcess() {
	var adoptions []*models.Adoption
	var users []*models.User

	result := db.DB.Find(&adoptions)

	if result.Error != nil {
		panic(result.Error)
	}

	for _, adoption := range adoptions {
		user, err := usersservices.GetUser(adoption.StepSonId)

		if err != nil {
			continue
		}

		user.God_father_id = adoption.GodFatherId
		users = append(users, user)
	}
	db.DB.Save(&users)

	// Drop table to flush them
	db.DB.Migrator().DropTable(&models.Parrainage{})
	db.DB.Migrator().DropTable(&models.Adoption{})

	db.DB.AutoMigrate(&models.Parrainage{})
	db.DB.AutoMigrate(&models.Adoption{})
}
