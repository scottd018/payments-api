package services

import (
	"fmt"

	"gorm.io/gorm"
)

type Service interface {
	Delete(int) error
}

type Base struct {
	db *gorm.DB
}

// GenericCreate is a generic function to create a model
// when no special logic is required.
func GenericCreate(db *gorm.DB, model interface{}) error {
	result := db.Create(model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GenericRead is a generic function to read a model
// when no special logic is required.
func GenericRead(db *gorm.DB, id int, model interface{}) error {
	result := db.Find(model, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("unable to find id [%d]", id)
	}

	return nil
}

// GenericDelete is a generic function to delete a model
// when no special logic is required.
func GenericDelete(db *gorm.DB, id int, model interface{}) error {
	result := db.Delete(model, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
