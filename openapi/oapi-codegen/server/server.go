package server

import (
	"cmp"
	"encoding/json"
	"net/http"
	"slices"
	"strings"

	"github.com/blck-snwmn/example-go/openapi/oapi-codegen/gen"
	"github.com/oapi-codegen/nullable"
)

func ptr[T any](v T) *T {
	return &v
}

type user struct {
	ID    string
	Name  string
	Email string
	Age   *int32
}

var users = map[string]user{
	"1": {ID: "1", Name: "Alice", Email: "alice@example.com", Age: ptr[int32](30)},
	"2": {ID: "2", Name: "Bob", Email: "bob@example.com", Age: ptr[int32](40)},
	"3": {ID: "3", Name: "Charlie", Email: "charlie@example.com", Age: ptr[int32](50)},
}

func NewServer() gen.ServerInterface {
	return &server{}
}

type server struct{}

// GetUserById implements ServerInterface.
func (s *server) GetUserById(w http.ResponseWriter, r *http.Request, userId string) {
	user, ok := users[userId]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	result := gen.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Age:   nullable.NewNullableWithValue(*user.Age),
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
}

// GetUsers implements ServerInterface.
func (s *server) GetUsers(w http.ResponseWriter, r *http.Request) {
	var result gen.Users
	for _, user := range users {
		result = append(result, gen.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Age:   nullable.NewNullableWithValue(*user.Age),
		})
	}
	slices.SortFunc(result, func(a, b gen.User) int {
		return cmp.Or(
			strings.Compare(a.Id, b.Id),
		)
	})

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
}

var _ gen.ServerInterface = (*server)(nil)
