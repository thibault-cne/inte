package services

import (
	"backend/models"
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewUser(email string, name string) *models.User {
	t := time.Now()

	year := 0

	if t.Month() < 9 {
		year = t.Year() - 1 + 3
	} else {
		year = t.Year() + 3
	}

	return &models.User{Name: name, Email: email, Current_year: 1, Promotion_year: year, Points: 0, User_type: "user"}
}

func AddUser(user *models.User) (int, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return 0, err
	}
	db.Create(user)

	return user.ID, nil
}

func GetUser(id int) (*models.User, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var user models.User

	db.First(&user, id)

	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var user models.User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserData(claims *models.Claims) (*models.User, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var user models.User

	if err := db.Where("id = ?", claims.User_id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func AddPoints(giver_id int, receiver_id int, points int) error {
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
	RegisterNewNotification(NewNotification(receiver_id, "points", message))

	return nil
}

// Get all users from the database
func GetAllUsers() ([]models.User, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var users []models.User

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
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

func ModifyProfilePicture(user_id int, picture_url string) error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	user, err := GetUser(user_id)

	if err != nil {
		return err
	}

	user.Profile_picture = picture_url

	db.Save(&user)

	return nil
}
