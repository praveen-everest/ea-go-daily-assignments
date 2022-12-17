package main

import "errors"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

var DivisionByZeroError = errors.New("division by zero is not allowed")

func div(x, y float64) (float64, error) {
	if y == 0 {
		return 0, DivisionByZeroError
	}
	return x / y, nil
}
