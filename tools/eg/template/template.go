package template

import (
	"fmt"

	"golang.org/x/xerrors"
)

func before(err error) error { //nolint: unused
	return fmt.Errorf("errors: %w", err)
}

func after(err error) error { //nolint: unused
	return xerrors.Errorf("errors: %w", err)
}
