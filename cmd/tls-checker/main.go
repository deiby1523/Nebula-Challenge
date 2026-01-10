package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"Nebula-Challenge/internal/ssllabs"
	"Nebula-Challenge/internal/ui"
	"Nebula-Challenge/internal/validator"
)


var (
	results     = []ssllabs.Response{}
	resultsFile = "results.json"
)

// Save result on json file
func saveResult() {
	data, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		log.Fatal("Error saving result:", err)
	}
	os.WriteFile(resultsFile, data, 0644)
}

// Load results from json file
func loadResults() {
	if _, err := os.Stat(resultsFile); err == nil {
		// checking if the file exists
		data, err := os.ReadFile(resultsFile)
		if err != nil {
			log.Fatal("Error loading domains:", err)
		}
		json.Unmarshal(data, &results)
	}

}

func main() {

	ui.PrintBanner()

	domain, err := readDomain()
	if err != nil {
		log.Fatal(err)
	}

	client := ssllabs.NewClient(10 * time.Second)

	result, err := client.RunAnalysis(domain, 20)
	if err != nil {
		log.Fatal(err)
	}

	printResult(result)

	save, err := readYesNo("¿Do you want to save the result? (y/n): ")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if save {
		results = append(results, *result)
		saveResult()
		fmt.Println("Saved.")
	}
}

func readDomain() (string, error) {
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

		if err := validator.ValidateDomain(domain); err != nil {
			fmt.Println("Invalid domain:", err)
		} else {
			break
		}
	}
	return domain,nil
}

func printResult(result *ssllabs.Response) {
	fmt.Println("\nTLS result:")

	for _, ep := range result.Endpoints {
		if ep.StatusMessage != "Ready" {
			continue
		}

		fmt.Printf("IP: %s | Grade: %s\n", ep.IPAddress, ep.Grade)

		if ep.Duration <= 0 {
			fmt.Println("assessment duration: unknown")
		} else {
			seconds := float64(ep.Duration) / 1000
			fmt.Printf("assessment duration: %.2f s\n", seconds)
		}		


		// Verificamos que los detalles y el certificado existan
		if ep.Details == nil {
			fmt.Println("  └─ Certificate details not available")
			continue
		}

		cert := ep.Details.Cert

		fmt.Println("  Certificate information:")
		fmt.Printf("    ├─ Subject: %s\n", cert.Subject)
		fmt.Printf("    ├─ Issuer: %s\n", cert.IssuerLabel)
		fmt.Printf("    ├─ Signature Algorithm: %s\n", cert.SigAlg)

		if len(cert.CommonNames) > 0 {
			fmt.Printf("    ├─ Common Names: %v\n", cert.CommonNames)
		}

		if len(cert.AltNames) > 0 {
			fmt.Printf("    ├─ Alternative Names: %v\n", cert.AltNames)
		}

		fmt.Printf(
			"    ├─ Validity: %s → %s\n",
			time.UnixMilli(cert.NotBefore).Format(time.RFC3339),
			time.UnixMilli(cert.NotAfter).Format(time.RFC3339),
		)

		if cert.ValidationType != "" {
			fmt.Printf("    ├─ Validation Type: %s\n", cert.ValidationType)
		}
		fmt.Printf("    ├─ Issues Bitmask: %d\n", cert.Issues)
		fmt.Printf("    └─ Embedded SCT: %t\n", cert.Sct)

		fmt.Println()
	}
}


func readYesNo(question string) (bool, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(question)
		input, err := reader.ReadString('\n')
		if err != nil {
			return false, err
		}

		response := strings.ToLower(strings.TrimSpace(input))

		switch response {
		case "y":
			return true, nil
		case "n":
			return false, nil
		default:
			fmt.Println("Invalid Option. Please enter 'y' or 'n'.")
		}
	}
}
