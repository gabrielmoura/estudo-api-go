package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gabrielmoura/estudo-api-go/infra/db"
	"net/http"
)

func getAllPerson(w http.ResponseWriter, r *http.Request) {
	person, err := db.GetAllPerson(db.Con)
	if err != nil {
		fmt.Fprint(w, "Erro: "+err.Error())
		return
	}
	json.NewEncoder(w).Encode(person)

}
func getAllUser(w http.ResponseWriter, r *http.Request) {

	users, err := db.GetAllUser(db.Con)
	if err != nil {
		fmt.Fprint(w, "Erro: "+err.Error())
		return
	}
	json.NewEncoder(w).Encode(users)
}
