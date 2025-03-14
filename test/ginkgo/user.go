package ginkgo

import (
	"context"
	"errors"
	"time"
)

// User is a struct representing user information
type User struct {
	ID       int
	Name     string
	Email    string
	Age      int
	IsActive bool
}

// UserRepository is an interface responsible for user data persistence
type UserRepository interface {
	FindByID(ctx context.Context, id int) (*User, error)
	Save(ctx context.Context, user *User) error
}

// UserService is a service that provides business logic for user-related operations
type UserService struct {
	repo UserRepository
}

// NewUserService creates a new UserService
func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// GetUser retrieves a user with the specified ID
func (s *UserService) GetUser(ctx context.Context, id int) (*User, error) {
	// Check for context timeout
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	return s.repo.FindByID(ctx, id)
}

// ActivateUser sets a user to active status
func (s *UserService) ActivateUser(ctx context.Context, id int) error {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if user.IsActive {
		return errors.New("user is already active")
	}

	user.IsActive = true
	return s.repo.Save(ctx, user)
}

// ValidateAge validates if an age is valid
// Age must be between 0 and 120 inclusive
func (s *UserService) ValidateAge(age int) bool {
	return age >= 0 && age <= 120
}

// ProcessUserDataAsync processes user data asynchronously
// This function sends the result to a channel when processing is complete
func (s *UserService) ProcessUserDataAsync(ctx context.Context, id int) <-chan *User {
	resultCh := make(chan *User, 1)

	go func() {
		defer close(resultCh)

		// Simulate a time-consuming process
		select {
		case <-time.After(100 * time.Millisecond):
			user, err := s.repo.FindByID(ctx, id)
			if err == nil && user != nil {
				resultCh <- user
			}
		case <-ctx.Done():
			// Do nothing if the context is canceled
			return
		}
	}()

	return resultCh
}
