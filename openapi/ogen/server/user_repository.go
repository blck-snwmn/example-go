package server

import (
	"cmp"
	"errors"
	"slices"
	"strings"
	"sync"
)

var (
	errUserNotFound = errors.New("user not found")
	errUserExists   = errors.New("user already exists")
)

type user struct {
	ID    string
	Name  string
	Email string
	Age   *int32
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: map[string]user{},
	}
}

type UserRepository struct {
	users map[string]user
	mx    sync.RWMutex
}

func (r *UserRepository) GetUserById(userId string) (user, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	u, ok := r.users[userId]
	if !ok {
		return user{}, errUserNotFound
	}

	return u, nil
}

func (r *UserRepository) ListUsers() []user {
	r.mx.RLock()
	defer r.mx.RUnlock()

	users := make([]user, 0, len(r.users))
	for _, u := range r.users {
		users = append(users, u)
	}
	slices.SortFunc(users, func(a, b user) int {
		return cmp.Or(
			strings.Compare(a.ID, b.ID),
		)
	})

	return users
}

func (r *UserRepository) AddUser(u user) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	if _, ok := r.users[u.ID]; ok {
		return errUserExists
	}

	r.users[u.ID] = u

	return nil
}
