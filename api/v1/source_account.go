package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/scottd018/payments-api/models"
	"github.com/scottd018/payments-api/services"
)

const (
	sourceAccountsPath = "source_accounts"
)

// ListCategories godoc
// @Summary      List Source Accounts
// @Description  List all source accounts
// @Tags         source_accounts
// @Produce      json
// @Success      200  {object}  models.SourceAccount
// @Failure      500  {string}  string "error"
// @Router       /source_accounts [get]
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

// CreateCategory godoc
// @Summary      Create Source Account
// @Description  Create a new source account
// @Tags         source_accounts
// @Produce      json
// @Success      200  {object}  models.SourceAccount
// @Failure      500  {string}  string "error"
// @Router       /source_accounts [post]
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

// ReadCategory godoc
// @Summary      Get Source Account
// @Description  Get an existing source account
// @Tags         source_accounts
// @Produce      json
// @Param        id   path      int     true   "ID"
// @Success      200  {object}  models.SourceAccount
// @Failure      500  {string}  string "error"
// @Router       /source_accounts/{id} [get]
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

// UpdateCategory godoc
// @Summary      Update Source Account
// @Description  Update an existing source account
// @Tags         source_accounts
// @Produce      json
// @Param        id   path      int     true   "ID"
// @Success      200  {object}  models.SourceAccount
// @Failure      500  {string}  string "error"
// @Router       /source_accounts/{id} [put]
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

// DeleteCategory godoc
// @Summary      Delete Source Account
// @Description  Delete a source account
// @Tags         source_accounts
// @Produce      json
// @Param        id   path      int     true   "ID"
// @Success      200  {object}  models.SourceAccount
// @Failure      500  {string}  string "error"
// @Router       /source_accounts/{id} [delete]
func (api *API) DeleteSourceAccount(c *gin.Context) {
	GenericDelete(services.GetSourceAccountService(), c)
}
