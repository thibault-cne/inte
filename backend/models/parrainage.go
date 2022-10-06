package models

import (
	"backend/db"

	"gorm.io/gorm"
)

type Parrainage struct {
	gorm.Model
	GodFatherId string `json:"godFatherId"`
	FirstWish   string `json:"firstWish"`
	SecondWish  string `json:"secondWish"`
	ThirdWish   string `json:"thirdWish"`
	FourthWish  string `json:"fourthWish"`
	FifthWish   string `json:"fifthWish"`
	IsGranted   bool   `json:"isGranted"`
}

type Adoption struct {
	gorm.Model
	GodFatherId string `json:"godFatherId"`
	StepSonId   string `json:"stepSonId"`
}

type ParrainageProcess struct {
	IsProcessOpen bool `json:"status"`
	IsRoundOpen   bool `json:"isRoundOpen"`
	CurrentRound  int  `json:"currentRound"`
}

type PendingParrainage struct {
	GodFatherName   string   `json:"godFatherName"`
	UserWishesNames []string `json:"userWishesNames"`
}

func (parrainage *Parrainage) AddParrainage() {
	db.DB.Create(parrainage)
}

func (parrainage *Parrainage) AddWhish(whishUserName string, wishOrder int) {
	var user *User

	db.DB.Where("name = ?", whishUserName).First(&user)

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

	db.DB.Save(parrainage)
}

func (parrainage *Parrainage) GrantWhish(wishUserId string) {
	if wishUserId == "" || parrainage.IsGranted {
		return
	}

	parrainage.IsGranted = true

	newAdoption := Adoption{GodFatherId: parrainage.GodFatherId, StepSonId: wishUserId}

	db.DB.Save(&newAdoption)
	db.DB.Save(&parrainage)
}

func (parrainage *Parrainage) CleanParrainageRound() {
	parrainage.FirstWish = ""
	parrainage.SecondWish = ""
	parrainage.ThirdWish = ""
	parrainage.FourthWish = ""
	parrainage.FifthWish = ""
	parrainage.IsGranted = false

	db.DB.Save(parrainage)
}

func (parr *Parrainage) CalculateWishes(roundNumber int) {
	var matchedWishes []Parrainage

	switch roundNumber {
	case 1:
		matchedResult := db.DB.Where("is_granted = ?", false).
			Where("first_wish = ?", parr.FirstWish).
			Or("second_wish = ?", parr.FirstWish).
			Or("third_wish = ?", parr.FirstWish).
			Or("fourth_wish = ?", parr.FirstWish).
			Or("fifth_wish = ?", parr.FirstWish).
			Find(&matchedWishes)

		if matchedResult.RowsAffected == 1 {
			parr.GrantWhish(parr.FirstWish)
		}
	case 2:
		matchedResult := db.DB.Where("is_granted = ?", false).
			Where("first_wish = ?", parr.SecondWish).
			Or("second_wish = ?", parr.SecondWish).
			Or("third_wish = ?", parr.SecondWish).
			Or("fourth_wish = ?", parr.SecondWish).
			Or("fifth_wish = ?", parr.SecondWish).
			Find(&matchedWishes)

		if matchedResult.RowsAffected == 1 {
			parr.GrantWhish(parr.SecondWish)
		}
	case 3:
		matchedResult := db.DB.Where("is_granted = ?", false).
			Where("first_wish = ?", parr.ThirdWish).
			Or("second_wish = ?", parr.ThirdWish).
			Or("third_wish = ?", parr.ThirdWish).
			Or("fourth_wish = ?", parr.ThirdWish).
			Or("fifth_wish = ?", parr.ThirdWish).
			Find(&matchedWishes)

		if matchedResult.RowsAffected == 1 {
			parr.GrantWhish(parr.ThirdWish)
		}
	case 4:
		matchedResult := db.DB.Where("is_granted = ?", false).
			Where("first_wish = ?", parr.FourthWish).
			Or("second_wish = ?", parr.FourthWish).
			Or("third_wish = ?", parr.FourthWish).
			Or("fourth_wish = ?", parr.FourthWish).
			Or("fifth_wish = ?", parr.FourthWish).
			Find(&matchedWishes)

		if matchedResult.RowsAffected == 1 {
			parr.GrantWhish(parr.FourthWish)
		}
	case 5:
		matchedResult := db.DB.Where("is_granted = ?", false).
			Where("first_wish = ?", parr.FifthWish).
			Or("second_wish = ?", parr.FifthWish).
			Or("third_wish = ?", parr.FifthWish).
			Or("fourth_wish = ?", parr.FifthWish).
			Or("fifth_wish = ?", parr.FifthWish).
			Find(&matchedWishes)

		if matchedResult.RowsAffected == 1 {
			parr.GrantWhish(parr.FifthWish)
		}
	}
}
