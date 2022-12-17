package main

import (
	"math"
	"testing"
)

type input[K int | float64] struct {
	first  K
	second K
	want   K
}

func TestSuccessfulAddition(t *testing.T) {
	for _, tc := range []input[int]{
		{1, 2, 3},
		{10, 15, 25},
		{25, -25, 0},
	} {
		result := add(tc.first, tc.second)
		if result != tc.want {
			t.Errorf("addition of %d and %d should be %d but got %d.", tc.first, tc.second, tc.want, result)
		}
	}
}

func TestSuccessfulSubtraction(t *testing.T) {
	for _, tc := range []input[int]{
		{1, 2, -1},
		{42, 24, 18},
		{25, -25, 50},
	} {
		result := sub(tc.first, tc.second)
		if result != tc.want {
			t.Errorf("subtraction of %d and %d should be %d but got %d.", tc.first, tc.second, tc.want, result)
		}
	}
}

func TestSuccessfulMultiplication(t *testing.T) {
	for _, tc := range []input[int]{
		{0, 2, 0},
		{42, 24, 1008},
		{25, -25, -625},
	} {
		result := mul(tc.first, tc.second)
		if result != tc.want {
			t.Errorf("multiplication of %d and %d should be %d but got %d.", tc.first, tc.second, tc.want, result)
		}
	}
}

func TestSuccessfulDivision(t *testing.T) {
	for _, tc := range []input[float64]{
		{3, 6, .5},
		{42, 21, 2},
		{1000, .1, 10000},
	} {
		result, _ := div(tc.first, tc.second)
		if result != tc.want {
			t.Errorf("division of %f and %f should be %f but got %f.", tc.first, tc.second, tc.want, result)
		}
	}
}

func TestDivisionWithZeroShouldReturnError(t *testing.T) {
	_, err := div(10, 0)
	if err != DivisionByZeroError {
		t.Errorf("should not perform division with zero.")
	}
}

func TestSin(t *testing.T) {
	for _, tc := range []struct {
		degree   float64
		expected float64
	}{
		{0, 0},
		{30, 0.5},
		{45, 0.707},
		{60, 0.866},
		{90, 1},
	} {
		value := sin(tc.degree)
		if math.Abs(tc.expected-value) > 1e-3 {
			t.Errorf("sine of %f degrees should be %f but got %f", tc.degree, tc.expected, value)
		}
	}
}

func TestCos(t *testing.T) {
	for _, tc := range []struct {
		degree   float64
		expected float64
	}{
		{0, 1},
		{30, 0.866},
		{45, 0.707},
		{60, 0.5},
		{90, 0},
	} {
		value := cos(tc.degree)
		if math.Abs(tc.expected-value) > 1e-3 {
			t.Errorf("cosine of %f degrees should be %f but got %f.", tc.degree, tc.expected, value)
		}
	}
}

func TestTan(t *testing.T) {
	for _, tc := range []struct {
		degree   float64
		expected float64
	}{
		{0, 0},
		{30, 0.577},
		{45, 1},
		{60, 1.732},
	} {
		value := tan(tc.degree)
		if math.Abs(tc.expected-value) > 1e-3 {
			t.Errorf("tan of %f degrees should be %f but got %f.", tc.degree, tc.expected, value)
		}
	}
}

func TestSqrt(t *testing.T) {
	for _, tc := range []struct {
		in       float64
		expected float64
	}{
		{0, 0},
		{2, 1.414},
		{16, 4},
		{625, 25},
	} {
		value, err := sqrt(tc.in)
		if math.Abs(value-tc.expected) > 1e-3 {
			t.Errorf("square root of %f should be %f but got %f.", tc.in, tc.expected, value)
		}
		if err != nil {
			t.Errorf("should not return error for valid input.")
		}
	}
}

func TestSqrtWithNegativeShouldReturnAnError(t *testing.T) {
	_, err := sqrt(-1)
	if err != SqrtForNegativeValueError {
		t.Errorf("should not calculate square root for negative numbers")
	}
}
