package main

import (
	"backend/models"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNewUser(t *testing.T) {
	user := models.NewUser("test@test.com", "test")

	if user.Name != "test" {
		t.Error("Expected name to be 'test' but got ", user.Name)
	}

	if user.Email != "test@test.com" {
		t.Error("Expected email to be 'test@test.com' but got ", user.Email)
	}

	if user.CurrentYear != 1 {
		t.Error("Expected current year to be 1 but got ", user.CurrentYear)
	}

	if user.PromotionYear != 2024 {
		t.Error("Expected promotion year to be 2024 but got ", user.PromotionYear)
	}

	if user.Points != 0 {
		t.Error("Expected points to be 0 but got ", user.Points)
	}

	if user.UserType != "user" {
		t.Error("Expected user type to be 'user' but got ", user.UserType)
	}

}

func TestAddUser(t *testing.T) {
	user := models.NewUser("test@test.com", "test")

	db, err := gorm.Open(sqlite.Open("test_database.db"), &gorm.Config{})

	// Add user to the test database
	if err := db.Create(&user).Error; err != nil {
		t.Error("Error while adding user to the database: ", err)
	}

	user_id := user.ID

	if err != nil {
		t.Error("Expected no error but got ", err)
	}

	if err := db.Where("id = ?", user_id).First(&user).Error; err != nil {
		t.Error("Expected no error but got ", err)
	}

	if err != nil {
		t.Error("Expected no error but got ", err)
	}

	if user.ID != user_id {
		t.Error("Expected user id to be ", user_id, " but got ", user.ID)
	}

	if user.Name != "test" {
		t.Error("Expected name to be 'test' but got ", user.Name)
	}

	if user.Email != "test@test.com" {
		t.Error("Expected email to be 'test@test.com' but got ", user.Email)
	}
}
