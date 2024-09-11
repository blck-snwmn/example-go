// Code generated by ogen, DO NOT EDIT.

package gen

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s GetEmployeesOK) Validate() error {
	switch s.Type {
	case MemberGetEmployeesOK:
		return nil // no validation needed
	case ManagerGetEmployeesOK:
		if err := s.Manager.Validate(); err != nil {
			return err
		}
		return nil
	default:
		return errors.Errorf("invalid type %q", s.Type)
	}
}

func (s *Manager) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if s.Teams == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "teams",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s Users) Validate() error {
	alias := ([]User)(s)
	if alias == nil {
		return errors.New("nil is invalid value")
	}
	return nil
}