package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gabrielmoura/estudo-api-go/infra/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAllPerson(c *gin.Context) {
	person, err := db.GetAllPerson(db.Con)
	if err != nil {
		fmt.Fprint(c.Writer, "Erro: "+err.Error())
		return
	}
	json.NewEncoder(c.Writer).Encode(person)

}

func getOnePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	person, err := db.GetPerson(db.Con, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
	}
	c.JSON(http.StatusOK, person)
}
