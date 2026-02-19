package day3

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Run(inputFile string) int64 {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	totalOutputJoltage := int64(0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		firstNum := byte('0')
		firstNumIndex := 0
		for i := 0; i < len(line)-1; i++ {
			currentNum := line[i]
			if currentNum > firstNum {
				firstNum = currentNum
				firstNumIndex = i
			}
		}
		secondNum := byte('0')
		for i := firstNumIndex + 1; i < len(line); i++ {
			currentNum := line[i]
			if currentNum > secondNum {
				secondNum = currentNum
			}
		}
		maxLineJoltage, err := strconv.ParseInt(string([]byte{firstNum, secondNum}), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		totalOutputJoltage += maxLineJoltage
	}

	return totalOutputJoltage
}
