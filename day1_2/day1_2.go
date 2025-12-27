package day1_2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Run(inputFile string) int {
	zeroes := 0
	currentPosition := 50

	f, _ := os.Open(inputFile)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var newZeroes int
		var err error
		currentPosition, newZeroes, err = getNewPosition(currentPosition, scanner.Text())
		if err != nil {
			panic(err)
		}
		zeroes += newZeroes
	}

	return zeroes
}

func getNewPosition(currentPosition int, action string) (int, int, error) {
	var newZeroes int

	direction := action[:1]
	number, err := strconv.Atoi(action[1:])
	if err != nil {
		return currentPosition, newZeroes, err
	}

	switch direction {
	case "L":
		for i := 0; i < number; i++ {
			currentPosition--
			if currentPosition == 0 {
				newZeroes++
			} else if currentPosition < 0 {
				currentPosition = currentPosition + 100
			}
		}
	case "R":
		for i := 0; i < number; i++ {
			currentPosition++
			if currentPosition == 100 {
				currentPosition = 0
				newZeroes++
			}
		}
	default:
		return currentPosition, newZeroes, fmt.Errorf("invalid direction: %s", direction)
	}

	return currentPosition, newZeroes, nil
}
