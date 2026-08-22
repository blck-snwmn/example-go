package main

import (
	"slices"
	"strconv"
	"testing"
)

func TestSeqFold(t *testing.T) {
	got := Values(1, 2, 3, 4).Fold(0, func(sum, value int) int {
		return sum + value
	})

	if want := 10; got != want {
		t.Errorf("Fold() = %d, want %d", got, want)
	}
}

func TestSeqMap(t *testing.T) {
	got := Values(1, 2, 3).Map(strconv.Itoa).Collect()
	want := []string{"1", "2", "3"}

	if !slices.Equal(got, want) {
		t.Errorf("Map() = %v, want %v", got, want)
	}
}

func TestSeqFilter(t *testing.T) {
	got := Values(1, 2, 3, 4).Filter(func(value int) bool {
		return value%2 == 0
	}).Collect()
	want := []int{2, 4}

	if !slices.Equal(got, want) {
		t.Errorf("Filter() = %v, want %v", got, want)
	}
}

func TestSeqFlatMap(t *testing.T) {
	got := Values(1, 2, 3).FlatMap(func(value int) Seq[int] {
		return Values(value, value*10)
	}).Collect()
	want := []int{1, 10, 2, 20, 3, 30}

	if !slices.Equal(got, want) {
		t.Errorf("FlatMap() = %v, want %v", got, want)
	}
}

func TestSeqIsLazy(t *testing.T) {
	converted := 0
	seq := Values(1, 2, 3).Map(func(value int) int {
		converted++
		return value * 2
	})

	if converted != 0 {
		t.Fatalf("Map() processed %d values before the sequence was consumed", converted)
	}

	seq.Collect()
	if converted != 3 {
		t.Errorf("Map() processed %d values, want 3", converted)
	}
}
