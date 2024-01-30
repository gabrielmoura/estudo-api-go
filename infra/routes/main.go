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
	user := auth.Group("/user")
	user.GET("/", getAllUser)
	user.GET("/:id", getOneUser)
	user.PUT("/:id", updateUser)
	user.POST("/", postUser)
	user.DELETE("/", deleteUser)

	///** Rotas de Pessoa **/
	//auth.GET("/person", getAllPerson)
	//auth.GET("/person/:id", getOnePerson)

	/** Rotas de Produtos **/
	product := auth.Group("/product")
	product.GET("/", getAllProduct)
	product.GET("/:id", getOneProduct)
	product.PUT("/:id", updateProduct)
	product.POST("/", postProduct)
	product.DELETE("/", deleteProduct)

	/** Rotas de Carrinho **/
	cart := auth.Group("/cart")
	cart.GET("/", getAllCart)
	cart.GET("/:id", getOneCart)
	cart.PUT("/:id", updateCart)
	cart.POST("/", postCart)
	cart.DELETE("/", deleteCart)

	r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Fatal(r.Run(addr))
}
