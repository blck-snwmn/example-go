package main

import (
	"cmp"
	"encoding/json"
	"log/slog"
	"maps"
	"net/http"
	"slices"
	"sync"
	"time"

	"github.com/blck-snwmn/example-go/test/runn/api"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type User struct {
	ID   string
	Name string
}

var store = make(map[string]User)

type badserver struct {
	mux sync.Mutex
}

// CreateReport implements api.ServerInterface.
func (s *badserver) CreateReport(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetReportById implements api.ServerInterface.
func (s *badserver) GetReportById(w http.ResponseWriter, r *http.Request, reportId string) {
	panic("unimplemented")
}

// GetReports implements api.ServerInterface.
func (s *badserver) GetReports(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
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
			Id   string `json:"idx"`
			Name string `json:"name"`
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
	json.NewEncoder(w).Encode(users) //nolint:errcheck,gosec // HTTP response encode errors aren't useful
}

func NewBadServer() api.ServerInterface {
	return &badserver{}
}

func main() {
	srv := NewBadServer()
	r := chi.NewMux()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			trace := r.Header.Get("X-Runn-Trace")
			slog.Info("access",
				slog.String("trace", trace),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
			)
			next.ServeHTTP(w, r)
		})
	})

	h := api.HandlerFromMux(srv, r)

	s := &http.Server{
		Handler:           h,
		Addr:              "0.0.0.0:8080",
		ReadHeaderTimeout: 5 * time.Second,
	}

	s.ListenAndServe() //nolint:errcheck,gosec // Error is handled by log.Fatal in production code
}
