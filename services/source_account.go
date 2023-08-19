package services

import (
	"fmt"
	"sync"

	"github.com/scottd018/payments-api/dbs"
	"github.com/scottd018/payments-api/models"
	"gorm.io/gorm"
)

type SourceAccount struct {
	Base
}

var sourceAccount *SourceAccount
var initSourceAccount sync.Once

func GetSourceAccountService() *SourceAccount {
	initSourceAccount.Do(initSourceAccountService)

	return sourceAccount
}

func initSourceAccountService() {
	db := dbs.GetDB()

	sourceAccount = NewSourceAccountService(db)
}

func NewSourceAccountService(db *gorm.DB) *SourceAccount {
	return &SourceAccount{
		Base: Base{
			db: db,
		},
	}
}

func (svc *SourceAccount) List() ([]models.SourceAccount, error) {
	var sourceAccounts []models.SourceAccount

	result := svc.db.Find(&sourceAccounts)
	if result.Error != nil {
		return nil, result.Error
	}

	return sourceAccounts, nil
}

func (svc *SourceAccount) Create(model *models.SourceAccount) error {
	return GenericCreate(svc.db, model)
}

func (svc *SourceAccount) Read(id int) (models.SourceAccount, error) {
	current := models.SourceAccount{}

	if err := GenericRead(svc.db, id, &current); err != nil {
		return current, fmt.Errorf("unable to read source account : %w", err)
	}

	return current, nil
}

func (svc *SourceAccount) Update(id int, updated models.SourceAccount) error {
	current := models.SourceAccount{}
	current.ID = id

	result := svc.db.Model(&current).Updates(updated)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (svc *SourceAccount) Delete(id int) error {
	return GenericDelete(svc.db, id, &models.SourceAccount{})
}
