package main


import (
	"testing"

	//"aoc2025/utils"
)

func TestRule2(t *testing.T) {
	// only invalid inputs
	inputs := []int {
		11,
		22,
		99,
		111,
		999,
		1010,
		222222,
		446446,
		565656,
		1111111,
		12341234,
		38593859,
		123123123,
		824824824,
		1212121212,
		1188511885,
		2121212121,
	}
	for _, input:= range inputs {
		if rule2(input) {
			t.Errorf("Input: %d", input)
		}
	}
}

func TestValid(t *testing.T) {
	// only valid inputs
		inputs := []int {
		33331,
		33332,
		33333,
		33334,
		33335,
		33336,
		33337,
		33338,
		33339,
		44440,
		44441,
		44442,
		44443,
		44444,
		44445,
		44446,
		44447,
		44448,
		44449,
		9990,
		9991,
		9992,
		9993,
		9994,
		9995,
		9996,
		9997,
		9998,
	}
	for _, input := range inputs {
		if !validID(input) {
			t.Errorf("%d", input)
		}
	}
}
