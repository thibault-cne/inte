package parrainageservices

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
