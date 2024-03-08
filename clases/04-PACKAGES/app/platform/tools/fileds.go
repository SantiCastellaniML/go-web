package tools

import (
	"errors"
	"fmt"
)

var (
	ErrFieldNotFound = errors.New("field not found")
)

type FieldError struct {
	Field string
	Msg   string
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

func CheckFieldExistence(fields map[string]any, requiredFields ...string) (err error) {
	for _, field := range requiredFields {
		if _, ok := fields[field]; !ok {
			err = &FieldError{
				Field: field,
				Msg:   "field not found",
			}
			return
		}
	}

	return
}
