package exportexample_test

import (
	"testing"

	"github.com/blck-snwmn/example-go/test/exportexample"
)

var XX = exportexample.X

func Test_x(t *testing.T) {
	if XX != "Hello, World" {
		t.Errorf("XX is not Hello, World")
	}
}

func Test_Y(t *testing.T) {
	if exportexample.S.Y() != "YYYYY" {
		t.Errorf("Y is not YYYYY")
	}
}
