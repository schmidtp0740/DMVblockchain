package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Vehicle containing ID, Year, Make, Model
type Vehicle struct {
	// new vehicle struct
	VIN           string `json:"id,omitempty"`
	Year          string `json:"year,omitempty"`
	Make          string `json:"make,omitempty"`
	Model         string `json:"model,omitempty"`
	Mileage       string `json:"mileage,omitempty"`
	Salvage       string `json:"salvage,omitempty"`
	PurchasePrice string `json:"purchasePrice,omitempty"`
	Owner         string `json:"owner,omitempty"`
	DOB           string `json:"dob,omitempty"`
	StreetAddress string `json:"streetAddress,omitempty"`
	City          string `json:"city,omitempty"`
	State         string `json:"state,omitempty"`
	Zip           string `json:"zip,omitempty"`
}

// NewVehicle ...
func NewVehicle(w http.ResponseWriter, r *http.Request) {
	// ...

	var vehicle Vehicle

	json.NewDecoder(r.Body).Decode(&vehicle)

	fmt.Println(vehicle.VIN)

	url := "http://129.146.106.151:4001/bcsgw/rest/v1/transaction/invocation"

	args := `["` +
		vehicle.VIN + `", "` +
		vehicle.Year + `", "` +
		vehicle.Make + `", "` +
		vehicle.Model + `", "` +
		vehicle.Mileage + `", "` +
		vehicle.Salvage + `", "` +
		vehicle.PurchasePrice + `", "` +
		vehicle.Owner + `", "` +
		vehicle.DOB + `", "` +
		vehicle.StreetAddress + `", "` +
		vehicle.City + `", "` +
		vehicle.State + `", "` +
		vehicle.Zip + `]`

	var jsonStr = []byte(`{
		"channel": "mychannel",
		"chaincode": "emrCC",
		"chaincodeVer": "v1",
		"method": "insertObject",
		"args": ` + args + `}`)

	fmt.Println("JSON to Blockchain: \n", string(jsonStr))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	json.NewEncoder(w).Encode(string(body))

	fmt.Println("Response from blockchain: ", string(body))

}

// ChangeOwner ...
func ChangeOwner(w http.ResponseWriter, r *http.Request) {
	// ...
	params := mux.Vars(r)

	var vehicle Vehicle

	json.NewDecoder(r.Body).Decode(&vehicle)

	fmt.Printf("vehicle vin and info: %v %v", params["vin"], vehicle)

}
