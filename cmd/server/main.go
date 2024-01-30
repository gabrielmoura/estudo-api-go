package main

import (
	"fmt"

	"github.com/gabrielmoura/estudo-api-go/configs"
	"github.com/gabrielmoura/estudo-api-go/infra/db"
	"github.com/gabrielmoura/estudo-api-go/infra/logger"
	"github.com/gabrielmoura/estudo-api-go/infra/routes"
)

//	@title			Estudo API Go
//	@version		1.0
//	@description	API de estudo em Go

//	@host	localhost:8001

//	@securityDefinitions.apiKey ApiKeyAuth
//	@in header
//	@name Authorization

// @externalDocs.description	GitHub Repository
// @externalDocs.url			https://github.com/gabrielmoura/estudo-api-go
func main() {
	// Define Configurações
	conf, _ := configs.LoadConfig(".")

	logger.InitLogger()

	db.Conn(conf.DBDriver, conf.DBName)

	port := fmt.Sprintf(":%d", conf.WebServerPort)
	routes.HandleRequest(port)
}
