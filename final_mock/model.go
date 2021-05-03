package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type BlueprintCreateRequest struct {
	Name            string          `json:"name"`
	ResourceUri     string          `json:"resourceUri"`
	Type            string          `json:"type"`
	Description     string          `json:"description"`
	Tag             string          `json:"tag"`
	AutoApply       bool            `json:"autoApply"`
	Network         Network         `json:"network"`
	NetworkServices NetworkServices `json:"networkServices"`
}

type Blueprint struct {
	ID            string    `json:"id"`
	CustomerId    string    `json:"customerId"`
	Generation    int       `json:"generation"`
	SchemaVersion string    `json:"schemaVersion"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	BlueprintCreateRequest
}

type Network struct {
	StartingIpv4Address      net.IP     `json:"startingIpv4Address"`
	EndingIpv4Address        net.IP     `json:"endingIpv4Address"`
	StartingIpv6AddressRange net.IP     `json:"startingIpv6AddressRange"`
	EndingIpv6AddressRange   net.IP     `json:"endingIpv6AddressRange"`
	Gateway                  net.IP     `json:"gateway"`
	SubnetMask               net.IPMask `json:"subnetMask"`
}

type NetworkServices struct {
	DnsServers []net.IP `json:"dnsServers`
	NtpServers []string `json:"ntpServers`
}

var Blueprints []Blueprint

func init() {

	bp1bytes := []byte(`
	{
	"customerId": "string",
	"generation": 0,
	"id": "1234",
	"name": "Customer11",
	"resourceUri": "http://example.com",
	"type": "DHCI",
	"description": "string",
	"tag": "string",
	"createdAt": "1970-01-01T00:00:00.000Z",
	"schemaVersion": "string",
	"updatedAt": "1970-01-01T00:00:00.000Z",
	"autoApply": false,
	"network": {
	  "endingIpv4Address": "192.168.100.100",
	  "endingIpv6AddressRange": "2001:0db8:5b96:0000:0000:426f:8e17:642a",
	  "gateway": "198.51.100.42",
	  "startingIpv4Address": "192.168.100.1",
	  "startingIpv6AddressRange": "2001:0db8:5b96:0000:0000:426f:8e17:642a",
	  "subnetMask": "255.255.0.0"
	},
	"networkServices": {
	  "dnsServers": [
		"198.51.100.42"
	  ],
	  "ntpServers": [
		"string"
	  ]
	}
  }`)
	var bp Blueprint
	fmt.Printf("Error: %v\n", json.Unmarshal(bp1bytes, &bp))
	Blueprints = append(Blueprints, bp)
	fmt.Printf("Error: %v\n", json.Unmarshal([]byte(`
	{
	"customerId": "Customer22",
	"generation": 0,
	"id": "9999",
	"name": "string",
	"resourceUri": "http://example.com",
	"type": "Nimble",
	"description": "string",
	"tag": "string",
	"createdAt": "1970-01-01T00:00:00.000Z",
	"schemaVersion": "string",
	"updatedAt": "1970-01-01T00:00:00.000Z",
	"autoApply": false,
	"network": {
	  "endingIpv4Address": "192.168.100.100",
	  "endingIpv6AddressRange": "2001:0db8:5b96:0000:0000:426f:8e17:642a",
	  "gateway": "198.51.100.42",
	  "startingIpv4Address": "192.168.100.1",
	  "startingIpv6AddressRange": "2001:0db8:5b96:0000:0000:426f:8e17:642a",
	  "subnetMask": "255.255.255.0"
	},
	"networkServices": {
	  "dnsServers": [
		"198.51.100.42"
	  ],
	  "ntpServers": [
		"string"
	  ]
	}
  }`), &bp))
	Blueprints = append(Blueprints, bp)
}
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func GetBlueprints(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Blueprints)
}

func GetBlueprint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range Blueprints {
		if item.ID == params["id"] {

			respondWithJSON(w, http.StatusOK, item)
			return
		}
	}
	respondWithError(w, http.StatusNotFound, "item not there")

}
func AddBlueprint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bp Blueprint
	_ = json.NewDecoder(r.Body).Decode(&bp)
	bp.ID = strconv.Itoa(rand.Intn(1000000))
	bp.CustomerId = strconv.Itoa(rand.Intn(1000000))
	bp.Generation = 0
	bp.CreatedAt = time.Now()
	bp.UpdatedAt = bp.CreatedAt
	bp.SchemaVersion = "1.0"
	Blueprints = append(Blueprints, bp)
	respondWithJSON(w, http.StatusOK, &bp)

}
