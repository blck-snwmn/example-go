package server

import (
	"cmp"
	"context"
	"slices"
	"strings"

	"github.com/blck-snwmn/example-go/openapi/oapi-codegen/strict/gen"
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

func NewServer() gen.StrictServerInterface {
	return &server{}
}

type server struct{}

// GetEmployees implements gen.StrictServerInterface.
func (s *server) GetEmployees(ctx context.Context, request gen.GetEmployeesRequestObject) (gen.GetEmployeesResponseObject, error) {
	// switch request.EmployeeId {
	// case "1":
	// 	_ = json.NewEncoder(w).Encode(gen.Manager{
	// 		Department: "Engineering",
	// 		Email:      "em@example.com",
	// 		Id:         "1",
	// 		Name:       "Alice",
	// 		Teams:      []string{"team1", "team2"},
	// 	})
	// 	return gen.GetEmployees200JSONResponse{}, nil
	// case "2":
	// 	m := gen.Member{
	// 		Department: "Sales",
	// 		Email:      "sales@example.com",
	// 		Id:         "2",
	// 		Name:       "Bob",
	// 		Team:       "team_sales",
	// 	}
	// 	raw, err := json.Marshal(m)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	return gen.GetEmployees200JSONResponse{}, nil
	// default:
	// }
	//
	// uinon in gen.GetEmployees200JSONResponse{} is Private field
	return gen.GetEmployees400Response{}, nil
}

// GetUserById implements gen.StrictServerInterface.
func (s *server) GetUserById(ctx context.Context, request gen.GetUserByIdRequestObject) (gen.GetUserByIdResponseObject, error) {
	user, ok := users[request.UserId]
	if !ok {
		return gen.GetUserById404Response{}, nil
	}
	result := gen.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Age:   nullable.NewNullableWithValue(*user.Age),
	}
	return gen.GetUserById200JSONResponse(result), nil
}

// GetUsers implements gen.StrictServerInterface.
func (s *server) GetUsers(ctx context.Context, request gen.GetUsersRequestObject) (gen.GetUsersResponseObject, error) {
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

	return gen.GetUsers200JSONResponse(result), nil
}

var _ gen.StrictServerInterface = (*server)(nil)
