package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type DomainCerts struct {
	CertName string    `json:"cert_name"`
	Domains  []Domains `json:"domains"`
}
type Domains struct {
	Base string   `json:"base"`
	Subs []string `json:"subs"`
}

func main() {
	testCert()
}

func testCert() {
	// Read the JSON file.
	data, err := os.ReadFile("domains.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Unmarshal the JSON data into the DomainCerts struct.
	var certDomains DomainCerts
	err = json.Unmarshal(data, &certDomains)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	command := "sudo certbot --nginx --cert-name " + certDomains.CertName
	for _, d := range certDomains.Domains {
		command += " -d " + d.Base
		for _, value := range d.Subs {
			command += " -d " + value + "." + d.Base
		}
	}

	// Print the output of the command.
	fmt.Println(string(command))
}
