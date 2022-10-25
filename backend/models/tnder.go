package models

import "gorm.io/gorm"

type Tnder struct {
	gorm.Model
	MatchedId  int `json:"matched_id"`
	MatchingId int `json:"matching_id"`
}
