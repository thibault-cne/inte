package planningservices

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (planning *Planning) ModifyPlanningDates(beginingDate string, endDate string) error {
	if beginingDate != "" {
		planning.Spawn_time = beginingDate
	}

	if endDate != "" {
		planning.End_time = endDate
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	db.Save(planning)
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

	// Save the planning in the database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	db.Save(planning)
	return nil
}
