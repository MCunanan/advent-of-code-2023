package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/MCunanan/advent-of-code-2023/util"
)

var numbersAndDot = regexp.MustCompile("[0-9.]")

var inputData []string

func getNeighboringNumbers(row, col int) []int {
	var neighboringNumbers []int

	// TODO: Scan around the symbol for numbers

	return neighboringNumbers
}

func SolvePuzzle(inputFile string) {
	sum := 0

	f, err := os.Open(inputFile)
	util.Check(err)
	defer f.Close()

	// Load the entire file into
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		inputData = append(inputData, line)
	}

	for row, element := range inputData {
		for col, char := range strings.Split(element, "") {
			if !numbersAndDot.MatchString(char) {
				schematicNumbers := getNeighboringNumbers(row, col)
			}
		}
	}

	fmt.Printf("Part 1 Schematic Sum: %d", sum)
}
