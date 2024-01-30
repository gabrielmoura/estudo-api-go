package routes

import (
	"github.com/gabrielmoura/estudo-api-go/infra/db"
	"github.com/gabrielmoura/estudo-api-go/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

// getAllCart godoc
// @Summary Get all carts
// @Description Get all carts
// @Tags carts
// @Accept  json
// @Produce  json
// @Success 200 {object} []entity.Cart
// @Router /cart [get]
func getAllCart(c *gin.Context) {
	cart, _ := db.GetAllCart(db.Con)
	c.JSON(http.StatusOK, cart)
}

// getOneCart godoc
// @Summary Get one cart
// @Description Get one cart
// @Tags carts
// @Accept  json
// @Produce  json
// @Param id path string true "Cart ID"
// @Success 200 {object} entity.Cart
// @Router /cart/{id} [get]
func getOneCart(c *gin.Context) {
	id := c.Params.ByName("id")
	cart, _ := db.GetOneCart(db.Con, id)
	c.JSON(http.StatusOK, cart)
}

// updateCart godoc
// @Summary Update a cart
// @Description Update a cart
// @Tags carts
// @Accept  json
// @Produce  json
// @Param id path string true "Cart ID"
// @Param cart body entity.Cart true "Cart"
// @Success 200 {object} entity.Cart
// @Router /cart/{id} [put]
func updateCart(c *gin.Context) {
	newCart := &entity.Cart{
		ID: c.Params.ByName("id"),
	}
_:
	c.BindJSON(&newCart)

_:
	db.UpdateCart(db.Con, newCart)
	c.JSON(http.StatusOK, newCart)
}

// postCart godoc
// @Summary Create a new cart
// @Description Create a new cart
// @Tags carts
// @Accept  json
// @Produce  json
// @Param cart body entity.Cart true "Cart"
// @Success 200 {object} entity.Cart
// @Router /cart [post]
func postCart(c *gin.Context) {
	cart := &entity.Cart{}
_:
	c.BindJSON(&cart)

	if _, err := db.InsertCart(db.Con, cart); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, cart)
}

// deleteCart godoc
// @Summary Delete a cart
// @Description Delete a cart
// @Tags carts
// @Accept  json
// @Produce  json
// @Param id path string true "Cart ID"
// @Success 200 {object} string
// @Router /cart/{id} [delete]
func deleteCart(c *gin.Context) {
	cart := &entity.Cart{
		ID: c.Param("id"),
	}
_:
	db.DeleteCart(db.Con, cart)
	c.JSON(200, gin.H{"message": "Cart deleted"})
}
