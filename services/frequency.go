package services

import (
	"fmt"
	"sync"

	"github.com/scottd018/payments-api/dbs"
	"github.com/scottd018/payments-api/models"
	"gorm.io/gorm"
)

type Frequency struct {
	Base
}

var frequency *Frequency
var initFrequency sync.Once

func GetFrequencyService() *Frequency {
	initFrequency.Do(initFrequencyService)

	return frequency
}

func initFrequencyService() {
	db := dbs.GetDB()

	frequency = NewFrequencyService(db)
}

func NewFrequencyService(db *gorm.DB) *Frequency {
	return &Frequency{
		Base: Base{
			db: db,
		},
	}
}

func (svc *Frequency) List() ([]models.Frequency, error) {
	var frequencies []models.Frequency

	result := svc.db.Find(&frequencies)
	if result.Error != nil {
		return nil, result.Error
	}

	return frequencies, nil
}

func (svc *Frequency) Create(model *models.Frequency) error {
	return GenericCreate(svc.db, model)
}

func (svc *Frequency) Read(id int) (models.Frequency, error) {
	current := models.Frequency{}

	if err := GenericRead(svc.db, id, &current); err != nil {
		return current, fmt.Errorf("unable to read frequency : %w", err)
	}

	return current, nil
}

func (svc *Frequency) Update(id int, updated models.Frequency) error {
	current := models.Frequency{}
	current.ID = id

	result := svc.db.Model(&current).Updates(updated)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (svc *Frequency) Delete(id int) error {
	return GenericDelete(svc.db, id, &models.Frequency{})
}
