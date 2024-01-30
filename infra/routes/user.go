package routes

import (
	"github.com/gabrielmoura/estudo-api-go/internal/dto"
	"net/http"

	"github.com/gabrielmoura/estudo-api-go/infra/db"
	"github.com/gabrielmoura/estudo-api-go/internal/entity"
	"github.com/gin-gonic/gin"
)

// getAllUser godoc
//
//	@Summary		Get all Users
//	@Description	Get all Users
//	@Tags			user
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entity.User
//	@Router			/user [get]
func getAllUser(c *gin.Context) {

	users, err := db.GetAllUser(db.Con)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	var uResp []*dto.UserResponse
	for _, u := range users {
		uResp = append(uResp, &dto.UserResponse{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, uResp)
}

// getOneUser godoc
//
//	@Summary		Get one User
//	@Description	Get one User
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"User ID"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	entity.User
//	@Router			/user/{id} [get]
func getOneUser(c *gin.Context) {
	id := c.Params.ByName("id")
	user, err := db.GetOneUser(db.Con, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error(), "id": id})
	}
	c.JSON(http.StatusOK, user)
}

// postUser godoc
//
//	@Summary		Create User
//	@Description	Create User
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param 			user body entity.User true "User"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	entity.User
//	@Router			/user [post]
func postUser(c *gin.Context) {
	reqUser := entity.User{}
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := reqUser.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	user, _ := entity.NewUser(reqUser.Name, reqUser.Email, reqUser.Password)
	_, err := db.InsertUser(db.Con, user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// deleteUser godoc
//
//	@Summary		Delete User
//	@Description	Delete User
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"User ID"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	string
//	@Router			/user/{id} [delete]
func deleteUser(c *gin.Context) {
	u := entity.User{
		ID: c.Params.ByName("id"),
	}
	err := db.DeleteUser(db.Con, &u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

// updateUser godoc
//
//	@Summary		Update User
//	@Description	Update User
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"User ID"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	entity.User
//	@Router			/user/{id} [put]
func updateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	user, err := db.GetOneUser(db.Con, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := db.UpdateUser(db.Con, user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
