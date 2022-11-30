package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maabodi/go-mux/controller/productcontroller"
	"github.com/maabodi/go-mux/controller/usercontroller"
	"github.com/maabodi/go-mux/models"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/products", productcontroller.Index).Methods("GET")
	r.HandleFunc("/product/{id}", productcontroller.Show).Methods("GET")
	r.HandleFunc("/product", productcontroller.Create).Methods("POST")
	r.HandleFunc("/product/{id}", productcontroller.Update).Methods("PUT")
	r.HandleFunc("/product", productcontroller.Delete).Methods("DELETE")

	r.HandleFunc("/users", usercontroller.Index).Methods("GET")
	r.HandleFunc("/user/{id}", usercontroller.Show).Methods("GET")
	r.HandleFunc("/user", usercontroller.Create).Methods("POST")

	log.Fatal(http.ListenAndServe(":1323", r))
}
