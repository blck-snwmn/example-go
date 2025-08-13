package main

import "fmt"

type Database interface {
	GetUsers() ([]string, error)
	AddUser(name string) error
}

type memoryDatabase struct {
	users []string
}

func (d *memoryDatabase) GetUsers() ([]string, error) {
	return d.users, nil
}

func (d *memoryDatabase) AddUser(name string) error {
	d.users = append(d.users, name)
	return nil
}

func NewDatabase(dsn DSN) (Database, error) {
	fmt.Printf("Initializing memory database with config: %s\n", dsn)
	db := &memoryDatabase{
		users: []string{"Alice", "Bob", "Charlie"},
	}
	return db, nil
}

type DSN string

func ProvideDSN() DSN {
	return DSN("memory://in-memory-db")
}