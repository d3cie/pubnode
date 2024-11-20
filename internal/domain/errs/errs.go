package errs

import (
	"fmt"
)

type Error interface {
	Error() string
}

type error struct {
	Status  int
	Message string
}

func (e *error) Error() string {
	return fmt.Sprintf("domain error - status: %d, message: %s", e.Status, e.Message)
}

func new(status int, message string) Error {
	return &error{
		Status:  status,
		Message: message,
	}
}
