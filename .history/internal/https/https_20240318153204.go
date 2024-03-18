package https

import (
	"fmt"
	"net/http"
)

// Process Requests
// Allow Generic Request Handler to be passed in
// Check Status Code
// Allow Generic PASSING of different unmarshaling logic
// Display Percentage of file processing

type HTTPS struct {
	ReqHandler  RequestHandler
	Unmarshaler Unmarshaler
}

// RequestHandler defines the interface for handling HTTP requests.
type RequestHandler interface {
	HandleRequest(req *http.Request) (*http.Response, error)
}

// Unmarshaler defines the interface for unmarshaling HTTP responses.
type Unmarshaler interface {
	UnmarshalResponse(resp *http.Response) (interface{}, error)
}

// NewHTTPS creates a new instance of the HTTPS package.
func NewHTTPS(reqHandler RequestHandler, unmarshaler Unmarshaler) *HTTPS {
	return &HTTPS{
		ReqHandler:  reqHandler,
		Unmarshaler: unmarshaler,
	}
}

// ProcessRequest sends an HTTP request and processes the response.
func (h *HTTPS) ProcessRequest(req *http.Request) (interface{}, error) {
	// Send the request using the provided handler
	resp, err := h.ReqHandler.HandleRequest(req)
	if err != nil {
		return nil, err
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Unmarshal the response using the provided unmarshaler
	data, err := h.Unmarshaler.UnmarshalResponse(resp)
	if err != nil {
		return nil, err
	}

	// Display processing percentage (for illustration purposes only)
	fmt.Println("Processing complete: 100%")

	return data, nil
}
