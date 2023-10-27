package seeds

import (
	"encoding/json"
	"os"

	"github.com/scottd018/payments-api/dbs"
	"github.com/scottd018/payments-api/models"
	"github.com/scottd018/payments-api/services"
)

type Seeds struct {
	Frequencies  []models.Frequency
	Categories   []models.Category
	AccountTypes []models.AccountType
}

func RunSeeds(seedFilePath string) error {

	frequencyService := services.GetFrequencyService()
	categoryService := services.GetCategoryService()
	accountTypesService := services.GetAccountTypeService()

	db := dbs.GetDB()

	seedFile, err := os.Open(seedFilePath)
	if err != nil {
		return err
	}
	defer seedFile.Close()

	var data Seeds
	if err := json.NewDecoder(seedFile).Decode(&data); err != nil {
		return err
	}

	for _, category := range data.Categories {
		findModel := &models.Category{}

		if result, err := services.GenericFindBy(db, "name", category.Name, findModel); err != nil {
			return err
		} else {
			// continue the loop if we found a model with the same name
			if result.RowsAffected > 0 {
				continue
			}
		}

		if err := categoryService.Create(&category); err != nil {
			return err
		}
	}

	for _, accountType := range data.AccountTypes {
		findModel := &models.AccountType{}

		if result, err := services.GenericFindBy(db, "name", accountType.Name, findModel); err != nil {
			return err
		} else {
			// continue the loop if we found a model with the same name
			if result.RowsAffected > 0 {
				continue
			}
		}

		if err := accountTypesService.Create(&accountType); err != nil {
			return err
		}
	}

	for _, frequency := range data.Frequencies {
		findModel := &models.Frequency{}

		if result, err := services.GenericFindBy(db, "name", frequency.Name, findModel); err != nil {
			return err
		} else {
			// continue the loop if we found a model with the same name
			if result.RowsAffected > 0 {
				continue
			}
		}

		if err := frequencyService.Create(&frequency); err != nil {
			return err
		}
	}

	return nil
}
