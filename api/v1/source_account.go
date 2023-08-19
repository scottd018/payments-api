package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/scottd018/payments-api/models"
	"github.com/scottd018/payments-api/services"
)

const (
	sourceAccountsPath = "payments/source_accounts"
)

func (api *API) ListSourceAccounts(c *gin.Context) {
	sourceAccounts, err := services.GetSourceAccountService().List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch sourceAccounts"})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": sourceAccounts,
	})
}

func (api *API) CreateSourceAccount(c *gin.Context) {
	var sourceAccount models.SourceAccount

	if err := c.ShouldBindJSON(&sourceAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := services.GetSourceAccountService().Create(&sourceAccount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add sourceAccount"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": sourceAccount.ID,
	})
}

func (api *API) ReadSourceAccount(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	sourceAccount, err := services.GetSourceAccountService().Read(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get sourceAccount"})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": sourceAccount,
	})
}

func (api *API) UpdateSourceAccount(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	var sourceAccount models.SourceAccount

	if err := c.ShouldBindJSON(&sourceAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := services.GetSourceAccountService().Update(id, sourceAccount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update sourceAccount"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (api *API) DeleteSourceAccount(c *gin.Context) {
	GenericDelete(services.GetSourceAccountService(), c)
}
