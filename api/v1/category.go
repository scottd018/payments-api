package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/scottd018/payments-api/models"
	"github.com/scottd018/payments-api/services"
)

const (
	categoriesPath = "categories"
)

// ListCategories godoc
// @Summary      List Categories
// @Description  List all categories
// @Tags         categories
// @Produce      json
// @Success      200  {object}  models.Category
// @Failure      500  {string}  string "error"
// @Router       /categories [get]
func (api *API) ListCategories(c *gin.Context) {
	categories, err := services.GetCategoryService().List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch categories"})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}

// CreateCategory godoc
// @Summary      Create Category
// @Description  Create a new category
// @Tags         categories
// @Produce      json
// @Success      200  {object}  models.Category
// @Failure      500  {string}  string "error"
// @Router       /categories [post]
func (api *API) CreateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := services.GetCategoryService().Create(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add category"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": category.ID,
	})
}

// ReadCategory godoc
// @Summary      Get Category
// @Description  Get an existing category
// @Tags         categories
// @Produce      json
// @Param        id   path      int     true   "ID"
// @Success      200  {object}  models.Category
// @Failure      500  {string}  string "error"
// @Router       /categories/{id} [get]
func (api *API) ReadCategory(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	category, err := services.GetCategoryService().Read(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

// UpdateCategory godoc
// @Summary      Update Category
// @Description  Update an existing category
// @Tags         categories
// @Produce      json
// @Param        id   path      int     true   "ID"
// @Success      200  {object}  models.Category
// @Failure      500  {string}  string "error"
// @Router       /categories/{id} [put]
func (api *API) UpdateCategory(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := services.GetCategoryService().Update(id, category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update category"})
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

// DeleteCategory godoc
// @Summary      Delete Category
// @Description  Delete a category
// @Tags         categories
// @Produce      json
// @Param        id   path      int     true   "ID"
// @Success      200  {object}  models.Category
// @Failure      500  {string}  string "error"
// @Router       /categories/{id} [delete]
func (api *API) DeleteCategory(c *gin.Context) {
	GenericDelete(services.GetCategoryService(), c)
}
