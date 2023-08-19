package services

import (
	"fmt"
	"sync"

	"github.com/scottd018/payments-api/dbs"
	"github.com/scottd018/payments-api/models"
	"gorm.io/gorm"
)

type Category struct {
	Base
}

var category *Category
var initCategory sync.Once

func GetCategoryService() *Category {
	initCategory.Do(initCategoryService)

	return category
}

func initCategoryService() {
	db := dbs.GetDB()

	category = NewCategoryService(db)
}

func NewCategoryService(db *gorm.DB) *Category {
	return &Category{
		Base: Base{
			db: db,
		},
	}
}

func (svc *Category) List() ([]models.Category, error) {
	var categories []models.Category

	result := svc.db.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}

func (svc *Category) Create(model *models.Category) error {
	return GenericCreate(svc.db, model)
}

func (svc *Category) Read(id int) (models.Category, error) {
	current := models.Category{}

	if err := GenericRead(svc.db, id, &current); err != nil {
		return current, fmt.Errorf("unable to read category : %w", err)
	}

	return current, nil
}

func (svc *Category) Update(id int, updated models.Category) error {
	current := models.Category{}
	current.ID = id

	result := svc.db.Model(&current).Updates(updated)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (svc *Category) Delete(id int) error {
	return GenericDelete(svc.db, id, &models.Category{})
}
