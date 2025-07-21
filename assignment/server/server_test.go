package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testCase struct {
	name            string
	path            string
	expectedCode    int
	expectedContent string
}

func setupAPI(t *testing.T) (url string, cleaner func()) {
	t.Helper()

	// Created server with newMux
	server := httptest.NewServer(newMux("/todo"))

	// Assigns the URL the server
	url = server.URL

	// Close the server
	cleaner = func() {
		server.Close()
	}

	return url, cleaner
}

func TestGet(t *testing.T) {
	cases := []testCase{
		{
			name:            "GET Root",
			path:            "/",
			expectedCode:    http.StatusOK,
			expectedContent: "Hello World!!",
		},
		{
			name:            "GET Not Found",
			path:            "/no-exists",
			expectedCode:    http.StatusNotFound,
			expectedContent: "404",
		},
	}

	// Start server of test
	url, cleaner := setupAPI(t)
	defer cleaner()

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := http.Get(url + tc.path)
			if err != nil {
				t.Fatalf("Error in GET: %v", err)
			}
			defer r.Body.Close()

			// Validates status code
			if r.StatusCode != tc.expectedCode {
				t.Errorf(
					"expected code %d (%s), but received %d (%s)",
					tc.expectedCode, http.StatusText(tc.expectedCode),
					r.StatusCode, http.StatusText(r.StatusCode),
				)
			}

			// Reads and validates the body content
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("error reading response body: %v", err)
			}
			body := string(bodyBytes)

			// Validates the type content
			switch r.Header.Get("Content-Type") {
			case "text/plain; charset=utf-8":
				if !strings.Contains(body, tc.expectedContent) {
					t.Errorf("content expected %q, received %q", tc.expectedContent, body)
				}
			default:
				t.Fatalf("Unsupported Content-Type: %q", r.Header.Get("Content-Type"))
			}
		})
	}
}
