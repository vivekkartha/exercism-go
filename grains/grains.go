package grains

import (
	"errors"
)

//Square gives the number of grains in a row
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("Does not exist")
	}
	return 1 << (uint64(n) - 1), nil
}

//Total returns the total grains in all squares
func Total() uint64 {
	return (1 << 64) - 1
}
