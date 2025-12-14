package main


import (
	"fmt"
	"strings"
	"strconv"

	"aoc2025/utils"
)

type Cell rune

const EMPTY Cell = '.'
const OCCUPIED Cell = '#'

type Axe rune
const XFLIP Axe = 'x'
const YFLIP Axe = 'y'

type Shape [][]Cell

type Region struct {
	grid [][]string
}

func NewRegion(x, y int) Region{
	surface := make([][]string, x)
	row := make([]string, y)
	for j:=0; j<y; j++ {
		row[j] = string(EMPTY)
	}
	for i:=0; i<x; i++ {
		surface[i] = row
	}
	return Region{surface}
}

func (r *Region) Plot() {
	for _, row := range r.grid {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println()
	}
}

func LoadRegions(lines []string) []Region {
	regions := make([]Region, len(lines))
	for i, line := range lines {
		elements := strings.Split(line, " ")
		coords := strings.Split(elements[0], "x")
		x,_ := strconv.Atoi(coords[0])
		ystr, _ := strings.CutSuffix(coords[1], ":")
		y, _ := strconv.Atoi(ystr)
		regions[i] = NewRegion(x, y)	
	}
	return regions
}

func yFlip(shape Shape) Shape {
	flipped := make(Shape, len(shape))
	nRow := len(shape)
	for i, row := range(shape) {
		flipped[nRow-1-i] = row
	}
	return flipped
}

func xFlip(shape Shape) Shape {
	flipped := make(Shape, len(shape))
	nCol := len(shape[0])
	flippedRow := make([]Cell, nCol)
	for i, row := range shape {
		for j, token := range row {
			flippedRow[nCol-1-j] = token
		}
		flipped[i] = flippedRow
	}
	return flipped
}

func Flip(shape Shape, orientation Axe) Shape {
	if orientation == XFLIP {
		return xFlip(shape)
	} else if orientation == YFLIP {
		return yFlip(shape)
	} else {
		panic("what you want me to do")
	}
}

func Rotate(s Shape) Shape {
	// 90 degree rotation <=> transpose + reverse row order
	tm := utils.Transpose(s)
	return yFlip(tm)
}




func main() {
	lines, err := utils.ReadFileToLines("inputs/examples/day12.txt")
	utils.CheckError(err)
	fmt.Println(lines)

	region := LoadRegions(lines[30:])[1]
	region.Plot()

}