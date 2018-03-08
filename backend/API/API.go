package API

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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

	fmt.Println(vehicle.ID)

	url := "http://129.146.106.151:4001/bcsgw/rest/v1/transaction/invocation"

	var jsonStr = []byte(`{
		"channel": "mychannel",
		"chaincode": "emrCC",
		"chaincodeVer": "v1",
		"method": "insertObject",
		"args": ["` + vehicle.ID + `", "` + vehicle.Make + `", "` + vehicle.Model + `", "nil", ` + strconv.Itoa(vehicle.Year) + `, "nil", "nil", "nil", "nil"]
	}`)

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
