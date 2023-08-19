package services

import (
	"fmt"
	"sync"

	"github.com/scottd018/payments-api/dbs"
	"github.com/scottd018/payments-api/models"
	"gorm.io/gorm"
)

type DestinationAccount struct {
	Base
}

var destinationAccount *DestinationAccount
var initDestinationAccount sync.Once

func GetDestinationAccountService() *DestinationAccount {
	initDestinationAccount.Do(initDestinationAccountService)

	return destinationAccount
}

func initDestinationAccountService() {
	db := dbs.GetDB()

	destinationAccount = NewDestinationAccountService(db)
}

func NewDestinationAccountService(db *gorm.DB) *DestinationAccount {
	return &DestinationAccount{
		Base: Base{
			db: db,
		},
	}
}

func (svc *DestinationAccount) List() ([]models.DestinationAccount, error) {
	var destinationAccounts []models.DestinationAccount

	result := svc.db.Find(&destinationAccounts)
	if result.Error != nil {
		return nil, result.Error
	}

	return destinationAccounts, nil
}

func (svc *DestinationAccount) Create(model *models.DestinationAccount) error {
	return GenericCreate(svc.db, model)
}

func (svc *DestinationAccount) Read(id int) (models.DestinationAccount, error) {
	current := models.DestinationAccount{}

	if err := GenericRead(svc.db, id, &current); err != nil {
		return current, fmt.Errorf("unable to read destination account : %w", err)
	}

	return current, nil
}

func (svc *DestinationAccount) Update(id int, updated models.DestinationAccount) error {
	current := models.DestinationAccount{}
	current.ID = id

	result := svc.db.Model(&current).Updates(updated)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (svc *DestinationAccount) Delete(id int) error {
	return GenericDelete(svc.db, id, &models.DestinationAccount{})
}
