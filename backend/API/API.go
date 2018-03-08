package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type vehicle struct {
	VIN           string `json:"vin,omitempty"`
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

type blockchainCall struct {
	channel      string
	chaincode    string
	chaincodeVer string
	method       string
	args         []string
}

// NewVehicle ...
func NewVehicle(w http.ResponseWriter, r *http.Request) {
	// ...
	handler(w, r, "insertObject")

}

// ChangeOwner ...
func ChangeOwner(w http.ResponseWriter, r *http.Request) {
	// ...
	handler(w, r, "modifyObject")

}

func getURL() (url string) {

	file, err := os.Open(".env")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]byte, 100)

	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	url = string(data[:count])

	fmt.Printf("url: %s\n\n", url)

	return
}

func blockchainRequest(m blockchainCall, c chan string) string {

	b, err := json.Marshal(m)

	url := <-c // Receive url from Channel
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func handler(w http.ResponseWriter, r *http.Request, action string) {
	var vehicle vehicle

	json.NewDecoder(r.Body).Decode(&vehicle)

	channelForURL := make(chan string)
	go func() {
		url := getURL() + "/bcsgw/rest/v1/transaction/invocation"
		channelForURL <- url
	}()

	m := blockchainCall{
		"mychannel",
		"emrCC",
		"v1",
		action,
		[]string{
			vehicle.VIN, vehicle.Year, vehicle.Make, vehicle.Model,
			vehicle.Mileage, vehicle.Salvage, vehicle.PurchasePrice, vehicle.Owner,
			vehicle.DOB, vehicle.StreetAddress, vehicle.City, vehicle.State, vehicle.Zip,
		},
	}

	body := blockchainRequest(m, channelForURL)

	json.NewEncoder(w).Encode(body)

	fmt.Printf("Response from blockchain: %s\n\n", body)
}
