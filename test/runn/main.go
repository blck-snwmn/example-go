package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"log/slog"
	"maps"
	"net/http"
	"slices"
	"sync"

	"github.com/blck-snwmn/example-go/test/runn/gen"
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

// CreateUser implements gen.ServerInterface.
func (s *server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u gen.CreateUser
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

// GetUserById implements gen.ServerInterface.
func (s *server) GetUserById(w http.ResponseWriter, r *http.Request, userId string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if user, ok := store[userId]; ok {
		_ = json.NewEncoder(w).Encode(gen.User{
			Id:   user.ID,
			Name: user.Name,
		})
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

// GetUsers implements gen.ServerInterface.
func (s *server) GetUsers(w http.ResponseWriter, r *http.Request) {
	s.mux.Lock()
	defer s.mux.Unlock()

	users := make([]gen.User, 0, len(store))
	sortedUsers := slices.SortedFunc(maps.Values(store), func(l, r User) int {
		return cmp.Compare(l.Name, r.Name)
	})
	for _, user := range sortedUsers {
		users = append(users, gen.User{
			Id:   user.ID,
			Name: user.Name,
		})
	}
	_ = json.NewEncoder(w).Encode(users)
}

func NewServer() gen.ServerInterface {
	return &server{}
}

func main() {
	srv := NewServer()
	r := http.NewServeMux()

	gen.HandlerFromMux(srv, r)

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	slog.Info(fmt.Sprintf("listening on %s", s.Addr))

	s.ListenAndServe()
}
