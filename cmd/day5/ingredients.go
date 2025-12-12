package main

import (
	"fmt"
	"strconv"

	"aoc2025/utils"
)


func splitInput(lines []string) ([]string, []string) {
	var split int
	for i, line := range lines {
		if line == "" {
			split = i
		}
	}
	return lines[:split], lines[split+1:]
}


func allRanges(lines []string) []utils.IdRanges {
	idRanges := make([]utils.IdRanges, len(lines))
	for i, line:= range lines {
		idRanges[i] = utils.GetIdRanges(line)[0]
	}
	return idRanges
}





func main() {
	lines, err := utils.ReadFileToLines("inputs/day5.txt")
	utils.CheckError(err)
	ranges, ids := splitInput(lines)
	bounds := allRanges(ranges)

	// task1
	totalFresh := 0
	for _, id:= range ids {
		checked := false
		ingredientID, _  := strconv.Atoi(id)
		for i:= 0; i< len(bounds) && !checked; i++ {
			boundaries := bounds[i]
			if boundaries.Min <= ingredientID && ingredientID <= boundaries.Max {
				//fmt.Println(ingredientID, " ", boundaries)
				checked = true
				totalFresh++
			}
		}
	}
	fmt.Println(totalFresh)

	//task2
	niceRanges := utils.NoOverlappingRanges(bounds)
	totalIds := 0
	for _, nr := range niceRanges {
		totalIds += nr.Max-nr.Min + 1
	}
	fmt.Println(totalIds)
}
