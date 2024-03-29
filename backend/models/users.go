package models

import (
	"backend/db"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                   string `json:"id" gorm:"primaryKey"`
	FirstName                 string `json:"first_name"`
	LastName			string 	`json:"last_name"`
	Email                string `json:"email"`
	CurrentYear         int    `json:"current_year"`
	PromotionYear       int    `json:"promotion_year"`
	Points               int    `json:"points"`
	GodFatherId        string `json:"god_father_id"`
	FacebookId          string `json:"facebook_id"`
	GoogleId            string `json:"google_id"`
	InstagramId         string `json:"instagram_id"`
	SnapchatId          string `json:"snapchat_id"`
	Hometown             string `json:"hometown"`
	Studies              string `json:"studies"`
	PersonalDescription string `json:"personal_description"`
	ProfilePicture      string `json:"profile_picture"`
	LastLogin           string `json:"last_login"`
	UserType            string `json:"user_type"`
	Color                string `json:"color"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index"`
	Stars []*Stars `json:"stars" gorm:"-"`
}

func (user *User) Save() error {
	fmt.Printf("%#v", user)
	return db.DB.Save(user).Error
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

	message := fmt.Sprintf("%s %s vous à donné(e) %d points", giver.FirstName, giver.LastName, points)
	n := NewNotification(receiver_id, "points", message)
	n.Create()

	return nil
}

func ModifyProfilePicture(user_id string, picture_extension string) error {
	user, err := GetUser(user_id)

	if err != nil {
		return err
	}

	user.ProfilePicture = "profile_picture_" + user_id + picture_extension

	db.DB.Save(&user)

	return nil
}

func GetProfilePicturePath(user_id string) (string, error) {
	user, err := GetUser(user_id)

	if err != nil {
		return "", err
	}

	return user.ProfilePicture, nil
}

func CheckAdmin(user_id int) (bool, error) {
	var user *User

	if err := db.DB.Where("id = ?", user_id).First(&user).Error; err != nil {
		return false, err
	}

	if user.UserType == "admin" {
		return true, nil
	} else {
		return false, nil
	}
}

func CheckUserByID(id string) bool {
	var user *User

	result := db.DB.Where("id = ?", id).Find(&user)

	return result.RowsAffected == 1
}

func NewUser(email string, firstName string, lastName string) *User {
	t := time.Now()

	year := 0

	if t.Month() < 9 {
		year = t.Year() - 1 + 3
	} else {
		year = t.Year() + 3
	}

	return &User{FirstName: firstName, LastName: lastName, Email: email, CurrentYear: 1, PromotionYear: year, Points: 0, UserType: "user"}
}

func (u *User) Create() error {
	if err := db.DB.Create(u).Error; err != nil {
		return err
	}

	return nil
}

func GetUser(id string) (*User, error) {
	user := &User{
		ID: id,
	}

	err := db.DB.Find(&user).Error

	return user, err
}

func GetUserByEmail(email string) (*User, error) {
	var user User

	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByLastName(lastName string) *User {
	var user User

	if err := db.DB.Where("lastName = ?", lastName).First(&user).Error; err != nil {
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
