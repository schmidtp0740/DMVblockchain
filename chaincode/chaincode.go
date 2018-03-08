package chaincode

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Chaincode ...
type Chaincode struct {
	// ...
}

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

func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		fmt.Printf("Error starting Chaincode: %s", err)
	}
}
