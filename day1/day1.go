package day1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/MCunanan/advent-of-code-2023/util"
)

var alphas = regexp.MustCompile(`[a-zA-Z]+`)

type NumericMapping struct {
	Text  string
	Digit string
}

var numericMap = map[int]NumericMapping{
	1: {"one", "on1e"},
	2: {"two", "tw2o"},
	3: {"three", "thre3e"},
	4: {"four", "4"},
	5: {"five", "fiv5e"},
	6: {"six", "6"},
	7: {"seven", "seve7n"},
	8: {"eight", "eigh8t"},
	9: {"nine", "nin9e"},
}

func convertStringDigits(str string) string {
	var keys []int
	for k := range numericMap {
		keys = append(keys, k)
	}

	for _, mapping := range numericMap {
		str = strings.ReplaceAll(str, mapping.Text, mapping.Digit)
	}
	return str
}

func cleanString(str string) string {
	return alphas.ReplaceAllString(str, "")
}

func SolvePuzzle(inputFile string) {
	f, err := os.Open(inputFile)
	util.Check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum, iter := 0, 0
	for scanner.Scan() {
		scanned := scanner.Text()
		converted := convertStringDigits(scanned)
		strNumber := cleanString(converted)

		// take first and take last
		trimmedNumber := string(strNumber[0]) + string(strNumber[len(strNumber)-1])
		iter++

		// convert
		i, err := strconv.Atoi(trimmedNumber)
		util.Check(err)
		sum += i
		fmt.Printf("%d: %s - %s - %s - %s - %d\n", iter, scanned, converted, strNumber, trimmedNumber, sum)
	}

	fmt.Printf("Sum of Configuration Values: %d", sum)
}
