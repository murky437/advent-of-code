package day7

import (
	"bufio"
	"log"
	"os"
)

func Run(inputFile string) int64 {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []rune(line))
	}

	numOfSplits := int64(0)

	for i, line := range lines {
		if i == 0 {
			for j, char := range line {
				if char == 'S' {
					lines[i+1][j] = '|'
				}
			}
		}
		if i == len(lines)-1 {
			continue
		}
		for j, char := range line {
			if char == '|' {
				if lines[i+1][j] == '^' {
					numOfSplits++
					if j-1 >= 0 {
						lines[i+1][j-1] = '|'
					}
					if j+1 < len(line) {
						lines[i+1][j+1] = '|'
					}
				} else {
					lines[i+1][j] = '|'
				}
			}
		}
	}

	return numOfSplits
}
