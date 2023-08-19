package services

import (
	"fmt"
	"sync"

	"github.com/scottd018/payments-api/dbs"
	"github.com/scottd018/payments-api/models"
	"gorm.io/gorm"
)

type Payment struct {
	Base
}

var payment *Payment
var initPayment sync.Once

func GetPaymentService() *Payment {
	initPayment.Do(initPaymentService)

	return payment
}

func initPaymentService() {
	db := dbs.GetDB()

	payment = NewPaymentService(db)
}

func NewPaymentService(db *gorm.DB) *Payment {
	return &Payment{
		Base: Base{
			db: db,
		},
	}
}

func (svc *Payment) List() ([]models.Payment, error) {
	var payments []models.Payment

	result := svc.db.Find(&payments)
	if result.Error != nil {
		return nil, result.Error
	}

	return payments, nil
}

func (svc *Payment) Create(model *models.Payment) error {
	return GenericCreate(svc.db, model)
}

func (svc *Payment) Read(id int) (models.Payment, error) {
	current := models.Payment{}

	if err := GenericRead(svc.db, id, &current); err != nil {
		return current, fmt.Errorf("unable to read payment : %w", err)
	}

	return current, nil
}

func (svc *Payment) Update(id int, updated models.Payment) error {
	current := models.Payment{}
	current.ID = id

	result := svc.db.Model(&current).Updates(updated)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (svc *Payment) Delete(id int) error {
	return GenericDelete(svc.db, id, &models.Payment{})
}
