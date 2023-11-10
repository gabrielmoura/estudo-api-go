package routes

import (
	"fmt"
	"log"

	"github.com/gabrielmoura/estudo-api-go/infra/logger"
	"github.com/gabrielmoura/estudo-api-go/infra/middleware"
	"github.com/gin-gonic/gin"
)

func HandleRequest(addr string) {
	r := gin.New()

	r.Use(middleware.GinLogger(logger.Logger), gin.Recovery())

	r.Use(middleware.CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		fmt.Fprint(c.Writer, "Bem Vindo")
	})
	/** Rotas de Autenticação **/
	r.POST("/login", LoginHandler)

	auth := r.Group("")
	auth.Use(middleware.JwtAuthMiddleware())

	/** Rotas de Usuário **/
	auth.GET("/user", getAllUser)
	auth.GET("/user/:id", getOneUser)
	auth.PUT("/user/:id", updateUser)
	auth.POST("/user", postUser)
	auth.DELETE("/user", deleteUser)

	/** Rotas de Pessoa **/
	auth.GET("/person", getAllPerson)
	auth.GET("/person/:id", getOnePerson)

	log.Fatal(r.Run(addr))
}
