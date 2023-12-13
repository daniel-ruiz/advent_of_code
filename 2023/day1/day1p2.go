package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	fmt.Println("Parameters:", parameters)
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
	for index := range line {
		newParameter, error := findNextParameterInLine(line, index)
		if error != nil {
			continue
		}
		if parameterCount == 2 {
			parameters[1] = newParameter
			continue
		}
		parameters[parameterCount] = newParameter
		parameterCount ++
	}
	if parameterCount == 0 { return 0 }
	if parameterCount == 1 {
		return parameters[0] * 10 + parameters[0]
	}
	return parameters[0] * 10 + parameters[1]
}

func findNextParameterInLine(line []byte, index int) (int, error) {
	if unicode.IsNumber(rune(line[index])) {
		parameter, _ := strconv.Atoi(string(line[index]))
		return parameter, nil
	}
	parameterLookup := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	for parameterIndicator, parameter := range parameterLookup {
		if strings.HasPrefix(string(line[index:]), parameterIndicator) {
			return parameter, nil
		}
	}
	return 0, errors.New("could not find parameter")
}
