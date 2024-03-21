package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"demola/veno/utils/errors"
	"math/rand"
	"net/http"
	"strconv"
)

const (
	Size1Kb  = 1024
	Size1Mb  = Size1Kb * 1024
	Size10Mb = Size1Mb * 10
)

var (
	ErrNoFormat = errors.New("no video format provided")
)

// DefaultClient type to use. No reason to change but you could if you wanted to.
var DefaultClient = AndroidClient

// Client offers methods to download video metadata and video streams.
type Client struct {
	// HTTPClient can be used to set a custom HTTP client.
	// If not set, http.DefaultClient will be used
	HTTPClient *http.Client

	// MaxRoutines to use when downloading a video.
	MaxRoutines int

	// ChunkSize to use when downloading videos in chunks. Default is Size10Mb.
	ChunkSize int64

	// playerCache caches the JavaScript code of a player response
	// playerCache playerCache

	// client *clientInfo

	consentID string
}

func (c *Client) httpDo(req *http.Request) (*http.Response, error) {
	client := c.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}

	req.Header.Set("User-Agent", "Go-Video-Downloader/1.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")

	if len(c.consentID) == 0 {
		c.consentID = strconv.Itoa(rand.Intn(1000000))
	}

	req.AddCookie(&http.Cookie{
		Name:     "consent",
		Value:    "YES+cb.20210328-17-p0.en+FX+" + c.consentID,
		Path:     "/",
		Domain:   ".youtube.com",
		HttpOnly: true,
	})

	res, err := client.Do(req)

	log := slog.With("method", req.Method, "url", req.URL)

	if err != nil {
		log.Debug(err)
	} else {
		log.Debug("consent cookie set")
	}

	return res, err
}

func (c *Client) httpGet(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.httpDo(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("http status %d", resp.StatusCode)
	}

	return resp, nil

}

func (c *Client) httpGetBodyBytes(ctx context.Context, url string) ([]byte, error) {
	resp, err := c.httpGet(ctx, url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)

}

// httpPost does a HTTP POST request with a body, checks the response to be a 200 OK and returns it
func (c *Client) httpPost(ctx context.Context, url string, body interface{}) (*http.Response, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Youtube-Client-Name", "3")
	req.Header.Set("X-Youtube-Client-Version", c.client.version)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	resp, err := c.httpDo(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, errors
	}

	return resp, nil
}

// httpPostBodyBytes reads the whole HTTP body and returns it
func (c *Client) httpPostBodyBytes(ctx context.Context, url string, body interface{}) ([]byte, error) {
	resp, err := c.httpPost(ctx, url, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
