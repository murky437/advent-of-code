package day2_2

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

		for id := startId; id <= stopId; id++ {
			iString := strconv.FormatInt(id, 10)

		CheckLoop:
			for partLen := 1; partLen <= len(iString)/2; partLen++ {
				if len(iString)%partLen != 0 {
					continue
				}

				numParts := len(iString) / partLen
				firstPart := iString[:partLen]

				for partIndex := 1; partIndex < numParts; partIndex++ {
					if iString[partIndex*partLen:partIndex*partLen+partLen] != firstPart {
						continue CheckLoop
					}
				}

				sumOfInvalidIds += id
				break
			}
		}
	}

	return sumOfInvalidIds
}
