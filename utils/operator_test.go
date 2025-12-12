package utils

import (
	"testing"
)

func TestModularArithmetic(t *testing.T) {
	tests := []struct {
		a, m, b, k  int
	} {
		{5, 3, 2, 1},
		{10, 4, 2, 2},
		{2, 5, 2, 0},
		{-5, 3, 1, -2},
		{0, 7, 0, 0},
		{-1, 7, 6, -1},
		{-500, 100, 0, -5},
		{500, 100, 0, 5},
	}
	for _, test := range tests {
		b, k := ModularAtithemtic(test.a, test.m)
		if b != test.b || k != test.k {
			t.Errorf("Modular(%d, %d) = %d, %d, expected %d, %d", test.a, test.m, b, k ,test.b, test.k)
		}
	}

}