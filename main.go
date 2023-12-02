package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// RProxy struct represents the in-house reverse proxy
type RProxy struct {
	// Start with single target that can be extended to multiple target servers
	targetURL *url.URL
	client    *http.Client
}

// ServeHTTP handles incoming HTTP requests and redirects them to the target server(s)
func (rp *RProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := r.Clone(r.Context())
	req.Host = rp.targetURL.Host
	req.URL = rp.targetURL
	req.RequestURI = ""

	// Perform the actual request
	resp, err := rp.client.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error redirecting request: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Construct the response to the client: header, status code, and body
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error copying response body: %v", err), http.StatusInternalServerError)
		return
	}
}

// NewRProxy creates a new reverse proxy instance
func NewRProxy(targetURL string) (*RProxy, error) {
	target, err := url.Parse(targetURL)
	if err != nil {
		fmt.Printf("Error parsing target URL: %v\n", err)
		return nil, err
	}

	return &RProxy{
		targetURL: target,
		client:    &http.Client{},
	}, nil
}

func main() {
	targetURL := "http://localhost:8888"
	rproxy, err := NewRProxy(targetURL)
	if err != nil {
		fmt.Printf("Error creating reverse proxy: %v\n", err)
		return
	}

	// Construct the middleware server to use the RProxy
	server := http.Server{
		Addr:    ":8080",
		Handler: rproxy,
	}

	fmt.Println("Reverse Proxy listening on :8080...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server on :8080: %v\n", err)
	}
}
