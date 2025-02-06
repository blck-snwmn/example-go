package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Server(t *testing.T) {
	sv := httptest.NewServer(genHandler())
	defer sv.Close()

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
		resp, err := sv.Client().Get(sv.URL + tt.path)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

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
