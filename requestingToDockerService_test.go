package main

import (
	"io"
	"net/http"
	"testing"
)

func TestSendRequest(t *testing.T) {
	res, err := http.DefaultClient.Get("http://localhost:8000")
	if err != nil {
		t.Errorf("Error sending the request: %s", err)
		return
	}

	text, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading the response: %s", err)
		return
	}

	if string(text) != "Hello, World from a container" {
		t.Errorf("Unexpected response, received: %s", string(text))
		return
	}
}
