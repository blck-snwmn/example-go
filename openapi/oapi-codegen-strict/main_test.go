package main

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/blck-snwmn/example-go/openapi/oapi-codegen-strict/gen"
	"github.com/blck-snwmn/example-go/openapi/oapi-codegen-strict/server"
	"github.com/go-chi/chi/v5"
)

var svURL string

func TestMain(m *testing.M) {
	svi := server.NewServer()

	handler := gen.NewStrictHandler(svi, nil)

	r := chi.NewRouter()

	gen.HandlerFromMux(handler, r)
	sv := httptest.NewServer(r)
	svURL = sv.URL

	m.Run()
}

func Test_GetUsers(t *testing.T) {
	client, err := gen.NewClient(svURL)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.GetUsers(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("status code: %d", resp.StatusCode)
	}

	var users []gen.User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		t.Fatal(err)
	}
	if len(users) != 3 {
		t.Errorf("expected 3 users, got %d", len(users))
	}
}

func Test_GetUser(t *testing.T) {
	client, err := gen.NewClient(svURL)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.GetUserById(context.Background(), "1")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("status code: %d", resp.StatusCode)
	}

	var user gen.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		t.Fatal(err)
	}
	if user.Id != "1" {
		t.Errorf("expected user ID 1, got %s", user.Id)
	}
}

func Test_GetEmployees_Manager(t *testing.T) {
	t.Skip("uinon in gen.GetEmployees200JSONResponse{} is private")

	client, err := gen.NewClient(svURL)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.GetEmployees(context.Background(), "1")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("status code: %d", resp.StatusCode)
	}

	var manager gen.Manager
	err = json.NewDecoder(resp.Body).Decode(&manager)
	if err != nil {
		t.Fatal(err)
	}
	if manager.Id != "1" {
		t.Errorf("expected manager ID 1, got %s", manager.Id)
	}
}

func Test_GetEmployees_Member(t *testing.T) {
	t.Skip("uinon in gen.GetEmployees200JSONResponse{} is private")

	client, err := gen.NewClient(svURL)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.GetEmployees(context.Background(), "2")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("status code: %d", resp.StatusCode)
	}

	var member gen.Member
	err = json.NewDecoder(resp.Body).Decode(&member)
	if err != nil {
		t.Fatal(err)
	}
	if member.Id != "2" {
		t.Errorf("expected member ID 2, got %s", member.Id)
	}
}
