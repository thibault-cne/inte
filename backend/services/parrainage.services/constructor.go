package parrainageservices

func createParrainage(godFatherId int) *Parrainage {
	return &Parrainage{
		GodFatherId: godFatherId,
		FirstWish:   0,
		SecondWish:  0,
		ThirdWish:   0,
		FourthWish:  0,
		FifthWish:   0,
		IsGranted:   false,
	}
}

func createPendingParrainage(userName string) *PendingParrainage {
	return &PendingParrainage{
		GodFatherName:   userName,
		UserWishesNames: make([]string, 0),
	}
}
