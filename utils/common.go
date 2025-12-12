package utils

import "log"

type Number interface {
	int | float32
}

func CheckError(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func Dial(n int) []int {
	r := make([]int, n)
	for i := range(n) {
		r[i] = i
	}
	return r
}

func Sum[N Number](list []N) N {
	sum := N(0)
	for _, value := range list {
		sum += value
	}
	return sum
}

func Transpose[T any](table [][]T) [][]T {
	y := len(table)
	x := len(table[0])
	new_table := make([][]T, x)
	for i:= range new_table {
		new_table[i] = make([]T, y)
	}
	for i, row := range table {
		for j:= range row {
			new_table[j][i] = table [i][j]
		}
	}
	return new_table
}