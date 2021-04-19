package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jacobintern/GoogleLogin/controllers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/home", controllers.Hello).Methods(http.MethodGet)
	r.HandleFunc("/login", controllers.LoginPage).Methods(http.MethodGet)

	http.Handle("/", r)
	err := http.ListenAndServe(":8822", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
