package gotestsum

import (
	"testing"
	"time"
)

func Test_Heavy(t *testing.T) {
	t.Parallel()
	t.Log("TestXxx1")
	time.Sleep(21 * time.Second)
}

func Test_Heavy2(t *testing.T) {
	t.Parallel()
	t.Log("TestXxx2")
	time.Sleep(22 * time.Second)
}

func Test_Heavy3(t *testing.T) {
	t.Parallel()
	t.Log("TestXxx3")
	time.Sleep(23 * time.Second)
}

func Test_Heavy4(t *testing.T) {
	t.Parallel()
	t.Log("TestXxx4")
	time.Sleep(24 * time.Second)
}

func Test_Heavy5(t *testing.T) {
	t.Parallel()
	t.Log("TestXxx5")
	time.Sleep(25 * time.Second)
}

func Test_Heavy6(t *testing.T) {
	t.Parallel()
	t.Log("TestXxx6")
	time.Sleep(26 * time.Second)
}

func Test_Heavy7(t *testing.T) {
	t.Parallel()
	t.Log("TestXxx7")
	time.Sleep(27 * time.Second)
}

func Test_Heavy8(t *testing.T) {
	t.Parallel()
	t.Log("TestXxx8")
	time.Sleep(28 * time.Second)
}

func Test_Heavy9(t *testing.T) {
	t.Parallel()
	t.Log("TestXxx9")
	time.Sleep(29 * time.Second)
}

func Test_Heavy10(t *testing.T) {
	t.Parallel()
	t.Log("TestXxx10")
	time.Sleep(30 * time.Second)
}

func Test_Heavy11(t *testing.T) {
	t.Log("TestXxx11")
	time.Sleep(31 * time.Second)
}

func Test_Heavy12(t *testing.T) {
	t.Parallel()
	t.Log("TestXxx12")
	time.Sleep(32 * time.Second)
}
