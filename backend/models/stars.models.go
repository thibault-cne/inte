package models

import "gorm.io/gorm"

type Stars struct {
	gorm.Model
	ID                        int    `json:"id"`
	Giver_id                  int    `json:"giver_id"`
	Receiver_id               int    `json:"receiver_id"`
	Type                      int    `json:"type"`
	Message                   string `json:"message"`
	Moderation_pending_status int    `json:"moderation_pending_status"`
	Moderation_status         bool   `json:"moderation_status"`
}
