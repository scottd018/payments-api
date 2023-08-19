package models

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&SourceAccount{},
		&DestinationAccount{},
		&Category{},
		&Frequency{},
		&Payment{},
	)
}
