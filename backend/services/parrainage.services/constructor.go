package parrainageservices

import "backend/models"

func createParrainage(godFatherId string) *models.Parrainage {
	return &models.Parrainage{
		GodFatherId: godFatherId,
		IsGranted:   false,
	}
}

func createPendingParrainage(userName string) *models.PendingParrainage {
	return &models.PendingParrainage{
		GodFatherName:   userName,
		UserWishesNames: make([]string, 0),
	}
}
