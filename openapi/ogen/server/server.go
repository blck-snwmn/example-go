package server

import (
	"context"

	"github.com/blck-snwmn/example-go/openapi/ogen/gen"
)

func NewServer(
	repository *UserRepository,
) (*gen.Server, error) {
	return gen.NewServer(&server{repository})
}

type server struct {
	repository *UserRepository
}

// CreateUser implements gen.Handler.
func (s *server) CreateUser(ctx context.Context, req *gen.User) (gen.CreateUserRes, error) {
	if _, err := s.repository.GetUserById(req.ID); err == nil {
		return &gen.CreateUserBadRequest{}, nil
	}
	var age *int32
	if v, ok := req.Age.Get(); ok {
		age = &v
	}
	s.repository.AddUser(user{
		ID:    req.ID,
		Name:  req.Name,
		Email: req.Email,
		Age:   age,
	})
	return &gen.CreateUserCreated{}, nil
}

// GetEmployees implements gen.Handler.
func (s *server) GetEmployees(ctx context.Context, params gen.GetEmployeesParams) (gen.GetEmployeesRes, error) {
	// Note: Return the value of `oneof` as a structure
	switch params.EmployeeID {
	case "1":
		resOk := gen.NewManagerGetEmployeesOK(gen.Manager{
			ID:         "1",
			Department: "Engineering",
			Email:      "em@example.com",
			Name:       "Alice",
			Teams:      []string{"team1", "team2"},
		})
		return &resOk, nil
	case "2":
		resOk := gen.NewMemberGetEmployeesOK(gen.Member{
			Department: "Sales",
			Email:      "sales@example.com",
			ID:         "2",
			Name:       "Bob",
			Team:       "team_sales",
		})
		return &resOk, nil
	default:
		// FIXME return 404
		return &gen.GetEmployeesBadRequest{}, nil
	}
}

// GetUserById implements gen.Handler.
func (s *server) GetUserById(ctx context.Context, params gen.GetUserByIdParams) (gen.GetUserByIdRes, error) {
	u, err := s.repository.GetUserById(params.UserID)
	if err != nil {
		return &gen.GetUserByIdNotFound{}, nil
	}
	resUser := &gen.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
	if u.Age != nil {
		resUser.Age = gen.NewOptNilInt32(*u.Age)
	}
	return resUser, nil
}

// GetUsers implements gen.Handler.
func (s *server) GetUsers(ctx context.Context) (gen.GetUsersRes, error) {
	users := s.repository.ListUsers()
	res := make(gen.Users, 0, len(users))
	for _, u := range users {
		gu := gen.User{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		}
		if u.Age != nil {
			gu.Age = gen.NewOptNilInt32(*u.Age)
		}
		res = append(res, gu)
	}
	return &res, nil
}
