package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/scottd018/payments-api/services"
)

const (
	apiPath = "/api/v1"
)

func path(endpoint string) string {
	return fmt.Sprintf("%s/%s", apiPath, endpoint)
}

func pathWithID(endpoint string) string {
	return fmt.Sprintf("%s/:id", path(endpoint))
}

type API struct {
	router *gin.Engine
}

func (api *API) Serve(addr string) error {
	return api.router.Run(addr)
}

func NewAPI() *API {
	router := gin.Default()

	api := &API{
		router,
	}

	// frequency
	router.GET(path(frequenciesPath), api.ListFrequencies)
	router.GET(pathWithID(frequenciesPath), api.ReadFrequency)

	// category
	router.POST(path(categoriesPath), api.CreateCategory)
	router.GET(path(categoriesPath), api.ListCategories)
	router.GET(pathWithID(categoriesPath), api.ReadCategory)
	router.PUT(pathWithID(categoriesPath), api.UpdateCategory)
	router.DELETE(pathWithID(categoriesPath), api.DeleteCategory)

	// account type
	router.POST(path(accountTypesPath), api.CreateAccountType)
	router.GET(path(accountTypesPath), api.ListAccountTypes)
	router.GET(pathWithID(accountTypesPath), api.ReadAccountType)
	router.PUT(pathWithID(accountTypesPath), api.UpdateAccountType)
	router.DELETE(pathWithID(accountTypesPath), api.DeleteAccountType)

	// source account
	router.POST(path(sourceAccountsPath), api.CreateSourceAccount)
	router.GET(path(sourceAccountsPath), api.ListSourceAccounts)
	router.GET(pathWithID(sourceAccountsPath), api.ReadSourceAccount)
	router.PUT(pathWithID(sourceAccountsPath), api.UpdateSourceAccount)
	router.DELETE(pathWithID(sourceAccountsPath), api.DeleteSourceAccount)

	// destination account
	router.POST(path(destinationAccountsPath), api.CreateDestinationAccount)
	router.GET(path(destinationAccountsPath), api.ListDestinationAccounts)
	router.GET(pathWithID(destinationAccountsPath), api.ReadDestinationAccount)
	router.PUT(pathWithID(destinationAccountsPath), api.UpdateDestinationAccount)
	router.DELETE(pathWithID(destinationAccountsPath), api.DeleteDestinationAccount)

	// payment
	router.POST(path(paymentsPath), api.CreatePayment)
	router.GET(path(paymentsPath), api.ListPayments)
	router.GET(pathWithID(paymentsPath), api.ReadPayment)
	router.PUT(pathWithID(paymentsPath), api.UpdatePayment)
	router.DELETE(pathWithID(paymentsPath), api.DeletePayment)

	return api
}

// GenericDelete is a delete function when no special logic is required.
func GenericDelete(svc services.Service, c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := svc.Delete(id); err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": fmt.Sprintf("failed to delete object with id [%d] - %s", id, err.Error())},
		)
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func getID(c *gin.Context) (int, error) {
	id := c.Param("id")

	parsed, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return -1, fmt.Errorf("invalid id, not a number")
	}

	return int(parsed), nil
}
