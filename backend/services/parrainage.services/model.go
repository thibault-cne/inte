package parrainageservices

import "gorm.io/gorm"

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
