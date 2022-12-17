package main

import (
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
			t.Errorf("addition of %d and %d should be %d.", tc.first, tc.second, tc.want)
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
			t.Errorf("subtraction of %d and %d should be %d.", tc.first, tc.second, tc.want)
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
			t.Errorf("multiplication of %d and %d should be %d.", tc.first, tc.second, tc.want)
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
			t.Errorf("division of %f and %f should be %f.", tc.first, tc.second, tc.want)
		}
	}
}

func TestDivisionWithZeroShouldReturnError(t *testing.T) {
	_, err := div(10, 0)
	if err != DivisionByZeroError {
		t.Errorf("should not perform division with zero.")
	}
}
