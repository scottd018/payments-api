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

// ListPayments godoc
// @Summary      List Payments
// @Description  List all payments
// @Tags         payments
// @Produce      json
// @Success      200  {object}  models.Payment
// @Failure      500  {string}  string "error"
// @Router       /payments [get]
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

// CreatePayment godoc
// @Summary      Create Payment
// @Description  Create a new payment
// @Tags         payments
// @Produce      json
// @Success      200  {object}  models.Payment
// @Failure      500  {string}  string "error"
// @Router       /payments [post]
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

// ReadPayment godoc
// @Summary      Get Payment
// @Description  Get an existing payment
// @Tags         payments
// @Produce      json
// @Param        id   path      int     true   "ID"
// @Success      200  {object}  models.Payment
// @Failure      500  {string}  string "error"
// @Router       /payments/{id} [get]
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

// UpdatePayment godoc
// @Summary      Update Payment
// @Description  Update an existing payment
// @Tags         payments
// @Produce      json
// @Param        id   path      int     true   "ID"
// @Success      200  {object}  models.Payment
// @Failure      500  {string}  string "error"
// @Router       /payments/{id} [put]
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

// DeletePayment godoc
// @Summary      Delete Payment
// @Description  Delete a payment
// @Tags         payments
// @Produce      json
// @Param        id   path      int     true   "ID"
// @Success      200  {object}  models.Payment
// @Failure      500  {string}  string "error"
// @Router       /payments/{id} [delete]
func (api *API) DeletePayment(c *gin.Context) {
	GenericDelete(services.GetPaymentService(), c)
}
