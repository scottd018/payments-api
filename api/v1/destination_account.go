package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/scottd018/payments-api/models"
	"github.com/scottd018/payments-api/services"
)

const (
	destinationAccountsPath = "payments/destination_accounts"
)

func (api *API) ListDestinationAccounts(c *gin.Context) {
	destinationAccounts, err := services.GetDestinationAccountService().List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch destinationAccounts"})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": destinationAccounts,
	})
}

func (api *API) CreateDestinationAccount(c *gin.Context) {
	var destinationAccount models.DestinationAccount

	if err := c.ShouldBindJSON(&destinationAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := services.GetDestinationAccountService().Create(&destinationAccount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add destinationAccount"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": destinationAccount.ID,
	})
}

func (api *API) ReadDestinationAccount(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	destinationAccount, err := services.GetDestinationAccountService().Read(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get destinationAccount"})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": destinationAccount,
	})
}

func (api *API) UpdateDestinationAccount(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	var destinationAccount models.DestinationAccount

	if err := c.ShouldBindJSON(&destinationAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := services.GetDestinationAccountService().Update(id, destinationAccount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update destinationAccount"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (api *API) DeleteDestinationAccount(c *gin.Context) {
	GenericDelete(services.GetDestinationAccountService(), c)
}
