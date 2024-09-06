package main

import (
	"fmt"
	"iter"
	"slices"

	"maps"
)

func main() {
	// maps/slices
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	// The order of the keys is undefined
	for k, v := range m {
		fmt.Printf("map: %v: %v\n", k, v)
	}

	// The order of the keys is undefined
	keys := maps.Keys(m)
	for k := range keys {
		fmt.Printf("maps.Keys: %v: %v\n", k, m[k])
	}

	// The order of the keys is sorted.
	sorted := slices.Sorted(keys)
	for _, k := range sorted {
		fmt.Printf("slices.Sorted: %v: %v\n", k, m[k])
	}

	// implement my iter
	numbers := generate(100)

	odd := filter(numbers, func(i int) bool { return i%2 == 1 })
	repeatByOdd := flatMap(odd, func(i int) iter.Seq[int] {
		return func(yield func(int) bool) {
			for j := 0; j < i; j++ {
				if !yield(j * i) {
					return
				}
			}
		}
	})

	evenTwices := filterMap(numbers, func(i int) (int, bool) {
		if i%2 == 0 {
			return i * 2, true
		}
		return 0, false
	})
	doubleTwice := mapf(evenTwices, func(i int) int { return i * 2 })

	concated := concat(repeatByOdd, doubleTwice)

	for i := range take(skip(concated, 5), 20) {
		fmt.Println(i)
	}

	fmt.Printf("Any odd number: %v\n", anyf(repeatByOdd, func(i int) bool { return i%2 == 1 }))
	fmt.Printf("All even number: %v\n", allf(repeatByOdd, func(i int) bool { return i%2 == 0 }))

	v, ok := findMap(numbers, func(i int) (int, error) {
		if i == 10 {
			return i, nil
		}
		return 0, fmt.Errorf("not found")
	})
	fmt.Printf("Find 10: %v, %v\n", v, ok)
}

func mapf[T, S any](seq iter.Seq[T], f func(T) S) iter.Seq[S] {
	return func(yield func(S) bool) {
		seq(func(v T) bool {
			return yield(f(v))
		})
	}
}

func filter[T any](seq iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		seq(func(v T) bool {
			if f(v) {
				return yield(v)
			}
			return true
		})
	}
}

func flatMap[T, S any](seq iter.Seq[T], f func(T) iter.Seq[S]) iter.Seq[S] {
	return func(yield func(S) bool) {
		for vs := range mapf(seq, f) {
			for v := range vs {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func filterMap[T, S any](seq iter.Seq[T], f func(T) (S, bool)) iter.Seq[S] {
	return func(yield func(S) bool) {
		seq(func(v T) bool {
			resultValue, ok := f(v)
			if ok {
				return yield(resultValue)
			}
			return true
		})
	}
}

func findMap[T, S any](seq iter.Seq[T], f func(T) (S, error)) (S, bool) {
	for i := range seq {
		v, err := f(i)
		if err == nil {
			return v, true
		}
	}
	var s S
	return s, false
}

type concated[L, R any] struct {
	left  L
	right R
}

func concat[L, R any](lseq iter.Seq[L], rseq iter.Seq[R]) iter.Seq[concated[L, R]] {
	// iter.Pull example
	return func(yield func(concated[L, R]) bool) {
		lv, lstop := iter.Pull(lseq)
		defer lstop()

		rv, rstop := iter.Pull(rseq)
		defer rstop()

		for {
			l, lmore := lv()
			r, rmore := rv()

			if !lmore || !rmore {
				// if either of the sequence is done, then we are done
				return
			}

			if !yield(concated[L, R]{l, r}) {
				return
			}
		}
	}
}

func take[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		count := 0
		seq(func(v T) bool {
			if count == n {
				return false
			}
			count++
			return yield(v)
		})
	}
}

func skip[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		count := 0
		seq(func(v T) bool {
			if count < n {
				count++
				return true
			}
			return yield(v)
		})
	}
}

func anyf[T any](seq iter.Seq[T], f func(T) bool) bool {
	for i := range seq {
		if f(i) {
			return true
		}
	}
	return false
}

func allf[T any](seq iter.Seq[T], f func(T) bool) bool {
	for i := range seq {
		if !f(i) {
			return false
		}
	}
	return true
}

func generate(end int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range end {
			if !yield(i) {
				return
			}
		}
	}
}
