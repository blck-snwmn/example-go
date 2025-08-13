package main

import (
	"errors"
	"fmt"
)

type Database interface {
	GetUsers() ([]string, error)
	AddUser(name string) error
}

type memoryDatabase struct {
	users []string
}

func (m *memoryDatabase) GetUsers() ([]string, error) {
	return m.users, nil
}

func (m *memoryDatabase) AddUser(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	m.users = append(m.users, name)
	return nil
}

func NewDatabase(dsn string) Database {
	fmt.Printf("Creating database with DSN: %s\n", dsn)
	return &memoryDatabase{
		users: []string{"alice", "bob", "charlie"},
	}
}

type DSN string

func ProvideDSN() DSN {
	return DSN("memory://default")
}