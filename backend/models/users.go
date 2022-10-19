package models

import (
	"backend/db"
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Current_year         int    `json:"current_year"`
	Promotion_year       int    `json:"promotion_year"`
	Points               int    `json:"points"`
	God_father_id        string `json:"god_father_id"`
	Facebook_id          string `json:"facebook_id"`
	Google_id            string `json:"google_id"`
	Instagram_id         string `json:"instagram_id"`
	Snapchat_id          string `json:"snapchat_id"`
	Hometown             string `json:"hometown"`
	Studies              string `json:"studies"`
	Personal_description string `json:"personal_description"`
	Profile_picture      string `json:"profile_picture"`
	Last_login           string `json:"last_login"`
	User_type            string `json:"user_type"`
	Color                string `json:"color"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}

func RetrieveAllLogs() []*Notifications {
	var logs []*Notifications

	result := db.DB.Find(&logs)

	if result.Error != nil {
		panic(result.Error)
	}

	return logs
}

func AddPoints(giver_id string, receiver_id string, points int) error {
	var user *User

	if err := db.DB.Where("id = ?", receiver_id).First(&user).Error; err != nil {
		return err
	}

	user.Points += points

	db.DB.Save(&user)

	// Also add a notification to the receiver
	giver, err := GetUser(giver_id)

	if err != nil {
		return err
	}

	message := fmt.Sprintf("%s vous à donné(e) %d points", giver.Name, points)
	AddNewNotification(NewNotification(receiver_id, "points", message))

	return nil
}

func ModifyProfileData(temp_user *User) error {
	// Modify the user in the database with the new data
	// We don't want to modify the ID and the fields that are equals to "" (empty) in the temp_user

	user, err := GetUser(temp_user.ID)

	if err != nil {
		return err
	}

	if temp_user.Personal_description != "" {
		db.DB.Model(&user).Update("personal_description", temp_user.Personal_description)
	}

	if temp_user.Facebook_id != "" {
		db.DB.Model(&user).Update("facebook_id", temp_user.Facebook_id)
	}

	if temp_user.Snapchat_id != "" {
		db.DB.Model(&user).Update("snapchat_id", temp_user.Snapchat_id)
	}

	if temp_user.Instagram_id != "" {
		db.DB.Model(&user).Update("instagram_id", temp_user.Instagram_id)
	}

	if temp_user.Google_id != "" {
		db.DB.Model(&user).Update("google_id", temp_user.Google_id)
	}

	if temp_user.Hometown != "" {
		db.DB.Model(&user).Update("hometown", temp_user.Hometown)
	}

	if temp_user.Studies != "" {
		db.DB.Model(&user).Update("studies", temp_user.Studies)
	}

	return nil
}

func ModifyProfilePicture(user_id string, picture_extension string) error {
	user, err := GetUser(user_id)

	if err != nil {
		return err
	}

	user.Profile_picture = "profile_picture_" + user_id + picture_extension

	db.DB.Save(&user)

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
	var user *User

	if err := db.DB.Where("id = ?", user_id).First(&user).Error; err != nil {
		return false, err
	}

	if user.User_type == "admin" {
		return true, nil
	} else {
		return false, nil
	}
}

func CheckUserByName(userName string) bool {
	var user *User

	result := db.DB.Where("name = ?", userName).Find(&user)

	return result.RowsAffected == 1
}

func RetrieveAllUsersData() []map[string]interface{} {
	var (
		usersData []map[string]interface{}
		users     []*User
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

func NewUser(email string, name string) *User {
	t := time.Now()

	year := 0

	if t.Month() < 9 {
		year = t.Year() - 1 + 3
	} else {
		year = t.Year() + 3
	}

	return &User{Name: name, Email: email, Current_year: 1, Promotion_year: year, Points: 0, User_type: "user"}
}

func AddUser(u *User) (string, error) {
	db.DB.Create(u)

	return u.ID, nil
}

func (u *User) Create() (string, error) {
	db.DB.Create(u)

	return u.ID, nil
}

func GetUser(id string) (*User, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var user User

	db.First(&user, id)

	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User

	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByName(name string) *User {
	var user User

	if err := db.DB.Where("name = ?", name).First(&user).Error; err != nil {
		panic(err)
	}

	return &user
}

// Get all users from the database
func GetAllUsers() ([]*User, error) {
	var users []*User

	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
