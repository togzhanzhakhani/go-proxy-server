package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
	"github.com/google/uuid"
)

// Request represents the incoming request structure
type Request struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

// Response represents the response structure to the client
type Response struct {
	ID      string            `json:"id"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Length  int               `json:"length"`
}

// Store for saving requests and responses
var store sync.Map

func handler(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate URL
	_, err := url.ParseRequestURI(req.URL)
	if err != nil {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	client := &http.Client{}
	proxyReq, err := http.NewRequest(req.Method, req.URL, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Set headers
	for key, value := range req.Headers {
		proxyReq.Header.Set(key, value)
	}

	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Failed to perform request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	// Generate unique request ID
	requestID := uuid.New().String()

	// Save the request and response to store
	store.Store(requestID, Response{
		ID:      requestID,
		Status:  resp.StatusCode,
		Headers: flattenHeaders(resp.Header),
		Length:  len(body),
	})

	// Create response to client
	clientResp := Response{
		ID:      requestID,
		Status:  resp.StatusCode,
		Headers: flattenHeaders(resp.Header),
		Length:  len(body),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientResp)
}

func flattenHeaders(headers http.Header) map[string]string {
	flat := make(map[string]string)
	for key, values := range headers {
		flat[key] = values[0]
	}
	return flat
}

func main() {
	http.HandleFunc("/proxy", handler)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
