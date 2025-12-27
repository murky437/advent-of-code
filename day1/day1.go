package day1

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
		var err error
		currentPosition, err = getNewPosition(currentPosition, scanner.Text())
		if err != nil {
			panic(err)
		}
		if currentPosition == 0 {
			zeroes++
		}
	}

	return zeroes
}

func getNewPosition(currentPosition int, action string) (int, error) {
	var newPosition int

	direction := action[:1]
	number, err := strconv.Atoi(action[1:])
	if err != nil {
		return newPosition, err
	}

	switch direction {
	case "L":
		newPosition = currentPosition - number
	case "R":
		newPosition = currentPosition + number
	default:
		return newPosition, fmt.Errorf("invalid direction: %s", direction)
	}

	if newPosition < 0 {
		newPosition = (newPosition%100 + 100) % 100
	} else if newPosition >= 100 {
		newPosition = newPosition % 100
	}

	return newPosition, nil
}
