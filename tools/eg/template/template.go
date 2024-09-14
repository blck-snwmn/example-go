package template

import (
	"fmt"

	"golang.org/x/xerrors"
)

func before(err error) error {
	return fmt.Errorf("errors: %w", err)
}

func after(err error) error {
	return xerrors.Errorf("errors: %w", err)
}
