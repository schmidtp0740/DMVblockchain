package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Vehicle containing ID, Year, Make, Model
type Vehicle struct {
	// new vehicle struct
	ID    string `json:"id,omitempty"`
	Year  int    `json:"year,omitempty"`
	Make  string `json:"make,omitempty"`
	Model string `json:"model,omitempty"`
}

func newVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle Vehicle
	json.NewDecoder(r.Body).Decode(&vehicle)
	fmt.Println(vehicle)
}

func changeOwner(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var vehicle Vehicle
	json.NewDecoder(r.Body).Decode(&vehicle)
	fmt.Printf("vehicle vin and info: %v %v", params["vin"], vehicle)
}
func main() {
	fmt.Println("hello")

	router := mux.NewRouter()
	router.HandleFunc("/newVehicle", newVehicle).Methods("POST")
	router.HandleFunc("/changeOwner/{vin}", changeOwner).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8000", router))
}
