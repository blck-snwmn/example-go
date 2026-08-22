package main

import (
	"slices"
	"strconv"
	"testing"
)

func TestListFold(t *testing.T) {
	got := List[int]{1, 2, 3, 4}.Fold(0, func(sum, value int) int {
		return sum + value
	})

	if want := 10; got != want {
		t.Errorf("Fold() = %d, want %d", got, want)
	}
}

func TestListMap(t *testing.T) {
	got := List[int]{1, 2, 3}.Map(strconv.Itoa)
	want := List[string]{"1", "2", "3"}

	if !slices.Equal(got, want) {
		t.Errorf("Map() = %v, want %v", got, want)
	}
}

func TestListFilter(t *testing.T) {
	got := List[int]{1, 2, 3, 4}.Filter(func(value int) bool {
		return value%2 == 0
	})
	want := List[int]{2, 4}

	if !slices.Equal(got, want) {
		t.Errorf("Filter() = %v, want %v", got, want)
	}
}

func TestListFlatMap(t *testing.T) {
	got := List[int]{1, 2, 3}.FlatMap(func(value int) List[int] {
		return List[int]{value, value * 10}
	})
	want := List[int]{1, 10, 2, 20, 3, 30}

	if !slices.Equal(got, want) {
		t.Errorf("FlatMap() = %v, want %v", got, want)
	}
}
