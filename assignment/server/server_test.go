package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupAPI(t *testing.T) (url string, cleaner func()) {
	t.Helper()

	// Created server with newMux
	server := httptest.NewServer(newMux())
	
	// Assings the URL the server
	url = server.URL

	// Close the server
	cleaner = func() {
		server.Close()
	}

	return url, cleaner
}

func TestGetRoot(t *testing.T) {
	path := "/"
	expectedCode := http.StatusOK
	expectedContent := "Hello World!!"
	
	// Start server of test
	url, cleaner := setupAPI(t)
	defer cleaner()
	
	// Generates the request Get
	res, err := http.Get(url + path)
	if err != nil {
		t.Fatalf("Error in Get: %v", err)
	}
	defer res.Body.Close()
	
	// Validates status code
	if res.StatusCode != expectedCode {
		t.Errorf("code expected %d (%s), but received %d (%s)", expectedCode, http.StatusText(expectedCode), res.StatusCode, http.StatusText(res.StatusCode))
	}

	// Reads and validates the body content
	bodyBytes, _ := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("error reading response body: %v", err)
	}
	body := string(bodyBytes)

	if !strings.Contains(body, expectedContent) {
		t.Errorf("expected content %q, but received %q", expectedContent, body)
	}
}

func TestGetNotFound(t *testing.T) {
	path := "/no-exists"
	expectedCode := http.StatusNotFound
	expectedContent := "404"
	
	// Start server of test
	url, cleaner := setupAPI(t)
	defer cleaner()
	
	// Genetares the request Get
	res, err := http.Get(url + path)
	if err != nil {
		t.Fatalf("Error in Get: %v", err)
	}
	defer res.Body.Close()
	
	// Validates status code
	if res.StatusCode != expectedCode {
		t.Errorf("code expected %d (%s), but received %d (%s)", expectedCode, http.StatusText(expectedCode), res.StatusCode, http.StatusText(res.StatusCode))
	}

	// Reads and validates the body content
	bodyBytes, _ := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("error reading response body: %v", err)
	}
	body := string(bodyBytes)
	
	if !strings.Contains(body, expectedContent) {
		t.Errorf("expected content %q, but received %q", expectedContent, body)
	}
}