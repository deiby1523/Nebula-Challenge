package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Base URL for SSL Labs API v2
const baseURL = "https://api.ssllabs.com/api/v2"

/*
Response represents the main response returned
by the SSL Labs analyze endpoint
*/
type Response struct {
	Host      string     `json:"host"`
	Status    string     `json:"status"`
	Endpoints []Endpoint `json:"endpoints"`
}

/*
Endpoint represents a single endpoint analyzed
by SSL Labs
*/
type Endpoint struct {
	IPAddress     string `json:"ipAddress"`
	StatusMessage string `json:"statusMessage"`
	Grade         string `json:"grade"`
}

/*
check() verifies basic connectivity with the
SSL Labs API
*/
func check() error {
	fmt.Println("Conecting to SSL Labs API . . .")
	_, err := http.Get(baseURL + "/info")
	if err != nil {
		return err
	}
	return nil
}

/*
analyze() triggers or retrieves a TLS analysis
for a given domain. If startNew is true, a new
analysis is explicitly started.
*/
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

	// 	The following status codes are used by SSL Labs:

	// 400 - invocation error (e.g., invalid parameters)
	// 429 - client request rate too high or too many new assessments too fast
	// 500 - internal error
	// 503 - the service is not available (e.g., down for maintenance)
	// 529 - the service is overloaded
	switch resp.StatusCode {
	case http.StatusOK:
		// Continue
	case http.StatusBadRequest:
		return nil, fmt.Errorf("invalid request (400): check domain or parameters")
	case http.StatusTooManyRequests:
		return nil, fmt.Errorf("rate limit exceeded (429): too many requests")
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("internal server error (500)")
	case http.StatusServiceUnavailable:
		return nil, fmt.Errorf("service unavailable (503)")
	default:
		if resp.StatusCode >= 500 {
			return nil, fmt.Errorf("server error (%d)", resp.StatusCode)
		}
		return nil, fmt.Errorf("unexpected HTTP status (%d)", resp.StatusCode)
	}

	var result Response
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func main() {

	// Validate command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <dominio> Example: go run main.go www.uts.edu.co")
		return
	}

	// Check API availability before starting the analysis
	if err := check(); err != nil {
		fmt.Println("Failed to connect.", err)
		return
	}

	domain := os.Args[1]
	fmt.Println("Starting TLS analisis for:", domain)

	// Start a new analysis request
	_, err := analyze(domain, true)
	if err != nil {
		fmt.Println("Error iniciando análisis:", err)
		return
	}

	// Poll the API until the analysis is completed
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

		// Wait 15 seconds before polling again to avoid excessive API requests
		time.Sleep(15 * time.Second)
	}
}
