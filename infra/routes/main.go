package routes

import (
	"fmt"
	"log"

	_ "github.com/gabrielmoura/estudo-api-go/docs"
	"github.com/gabrielmoura/estudo-api-go/infra/logger"
	"github.com/gabrielmoura/estudo-api-go/infra/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
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
	r.POST("/register", RegisterHandler)

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

	/** Rotas de Produtos **/
	auth.GET("/product", getAllProduct)
	auth.GET("/product/:id", getOneProduct)
	auth.PUT("/product/:id", updateProduct)
	auth.POST("/product", postProduct)
	auth.DELETE("/product", deleteProduct)

	/** Rotas de Carrinho **/
	auth.GET("/cart", getAllCart)
	auth.GET("/cart/:id", getOneCart)
	auth.PUT("/cart/:id", updateCart)
	auth.POST("/cart", postCart)
	auth.DELETE("/cart", deleteCart)

	r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(r.Run(addr))
}
