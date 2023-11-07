package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func HandleRequest(addr string) {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		fmt.Fprint(c.Writer, "Bem Vindo")
	})
	/** Rotas de Usuário **/
	r.GET("/user", getAllUser)
	r.GET("/user/:id", getOneUser)
	r.PUT("/user/:id", updateUser)
	r.POST("/user", postUser)
	r.DELETE("/user", deleteUser)

	/** Rotas de Pessoa **/
	r.GET("/person", getAllPerson)
	r.GET("/person/:id", getOnePerson)

	log.Fatal(r.Run(addr))
}
