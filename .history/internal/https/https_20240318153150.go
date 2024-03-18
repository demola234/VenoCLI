package https

import "net/http"

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


