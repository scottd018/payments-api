package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/scottd018/payments-api/services"
)

const (
	frequenciesPath = "frequencies"
)

// ListFrequencies godoc
// @Summary      List Frequencies
// @Description  List all frequencies
// @Tags         frequencies
// @Produce      json
// @Success      200  {object}  models.Frequency
// @Failure      500  {string}  string "error"
// @Router       /frequencies [get]
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

// ReadFrequency godoc
// @Summary      Get Frequency
// @Description  Get an existing frequency
// @Tags         frequencies
// @Produce      json
// @Param        id   path      int     true   "ID"
// @Success      200  {object}  models.Frequency
// @Failure      500  {string}  string "error"
// @Router       /frequencies/{id} [get]
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
