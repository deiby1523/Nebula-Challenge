package validator

import (
	"fmt"
	"regexp"
	"strings"
)

func ValidateDomain(domain string) error {
	if domain == "" {
		return fmt.Errorf("domain cannot be empty")
	}

	// Reject URLs with scheme
	if strings.Contains(domain, "://") {
		return fmt.Errorf("domain must not include scheme (http:// or https://)")
	}

	// Regex para validar dominios: permite subdominios, letras, números y guiones
	regex := regexp.MustCompile(`^(?:[a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`)
	
	if !regex.MatchString(domain) {
		return fmt.Errorf("invalid domain format")
	}

	return nil
}