package main

import (
	"fmt"
	"math"
	//"strings"

	"aoc2025/utils"
)

/*func maxPower(bank string) (string, int) {
	fmt.Println("----\nmax power: ", bank)
	powered := 0
	var consumed []rune

	for i, battery := range bank {
		p := int(battery-'0')
		consumed = append(consumed, battery)
		_, check := strings.CutPrefix(bank, string(consumed))
		//fmt.Println(p, " ", string(consumed), " ")
		if !check {
			panic("OMG")
		}
		//fmt.Println(i, bank[i], battery, p)
		if p > powered {
			powered = p
			fmt.Println("selection: ", i, " ", p, " ", powered)
		}
		if powered == 9 {
			return string(consumed), powered
		}
	}
	return string(consumed), powered
}*/

func maxPower(bank string) (int, int) {
	power := 0
	index := -1
	for i, battery := range bank {
		p := int(battery-'0')
		if p > power {
			power = p
			index = i
			if power == 9 {
				return power, index
			}
		}
	}
	return power, index
}


func highestJoltage(bank string, nb int) int {
	joltage := 0
	start := 0
	for i:=nb-1;  i >= 0; i-- {
		power, index := maxPower(bank[start:len(bank)-i])
		//fmt.Println(selected, bank)
		joltage += power * int(math.Pow(10, float64(i)))
		//fmt.Println(selected, joltage)
		start += (index + 1)
	}
	return joltage
}


func main() {
	banks, err := utils.ReadFileToLines("inputs/day3.txt")
	utils.CheckError(err)
	joltage := 0
	for _, bank := range banks {
		fmt.Println(bank)
		power := highestJoltage(bank, 12)
		fmt.Println("Selected Power: ",power)
		joltage += power
	}
	fmt.Println("JOLTAGE :", joltage)
}