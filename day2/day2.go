package day2

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
	scanner.Scan()
	line := scanner.Text()

	var sumOfInvalidIds int64 = 0

	idRanges := strings.Split(line, ",")

	for _, idRange := range idRanges {
		startAndStop := strings.Split(idRange, "-")
		startId, err := strconv.ParseInt(startAndStop[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		stopId, err := strconv.ParseInt(startAndStop[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		for i := startId; i <= stopId; i++ {
			iString := strconv.FormatInt(i, 10)
			if len(iString)%2 != 0 {
				continue
			}
			mid := len(iString) / 2
			left := iString[:mid]
			right := iString[mid:]
			if left == right {
				sumOfInvalidIds += i
			}
		}
	}

	return sumOfInvalidIds
}
