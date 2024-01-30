package routes

import (
	"github.com/gabrielmoura/estudo-api-go/internal/dto"
	"net/http"

	"github.com/gabrielmoura/estudo-api-go/infra/db"
	"github.com/gabrielmoura/estudo-api-go/internal/entity"
	"github.com/gin-gonic/gin"
)

// getAllProduct godoc
//
//	@Summary		Get all Products
//	@Description	Get all Products
//	@Tags			product
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entity.Product
//	@Router			/product [get]
func getAllProduct(c *gin.Context) {
	products, err := db.GetAllProduct(db.Con)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// getOneProduct godoc
//
//	@Summary		Get one Product
//	@Description	Get one Product
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Product ID"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	entity.Product
//	@Router			/product/{id} [get]
func getOneProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	product, err := db.GetOneProduct(db.Con, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error(), "id": id})
	}
	c.JSON(http.StatusOK, product)
}

// updateProduct godoc
//
//	@Summary		Update a Product
//	@Description	Update a Product
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Product ID"
//	@Param			product	body	object	true	"Product"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	entity.Product
//	@Router			/product/{id} [put]
func updateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	product, err := db.GetOneProduct(db.Con, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := db.UpdateProduct(db.Con, product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// postProduct godoc
// @Summary Create a new Product
// @Description Create a new Product
// @Tags product
// @Accept  json
// @Produce  json
// @Param product body dto.ProductRequest true "Product"
// @Security ApiKeyAuth
// @Success 200 {object} entity.Product
// @Failure 400 {object} dto.ErrorResponse
// @Router /product [post]
func postProduct(c *gin.Context) {
	reqProd := dto.ProductRequest{}
	if err := c.ShouldBindJSON(&reqProd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	product, _ := entity.NewProduct(reqProd.Name, reqProd.Price)
	_, err := db.InsertProduct(db.Con, product)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Erro ao inserir produto", Stack: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

// deleteProduct godoc
//
//	@Summary		Delete a Product
//	@Description	Delete a Product
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"Product ID"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	string
//	@Router			/product/{id} [delete]
func deleteProduct(c *gin.Context) {
	p := entity.Product{
		ID: c.Params.ByName("id"),
	}
	err := db.DeleteProduct(db.Con, &p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
