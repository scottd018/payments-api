package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/scottd018/payments-api/models"
	"github.com/scottd018/payments-api/services"
)

const (
	accountTypesPath = "payments/account_types"
)

func (api *API) ListAccountTypes(c *gin.Context) {
	accountTypes, err := services.GetAccountTypeService().List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch accountTypes"})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": accountTypes,
	})
}

func (api *API) CreateAccountType(c *gin.Context) {
	var accountType models.AccountType

	if err := c.ShouldBindJSON(&accountType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := services.GetAccountTypeService().Create(&accountType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add accountType"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": accountType.ID,
	})
}

func (api *API) ReadAccountType(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	accountType, err := services.GetAccountTypeService().Read(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get accountType"})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": accountType,
	})
}

func (api *API) UpdateAccountType(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	var accountType models.AccountType

	if err := c.ShouldBindJSON(&accountType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := services.GetAccountTypeService().Update(id, accountType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update accountType"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (api *API) DeleteAccountType(c *gin.Context) {
	GenericDelete(services.GetAccountTypeService(), c)
}
