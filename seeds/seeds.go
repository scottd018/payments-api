package seeds

import (
	"encoding/json"
	"os"

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

	// dbTasks, err := taskService.GetAllTasks()
	// if err != nil {
	// 	return err
	// }
	// if len(dbTasks) > 0 {
	// 	logrus.Info("Tasks already exist, skipping all migrations")
	// 	return nil
	// }

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
		if err := categoryService.Create(&category); err != nil {
			return err
		}
	}

	for _, accountType := range data.AccountTypes {
		if err := accountTypesService.Create(&accountType); err != nil {
			return err
		}
	}

	for _, frequency := range data.Frequencies {
		if err := frequencyService.Create(&frequency); err != nil {
			return err
		}
	}

	return nil
}
