package parrainageservices

import (
	"backend/models"
	usersservices "backend/services/users.services"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func EndParrainageRound() {
	var parrainages []Parrainage

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	loopState := true

	for loopState {
		loopState = false

		result := db.Find(&parrainages)

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

func (parr *Parrainage) CalculateWishes(roundNumber int) {
	var matchedWishes []Parrainage

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	switch roundNumber {
	case 1:
		matchedResult := db.Where("is_granted = ?", false).
			Where("first_wish = ?", parr.FirstWish).
			Or("second_wish = ?", parr.FirstWish).
			Or("third_wish = ?", parr.FirstWish).
			Or("fourth_wish = ?", parr.FirstWish).
			Or("fifth_wish = ?", parr.FirstWish).
			Find(&matchedWishes)

		if matchedResult.RowsAffected == 1 {
			parr.grantWhish(parr.FirstWish)
		}
	case 2:
		matchedResult := db.Where("is_granted = ?", false).
			Where("first_wish = ?", parr.SecondWish).
			Or("second_wish = ?", parr.SecondWish).
			Or("third_wish = ?", parr.SecondWish).
			Or("fourth_wish = ?", parr.SecondWish).
			Or("fifth_wish = ?", parr.SecondWish).
			Find(&matchedWishes)

		if matchedResult.RowsAffected == 1 {
			parr.grantWhish(parr.SecondWish)
		}
	case 3:
		matchedResult := db.Where("is_granted = ?", false).
			Where("first_wish = ?", parr.ThirdWish).
			Or("second_wish = ?", parr.ThirdWish).
			Or("third_wish = ?", parr.ThirdWish).
			Or("fourth_wish = ?", parr.ThirdWish).
			Or("fifth_wish = ?", parr.ThirdWish).
			Find(&matchedWishes)

		if matchedResult.RowsAffected == 1 {
			parr.grantWhish(parr.ThirdWish)
		}
	case 4:
		matchedResult := db.Where("is_granted = ?", false).
			Where("first_wish = ?", parr.FourthWish).
			Or("second_wish = ?", parr.FourthWish).
			Or("third_wish = ?", parr.FourthWish).
			Or("fourth_wish = ?", parr.FourthWish).
			Or("fifth_wish = ?", parr.FourthWish).
			Find(&matchedWishes)

		if matchedResult.RowsAffected == 1 {
			parr.grantWhish(parr.FourthWish)
		}
	case 5:
		matchedResult := db.Where("is_granted = ?", false).
			Where("first_wish = ?", parr.FifthWish).
			Or("second_wish = ?", parr.FifthWish).
			Or("third_wish = ?", parr.FifthWish).
			Or("fourth_wish = ?", parr.FifthWish).
			Or("fifth_wish = ?", parr.FifthWish).
			Find(&matchedWishes)

		if matchedResult.RowsAffected == 1 {
			parr.grantWhish(parr.FifthWish)
		}
	}
}

func CleanAllParrainage() {
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
		parr.cleanParrainageRound()
	}
}

func EndParrainageProcess() {
	var adoptions []Adoption
	users := make([]*models.User, 0)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	result := db.Find(&adoptions)

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
	db.Save(&users)

	// Drop table to flush them
	db.Migrator().DropTable(&Parrainage{})
	db.Migrator().DropTable(&Adoption{})

	db.AutoMigrate(&Parrainage{})
	db.AutoMigrate(&Adoption{})
}
