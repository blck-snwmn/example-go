package standard

import (
	"fmt"
	"testing"
)

func Test_Subtest(t *testing.T) {
	t.Skip("Skipping this test as it contains failing subtests as examples")
	ok := t.Run("sub test 1", func(t *testing.T) {
		fmt.Println("Test_Subtest: sub test 1")
		t.Fatal()
	})
	fmt.Printf("result: %v\n", ok)

	if t.Failed() {
		fmt.Println("Test_Subtest: sub test 1 failed")
	}
	ok2 := t.Run("sub test 2", func(t *testing.T) {
		if t.Failed() {
			fmt.Println("Test_Subtest: sub test 2 in sub test 1 failed")
		}
		fmt.Println("Test_Subtest: sub test 2")
	})
	fmt.Printf("result: %v\n", ok2)
}

func Test_Subtest_runStep(t *testing.T) {
	t.Skip("Skipping this test as it contains failing subtests as examples")

	runStep(t, "sub test 1", func(t *testing.T) {
		fmt.Println("Test_Subtest_runStep: sub test 1")
		t.Fatalf("Test_Subtest_runStep: sub test 1 failed")
	})
	runStep(t, "sub test 2", func(t *testing.T) {
		fmt.Println("Test_Subtest_runStep: sub test 2")
	})
}

func runStep(t *testing.T, name string, fn func(t *testing.T)) bool {
	ok := t.Run(name, fn)
	if !ok {
		t.Fatalf("step '%s' failed", name)
	}
	return ok
}
