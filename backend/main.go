package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/schmidtp0740/DMVblockchain/backend/api"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/newVehicle", api.NewVehicle).Methods("POST")
	router.HandleFunc("/changeOwner", api.ChangeOwner).Methods("PATCH")
	fmt.Println("Listening on port: 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
