package main

import (
	"testing"
	"unique"
)

func Test_unique(t *testing.T) {
	empty1 := unique.Make("")
	empty2 := unique.Make("")
	value := unique.Make("abcdefg")

	if empty1 != empty2 {
		t.Errorf("empty1 != empty2")
	}

	if empty1 == value {
		t.Errorf("empty1 == value")
	}
}

func Test_unique2(t *testing.T) {
	var (
		zeroNonMake unique.Handle[User]
		zero        = unique.Make(User{})
	)

	if zeroNonMake == zero {
		t.Errorf("zeroNonMake == zero")
	}

	u1l := unique.Make(User{ID: 1, Name: "Alice"})
	u1r := unique.Make(User{ID: 1, Name: "Alice"})
	if u1l != u1r {
		t.Errorf("u1l != u1r")
	}

	u1lp := unique.Make(&User{ID: 1, Name: "Alice"})
	u1rp := unique.Make(&User{ID: 1, Name: "Alice"})
	if u1lp == u1rp {
		t.Errorf("u1lp == u1rp")
	}
}
