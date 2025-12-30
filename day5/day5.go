package day5

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run(inputFile string) int64 {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var ranges []Range

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		startAndStop := strings.Split(line, "-")
		start, err := strconv.ParseInt(startAndStop[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		stop, err := strconv.ParseInt(startAndStop[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		ranges = append(ranges, Range{start, stop})
	}

	var numOfFreshIngredients int64

	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if Contains(ranges, number) {
			numOfFreshIngredients++
		}
	}

	return numOfFreshIngredients
}

type Range struct {
	Start int64
	End   int64
}

func Contains(ranges []Range, x int64) bool {
	for _, r := range ranges {
		if x >= r.Start && x <= r.End {
			return true
		}
	}
	return false
}
