package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

//
const baseURL = "https://api.ssllabs.com/api/v2"

type Response struct {
	Host      string     `json:"host"`
	Status    string     `json:"status"`
	Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	IPAddress     string `json:"ipAddress"`
	StatusMessage string `json:"statusMessage"`
	Grade         string `json:"grade"`
}

func check() error {
	fmt.Println("Conecting to SSL Labs API . . .")
	_,err := http.Get(baseURL+"/info")
	if err != nil {
		return err
	}
	return nil
}

func analyze(domain string, startNew bool) (*Response, error) {
	url := fmt.Sprintf("%s/analyze?host=%s&all=done", baseURL, domain)
	if startNew {
		url += "&startNew=on"
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result Response
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <dominio> Example: go run main.go www.uts.edu.co")
		return
	}

	if err := check(); err != nil {
		fmt.Println("Failed to connect.",err)
		return
	}

	domain := os.Args[1]
	fmt.Println("Starting TLS analisis for:", domain)

	// Start for the first time
	_, err := analyze(domain, true)
	if err != nil {
		fmt.Println("Error iniciando análisis:", err)
		return
	}

	// Call the API until finish
	for {
		result, err := analyze(domain, false)
		if err != nil {
			fmt.Println("Error consultando análisis:", err)
			return
		}

		fmt.Println("Estado:", result.Status)

		if result.Status == "READY" {
			fmt.Println("\nResultado TLS:")
			for _, ep := range result.Endpoints {
				if ep.StatusMessage == "Ready" {
					fmt.Printf("IP: %s | Calificación: %s\n", ep.IPAddress, ep.Grade)
				}
			}
			break
		}

		if result.Status == "ERROR" {
			fmt.Println("El análisis falló.")
			break
		}

		// Wait 15 seconds
		time.Sleep(15 * time.Second)
	}
}
