package v1

// @title           Payments API V1
// @version         1
// @description     This is the Payments API V1.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"

	_ "github.com/scottd018/payments-api/api/v1/docs"
	"github.com/scottd018/payments-api/services"
)

const (
	apiPath = "/api/v1"
)

func pathWithID(endpoint string) string {
	return fmt.Sprintf("%s/:id", endpoint)
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

	v1 := router.Group(apiPath)

	// frequency
	v1.GET(frequenciesPath, api.ListFrequencies)
	v1.GET(pathWithID(frequenciesPath), api.ReadFrequency)

	// category
	v1.POST(categoriesPath, api.CreateCategory)
	v1.GET(categoriesPath, api.ListCategories)
	v1.GET(pathWithID(categoriesPath), api.ReadCategory)
	v1.PUT(pathWithID(categoriesPath), api.UpdateCategory)
	v1.DELETE(pathWithID(categoriesPath), api.DeleteCategory)

	// account type
	v1.POST(accountTypesPath, api.CreateAccountType)
	v1.GET(accountTypesPath, api.ListAccountTypes)
	v1.GET(pathWithID(accountTypesPath), api.ReadAccountType)
	v1.PUT(pathWithID(accountTypesPath), api.UpdateAccountType)
	v1.DELETE(pathWithID(accountTypesPath), api.DeleteAccountType)

	// source account
	v1.POST(sourceAccountsPath, api.CreateSourceAccount)
	v1.GET(sourceAccountsPath, api.ListSourceAccounts)
	v1.GET(pathWithID(sourceAccountsPath), api.ReadSourceAccount)
	v1.PUT(pathWithID(sourceAccountsPath), api.UpdateSourceAccount)
	v1.DELETE(pathWithID(sourceAccountsPath), api.DeleteSourceAccount)

	// destination account
	v1.POST(destinationAccountsPath, api.CreateDestinationAccount)
	v1.GET(destinationAccountsPath, api.ListDestinationAccounts)
	v1.GET(pathWithID(destinationAccountsPath), api.ReadDestinationAccount)
	v1.PUT(pathWithID(destinationAccountsPath), api.UpdateDestinationAccount)
	v1.DELETE(pathWithID(destinationAccountsPath), api.DeleteDestinationAccount)

	// payment
	v1.POST(paymentsPath, api.CreatePayment)
	v1.GET(paymentsPath, api.ListPayments)
	v1.GET(pathWithID(paymentsPath), api.ReadPayment)
	v1.PUT(pathWithID(paymentsPath), api.UpdatePayment)
	v1.DELETE(pathWithID(paymentsPath), api.DeletePayment)

	// documentation
	v1.GET("docs", func(c *gin.Context) { c.Redirect(http.StatusFound, "docs/index.html") })
	v1.GET("docs/*any", swagger.WrapHandler(files.Handler))

	v1.GET("/dashboard", func(c *gin.Context) { c.Redirect(http.StatusFound, "docs/index.html") })

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
