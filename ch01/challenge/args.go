package main

import (
	"fmt"
	"strconv"
)

// Width is a positive number less than 250 flag.Value.
type Width struct {
	val *int
}

// String returns the string representation of the current value.
func (w *Width) String() string {
	if w.val == nil {
		return ""
	}

	return fmt.Sprintf("%d", *w.val)
}

// Set sets the value from command line.
func (w *Width) Set(val string) error {
	n, err := strconv.Atoi(val)
	if err != nil {
		return fmt.Errorf("bad number: %s", err)
	}

	if n < 0 || n > 250 {
		return fmt.Errorf("width %d out of range [1:250]", n)
	}

	*w.val = n
	return nil
}
