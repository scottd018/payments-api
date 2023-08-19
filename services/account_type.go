package services

import (
	"fmt"
	"sync"

	"github.com/scottd018/payments-api/dbs"
	"github.com/scottd018/payments-api/models"
	"gorm.io/gorm"
)

type AccountType struct {
	Base
}

var accountType *AccountType
var initAccountType sync.Once

func GetAccountTypeService() *AccountType {
	initAccountType.Do(initAccountTypeService)

	return accountType
}

func initAccountTypeService() {
	db := dbs.GetDB()

	accountType = NewAccountTypeService(db)
}

func NewAccountTypeService(db *gorm.DB) *AccountType {
	return &AccountType{
		Base: Base{
			db: db,
		},
	}
}

func (svc *AccountType) List() ([]models.AccountType, error) {
	var accountTypes []models.AccountType

	result := svc.db.Find(&accountTypes)
	if result.Error != nil {
		return nil, result.Error
	}

	return accountTypes, nil
}

func (svc *AccountType) Create(model *models.AccountType) error {
	return GenericCreate(svc.db, model)
}

func (svc *AccountType) Read(id int) (models.AccountType, error) {
	current := models.AccountType{}

	if err := GenericRead(svc.db, id, &current); err != nil {
		return current, fmt.Errorf("unable to read account type : %w", err)
	}

	return current, nil
}

func (svc *AccountType) Update(id int, updated models.AccountType) error {
	current := models.AccountType{}
	current.ID = id

	result := svc.db.Model(&current).Updates(updated)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (svc *AccountType) Delete(id int) error {
	return GenericDelete(svc.db, id, &models.AccountType{})
}
