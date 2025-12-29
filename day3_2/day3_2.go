package day3_2

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

	maxDigits := 12

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		digits := make([]byte, maxDigits)
		digitIndexes := make([]int, maxDigits)
		nextDigitStartIndex := 0
		for digitIndex := 0; digitIndex < maxDigits; digitIndex++ {
			for lineIndex := nextDigitStartIndex; lineIndex <= len(line)-maxDigits+digitIndex; lineIndex++ {
				digit := line[lineIndex]
				if digit > digits[digitIndex] {
					digits[digitIndex] = digit
					digitIndexes[digitIndex] = lineIndex
					nextDigitStartIndex = lineIndex + 1
				}
			}
		}
		maxLineJoltage, err := strconv.ParseInt(string(digits), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		totalOutputJoltage += maxLineJoltage
	}

	return totalOutputJoltage
}
