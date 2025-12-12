package main


import (
	"fmt"
	"strings"
	"strconv"
	"errors"

	"aoc2025/utils"
)


func solveProblem(operands []int, operator string) (int, error) {
	result := 0
	var oper func (x, y int) int

	switch operator{
		case "*" : {
			oper = func (x, y int) int {return x*y}
			result = 1
		}
		case "+": oper = func (x, y int) int {return x+y}
		default: return -100, errors.New("unknown operator")
	}
 	for _, o:= range operands {
		result = oper(result, o)
	}
	return result, nil
}


func main() {
	lines, err := utils.ReadFileToLines("inputs/day6.txt")
	utils.CheckError(err)

	problems := [][]string{}
	for _, line := range lines {
		splited := strings.Split(line, " ")
		values := []string{}
		for _, v :=range splited {
			v = strings.TrimSpace(v)
			if v != "" {
				values = append(values, v)
			}
		}
		problems = append(problems, values)
	}

	problems = utils.Transpose(problems)
	fmt.Println(problems)
	solutions := make([]int, len(problems))
	for i, problem := range problems {
		operands, operator := problem[:len(problem)-1], problem[len(problem)-1]
		values := make([]int, len(operands))
		for i, v := range operands {
			values[i], _ = strconv.Atoi(v)
		}
		solutions[i], err = solveProblem(values, operator)
		utils.CheckError(err)
	}

	fmt.Println(solutions)
	answer, err := solveProblem(solutions, "+")
	utils.CheckError(err)
	fmt.Println(answer)
}