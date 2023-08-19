package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/scottd018/payments-api/services"
)

const (
	frequenciesPath = "payments/frequencies"
)

func (api *API) ListFrequencies(c *gin.Context) {
	frequencies, err := services.GetFrequencyService().List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch frequencies"})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": frequencies,
	})
}

func (api *API) ReadFrequency(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	frequency, err := services.GetFrequencyService().Read(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get frequency"})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": frequency,
	})
}
