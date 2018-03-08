package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/schmidtp0740/DMVblockchain/backend/API"
)

func main() {
	fmt.Println("hello")

	router := mux.NewRouter()
	router.HandleFunc("/newVehicle", API.NewVehicle).Methods("POST")
	router.HandleFunc("/changeOwner/{vin}", API.ChangeOwner).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8000", router))
}
