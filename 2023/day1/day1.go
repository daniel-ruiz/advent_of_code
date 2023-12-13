package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fileName := os.Args[1]
	data, _ := os.ReadFile(fileName)
	lines := bytes.Split(data, []byte("\n"))
	parameters := make([]int, len(lines))
	for index, line := range lines {
		parameters[index] = extractParameterFromLine(line)
	}

	fmt.Println("Result:", sum(parameters))
}

func sum(parameters []int) int {
	result := 0
	for _, number := range parameters {
		result += number
	}
	return result
}

func extractParameterFromLine(line []byte) int {
	var parameters [2]int
	parameterCount := 0
	for _, char := range line {
		if unicode.IsNumber(rune(char)) {
			newParameter, _ := strconv.Atoi(string(char))
			if parameterCount == 2 {
				parameters[1] = newParameter
				continue
			}
			parameters[parameterCount] = newParameter
			parameterCount ++
		}
	}
	if parameterCount == 0 { return 0 }
	if parameterCount == 1 {
		return parameters[0] * 10 + parameters[0]
	}
	return parameters[0] * 10 + parameters[1]
}
