package main

import (
	"fmt"
	"aoc2025/utils"
)



func downwardBeam(manifold [][]rune) ([][]rune, int) {
	splits := 0
	// start the beam
	start := 'S'
	split := '^'
	beam := '|'
	space := '.'
	for i, pos := range manifold[0] {
		if pos == start {
			manifold[0][i] = beam
		}
	}
	for i:=1; i <len(manifold); i++ {
		for j, pos := range(manifold[i]) {
			if pos == split && manifold[i-1][j]== beam {
				splits++
				if j > 0 {
					manifold[i][j-1] = beam
				}
				if j < len(manifold[0])-1 {
					manifold[i][j+1] = beam
				}
			} else if pos == space && manifold[i-1][j] == beam{
				// coninue beam
				manifold[i][j] = beam
			}
		}
	}
	return manifold, splits
}

func main() {
	lines, err := utils.ReadFileToLines("inputs/day7.txt")
	tables := make([][]rune, len(lines))
	for i, line :=range(lines) {
		tables[i] = []rune(line)
	}
	utils.CheckError(err)
	fmt.Println(lines)
	fmt.Println()
	result, total := downwardBeam(tables)
	for _, r := range result {
		fmt.Println(string(r))
	}
	fmt.Println(total)
}