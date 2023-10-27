package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/scottd018/payments-api/models"
	"github.com/scottd018/payments-api/services"
)

const (
	accountTypesPath = "account_types"
)

// ListAccountTypes godoc
// @Summary      List Account Types
// @Description  List all account types
// @Tags         account_types
// @Produce      json
// @Success      200  {object}  models.AccountType
// @Failure      500  {string}  string "error"
// @Router       /account_types [get]
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

// CreateAccount Type godoc
// @Summary      Create Account Type
// @Description  Create a new account type
// @Tags         account_types
// @Produce      json
// @Success      200  {object}  models.AccountType
// @Failure      500  {string}  string "error"
// @Router       /account_types [post]
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

// ReadAccount Type godoc
// @Summary      Get Account Type
// @Description  Get an existing account type
// @Tags         account_types
// @Produce      json
// @Param        id   path      int     true   "ID"
// @Success      200  {object}  models.AccountType
// @Failure      500  {string}  string "error"
// @Router       /account_types/{id} [get]
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

// UpdateAccount Type godoc
// @Summary      Update Account Type
// @Description  Update an existing account type
// @Tags         account_types
// @Produce      json
// @Param        id   path      int     true   "ID"
// @Success      200  {object}  models.AccountType
// @Failure      500  {string}  string "error"
// @Router       /account_types/{id} [put]
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

// DeleteAccount Type godoc
// @Summary      Delete Account Type
// @Description  Delete an account type
// @Tags         account_types
// @Produce      json
// @Param        id   path      int     true   "ID"
// @Success      200  {object}  models.AccountType
// @Failure      500  {string}  string "error"
// @Router       /account_types/{id} [delete]
func (api *API) DeleteAccountType(c *gin.Context) {
	GenericDelete(services.GetAccountTypeService(), c)
}
