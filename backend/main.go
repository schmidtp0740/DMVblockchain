package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func newVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	fmt.Println(params)
}

func main() {
	fmt.Println("hello")

	router := mux.NewRouter()
	router.HandleFunc("/newVehicle", newVehicle).Methods("POST")
}
