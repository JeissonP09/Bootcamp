package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"
)

type testResponse struct {
	Results []struct {
		Task string `json:"task"`
	} `json:"results"`
	Date         time.Time `json:"date"`
	TotalResults int       `json:"total_results"`
}

func setupAPI(t *testing.T) (server *httptest.Server, clean func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "todo_test_*.json")
	if err != nil {
		t.Fatal(err)
	}

	server = httptest.NewServer(newMux(tmpFile.Name()))

	for i := 1; i <= 3; i++ {
		var body bytes.Buffer
		task := struct {
			Task string `json:"task"`
		}{
			Task: "Task " + strconv.Itoa(i),
		}

		if err := json.NewEncoder(&body).Encode(task); err != nil {
			t.Fatalf("encoding task %d: %v", i, err)
		}

		resp, err := http.Post(server.URL+"/todo", "application/json", &body)
		if err != nil {
			t.Fatalf("posting task %d: %v", i, err)
		}
		if resp.StatusCode != http.StatusCreated {
			t.Fatalf("expected status 201, got %d", resp.StatusCode)
		}
	}

	clean = func() {
		server.Close()
		os.Remove(tmpFile.Name())
	}
	return
}

func TestGet(t *testing.T) {
	srv, clean := setupAPI(t)
	defer clean()

	tests := []struct {
		name     string
		path     string
		expItems int
		expTask  string
	}{
		{
			name:     "GetAll",
			path:     "/todo",
			expItems: 3,
			expTask:  "Task 1",
		},
		{
			name:     "GetOne",
			path:     "/todo/1",
			expItems: 1,
			expTask:  "Task 2",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(srv.URL + tc.path)
			if err != nil {
				t.Fatalf("GET %s: %v", tc.path, err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("expected status 200, got %d", resp.StatusCode)
			}

			var data testResponse
			if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
				t.Fatalf("decoding response: %v", err)
			}

			if len(data.Results) != tc.expItems {
				t.Errorf("expected %d items, got %d", tc.expItems, len(data.Results))
			}

			if data.Results[0].Task != tc.expTask {
				t.Errorf("expected first task to be '%s', got '%s'", tc.expTask, data.Results[0].Task)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	srv, clean := setupAPI(t)
	defer clean()

	t.Run("Add", func(t *testing.T) {
		var body bytes.Buffer
		task := struct {
			Task string `json:"task"`
		}{
			Task: "Task 4",
		}

		if err := json.NewEncoder(&body).Encode(task); err != nil {
			t.Fatalf("encoding task: %v", err)
		}

		resp, err := http.Post(srv.URL+"/todo", "application/json", &body)
		if err != nil {
			t.Fatalf("POST error: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("expected 201 Created, got %d", resp.StatusCode)
		}
	})

	t.Run("CheckAdd", func(t *testing.T) {
		resp, err := http.Get(srv.URL + "/todo/3")
		if err != nil {
			t.Fatalf("GET error: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected 200 OK, got %d", resp.StatusCode)
		}

		var data testResponse
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			t.Fatalf("decoding response: %v", err)
		}

		if len(data.Results) != 1 {
			t.Fatalf("expected 1 result, got %d", len(data.Results))
		}

		if data.Results[0].Task != "Task 4" {
			t.Errorf("expected 'Task 4', got '%s'", data.Results[0].Task)
		}
	})
}
