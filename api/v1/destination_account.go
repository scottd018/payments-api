package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/scottd018/payments-api/models"
	"github.com/scottd018/payments-api/services"
)

const (
	destinationAccountsPath = "destination_accounts"
)

// ListDestinationAccounts godoc
// @Summary      List Categories
// @Description  List all destination accounts
// @Tags         destination_accounts
// @Produce      json
// @Success      200  {object}  models.DestinationAccount
// @Failure      500  {string}  string "error"
// @Router       /destination_accounts [get]
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

// CreateDestinationAccount godoc
// @Summary      Create Destination Account
// @Description  Create a new destination account
// @Tags         destination_accounts
// @Produce      json
// @Success      200  {object}  models.DestinationAccount
// @Failure      500  {string}  string "error"
// @Router       /destination_accounts [post]
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

// ReadDestinationAccount godoc
// @Summary      Get Destination Account
// @Description  Get an existing destination account
// @Tags         destination_accounts
// @Produce      json
// @Param        id   path      int     true   "ID"
// @Success      200  {object}  models.DestinationAccount
// @Failure      500  {string}  string "error"
// @Router       /destination_accounts/{id} [get]
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

// UpdateDestinationAccount godoc
// @Summary      Update Destination Account
// @Description  Update an existing destination account
// @Tags         destination_accounts
// @Produce      json
// @Param        id   path      int     true   "ID"
// @Success      200  {object}  models.DestinationAccount
// @Failure      500  {string}  string "error"
// @Router       /destination_accounts/{id} [put]
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

// DeleteDestinationAccount godoc
// @Summary      Delete Destination Account
// @Description  Delete a destination account
// @Tags         destination_accounts
// @Produce      json
// @Param        id   path      int     true   "ID"
// @Success      200  {object}  models.DestinationAccount
// @Failure      500  {string}  string "error"
// @Router       /destination_accounts/{id} [delete]
func (api *API) DeleteDestinationAccount(c *gin.Context) {
	GenericDelete(services.GetDestinationAccountService(), c)
}
