package https

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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

type PassThru struct {
	io.Reader
	total int64 // Total # of bytes transferred
}

// Read 'overrides' the underlying io.Reader's Read method.
// This is the one that will be called by io.Copy(). We simply
// use it to keep track of byte counts and then forward the call.
func (pt *PassThru) Read(p []byte) (int, error) {

	n, err := pt.Reader.Read(p)
	pt.total += int64(n)

	if err == nil {
		fmt.Println("Read", n, "bytes for a total of", pt.total)
	}

	return n, err
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
	var src io.Reader    // Source file/url/etc
	var dst bytes.Buffer // Destination file/buffer/etc

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

	// Example data source (can be replaced with actual source)
	src = bytes.NewBufferString(strings.Repeat("Some random input data", 1000))

	// Wrap it with our custom io.Reader.
	src = &PassThru{Reader: src}

	// Copy data from source to destination
	count, err := io.Copy(&dst, src)
	if err != nil {
		return nil, err
	}

	// Display processing percentage (for illustration purposes only)
	fmt.Printf("Processing complete: %.2f%%\n", float64(count)/float64(dst.Len())*100)

	return data, nil
}


\