// Code generated by ogen, DO NOT EDIT.

package genogen

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateUser implements createUser operation.
	//
	// Create a new user.
	//
	// POST /v1/users
	CreateUser(ctx context.Context, req *User) (CreateUserRes, error)
	// GetUserById implements getUserById operation.
	//
	// Get user by id.
	//
	// GET /v1/users/{user_id}
	GetUserById(ctx context.Context, params GetUserByIdParams) (GetUserByIdRes, error)
	// GetUsers implements getUsers operation.
	//
	// Get all users.
	//
	// GET /v1/users
	GetUsers(ctx context.Context) (GetUsersRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}