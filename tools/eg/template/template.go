package template

import (
	"fmt"
)

func before(s string) error {
	return fmt.Errorf("%s", s)
}

func after(s string) string {
	return s
}
