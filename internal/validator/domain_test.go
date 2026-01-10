package validator

import "testing"

func TestValidateDomain_ValidDomains(t *testing.T) {
	validDomains := []string{
		"example.com",
		"www.example.com",
		"sub.domain.co",
	}

	for _, domain := range validDomains {
		if err := ValidateDomain(domain); err != nil {
			t.Errorf("expected domain %s to be valid, got error: %v", domain, err)
		}
	}
}

func TestValidateDomain_InvalidDomains(t *testing.T) {
	invalidDomains := []string{
		"",
		"http://example.com",
		"https://example.com",
		"example",
		"example.",
		".com",
		"exa_mple.com",
	}

	for _, domain := range invalidDomains {
		if err := ValidateDomain(domain); err == nil {
			t.Errorf("expected domain %s to be invalid", domain)
		}
	}
}