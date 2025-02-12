package server

import (
	"encoding/json"
	"net/http"

	"github.com/blck-snwmn/example-go/openapi/oapi-codegen/stdhttp/gen"
	"github.com/oapi-codegen/nullable"
)

func ptr[T any](v T) *T {
	return &v
}

func NewServer(repository *UserRepository) gen.ServerInterface {
	return &server{
		repository: repository,
	}
}

type server struct {
	repository *UserRepository
}

// CreateUser implements gen.ServerInterface.
func (s *server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u gen.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, err := s.repository.GetUserById(u.Id); err == nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	var age *int32
	tmp, err := u.Age.Get()
	if err == nil {
		age = ptr(tmp)
	}
	err = s.repository.AddUser(user{
		ID:    u.Id,
		Name:  u.Name,
		Email: u.Email,
		Age:   age,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetEmployees implements gen.ServerInterface.
func (s *server) GetEmployees(w http.ResponseWriter, r *http.Request, employeeId string) {
	switch employeeId {
	case "1":
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(gen.Manager{
			Department: "Engineering",
			Email:      "em@example.com",
			Id:         "1",
			Name:       "Alice",
			Teams:      []string{"team1", "team2"},
		})
	case "2":
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(gen.Member{
			Department: "Sales",
			Email:      "sales@example.com",
			Id:         "2",
			Name:       "Bob",
			Team:       "team_sales",
		})
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// GetUserById implements gen.ServerInterface.
func (s *server) GetUserById(w http.ResponseWriter, r *http.Request, userId string) {
	u, err := s.repository.GetUserById(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	age := nullable.NewNullNullable[int32]()
	if u.Age != nil {
		age.Set(*u.Age)
	}
	_ = json.NewEncoder(w).Encode(gen.User{
		Id:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Age:   age,
	})
}

// GetUsers implements gen.ServerInterface.
func (s *server) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := s.repository.ListUsers()

	var result gen.Users
	for _, user := range users {
		age := nullable.NewNullNullable[int32]()
		if user.Age != nil {
			age.Set(*user.Age)
		}
		result = append(result, gen.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Age:   age,
		})
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
}

var _ gen.ServerInterface = (*server)(nil)
