package standard

import "testing"

var counter = 0

func Test_Shuffle1(t *testing.T) {
	counter++
	t.Logf("Test_Shuffle1: %d", counter)
}

func Test_Shuffle2(t *testing.T) {
	counter++
	t.Logf("Test_Shuffle2: %d", counter)
}

func Test_Shuffle3(t *testing.T) {
	counter++
	t.Logf("Test_Shuffle3: %d", counter)
}

func Test_Shuffle4_subtest(t *testing.T) {
	// t.Parallel()
	t.Run("sub test 1", func(t *testing.T) {
		t.Log("Test_Shuffle4_subtest: sub test 1")
	})
	t.Run("sub test 2", func(t *testing.T) {
		t.Log("Test_Shuffle4_subtest: sub test 2")
	})
	t.Run("sub test 3", func(t *testing.T) {
		t.Log("Test_Shuffle4_subtest: sub test 3")
	})
}

func Test_Shuffle5_subtest_parallel(t *testing.T) {
	t.Parallel()
	t.Run("sub test 1", func(t *testing.T) {
		t.Parallel()
		t.Log("Test_Shuffle5_subtest_parallel: sub test 1")
	})
	t.Run("sub test 2", func(t *testing.T) {
		t.Parallel()
		t.Log("Test_Shuffle5_subtest_parallel: sub test 2")
	})
	t.Run("sub test 3", func(t *testing.T) {
		t.Parallel()
		t.Log("Test_Shuffle5_subtest_parallel: sub test 3")
	})
}
