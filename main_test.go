package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBitcoinPrice(t *testing.T) {
	req, err := http.NewRequest("GET", "/bitcoin-price", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(bitcoinPrice)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("BitcoinPrice handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedPrefix := "Bitcoin price: "
	if !strings.HasPrefix(rr.Body.String(), expectedPrefix) {
		t.Errorf("BitcoinPrice handler returned unexpected body prefix: got %v want prefix %v",
			rr.Body.String(), expectedPrefix)
	}
}

// Similarly, write test functions for ethereumPrice and tetherPrice...
