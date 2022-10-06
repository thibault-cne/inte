package models

import (
	"backend/db"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"gorm.io/gorm"
)

type Planning struct {
	gorm.Model
	ID         int    `json:"id"`
	Picture    string `json:"picture"`
	Spawn_time string `json:"spawn_time"`
	End_time   string `json:"end_time"`
}

func (planning *Planning) ModifyPlanningDates(beginingDate string, endDate string) error {
	if beginingDate != "" {
		planning.Spawn_time = beginingDate
	}

	if endDate != "" {
		planning.End_time = endDate
	}

	db.DB.Save(planning)
	return nil
}

func (planning *Planning) ModifyPlanningPicture(file *multipart.FileHeader) error {
	// Get the file extension
	fileExtension := filepath.Ext(file.Filename)

	// Get the file content
	fileContent, err := file.Open()

	if err != nil {
		return err
	}

	// Create a new file
	pictureName := "planning_pictures_" + strings.Replace(planning.Spawn_time, "/", "-", -1) + "-" + strings.Replace(planning.End_time, "/", "-", -1) + fileExtension
	newFile, err := os.Create("static/images/planning_pictures/" + pictureName)

	if err != nil {
		return err
	}

	// Copy the file content to the new file
	_, err = io.Copy(newFile, fileContent)

	if err != nil {
		return err
	}

	// Close the file
	err = newFile.Close()

	if err != nil {
		return err
	}

	db.DB.Save(planning)
	return nil
}
