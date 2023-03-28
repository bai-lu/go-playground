// package foo
package err

import (
	"errors"
	"fmt"
)

type NotFoundError struct {
	File string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("file %q not found", e.File)
}

func Open2(file string) error {
	return &NotFoundError{File: file}
}

// package bar
func F2() {
	if err := Open2("testfile.txt"); err != nil {
		// var notFound
		if errors.As(err, new(*NotFoundError)) {
			print(true)
			// handle the error
		} else {
			panic("unknown error")
		}
	}
}
