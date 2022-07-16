package parrainageservices

import (
	"backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Function to retrieve all unadopted users to display them during the parrainge process
func RetrieveUnadoptedUsers() []string {
	var adoptions []Adoption
	var users []models.User

	unadoptedUsers := make([]string, 0)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	usersResult := db.Where("current_year = ?", 1).Find(&users)

	if usersResult.Error != nil {
		panic(usersResult.Error)
	}

	for _, user := range users {
		adoptions := db.Where("step_son_id = ?", user.ID).Find(&adoptions)

		if adoptions.Error != nil {
			panic(adoptions.Error)
		}

		if adoptions.RowsAffected == 0 {
			unadoptedUsers = append(unadoptedUsers, user.Name)
		}
	}

	return unadoptedUsers
}
