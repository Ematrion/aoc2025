package main

import (
	"testing"
	"strconv"

	"aoc2025/utils"
)


func TestSafeClick (t *testing.T) {
	inputs := []struct {
		action string
		current, clicks, stops int
	} {
		{"L68", 82, 1, 0},	// The dial is rotated L68 to point at 82.
		{"L30", 52, 1, 0}, 	// The dial is rotated L30 to point at 52.
		{"R48", 0, 2, 1},	// The dial is rotated R48 to point at 0.
		{"L5", 95, 2, 1},	// The dial is rotated L5 to point at 95.
		{"R60", 55, 3, 1},	// The dial is rotated R60 to point at 55.
		{"L55", 0, 4, 2},	// The dial is rotated L55 to point at 0.
		{"L1", 99, 4, 2},	// The dial is rotated L1 to point at 99.
		{"L99", 0, 5, 3}, 	// The dial is rotated L99 to point at 0.
		{"R14", 14, 5, 3},	// The dial is rotated R14 to point at 14.
		{"L82", 32, 6, 3},	// The dial is rotated L82 to point at 32.
		{"L500", 32, 11, 3},
		{"R500", 32, 16, 3},
		{"R68", 0, 17, 4},
		{"L500", 0, 22, 5},
		{"R500", 0, 27, 6},
	}

	safe := NewSafe(50, 0, 100) // The dial starts by pointing at 50, checks 0, size 100
	for _, input := range inputs {
		r := []rune(input.action)
		char := r[0]
		num, err := strconv.Atoi(string(r[1:]))
		utils.CheckError(err)
		prev, click, stop := safe.dial[safe.pointer], safe.countValueClick, safe.countValueStop
		safe.Rotate(char, num)
		if safe.dial[safe.pointer] != input.current || safe.countValueClick != input.clicks || safe.countValueStop != input.stops {
			
			t.Errorf("Input: (%s, %d, %d, %d) - (%d, %d, %d) <- (%d, %d, %d)",
			input.action, input.current, input.clicks, input.stops,
			safe.dial[safe.pointer], safe.countValueClick, safe.countValueStop,
			prev, click, stop)
		}
	}
}