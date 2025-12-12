package main

import (
	"fmt"
	"math"
	"strings"

	//"aoc2025/utils"
)

func maxPower(bank string) (string, int) {
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
}

func highestJoltage(bank string, nb int) int {
	joltage := 0
	for i:=nb-1;  i >= 0; i-- {
		consumed, selected := maxPower(bank[:len(bank)-i])
		//fmt.Println(selected, bank)
		joltage += selected * int(math.Pow(10, float64(i)))
		//fmt.Println(selected, joltage)
		bank, _ = strings.CutPrefix(bank, consumed)
	}
	return joltage
}


func main() {
	//banks, err := utils.ReadFileToLines("inputs/day3.txt")
	//utils.CheckError(err)
	banks := []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
		//"2274342251334447452432195314423334442326243253434454374212225244643223361286134343534441232331432424",
		//Selected Power:  98
	}
	joltage := 0
	for _, bank := range banks {
		fmt.Println(bank)
		power := highestJoltage(bank, 2)
		fmt.Println("Selected Power: ",power)
		joltage += power
	}
	fmt.Println("JOLTAGE :", joltage)
}