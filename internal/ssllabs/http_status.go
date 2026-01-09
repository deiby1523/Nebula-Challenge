package ssllabs

import (
	"fmt"
	"net/http"
)

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