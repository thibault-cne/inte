package parrainageservices

import "gorm.io/gorm"

type Parrainage struct {
	gorm.Model
	GodFatherId int  `json:"godFatherId"`
	FirstWish   int  `json:"firstWish"`
	SecondWish  int  `json:"secondWish"`
	ThirdWish   int  `json:"thirdWish"`
	FourthWish  int  `json:"fourthWish"`
	FifthWish   int  `json:"fifthWish"`
	IsGranted   bool `json:"isGranted"`
}

type Adoption struct {
	gorm.Model
	GodFatherId int `json:"godFatherId"`
	StepSonId   int `json:"stepSonId"`
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
