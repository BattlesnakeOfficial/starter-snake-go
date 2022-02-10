package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithServerID(t *testing.T) {

	// Create a simple test request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a simple test handler
	testHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "")
	}
	// Wrap with middleware
	testHandler = withServerID(testHandler)

	// create a http response recorder
	rr := httptest.NewRecorder()

	// invoke the test handler
	testHandler(rr, req)

	// ensure the middleware worked
	sh := rr.Result().Header.Get("Server")
	if sh != ServerID {
		t.Error("middleware did not set header")
	}
}
