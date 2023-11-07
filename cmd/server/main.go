package main

import (
	"fmt"
	"github.com/gabrielmoura/estudo-api-go/configs"
	"github.com/gabrielmoura/estudo-api-go/infra/db"
	"github.com/gabrielmoura/estudo-api-go/infra/routes"
	"github.com/gabrielmoura/estudo-api-go/internal/entity"
)

func main() {
	// Define Configurações
	conf, _ := configs.LoadConfig(".")
	// Inicia conexão com banco de dados
	//conn := db.Conn(conf.DBDriver, conf.DBName)
	//app := &App{
	//	conn: db.Conn(conf.DBDriver, conf.DBName),
	//}
	db.Con = db.Conn(conf.DBDriver, conf.DBName)

	person := entity.NewPerson("John", 30)
	user, _ := entity.NewUser("Administrador", "admin@example.com", "123456")

	db.InsertPerson(db.Con, person)
	db.InsertUser(db.Con, user)

	println(person.Name)

	println(fmt.Sprintf("Nome do Usuário:%s, email: %s, senha: %s", user.Name, user.Email, user.Password))

	port := fmt.Sprintf(":%d", conf.WebServerPort)
	routes.HandleRequest(port)
}
