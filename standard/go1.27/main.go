package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

func main() {
	items := Values(1, 2, 3, 4).
		Filter(func(value int) bool {
			return value%2 == 0
		}).
		Map(func(value int) string {
			return fmt.Sprintf("item-%d", value)
		}).
		FlatMap(func(value string) Seq[string] {
			return Values(value, strings.ToUpper(value))
		}).
		Collect()
	fmt.Println(items)

	// Go 1.27 allows promoted fields to be used as keys in a struct literal.
	config := Config{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("%s:%d\n", config.Host, config.Port)

	// Go 1.27 infers identity[int] from the field's function type.
	functions := Functions{
		Identity: identity,
	}
	fmt.Println(functions.Identity(27))
}

// Seq is a lazy sequence that passes values of type T to a consumer.
type Seq[T any] iter.Seq[T]

// Values returns a sequence that yields the supplied values in order.
func Values[T any](values ...T) Seq[T] {
	return Seq[T](slices.Values(values))
}

// Map returns a sequence that converts each value from T to U.
func (seq Seq[T]) Map[U any](convert func(T) U) Seq[U] {
	return func(yield func(U) bool) {
		for value := range seq {
			if !yield(convert(value)) {
				return
			}
		}
	}
}

// Filter returns a sequence containing only values accepted by keep.
func (seq Seq[T]) Filter(keep func(T) bool) Seq[T] {
	return func(yield func(T) bool) {
		for value := range seq {
			if keep(value) && !yield(value) {
				return
			}
		}
	}
}

// FlatMap converts each value to a sequence and flattens the results.
func (seq Seq[T]) FlatMap[U any](convert func(T) Seq[U]) Seq[U] {
	return func(yield func(U) bool) {
		for value := range seq {
			for converted := range convert(value) {
				if !yield(converted) {
					return
				}
			}
		}
	}
}

// Fold consumes the sequence and combines its values into a single result.
func (seq Seq[T]) Fold[U any](initial U, combine func(U, T) U) U {
	result := initial
	for value := range seq {
		result = combine(result, value)
	}
	return result
}

// Collect consumes the sequence and returns its values as a slice.
func (seq Seq[T]) Collect() []T {
	return slices.Collect(iter.Seq[T](seq))
}

// Config embeds the server settings.
type Config struct {
	Server
}

// Server contains host and port settings.
type Server struct {
	Host string
	Port int
}

// Functions groups functions that operate on integers.
type Functions struct {
	Identity func(int) int
}

// identity returns value unchanged.
func identity[T any](value T) T {
	return value
}
