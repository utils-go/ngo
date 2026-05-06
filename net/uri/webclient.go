// Package uri provides .NET System.Net-like WebClient and Uri utilities for Go.
package uri

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// WebClient provides common methods for sending data to and receiving data from a resource.
// Equivalent to System.Net.WebClient in .NET.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.net.webclient?view=netframework-4.7.2
type WebClient struct {
	client  *http.Client
	headers map[string]string
}

// NewWebClient creates a new WebClient.
func NewWebClient() *WebClient {
	return &WebClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		headers: make(map[string]string),
	}
}

// SetTimeout sets the request timeout.
func (wc *WebClient) SetTimeout(d time.Duration) {
	wc.client.Timeout = d
}

// SetHeader sets a header for all requests.
func (wc *WebClient) SetHeader(name, value string) {
	wc.headers[name] = value
}

// DownloadString downloads the requested resource as a string.
func (wc *WebClient) DownloadString(address string) (string, error) {
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return "", err
	}
	wc.applyHeaders(req)

	resp, err := wc.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// DownloadData downloads the requested resource as a byte array.
func (wc *WebClient) DownloadData(address string) ([]byte, error) {
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return nil, err
	}
	wc.applyHeaders(req)

	resp, err := wc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// UploadString uploads a string to the specified resource.
func (wc *WebClient) UploadString(address, method, data string) (string, error) {
	req, err := http.NewRequest(method, address, strings.NewReader(data))
	if err != nil {
		return "", err
	}
	wc.applyHeaders(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := wc.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// UploadData uploads a byte array to the specified resource.
func (wc *WebClient) UploadData(address, method string, data []byte) ([]byte, error) {
	req, err := http.NewRequest(method, address, strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	wc.applyHeaders(req)

	resp, err := wc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (wc *WebClient) applyHeaders(req *http.Request) {
	for k, v := range wc.headers {
		req.Header.Set(k, v)
	}
}

// ---- Uri ----

// Uri provides an object representation of a uniform resource identifier (URI).
// Equivalent to System.Uri in .NET.
// Reference: https://learn.microsoft.com/en-us/dotnet/api/system.uri?view=netframework-4.7.2
type Uri struct {
	raw      string
	parsed   *url.URL
}

// ParseUri creates a new Uri from a URI string.
func ParseUri(uriString string) (*Uri, error) {
	parsed, err := url.Parse(uriString)
	if err != nil {
		return nil, err
	}
	return &Uri{raw: uriString, parsed: parsed}, nil
}

// Scheme gets the scheme name for this URI.
func (u *Uri) Scheme() string {
	return u.parsed.Scheme
}

// Host gets the host component of this instance.
func (u *Uri) Host() string {
	return u.parsed.Host
}

// Port gets the port number.
func (u *Uri) Port() string {
	return u.parsed.Port()
}

// Path gets the path of the URI.
func (u *Uri) Path() string {
	return u.parsed.Path
}

// Query gets the query string.
func (u *Uri) Query() string {
	return u.parsed.RawQuery
}

// Fragment gets the fragment portion.
func (u *Uri) Fragment() string {
	return u.parsed.Fragment
}

// AbsoluteUri gets the absolute URI.
func (u *Uri) AbsoluteUri() string {
	return u.parsed.String()
}

// OriginalString gets the original URI string that was passed to the constructor.
func (u *Uri) OriginalString() string {
	return u.raw
}

// IsAbsoluteUri indicates whether the Uri instance is absolute.
func (u *Uri) IsAbsoluteUri() bool {
	return u.parsed.IsAbs()
}

// String returns the canonical string representation.
func (u *Uri) String() string {
	return u.parsed.String()
}

// EscapeUriString converts a URI string to its escaped representation.
func EscapeUriString(str string) string {
	return url.PathEscape(str)
}

// EscapeDataString converts a string to its escaped representation for use in URI query data.
func EscapeDataString(str string) string {
	return url.QueryEscape(str)
}

// UnescapeDataString converts an escaped URI data string to its original form.
func UnescapeDataString(str string) (string, error) {
	return url.QueryUnescape(str)
}
