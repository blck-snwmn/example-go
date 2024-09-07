package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var svURL string

func TestMain(m *testing.M) {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All users"))
	})
	mux.HandleFunc("GET /users/{name}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User name: "))
		w.Write([]byte(r.PathValue("name")))
	})
	mux.HandleFunc("GET /wild/{path...}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.PathValue("path")))
	})
	sv := httptest.NewServer(mux)
	svURL = sv.URL

	m.Run()
}

func TestGreet(t *testing.T) {
	t.Run("GET /greet", func(t *testing.T) {
		resp, err := http.Get(svURL + "/greet")
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("status code: %d", resp.StatusCode)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != "Hello, World!" {
			t.Errorf("body: %s", body)
		}
	})
	t.Run("POST /greet", func(t *testing.T) {
		resp, err := http.Post(svURL+"/greet", "", nil)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("status code: %d", resp.StatusCode)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != "Hello, World!" {
			t.Errorf("body: %s", body)
		}
	})
}

func TestUsers(t *testing.T) {
	t.Run("GET /users", func(t *testing.T) {
		resp, err := http.Get(svURL + "/users")
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("status code: %d", resp.StatusCode)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != "All users" {
			t.Errorf("body: %s", body)
		}
	})
	t.Run("POST /users", func(t *testing.T) {
		resp, err := http.Post(svURL+"/users", "", nil)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("status code: %d", resp.StatusCode)
		}
	})
}

func TestUser(t *testing.T) {
	t.Run("GET /users/alice", func(t *testing.T) {
		resp, err := http.Get(svURL + "/users/alice")
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("status code: %d", resp.StatusCode)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != "User name: alice" {
			t.Errorf("body: %s", body)
		}
	})
	t.Run("GET /users/bob", func(t *testing.T) {
		resp, err := http.Get(svURL + "/users/bob")
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("status code: %d", resp.StatusCode)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != "User name: bob" {
			t.Errorf("body: %s", body)
		}
	})
}

func TestWild(t *testing.T) {
	t.Run("GET /wild/1/2/3", func(t *testing.T) {
		resp, err := http.Get(svURL + "/wild/1/2/3")
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("status code: %d", resp.StatusCode)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != "1/2/3" {
			t.Errorf("body: %s", body)
		}
	})
	t.Run("GET /wild/1/2/3/4", func(t *testing.T) {
		resp, err := http.Get(svURL + "/wild/1/2/3/4")
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("status code: %d", resp.StatusCode)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != "1/2/3/4" {
			t.Errorf("body: %s", body)
		}
	})
}
