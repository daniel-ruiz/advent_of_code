package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type cubeSet struct {
	redCount int
	greenCount int
	blueCount int
}

func (set *cubeSet) isSubsetOf(anotherSet *cubeSet) bool {
	return set.redCount <= anotherSet.redCount && set.greenCount <= anotherSet.greenCount && set.blueCount <= anotherSet.blueCount
}

func (set *cubeSet) power() int {
	return set.redCount * set.greenCount * set.blueCount
}

type game struct {
	id int
	cubeSets []cubeSet
}

func (game *game) minimumCubeSet() cubeSet {
	minimumCubeSet := cubeSet{0, 0, 0}
	for _, cubeSet := range game.cubeSets {
		if cubeSet.redCount > minimumCubeSet.redCount {
			minimumCubeSet.redCount = cubeSet.redCount
		}
		if cubeSet.greenCount > minimumCubeSet.greenCount {
			minimumCubeSet.greenCount = cubeSet.greenCount
		}
		if cubeSet.blueCount > minimumCubeSet.blueCount {
			minimumCubeSet.blueCount = cubeSet.blueCount
		}
	}
	return minimumCubeSet
}


func main() {
	var minimumCubeSetPowers []int
	fileName := os.Args[1]
	data, _ := os.ReadFile(fileName)
	games := parseGamesFromFileData(data)
	fmt.Println("Games:", games)
	for _, game := range games {
		minimumCubeSet := game.minimumCubeSet()
		minimumCubeSetPowers = append(minimumCubeSetPowers, minimumCubeSet.power())
	}
	fmt.Println("Powers of the minimium cube sets: ", minimumCubeSetPowers)
	fmt.Println("Result: ", sum(minimumCubeSetPowers))
}

func isImpossibleGame(game game) bool {
	goalCubeSet := cubeSet{12, 13, 14}
	for _, cubeSet := range game.cubeSets {
		if !cubeSet.isSubsetOf(&goalCubeSet) {
			return false
		}
	}
	return true
}

func sum(parameters []int) int {
	result := 0
	for _, number := range parameters {
		result += number
	}
	return result
}

func parseGamesFromFileData(data []byte) []game {
	var games []game
	lines := bytes.Split(data, []byte("\n"))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		newGame := parseGameFromLine(line)
		games = append(games, newGame)
	}
	return games
}

func parseGameFromLine(line []byte) game {
	gameParts := bytes.Split(line, []byte(": "))
	idPart := gameParts[0]
	setParts := gameParts[1]
	gameId, _ := strconv.Atoi(string(bytes.Split(idPart, []byte(" "))[1]))
	cubeSets := parseCubeSetsFromSetParts(setParts)
	return game{id: gameId, cubeSets: cubeSets}
}

func parseCubeSetsFromSetParts(setParts []byte) []cubeSet {
	var cubeSets []cubeSet
	sets := bytes.Split(setParts, []byte("; "))

	for _, setPart := range sets {
		newCubeSet := parseCubeSetFromSetPart(setPart)
		cubeSets = append(cubeSets, newCubeSet)
	}
	return cubeSets
}

func parseCubeSetFromSetPart(setPart []byte) cubeSet {
	newCubeSet := cubeSet{}
	setCounts := bytes.Split(setPart, []byte(", "))
	for _, counts := range setCounts {
		amountPart := bytes.Split(counts, []byte(" "))[0]
		cubeColor := string(bytes.Split(counts, []byte(" "))[1])
		amount, _ := strconv.Atoi(string(amountPart))
		if cubeColor == "red" {
			newCubeSet.redCount = amount
		}
		if cubeColor == "green" {
			newCubeSet.greenCount = amount
		}
		if cubeColor == "blue" {
			newCubeSet.blueCount = amount
		}
	}
	return newCubeSet
}
