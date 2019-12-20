package fileserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateHandler(t *testing.T) {
	handler := CreateHandler("/")
	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := http.Get(server.URL)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}
}
