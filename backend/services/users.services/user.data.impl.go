package usersservices

import (
	"backend/models"
	notifications_services "backend/services/notification.services"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func AddPoints(giver_id string, receiver_id string, points int) error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	var user models.User

	if err := db.Where("id = ?", receiver_id).First(&user).Error; err != nil {
		return err
	}

	user.Points += points

	db.Save(&user)

	// Also add a notification to the receiver
	giver, err := GetUser(giver_id)

	if err != nil {
		return err
	}

	message := fmt.Sprintf("%s vous à donné(e) %d points", giver.Name, points)
	notifications_services.AddNewNotification(notifications_services.NewNotification(receiver_id, "points", message))

	return nil
}

func ModifyProfileData(temp_user *models.User) error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	// Modify the user in the database with the new data
	// We don't want to modify the ID and the fields that are equals to "" (empty) in the temp_user

	user, err := GetUser(temp_user.ID)

	if err != nil {
		return err
	}

	if temp_user.Personal_description != "" {
		db.Model(&user).Update("personal_description", temp_user.Personal_description)
	}

	if temp_user.Facebook_id != "" {
		db.Model(&user).Update("facebook_id", temp_user.Facebook_id)
	}

	if temp_user.Snapchat_id != "" {
		db.Model(&user).Update("snapchat_id", temp_user.Snapchat_id)
	}

	if temp_user.Instagram_id != "" {
		db.Model(&user).Update("instagram_id", temp_user.Instagram_id)
	}

	if temp_user.Google_id != "" {
		db.Model(&user).Update("google_id", temp_user.Google_id)
	}

	if temp_user.Hometown != "" {
		db.Model(&user).Update("hometown", temp_user.Hometown)
	}

	if temp_user.Studies != "" {
		db.Model(&user).Update("studies", temp_user.Studies)
	}

	return nil
}

func ModifyProfilePicture(user_id string, picture_extension string) error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	user, err := GetUser(user_id)

	if err != nil {
		return err
	}

	user.Profile_picture = "profile_picture_" + user_id + picture_extension

	db.Save(&user)

	return nil
}

func GetProfilePicturePath(user_id string) (string, error) {
	user, err := GetUser(user_id)

	if err != nil {
		return "", err
	}

	return user.Profile_picture, nil
}

func CheckAdmin(user_id int) (bool, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return false, err
	}

	var user models.User

	if err := db.Where("id = ?", user_id).First(&user).Error; err != nil {
		return false, err
	}

	if user.User_type == "admin" {
		return true, nil
	} else {
		return false, nil
	}
}

func CheckUserByName(userName string) bool {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return false
	}

	var user models.User

	result := db.Where("name = ?", userName).Find(&user)

	return result.RowsAffected == 1
}

func RetrieveAllUsersData() []map[string]interface{} {
	var (
		usersData []map[string]interface{}
		users     []models.User
	)

	users, err := GetAllUsers()

	if err != nil {
		panic(err)
	}

	for _, user := range users {
		usersData = append(usersData, map[string]interface{}{
			"userName": user.Name,
			"year":     user.Current_year,
			"color":    user.Color,
		})
	}

	return usersData
}
