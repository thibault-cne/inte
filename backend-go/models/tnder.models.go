package models

import "gorm.io/gorm"

type Tnder struct {
	gorm.Model
	ID          int `json:"id"`
	Matched_id  int `json:"matched_id"`
	Matching_id int `json:"matching_id"`
}
