package main

import "fmt"

type UserService interface {
	ListUsers() ([]string, error)
	CreateUser(name string) error
}

type userService struct {
	db Database
}

func (u *userService) ListUsers() ([]string, error) {
	fmt.Println("UserService: Listing users...")
	return u.db.GetUsers()
}

func (u *userService) CreateUser(name string) error {
	fmt.Printf("UserService: Creating user %s...\n", name)
	return u.db.AddUser(name)
}

func NewUserService(db Database) UserService {
	fmt.Println("Creating UserService with Database dependency")
	return &userService{db: db}
}