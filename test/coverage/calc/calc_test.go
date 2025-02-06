package calc

import "testing"

func Test_Add(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Errorf("Add(1, 2) is not 3")
	}
}

func Test_Sub(t *testing.T) {
	if Sub(1, 2) != -1 {
		t.Errorf("Sub(1, 2) is not -1")
	}
}
