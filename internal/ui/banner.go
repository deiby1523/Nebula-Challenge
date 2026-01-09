package ui

import "fmt"

func PrintBanner() {
	fmt.Println(`===============================================================================================`)
	fmt.Println(`                                                                                                                                                                    
███  ██ ▄▄▄▄▄ ▄▄▄▄  ▄▄ ▄▄ ▄▄     ▄▄▄    ▄█████ ▄▄ ▄▄  ▄▄▄  ▄▄    ▄▄    ▄▄▄▄▄ ▄▄  ▄▄  ▄▄▄▄ ▄▄▄▄▄ 
██ ▀▄██ ██▄▄  ██▄██ ██ ██ ██    ██▀██   ██     ██▄██ ██▀██ ██    ██    ██▄▄  ███▄██ ██ ▄▄ ██▄▄  
██   ██ ██▄▄▄ ██▄█▀ ▀███▀ ██▄▄▄ ██▀██   ▀█████ ██ ██ ██▀██ ██▄▄▄ ██▄▄▄ ██▄▄▄ ██ ▀██ ▀███▀ ██▄▄▄ 
	`)
	fmt.Println(`===============================================================================================`)

	fmt.Println()
	fmt.Println("Nebula Challenge - TLS Security Checker")
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
