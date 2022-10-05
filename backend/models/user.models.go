package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
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
}
