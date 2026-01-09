package ssllabs

import (
	"fmt"
	"time"
)

func (c *Client) RunAnalysis(domain string, maxAttempts int) (*Response, error) {

	// Check API availability before starting the analysis
	if err := c.check(); err != nil {
		return nil, fmt.Errorf("Failed to conect: %w", err)
	}

	// Start a new analysis request
	_, err := c.analyze(domain, true)
	if err != nil {
		return nil, err
	}

	// Poll the API until the analysis is completed
	for maxAttempts > 0 {

		result, err := c.analyze(domain, false)
		maxAttempts--

		if err != nil {
			return nil, err
		}

		fmt.Println("Status:", result.Status)

		switch result.Status {
		case "READY":
			return result, nil
		case "ERROR":
			return nil, fmt.Errorf("Analisis failed.")
		}

		// Wait 15 seconds before polling again to avoid excessive API requests
		time.Sleep(15 * time.Second)
	}

	return nil, fmt.Errorf("Analisis not finished, timeout.")
}