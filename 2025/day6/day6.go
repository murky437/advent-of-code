package day6

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

	grandTotal := int64(0)

	var numbers [][]string
	var operators []string

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0 {
			log.Fatal("Empty line")
		}
		_, err := strconv.Atoi(parts[0])
		if err != nil {
			operators = parts
			break
		}
		numbers = append(numbers, parts)
	}

	for i, operator := range operators {
		var total int64
		if operator == "*" {
			total = 1
		}
		for j := 0; j < len(numbers); j++ {
			number, err := strconv.ParseInt(numbers[j][i], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			if operator == "*" {
				total *= number
			} else if operator == "+" {
				total += number
			}
		}
		grandTotal += total
	}

	return grandTotal
}
