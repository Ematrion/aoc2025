package utils

import (
	"os"
	"bufio"
)


func ReadFileToLines(path string) ([]string, error) {
	file, err := os.Open(path)
	CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	CheckError(scanner.Err())

	return lines, nil
}