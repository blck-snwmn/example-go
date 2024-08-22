package main

import "fmt"

func main() {
	print(1)
	print("Hello, World!")

	print(Item[int]{ID: "1", Name: "Alice", Value: 1})
	print(Item[string]{ID: "1", Name: "Alice", Value: "xxxx"})

	books := []Book{
		{ID: "1", Title: "Alice"},
		{ID: "2", Title: "Bob"},
		{ID: "3", Title: "Charlie"},
	}
	print(Names(books))
}

func print[T any](t T) {
	fmt.Println(t)
}

type Item[Value any] struct {
	ID    string
	Name  string
	Value Value
}

type Book struct {
	ID    string
	Title string
}

func (b Book) Name() string {
	return b.Title
}

func Names[T Namer](items []T) []string {
	names := make([]string, 0, len(items))
	for _, item := range items {
		names = append(names, item.Name())
	}
	return names
}

type Namer interface {
	Name() string
}
