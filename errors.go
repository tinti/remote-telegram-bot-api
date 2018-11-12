package rbot

import (
	"fmt"
)

type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return e.s
}

type ErrorRemoteBot struct {
	s string
	i error
}

func (e *ErrorRemoteBot) Error() string {
	if e.i != nil {
		return fmt.Sprintf("%s: %s", e.s, e.i)
	}

	return e.s
}

func NewErrorRemoteBot(s string, e error) error {
	return &ErrorRemoteBot{s, e}
}
