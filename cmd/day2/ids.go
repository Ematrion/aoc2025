package main

import (
	"fmt"
	"strconv"

	"aoc2025/utils"
)



func evenLen(id int) bool {
	word := strconv.Itoa(id)
	return len(word) % 2 == 0
}

func rule1(id int) bool {
	valid := true
	word := strconv.Itoa(id)
	l := len(word)
	if l % 2 == 0 {
		start:= word[:l/2]
		end := word[l/2:]
		//print(word, ", ", start,", ", end, "\n")
		if start == end {
			//print(start, end)
			valid = false
		}
	}
	return valid
}

func rule2(id int) bool {
	valid := true
	word := strconv.Itoa(id)
	for i := 1; i <= len(word)/2; i++ {
		//b, _ := utils.ModularAtithemtic(len(word), i)
		//if b == 0 {
		if len(word) % i == 0 {
			pattern := word[:i]
			allPartMatch := true
			//fmt.Println("Testing pattern:", pattern)
			for j := i; j+i<=len(word) && allPartMatch; j+=i {
				//fmt.Println(word[j:j+i])
				if pattern != word[j:j+i] {
					//fmt.Println(word, i, pattern, word[j:j+i])
					allPartMatch = false
				} else {
					;//fmt.Println(word, pattern, word[j:j+i])
				}
			}
			if allPartMatch {
				return false
			}
		}
	}
	return valid
}

func validID(id int) bool {
	if !evenLen(id) {
		//fmt.Println(id, "valid")
		return true
	}
	return rule1(id) && rule2(id)
}

func main() {
	lines, err := utils.ReadFileToLines("inputs/day2.txt")
	utils.CheckError(err)
	//lines := []string{"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"}
	//lines := []string{"95-115,998-1012,565653-565659,824824821-824824827,2121212118-2121212124"}
	var invalidIDS []int
	for _, bounds := range(utils.GetIdRanges(lines[0])) {
		for id:=bounds.Min; id <= bounds.Max; id++ {
			if !validID(id) {
				fmt.Println(id)
				invalidIDS = append(invalidIDS, id)
			}
		}
	}
	result := utils.Sum(invalidIDS)
	fmt.Println(result)
}