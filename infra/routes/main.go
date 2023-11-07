package routes

import (
	"fmt"
	"log"
	"net/http"
)

func HandleRequest(addr string) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Bem Vindo")
	})
	http.HandleFunc("/user", getAllUser)
	http.HandleFunc("/person", getAllPerson)
	log.Fatal(http.ListenAndServe(addr, nil))
}
