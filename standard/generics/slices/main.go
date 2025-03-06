package main

import (
	"fmt"
	"math/rand/v2"
)

func randSlice[T any](slice []T) T {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice[0]
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(randSlice(slice))

	slice2 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(randSlice(slice2))

	type Person struct {
		Name string
		Age  int
	}

	slice3 := []Person{{Name: "John", Age: 20}, {Name: "Jane", Age: 21}, {Name: "Jim", Age: 22}}
	fmt.Println(randSlice(slice3))
}
