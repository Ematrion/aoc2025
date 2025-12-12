package main

import (
	"fmt"
	"strconv"

	"aoc2025/utils"
)

// method 0x434C49434B


type Safe struct {
	pointer int
	value int
	countValueStop int
	countValueClick int
	dial []int
}

func (s *Safe) updateStop() {
	if s.dial[s.pointer] == s.value {
		s.countValueStop++
		s.countValueClick++
	}
}

func (s *Safe) updateClick() {
	if s.dial[s.pointer] == s.value {
		s.countValueClick++
	}
}

func (s *Safe) updatePointer() {
	s.pointer, _ = utils.ModularAtithemtic(s.pointer, len(s.dial))
}

func (s *Safe) clickLeft() {
	s.pointer--
	s.updatePointer()
	s.updateClick()
}

func (s *Safe) clickRight() {
	s.pointer++
	s.updatePointer()
	s.updateClick()
}

func (s *Safe) Rotate(c rune, inc int) {
	var f func()
	switch c {
		case 'L': f = s.clickLeft
		case 'R': f = s.clickRight
		default: panic("Unknow action")
	}
	for i:=0; i<inc; i++ {
		f()
	}
	s.updateStop()
}

func NewSafe(current, value int, n int) Safe {
	return Safe{
		pointer: current,
		value: value,
		countValueStop: 0,
		countValueClick: 0,
		dial: utils.Dial(n),
	}
}


func main() {
	lines, err := utils.ReadFileToLines("inputs/day1.txt")
	utils.CheckError(err)
	safeA := NewSafe(50, 0, 100)
	fmt.Println(safeA.dial)
	for _, line := range(lines) {
		r := []rune(line)
		char := r[0]
		num, err := strconv.Atoi(string(r[1:]))
		utils.CheckError(err)
		safeA.Rotate(char, num)
	}
	fmt.Println(safeA.countValueStop, safeA.countValueClick)
}