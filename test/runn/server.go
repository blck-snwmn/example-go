package main

import (
	"cmp"
	"encoding/json"
	"maps"
	"net/http"
	"slices"
	"sync"
	"time"

	"github.com/blck-snwmn/example-go/test/runn/api"
	"github.com/google/uuid"
)

type User struct {
	ID   string
	Name string
}

var store = make(map[string]User)

type server struct {
	mux sync.Mutex
}

// GetHeavy implements api.ServerInterface.
func (s *server) GetHeavy(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	w.WriteHeader(http.StatusOK)
}

// CreateUser implements api.ServerInterface.
func (s *server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u api.CreateUser
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	s.mux.Lock()
	defer s.mux.Unlock()

	id := uuid.Must(uuid.NewV7()).String()
	store[id] = User{id, u.Name}

	w.WriteHeader(http.StatusCreated)
}

// GetUserById implements api.ServerInterface.
func (s *server) GetUserById(w http.ResponseWriter, r *http.Request, userId string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if user, ok := store[userId]; ok {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(api.User{
			Id:   user.ID,
			Name: user.Name,
		})
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

// GetUsers implements api.ServerInterface.
func (s *server) GetUsers(w http.ResponseWriter, r *http.Request) {
	s.mux.Lock()
	defer s.mux.Unlock()

	users := make([]api.User, 0, len(store))
	sortedUsers := slices.SortedFunc(maps.Values(store), func(l, r User) int {
		return cmp.Compare(l.Name, r.Name)
	})
	for _, user := range sortedUsers {
		users = append(users, api.User{
			Id:   user.ID,
			Name: user.Name,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func NewServer() api.ServerInterface {
	return &server{}
}

type badserver struct {
	mux sync.Mutex
}

// GetHeavy implements api.ServerInterface.
func (s *badserver) GetHeavy(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// CreateUser implements api.ServerInterface.
func (s *badserver) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u api.CreateUser
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	s.mux.Lock()
	defer s.mux.Unlock()

	id := uuid.Must(uuid.NewV7()).String()
	store[id] = User{id, u.Name}

	w.WriteHeader(http.StatusCreated)
}

// GetUserById implements api.ServerInterface.
func (s *badserver) GetUserById(w http.ResponseWriter, r *http.Request, userId string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if user, ok := store[userId]; ok {
		w.Header().Set("Content-Type", "application/json")
		type User struct {
			Id   string `json:"id"`
			Name string `json:"namex"`
		}
		_ = json.NewEncoder(w).Encode(User{
			Id:   user.ID,
			Name: user.Name,
		})
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

// GetUsers implements api.ServerInterface.
func (s *badserver) GetUsers(w http.ResponseWriter, r *http.Request) {
	s.mux.Lock()
	defer s.mux.Unlock()

	users := make([]api.User, 0, len(store))
	sortedUsers := slices.SortedFunc(maps.Values(store), func(l, r User) int {
		return cmp.Compare(l.Name, r.Name)
	})
	for _, user := range sortedUsers {
		users = append(users, api.User{
			Id:   user.ID,
			Name: user.Name,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func NewBadServer() api.ServerInterface {
	return &badserver{}
}
