package sample

import (
	"fmt"

	"golang.org/x/xerrors"
)

func sample() {
	err := xerrors.New("sample")
	err = fmt.Errorf("errors: %w", err)
	fmt.Println(err)
}
