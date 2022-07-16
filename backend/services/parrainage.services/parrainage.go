package parrainageservices

import (
	"backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RetrieveCurrentParrainage(userId int) *Parrainage {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	var parrainage Parrainage

	result := db.Where("god_father_id = ?", userId).Find(&parrainage)

	if result.RowsAffected == 0 {
		newParrainage := createParrainage(userId)
		newParrainage.AddParrainage()

		return newParrainage
	}

	db.Where("godFatherId = ?", userId).First(parrainage)

	return &parrainage
}

func (parrainage *Parrainage) AddParrainage() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.Create(parrainage)
}

func (parrainage *Parrainage) AddWhish(whishUserName string, wishOrder int) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	var user models.User

	db.Where("name = ?", whishUserName).First(&user)

	switch wishOrder {
	case 1:
		parrainage.FirstWish = user.ID
	case 2:
		parrainage.SecondWish = user.ID
	case 3:
		parrainage.ThirdWish = user.ID
	case 4:
		parrainage.FourthWish = user.ID
	case 5:
		parrainage.FifthWish = user.ID
	}

	db.Save(parrainage)
}

func (parrainage *Parrainage) grantWhish(wishUserId int) {
	if wishUserId == 0 || parrainage.IsGranted {
		return
	}

	parrainage.IsGranted = true

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	newAdoption := Adoption{GodFatherId: parrainage.GodFatherId, StepSonId: wishUserId}

	db.Save(&newAdoption)
	db.Save(&parrainage)
}

func (parrainage *Parrainage) cleanParrainageRound() {
	parrainage.FirstWish = 0
	parrainage.SecondWish = 0
	parrainage.ThirdWish = 0
	parrainage.FourthWish = 0
	parrainage.FifthWish = 0
	parrainage.IsGranted = false

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.Save(parrainage)
}
