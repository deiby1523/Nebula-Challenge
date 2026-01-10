package ssllabs

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRunAnalysis_Timeout(t *testing.T) {
	
	// Servidor falso que siempre responde IN_PROGRESS
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"status": "IN_PROGRESS",
			"endpoints": []
		}`))
	}))
	defer server.Close()

	// Crear cliente apuntando al servidor falso
	client := &Client{
		httpClient: &http.Client{
			Timeout: 2 * time.Second,
		},
	}

	// Sobrescribimos baseURL solo para el test
	originalBaseURL := baseURL
	baseURL = server.URL
	
	defer func() { baseURL = originalBaseURL }()

	_, err := client.RunAnalysis("example.com", 3)

	if err == nil {
		t.Errorf("expected timeout error, got nil")
	}
}