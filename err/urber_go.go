package err

import "errors"

// package foo

var ErrCouldNotOpen = errors.New("could not open")

func Open() error {
	return ErrCouldNotOpen
}

// package bar

func F1() {
	if err := Open(); err != nil {
		if errors.Is(err, ErrCouldNotOpen) {
			// handle the error
		} else {
			panic("unknown error")
		}
	}
}
