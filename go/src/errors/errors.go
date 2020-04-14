package errors

import (
	"errors"
)

func NewError(s string) error {
	return errors.New(s)
}
