package day11

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run(inputFile string) int {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	presentSizes := make(map[int]int)
	var currentPresentIndex int
	var currentPresentSize int

	var okRegions int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		headerParts := strings.Split(parts[0], "x")

		if len(parts) == 2 {
			if len(headerParts) == 1 {
				currentPresentIndex, err = strconv.Atoi(headerParts[0])
				if err != nil {
					log.Fatal(err)
				}
			} else {
				width, err := strconv.Atoi(headerParts[1])
				if err != nil {
					log.Fatal(err)
				}
				height, err := strconv.Atoi(headerParts[0])
				if err != nil {
					log.Fatal(err)
				}

				regionSize := width * height

				presentRequirements := strings.Fields(parts[1])
				var naiveSize int
				for i := 0; i < len(presentRequirements); i++ {
					presentRequirement, err := strconv.Atoi(presentRequirements[i])
					if err != nil {
						log.Fatal(err)
					}
					naiveSize += presentRequirement * presentSizes[i]
				}
				if naiveSize <= regionSize {
					okRegions++
				}
			}
		} else if line == "" {
			presentSizes[currentPresentIndex] = currentPresentSize
			currentPresentSize = 0
		} else {
			currentPresentSize += len(line)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return okRegions
}
