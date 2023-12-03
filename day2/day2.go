package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/MCunanan/advent-of-code-2023/util"
)

func isGamePossible(game string, red, green, blue int) bool {
	restrictions := map[string]int{
		"red":   red,
		"green": green,
		"blue":  blue,
	}
	game = strings.TrimSpace(game)
	rounds := strings.Split(game, ";")
	for _, round := range rounds {
		cubes := strings.Split(round, ",")
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			count := strings.Split(cube, " ")

			value, err := strconv.Atoi(count[0])
			util.Check(err)

			key := count[1]

			if restrictions[key] < value {
				return false
			}
		}
	}

	return true
}

func calculatePowerOfGame(game string) int {
	maxCubes := make(map[string]int)

	game = strings.TrimSpace(game)
	rounds := strings.Split(game, ";")
	for _, round := range rounds {
		cubes := strings.Split(round, ",")
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			count := strings.Split(cube, " ")

			value, err := strconv.Atoi(count[0])
			util.Check(err)

			key := count[1]

			if maxCubes[key] < value {
				maxCubes[key] = value
			}
		}
	}

	power := 1
	for _, value := range maxCubes {
		power = power * value
	}
	return power
}

func SolvePuzzle(inputFile string, red, green, blue int) {
	sum := 0
	powers := 0

	f, err := os.Open(inputFile)
	util.Check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		scanned := scanner.Text()
		gameSplit := strings.Split(scanned, ":")
		gameString := util.RemoveAlphas(gameSplit[0])

		gameNumber, err := strconv.Atoi(strings.TrimSpace(gameString))
		util.Check(err)

		if isGamePossible(gameSplit[1], red, green, blue) {
			sum += gameNumber
		}
		powers += calculatePowerOfGame(gameSplit[1])
	}

	fmt.Printf("Sum of possible game IDs: %d\nSum of powers: %d", sum, powers)
}
