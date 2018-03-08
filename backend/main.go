package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/schmidtp0740/DMVblockchain/backend/api"
)

func main() {
	fmt.Println("hello")

	router := mux.NewRouter()

	router.HandleFunc("/newVehicle", api.NewVehicle).Methods("POST")
	router.HandleFunc("/changeOwner/{vin}", api.ChangeOwner).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8000", router))
}
