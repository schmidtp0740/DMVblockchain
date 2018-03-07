package main

import (
	"bytes"
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

// NewVehicle ...
func NewVehicle(w http.ResponseWriter, r *http.Request) {
	// ...
	var vehicle Vehicle
	json.NewDecoder(r.Body).Decode(&vehicle)
	fmt.Println(vehicle)

	var jsonStr = []byte(`{
		"channel": "mychannel",
		"chaincode": "emrCC",
		"chaincodeVer": "v1",
		"args": ["", "", "", 123, "", "", "", ""]
	}`)
	resp, err := http.NewRequest("POST",
		"http:129.0.0.1:4001/bcsgw/rest/v1/transaction/invocation",
		bytes.NewBuffer(jsonStr))

	fmt.Printf("Response from blockchain: %v\n%v", resp, err)
}

// ChangeOwner ...
func ChangeOwner(w http.ResponseWriter, r *http.Request) {
	// ...
	params := mux.Vars(r)

	var vehicle Vehicle
	json.NewDecoder(r.Body).Decode(&vehicle)
	fmt.Printf("vehicle vin and info: %v %v", params["vin"], vehicle)
}

func main() {
	fmt.Println("hello")

	router := mux.NewRouter()
	router.HandleFunc("/newVehicle", NewVehicle).Methods("POST")
	router.HandleFunc("/changeOwner/{vin}", ChangeOwner).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8000", router))
}
