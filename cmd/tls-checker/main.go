package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"Nebula-Challenge/internal/ssllabs"
	"Nebula-Challenge/internal/ui"
	"Nebula-Challenge/internal/validator"
)

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

	fmt.Println("\nPress Enter to exit...")
	bufio.NewReader(os.Stdin).ReadString('\n')
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
		if ep.StatusMessage == "Ready" {
			fmt.Printf("IP: %s | Grade: %s\n", ep.IPAddress, ep.Grade)
		}
	}
}
