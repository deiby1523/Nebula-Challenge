package ssllabs

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