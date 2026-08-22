package main

import (
	"fmt"
	"strings"
)

func main() {
	items := List[int]{1, 2, 3, 4}.
		Filter(func(value int) bool {
			return value%2 == 0
		}).
		Map(func(value int) string {
			return fmt.Sprintf("item-%d", value)
		}).
		FlatMap(func(value string) List[string] {
			return List[string]{value, strings.ToUpper(value)}
		})
	fmt.Println(items)

	config := Config{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("%s:%d\n", config.Host, config.Port)

	functions := Functions{
		Identity: identity,
	}
	fmt.Println(functions.Identity(27))
}

// List demonstrates methods with their own type parameters. Map and FlatMap
// can return a List whose element type differs from the receiver's type.
type List[T any] []T

func (list List[T]) Map[U any](convert func(T) U) List[U] {
	return list.Fold(make(List[U], 0, len(list)), func(result List[U], value T) List[U] {
		return append(result, convert(value))
	})
}

func (list List[T]) Filter(keep func(T) bool) List[T] {
	return list.Fold(make(List[T], 0, len(list)), func(result List[T], value T) List[T] {
		if keep(value) {
			return append(result, value)
		}
		return result
	})
}

func (list List[T]) FlatMap[U any](convert func(T) List[U]) List[U] {
	return list.Fold(List[U]{}, func(result List[U], value T) List[U] {
		return append(result, convert(value)...)
	})
}

func (list List[T]) Fold[U any](initial U, combine func(U, T) U) U {
	result := initial
	for _, value := range list {
		result = combine(result, value)
	}
	return result
}

// Config demonstrates nested field selectors as struct literal keys.
type Config struct {
	Server
}

type Server struct {
	Host string
	Port int
}

// Functions demonstrates inferring a generic function's type arguments inside
// a struct literal, an assignment context added in Go 1.27.
type Functions struct {
	Identity func(int) int
}

func identity[T any](value T) T {
	return value
}
