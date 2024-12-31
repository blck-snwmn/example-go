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
	"github.com/google/uuid"
)

type User struct {
	ID   string
	Name string
}

type Report struct {
	ID       string
	UserID   string
	Category string
	Title    string
	Content  string
}

var (
	storeUser   = make(map[string]User)
	storeReport = make(map[string]Report)
)

type server struct {
	mux sync.Mutex
}

// CreateReport implements api.ServerInterface.
func (s *server) CreateReport(w http.ResponseWriter, r *http.Request) {
	var report api.CreateReport
	if err := json.NewDecoder(r.Body).Decode(&report); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	s.mux.Lock()
	defer s.mux.Unlock()

	id := uuid.Must(uuid.NewV7()).String()
	storeReport[id] = Report{
		ID:       id,
		UserID:   report.UserId,
		Category: string(report.Category),
		Title:    report.Title,
		Content:  report.Content,
	}

	slog.Info("create report",
		slog.String("id", id),
		slog.String("user_id", report.UserId),
		slog.String("category", string(report.Category)),
		slog.String("title", report.Title),
		slog.String("content", report.Content),
	)

	w.WriteHeader(http.StatusCreated)
}

// GetReportById implements api.ServerInterface.
func (s *server) GetReportById(w http.ResponseWriter, r *http.Request, reportId string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if report, ok := storeReport[reportId]; ok {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(api.Report{
			Id:       report.ID,
			UserId:   report.UserID,
			Category: report.Category,
			Title:    report.Title,
			Content:  report.Content,
		})
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

// GetReports implements api.ServerInterface.
func (s *server) GetReports(w http.ResponseWriter, r *http.Request) {
	s.mux.Lock()
	defer s.mux.Unlock()

	reports := make([]api.Report, 0, len(storeReport))
	sortedReports := slices.SortedFunc(maps.Values(storeReport), func(l, r Report) int {
		return cmp.Compare(l.Title, r.Title)
	})
	for _, report := range sortedReports {
		reports = append(reports, api.Report{
			Id:       report.ID,
			UserId:   report.UserID,
			Category: report.Category,
			Title:    report.Title,
			Content:  report.Content,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
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
	storeUser[id] = User{id, u.Name}

	slog.Info("create user",
		slog.String("id", id),
		slog.String("name", u.Name),
	)

	w.WriteHeader(http.StatusCreated)
}

// GetUserById implements api.ServerInterface.
func (s *server) GetUserById(w http.ResponseWriter, r *http.Request, userId string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if user, ok := storeUser[userId]; ok {
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

	users := make([]api.User, 0, len(storeUser))
	sortedUsers := slices.SortedFunc(maps.Values(storeUser), func(l, r User) int {
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
	storeUser[id] = User{id, u.Name}

	w.WriteHeader(http.StatusCreated)
}

// GetUserById implements api.ServerInterface.
func (s *badserver) GetUserById(w http.ResponseWriter, r *http.Request, userId string) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if user, ok := storeUser[userId]; ok {
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

	users := make([]api.User, 0, len(storeUser))
	sortedUsers := slices.SortedFunc(maps.Values(storeUser), func(l, r User) int {
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
