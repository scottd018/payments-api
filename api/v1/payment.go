package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/scottd018/payments-api/models"
	"github.com/scottd018/payments-api/services"
)

const (
	paymentsPath = "payments"
)

func (api *API) ListPayments(c *gin.Context) {
	payments, err := services.GetPaymentService().List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch payments"})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": payments,
	})
}

func (api *API) CreatePayment(c *gin.Context) {
	var payment models.Payment

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := services.GetPaymentService().Create(&payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add payment"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": payment.ID,
	})
}

func (api *API) ReadPayment(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	payment, err := services.GetPaymentService().Read(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get payment"})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": payment,
	})
}

func (api *API) UpdatePayment(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	var payment models.Payment

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := services.GetPaymentService().Update(id, payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update payment"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (api *API) DeletePayment(c *gin.Context) {
	GenericDelete(services.GetPaymentService(), c)
}
