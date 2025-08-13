package main

import "fmt"

type UserService interface {
	ListUsers() ([]string, error)
}

type userService struct {
	db Database
}

func (s *userService) ListUsers() ([]string, error) {
	fmt.Println("UserService: Getting users from database")
	return s.db.GetUsers()
}

func NewUserService(db Database) UserService {
	return &userService{db: db}
}