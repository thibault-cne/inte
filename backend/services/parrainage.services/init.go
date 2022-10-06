package parrainageservices

import "backend/models"

var Process = &models.ParrainageProcess{
	IsProcessOpen: false,
	CurrentRound:  1,
	IsRoundOpen:   true,
}
