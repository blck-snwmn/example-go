package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var url string

func TestMain(m *testing.M) {
	if os.Getenv("LOCAL") != "" {
		url = "http://localhost:8080"
	} else {
		sv := httptest.NewServer(GenHandler())
		defer sv.Close()

		url = sv.URL
	}

	m.Run()
}

func Test_Server(t *testing.T) {

	tests := []struct {
		path   string
		expect string
	}{
		{
			path:   "/hello",
			expect: "Hello, World!",
		},
	}

	for _, tt := range tests {
		resp, err := http.Get(url + tt.path)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close() //nolint:errcheck // Closing response body on defer is standard practice

		if resp.StatusCode != http.StatusOK {
			t.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		if string(b) != tt.expect {
			t.Errorf("unexpected response: %s", string(b))
		}
	}
}
