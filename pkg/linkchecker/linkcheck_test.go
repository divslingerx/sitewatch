package linkchecker

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestCheckURL_GoodURL(t *testing.T) {
	// Create a test server that always responds with 200 OK
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create a WaitGroup and a channel to use in the test
	wg := new(sync.WaitGroup)
	ch := make(chan URLStatus)

	// Add one to the WaitGroup counter and start the checkURL goroutine
	wg.Add(1)
	go CheckURL(server.URL, wg, ch)

	// Start another goroutine to wait for the checkURL goroutine to finish and close the channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Read the URL status from the channel
	urlStatus := <-ch

	// Check that the URL and status are correct
	if urlStatus.URL != server.URL {
		t.Errorf("got URL %s, want %s", urlStatus.URL, server.URL)
	}
	if urlStatus.Status != "🟢" {
		t.Errorf("got status %s, want 🟢", urlStatus.Status)
	}
}

func TestCheckURL_BadURL(t *testing.T) {
	// Create a test server that always responds with 404 Not Found
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	// Create a WaitGroup and a channel to use in the test
	wg := new(sync.WaitGroup)
	ch := make(chan URLStatus)

	// Add one to the WaitGroup counter and start the checkURL goroutine
	wg.Add(1)
	go CheckURL(server.URL, wg, ch)

	// Start another goroutine to wait for the checkURL goroutine to finish and close the channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Read the URL status from the channel
	urlStatus := <-ch

	// Check that the URL and status are correct
	if urlStatus.URL != server.URL {
		t.Errorf("got URL %s, want %s", urlStatus.URL, server.URL)
	}
	if urlStatus.Status != "🔴" {
		t.Errorf("got status %s, want 🔴", urlStatus.Status)
	}
}
