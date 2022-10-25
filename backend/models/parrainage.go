package models

import (
	"backend/db"

	"gorm.io/driver/sqlite"
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

var Process = &ParrainageProcess{
	IsProcessOpen: false,
	CurrentRound:  1,
	IsRoundOpen:   true,
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

func createParrainage(godFatherId string) *Parrainage {
	return &Parrainage{
		GodFatherId: godFatherId,
		IsGranted:   false,
	}
}

func createPendingParrainage(userName string) *PendingParrainage {
	return &PendingParrainage{
		GodFatherName:   userName,
		UserWishesNames: make([]string, 0),
	}
}

func RetrieveCurrentParrainage(userId string) *Parrainage {
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

// Function to retrieve all unadopted users to display them during the parrainge process
func RetrieveUnadoptedUsers() []string {
	var adoptions []*Adoption
	var users []*User

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

func RetrievePendingWishes() []*PendingParrainage {
	var pendingWishes []*PendingParrainage
	var parrainages []*Parrainage

	result := db.DB.Find(&parrainages)

	if result.Error != nil {
		panic(result.Error)
	}

	for _, parr := range parrainages {
		user, _ := GetUser(parr.GodFatherId)

		pendingWish := createPendingParrainage(user.Name)

		if parr.FirstWish != "" {
			user, _ = GetUser(parr.FirstWish)
			pendingWish.UserWishesNames = append(pendingWish.UserWishesNames, user.Name)
		}

		if parr.SecondWish != "" {
			user, _ = GetUser(parr.SecondWish)
			pendingWish.UserWishesNames = append(pendingWish.UserWishesNames, user.Name)
		}

		if parr.ThirdWish != "" {
			user, _ = GetUser(parr.ThirdWish)
			pendingWish.UserWishesNames = append(pendingWish.UserWishesNames, user.Name)
		}

		if parr.FourthWish != "" {
			user, _ = GetUser(parr.FourthWish)
			pendingWish.UserWishesNames = append(pendingWish.UserWishesNames, user.Name)
		}

		if parr.FifthWish != "" {
			user, _ = GetUser(parr.FifthWish)
			pendingWish.UserWishesNames = append(pendingWish.UserWishesNames, user.Name)
		}

		pendingWishes = append(pendingWishes, pendingWish)
	}

	return pendingWishes
}

func GrantWishByNames(godFatherName string, stepSonName string) {
	godFather := GetUserByName(godFatherName)
	stepSon := GetUserByName(stepSonName)

	parr := RetrieveCurrentParrainage(godFather.ID)

	parr.GrantWhish(stepSon.ID)
}

func EndParrainageRound() {
	var parrainages []*Parrainage

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
	var parrainages []*Parrainage

	result := db.DB.Find(&parrainages)

	if result.Error != nil {
		panic(result.Error)
	}

	for _, parr := range parrainages {
		parr.CleanParrainageRound()
	}
}

func EndParrainageProcess() {
	var adoptions []*Adoption
	var users []*User

	result := db.DB.Find(&adoptions)

	if result.Error != nil {
		panic(result.Error)
	}

	for _, adoption := range adoptions {
		user, err := GetUser(adoption.StepSonId)

		if err != nil {
			continue
		}

		user.GodFatherId = adoption.GodFatherId
		users = append(users, user)
	}
	db.DB.Save(&users)

	// Drop table to flush them
	db.DB.Migrator().DropTable(&Parrainage{})
	db.DB.Migrator().DropTable(&Adoption{})

	db.DB.AutoMigrate(&Parrainage{})
	db.DB.AutoMigrate(&Adoption{})
}
