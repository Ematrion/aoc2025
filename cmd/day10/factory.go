package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2025/utils"
)

const ON = '#'
const OFF = '.'
const LIGHT = "light"
const POWER = "power"

type Machine struct {
	goal []rune
	light []rune
	buttons [][]int
	joltage []int
	powers []int
	mode string
}

func LoadMachine(desc string) Machine {
	//fmt.Println("---")
	tokens := strings.Split(desc, " ")
	//fmt.Println(tokens)

	pattern := tokens[0]
	pattern, _ = strings.CutPrefix(pattern, "[")
	pattern, _ = strings.CutSuffix(pattern, "]")
	goal :=  []rune(pattern)
	light := make([]rune, len(goal))
	for i:= range light {light[i] = OFF}

	buttons := make([][]int, len(tokens)-2) 
	for i, token := range tokens[1: len(tokens)-1] {
		element, _ := strings.CutPrefix(token, "(")
		element, _ = strings.CutSuffix(element, ")")
		values := strings.Split(element, ",")
		button := make([]int, len(values))
		for j, v := range values {
			button[j], _ = strconv.Atoi(v)
		}
		buttons[i] = button
	}

	element, _ := strings.CutPrefix(tokens[len(tokens)-1], "{")
	element, _ = strings.CutSuffix(element, "}")
	values := strings.Split(element, ",")
	joltage := make([]int, len(values))
	power := make([]int, len(values))
	for i, v := range values {
		joltage[i], _ = strconv.Atoi(v)
		power[i] = 0 // unecessary
	}

	//fmt.Println(goal,  " ", light, " ", buttons, " ", joltage)
	return Machine{goal, light, buttons, joltage, power, LIGHT}
}

func (m *Machine) Clone() Machine {
	np := make([]int, len(m.powers))
	nl := make([]rune, len(m.light))
	copy(nl, m.light)
	copy(np, m.powers)
	return Machine{m.goal, nl, m.buttons, m.joltage, np, m.mode}
}

func (m *Machine) SwitchLight(l int) {
	if m.light[l] == ON {
		m.light[l] = OFF
	} else {
		m.light[l] = ON
	}
}

func (m *Machine) MorePower(b int) {
	m.powers[b]++
}

func (m *Machine) CheckGoal() bool {
	switch m.mode {
	case LIGHT: return m.checkLight()
	case POWER: return m.checkPower()
	default: panic("unknown mode")
	}
}

func (m *Machine) checkLight() bool {
	success := true
	for i, light := range m.goal {
		if m.light[i] != light {
			success = false
		}
	}
	return success
}

func (m *Machine) checkPower() bool {
	for i, power := range m.powers {
		if m.joltage[i] != power {
			return false
		}
	}
	return true
}

func (m *Machine) PressButton(b int) {
	switch m.mode {
	case LIGHT: {
		for _, button := range m.buttons[b] {
			m.SwitchLight(button)
		}
	}
	case POWER: {
		for _, button := range m.buttons[b] {
			m.MorePower(button)
		}
	}
	default:
		panic("unknow mode")
	}

}


func (m *Machine) Extend() []*Machine {
	switch m.mode {
	case LIGHT: return m.extendLight()
	case POWER: return m.extendPower()
	default: panic("unknown mode")
	}
}

func (m *Machine) extendLight() []*Machine {
	if !(m.mode==LIGHT){ panic("cannot extendLight when not in mode light") }
	machines := make([]*Machine, len(m.buttons))
	for i := range m.buttons {
		new_machine := (*m).Clone()
		new_machine.PressButton(i)
		machines[i] = &new_machine
	}
	return machines
}

func (m *Machine) extendPower() []*Machine {
	if !(m.mode==POWER){ panic("cannot extendLight when not in mode light") }
	var machines []*Machine 
	for i := range m.buttons {
		new_machine := (*m).Clone()
		new_machine.PressButton(i)
		overpowered := false
		for i, power := range m.powers {
			if power > m.joltage[i] {
				overpowered = true
			}
		}
		if !overpowered {
			machines = append(machines, &new_machine)
		}
	}
	return machines
}




func main() {
	lines, err := utils.ReadFileToLines("inputs/day10.txt")
	utils.CheckError(err)

	machines := make([]Machine, len(lines))
	for i, line := range lines {
		machines[i] = LoadMachine(line)
	}

	
	//fmt.Println(machine)
	//fmt.Println(machine)
	
	steps := 0
	for i, m := range machines {
		m.mode = POWER
		fmt.Println(i)
		//fmt.Println(m)
		_, step := utils.BFS(&m)
		//fmt.Println(solution)
		fmt.Println(i, " in ",step)
		//fmt.Println("---")
		steps += step
	}
	fmt.Println("total: ", steps)



}