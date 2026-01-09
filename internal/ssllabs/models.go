package ssllabs

/*
Response represents the main response returned
by the SSL Labs analyze endpoint
*/
type Response struct {
	Host      string     `json:"host"`
	Status    string     `json:"status"`
	Port      int	 	 `json:"port"`
	Protocol  string	 `json:"protocol"`
	Endpoints []Endpoint `json:"endpoints"`
}

/*
Endpoint represents a single endpoint analyzed
by SSL Labs
*/
type Endpoint struct {
	IPAddress     string `json:"ipAddress"`
	ServerName    string `json:"serverName"`
	StatusMessage string `json:"statusMessage"`
	Grade         string `json:"grade"`
	Duration      int `json:"duration"`
	Details       *EndpointDetails `json:"details,omitempty"`
}

/*
EndpointDetails contains deep technical details
about the analyzed endpoint
*/
type EndpointDetails struct {
	HostStartTime int64 `json:"hostStartTime"`
	Cert          Cert  `json:"cert"`
}

/*
Cert represents the SSL certificate information
*/
type Cert struct {
	Subject            string   `json:"subject"`
	CommonNames        []string `json:"commonNames"`
	AltNames           []string `json:"altNames"`
	NotBefore          int64    `json:"notBefore"`
	NotAfter           int64    `json:"notAfter"`
	IssuerSubject      string   `json:"issuerSubject"`
	IssuerLabel        string   `json:"issuerLabel"`
	SigAlg             string   `json:"sigAlg"`
	ValidationType     string   `json:"validationType"`
	Issues             int      `json:"issues"`
	Sct                bool     `json:"sct"`
}