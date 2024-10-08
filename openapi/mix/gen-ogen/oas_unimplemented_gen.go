// Code generated by ogen, DO NOT EDIT.

package genogen

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateUser implements createUser operation.
//
// Create a new user.
//
// POST /v1/users
func (UnimplementedHandler) CreateUser(ctx context.Context, req *User) (r CreateUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserById implements getUserById operation.
//
// Get user by id.
//
// GET /v1/users/{user_id}
func (UnimplementedHandler) GetUserById(ctx context.Context, params GetUserByIdParams) (r GetUserByIdRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUsers implements getUsers operation.
//
// Get all users.
//
// GET /v1/users
func (UnimplementedHandler) GetUsers(ctx context.Context) (r GetUsersRes, _ error) {
	return r, ht.ErrNotImplemented
}
