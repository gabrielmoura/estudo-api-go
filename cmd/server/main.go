package main

import (
	"fmt"
	"github.com/gabrielmoura/estudo-api-go/configs"
	"github.com/gabrielmoura/estudo-api-go/infra/db"
	"github.com/gabrielmoura/estudo-api-go/infra/routes"
)

func main() {
	// Define Configurações
	conf, _ := configs.LoadConfig(".")
	db.Con = db.Conn(conf.DBDriver, conf.DBName)

	port := fmt.Sprintf(":%d", conf.WebServerPort)
	routes.HandleRequest(port)
}
