package day1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var alphas = regexp.MustCompile(`[a-zA-Z]+`)

func cleanString(str string) string {
	return alphas.ReplaceAllString(str, "")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SolvePuzzle(inputFile string) {
	f, err := os.Open(inputFile)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum, iter := 0, 1
	for scanner.Scan() {
		strNumber := cleanString(scanner.Text())

		// take first and take last
		trimmedNumber := string(strNumber[0]) + string(strNumber[len(strNumber)-1])
		fmt.Printf("%d: %s\n", iter, trimmedNumber)
		iter++

		// convert
		i, err := strconv.Atoi(trimmedNumber)
		check(err)
		sum += i
	}

	fmt.Printf("Sum of Configuration Values: %d", sum)
}
