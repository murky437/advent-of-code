package day7_2

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

	counts := make([][]int64, len(lines))
	for i := range lines {
		counts[i] = make([]int64, len(lines[i]))
	}

	for i, line := range lines {
		if i == 0 {
			for j, char := range line {
				if char == 'S' {
					lines[i+1][j] = '|'
					counts[i+1][j]++
				}
			}
		}
		if i == len(lines)-1 {
			continue
		}
		for j, char := range line {
			if char == '|' {
				prevCount := counts[i][j]
				if lines[i+1][j] == '^' {
					if j-1 >= 0 {
						lines[i+1][j-1] = '|'
						counts[i+1][j-1] += prevCount
					}
					if j+1 < len(line) {
						lines[i+1][j+1] = '|'
						counts[i+1][j+1] += prevCount
					}
				} else {
					lines[i+1][j] = '|'
					counts[i+1][j] += prevCount
				}
			}
		}
	}

	var numOfTimelines int64

	for _, count := range counts[len(counts)-1] {
		numOfTimelines += count
	}

	return numOfTimelines
}
