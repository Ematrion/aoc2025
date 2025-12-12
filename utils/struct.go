package utils

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"math"
)

type Box struct {
	X float64
	Y float64
	Z float64
}

func Distance(a, b Box) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2) + math.Pow(a.Z-b.Z, 2))
}

type IdRanges struct {
	Min int
	Max int
}

func GetIdRanges(line string) []IdRanges {
	ranges := strings.Split(line, ",")
	idRanges := make([]IdRanges, len(ranges))
	for i := range(len(ranges)) {
		extremes := strings.Split(ranges[i], "-")
		min, _ := strconv.Atoi(extremes[0])
		max, _ := strconv.Atoi(extremes[1])
		idRanges[i] = IdRanges{Min: min, Max: max}
	}
	return idRanges
}

func NoOverlappingRanges(ranges []IdRanges) []IdRanges {
	bounds := []IdRanges{ranges[0]}
	sorted_ranges := ranges[1:]
	sort.Slice(sorted_ranges, func (i, j int) bool {
		return sorted_ranges[i].Min < sorted_ranges[j].Min
	})
	//fmt.Println("---", bounds, sorted_ranges)
	for _, r:= range sorted_ranges {
		processed := false
		for i:= 0;  i <len(bounds) && !processed; i++ {
			//fmt.Println(bounds[i], r)
			// fix lower bound of range
			if r.Min < bounds[i].Min && r.Max >= bounds[i].Min {
				bounds[i] = IdRanges{r.Min, bounds[i].Max}
				processed = true
			}
			// fix upper bound of range
			if r.Min <= bounds[i].Max && r.Max > bounds[i].Max {
				bounds[i] = IdRanges{bounds[i].Min, r.Max}
				processed = true
			}
		}
		if !processed {
			bounds = append(bounds, r)
		}
	}

	return bounds
}



/*
Queue Implementation From:
https://medium.com/@danielabatibabatunde1/mastering-queues-in-golang-be77414abe9e

Modification To make it Generic

*/

// Queue is a type alias for a slice of integers
type Queue[T any] []T

// Enqueue adds an element to the rear of the queue
func (q *Queue[T]) Enqueue(value T) {
	*q = append(*q, value)
}

// Dequeue removes and returns an element from the front of the queue
func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, fmt.Errorf("empty queue")
	}
	value := (*q)[0]
	(*q)[0] = zero // Zero out the element (optional)
	*q = (*q)[1:]
	return value, nil
}

// CheckFront returns the front element without removing it
func (q *Queue[T]) CheckFront() (T, error) {
	var zero T
	if q.IsEmpty() {
	return zero, fmt.Errorf("empty queue")
	}
	return (*q)[0], nil
}

// IsEmpty checks if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

// Size returns the number of elements in the queue
func (q *Queue[T]) Size() int {
	return len(*q)
}
