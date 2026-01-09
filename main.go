package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
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

	if err := validateHTTPStatus(resp); err != nil {
		return nil, err
	}

	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*
validateHTTPStatus() checks the status code of
the API response, to handle different types
of errors and responses
*/
func validateHTTPStatus(resp *http.Response) error {

	/*
		he following status codes are used in the SSL Labs API:

		400 - invocation error (e.g., invalid parameters)
		429 - client request rate too high or too many new assessments too fast
		500 - internal error
		503 - the service is not available (e.g., down for maintenance)
		529 - the service is overloaded
	*/
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusBadRequest:
		return fmt.Errorf("invalid request (400): check domain or parameters")
	case http.StatusTooManyRequests:
		return fmt.Errorf("rate limit exceeded (429): too many requests")
	case http.StatusInternalServerError:
		return fmt.Errorf("internal server error (500)")
	case http.StatusServiceUnavailable:
		return fmt.Errorf("service unavailable (503)")
	case 529:
		return fmt.Errorf("service is overloaded (529)")
	default:
		if resp.StatusCode >= 500 {
			return fmt.Errorf("server error (%d)", resp.StatusCode)
		}
		return fmt.Errorf("unexpected HTTP status (%d)", resp.StatusCode)
	}
}

func validateDomain(domain string) error {
	if domain == "" {
		return fmt.Errorf("domain cannot be empty")
	}

	// Reject URLs with scheme
	if strings.Contains(domain, "://") {
		return fmt.Errorf("domain must not include scheme (http:// or https://)")
	}

	// Regex para validar dominios: permite subdominios, letras, números y guiones
	var domainRegex = regexp.MustCompile(`^(?:[a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`)

	if !domainRegex.MatchString(domain) {
		return fmt.Errorf("invalid domain format")
	}

	return nil
}

func main() {

	printBanner()

	var domain string

	for {
		if len(os.Args) > 1 {
			domain = os.Args[1]
		} else {
			fmt.Print("Enter domain: ")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			domain = strings.TrimSpace(input)
		}

		if err := validateDomain(domain); err != nil {
			fmt.Println("Invalid domain:", err)
		} else {
			break
		}
	}

	// Check API availability before starting the analysis
	if err := check(); err != nil {
		fmt.Println("Failed to connect.", err)
		return
	}

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

	fmt.Println("\nPress Enter to exit...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func printBanner() {
	fmt.Println(`===============================================================================================`)
	fmt.Println(`                                                                                                                                                                    
███  ██ ▄▄▄▄▄ ▄▄▄▄  ▄▄ ▄▄ ▄▄     ▄▄▄    ▄█████ ▄▄ ▄▄  ▄▄▄  ▄▄    ▄▄    ▄▄▄▄▄ ▄▄  ▄▄  ▄▄▄▄ ▄▄▄▄▄ 
██ ▀▄██ ██▄▄  ██▄██ ██ ██ ██    ██▀██   ██     ██▄██ ██▀██ ██    ██    ██▄▄  ███▄██ ██ ▄▄ ██▄▄  
██   ██ ██▄▄▄ ██▄█▀ ▀███▀ ██▄▄▄ ██▀██   ▀█████ ██ ██ ██▀██ ██▄▄▄ ██▄▄▄ ██▄▄▄ ██ ▀██ ▀███▀ ██▄▄▄ 
	`)
	fmt.Println(`===============================================================================================`)

																	            
	fmt.Println()
	fmt.Println("Nebula Challenge - TLS Security Checker (Go)")
	fmt.Println("A high-performance command-line tool for analyzing TLS/SSL security of domains.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  go run main.go               -> prompts for a domain")
	fmt.Println("  go run main.go <domain.com>  -> checks the specified domain")
	fmt.Println()
	fmt.Println("Domain format:")
	fmt.Println("  - Must be a valid hostname, e.g., www.example.com")
	fmt.Println("  - Can include subdomains")
	fmt.Println("  - Only letters, numbers, and hyphens allowed in labels")
	fmt.Println("  - Must have a valid TLD (e.g., .com, .edu, .co)")
	fmt.Println("===================================================")
	fmt.Println()
}
