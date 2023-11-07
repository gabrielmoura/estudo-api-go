package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gabrielmoura/estudo-api-go/infra/db"
	"github.com/gabrielmoura/estudo-api-go/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAllUser(c *gin.Context) {

	users, err := db.GetAllUser(db.Con)
	if err != nil {
		fmt.Fprint(c.Writer, "Erro: "+err.Error())
		return
	}
	json.NewEncoder(c.Writer).Encode(users)
}
func getOneUser(c *gin.Context) {
	id := c.Params.ByName("id")
	user, err := db.GetOneUser(db.Con, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error(), "id": id})
	}
	c.JSON(http.StatusOK, user)
}

func postUser(c *gin.Context) {
	reqUser := entity.User{}
	if err := c.ShouldBindJSON(&reqUser); err != nil {
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
