package main

import (
	"fmt"
	"math"
	"strings"
	"strconv"

	"aoc2025/utils"
)

type Circuits struct {
	CircuitsArray [][]int
}

func (c *Circuits) NewCircuit(values []int) bool {
	newCircuit := make([]int, len(values))
	for i, v := range values {
		_, check := c.GetCircuit(v)
		if check {
			return false
		} else {
			newCircuit[i] = v
		}
	}
	c.CircuitsArray = append(c.CircuitsArray, newCircuit)
	return true
}

func (c *Circuits) GetCircuit(a int) (int, bool) {
	for i, circuit := range c.CircuitsArray {
		for _, v := range circuit {
			if a == v {
				return i, true
			}
		}
	}
	return -1, false
}

func (c *Circuits) SetCircuit(a, circuit int) bool {
	_, found := c.GetCircuit(a)
	if found {
		return false
	} else {
		c.CircuitsArray[circuit] = append(c.CircuitsArray[circuit], a)
		return true
	}
}

func (c *Circuits) PopCircuit(circuit int) ([]int, bool) {
	if !(circuit < len(c.CircuitsArray)) {
		return []int{}, false
	}
	popped := c.CircuitsArray[circuit]
	newCircuit := make([][]int, len(c.CircuitsArray)-1)
	for i, circ := range c.CircuitsArray {
		if i < circuit{
		newCircuit[i] = circ
		} else if i > circuit {
			newCircuit[i-1] = circ
		}
	}
	c.CircuitsArray = newCircuit
	return popped, true
}

func (c *Circuits) MergeCircuits(a, b int) bool {
	if !(a<len(c.CircuitsArray)) && !(b<len(c.CircuitsArray)) {
		return false
	}
	if a == b {
		return false
	}
	ext, rem := a, b
	if b < a {
		ext = b
		rem = a
	}
	removed, check := c.PopCircuit(rem)
	if !check {
		return false
	}
	extended := c.CircuitsArray[ext]
	extended = append(extended, removed...)
	c.CircuitsArray[ext] = extended
	return true
}

func initCircuit(boxes []utils.Box) ([][]float64, Circuits) {
	distances := make([][]float64, len(boxes))
	var circuits Circuits
	for i:= range distances {
		distances[i] = make([]float64, len(boxes))
		for j:= range boxes {
			if i==j {
				distances[i][j] = math.NaN()
			} else {
			distances[i][j] = utils.Distance(boxes[i], boxes[j])
			}
		}
	}
	return distances, circuits
}


func smallestDistance(distances [][]float64) (int, int) {
	//fmt.Println("--------------------")
	a, b := 0, 0
	minDist := -1.0
	for i := range distances {
		for j := range distances {
			//fmt.Println("---")
			dist := distances[i][j]
			//fmt.Println(dist)
			if  !math.IsNaN(dist) && (minDist < 0 || dist < minDist) {
				a = i
				b = j
				minDist = dist
				//fmt.Println(a, b, dist)
			}
		}
	}
	return a, b
}

func main() {
	lines, err := utils.ReadFileToLines("inputs/day8.txt")
	utils.CheckError(err)
	//fmt.Println(lines)

	boxes := make([]utils.Box, len(lines))
	for i, line := range lines {
		sv := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(sv[0], 64)
		y, _ := strconv.ParseFloat(sv[1], 64)
		z, _ := strconv.ParseFloat(sv[2], 64)
		boxes[i] = utils.Box{X: x, Y: y, Z: z}
	}

	distances, circuits := initCircuit(boxes)

	for i:=0; i<1000; i++ {
		boxA, boxB:= smallestDistance(distances)
		//dist := distances[boxA][boxB]
		distances[boxA][boxB] = math.NaN()
		distances[boxB][boxA] = math.NaN()
		cA, checkA := circuits.GetCircuit(boxA)
		cB, checkB := circuits.GetCircuit(boxB)
		if checkA && checkB {
			circuits.MergeCircuits(cA, cB)
		} else if checkA && ! checkB {
			circuits.SetCircuit(boxB, cA)
		} else if !checkA && checkB {
			circuits.SetCircuit(boxA, cB)
		} else {
			circuits.NewCircuit([]int{boxA, boxB})
		}
		//cA, _ = circuits.GetCircuit(boxA)
		//cB, _ = circuits.GetCircuit(boxB)
		//fmt.Println(i, "  :  ", boxes[boxA], boxes[boxB], dist)
		//fmt.Println(boxA, boxB, " | ", cA, cB)
		//fmt.Println(circuits.CircuitsArray)
	}

	top1, top2, top3 := 0, 0, 0
	for _, circuit := range circuits.CircuitsArray {
		n:= len(circuit)
		if n > top1 {
			top3 = top2
			top2 = top1
			top1 = n
		} else if n > top2 {
			top3 = top2
			top2 = n
		} else if n > top3 {
			top3 = n
		}
	}
	fmt.Println(top1, top2, top3)
	fmt.Println(top1*top2*top3)

	//task2
	var boxA, boxB int
	for len(circuits.CircuitsArray[0])!=len(boxes) {
		boxA, boxB = smallestDistance(distances)
		//dist := distances[boxA][boxB]
		distances[boxA][boxB] = math.NaN()
		distances[boxB][boxA] = math.NaN()
		cA, checkA := circuits.GetCircuit(boxA)
		cB, checkB := circuits.GetCircuit(boxB)
		if checkA && checkB {
			circuits.MergeCircuits(cA, cB)
		} else if checkA && ! checkB {
			circuits.SetCircuit(boxB, cA)
		} else if !checkA && checkB {
			circuits.SetCircuit(boxA, cB)
		} else {
			circuits.NewCircuit([]int{boxA, boxB})
		}
	}
	fmt.Println(boxes[boxA], boxes[boxB])
	fmt.Println(boxes[boxA].X * boxes[boxB].X)
}
