package main

import (
	"fmt"
	"strconv"
	"strings"
	"math"

	"aoc2025/utils"
)




func main() {
	lines, err := utils.ReadFileToLines("inputs/day9.txt")
	utils.CheckError(err)
	tiles := make([]utils.Box, len(lines))
	for i, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		tiles[i] = utils.Box{X:float64(x), Y:float64(y), Z:0}
	}
	largest := 0.0
	for i, t1 := range tiles {
		for _, t2 := range tiles[i:] {
			area := (math.Abs(t1.X-t2.X)+1) * (math.Abs(t1.Y-t2.Y)+1)
			if area > largest {
				//fmt.Println(t1, t2)
				largest = area
			}
		}
	}


	fmt.Printf("%d", int(largest))
	
}