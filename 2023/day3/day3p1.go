package main

import (
	"bytes"
	"os"
)

type location struct {
	row int
	column int
}

type engineSymbol struct {
	location location
}

type engineNumber struct {
	amount int
	startLocation location
	endLocation location
}

func (number *engineNumber) allLocations() []location {
	var allLocations []location

}

func main() {
	fileName := os.Args[1]
	fileContents, _ := os.ReadFile(fileName)
	fileLines := bytes.Split(fileContents, []byte("\n"))
}
