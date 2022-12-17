package main

import (
	"errors"
	"fmt"
	"math"
)

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

const DegreeToRadians = math.Pi / 180

func sin(degree float64) float64 {
	return math.Sin(DegreeToRadians * degree)
}

func cos(degree float64) float64 {
	return math.Cos(DegreeToRadians * degree)
}

func tan(degree float64) float64 {
	return math.Tan(DegreeToRadians * degree)
}

var SqrtForNegativeValueError = errors.New("square root for negative numbers is not allowed")

func sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, SqrtForNegativeValueError
	}
	return math.Sqrt(value), nil
}

func main() {
	for i := 1; i <= 8; i++ {
		fmt.Println(math.Sqrt(float64(i)))
	}
}
