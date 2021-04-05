package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

//func TestWithinTime(t *testing.T) {
//	fastServer := createServer(7 * time.Millisecond)
//	slowServer := createServer(9 * time.Millisecond)
//
//	defer slowServer.Close()
//	defer fastServer.Close()
//
//	slowURL := slowServer.URL
//	fastURL := fastServer.URL
//
//	want := fastURL
//	got, err := CustomRacer(slowURL, fastURL, time.Millisecond*10)
//
//	if err != nil {
//		t.Fatalf("Expected no error got: %q", err)
//	}
//	if got != want {
//		t.Errorf("GOT %q, want %q", got, want)
//	}
//}

func TestTimeout(t *testing.T) {
	server1 := createServer(11 * time.Millisecond)
	server2 := createServer(12 * time.Millisecond)

	defer server1.Close()
	defer server2.Close()

	_, err := CustomRacer(server1.URL, server2.URL, time.Millisecond*10)

	if err == nil {
		t.Errorf("Expected error but didn't get one")
	}
}

func createServer(resTime time.Duration) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(resTime)
			w.WriteHeader(http.StatusOK)
		}),
	)
}
