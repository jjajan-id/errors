package errors

import (
	def "errors"
	"fmt"
	"strings"
)

// New is just a wrapper for default package errors.New
func New(text string) error {
	return def.New(text)
}

// As is just a wrapper for default package errors.As
func As(err error, target interface{}) bool {
	return def.As(err, target)
}

// Is is just a wrapper for default package errors.Is
func Is(err error, target error) bool {
	return def.Is(err, target)
}

// Unwrap is just a wrapper for default package errors.Unwrap
func Unwrap(err error) error {
	return def.Unwrap(err)
}

// RootCause will get the first error
func RootCause(err error) error {
	root := err

	temp := err

	for {
		temp = def.Unwrap(temp)

		if temp == nil {
			break
		}

		root = temp
	}

	return root
}

// Traces will print the wrapped error stack
func Traces(err error) error {
	traces := []string{err.Error()}

	temp := err
	i := 0
	for {
		i = i + 1
		temp = def.Unwrap(temp)

		if temp == nil {
			break
		}

		pre := ""
		for j := 0; j < i; j++ {
			pre = fmt.Sprintf("%s ", pre)
		}

		traces = append(traces, fmt.Sprintf("%s-%s", pre, temp.Error()))
	}

	return def.New(strings.Join(traces, "\n"))
}
