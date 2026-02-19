package day6_2

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

	var numberLines []string
	var columnGroups []ColumnGroup

	for scanner.Scan() {
		line := scanner.Text()
		firstChar := line[0]
		if firstChar == '+' || firstChar == '*' {
			columnGroups = parseOperatorLine(line)
			break
		}
		numberLines = append(numberLines, line)
	}

	lastLen := 0
	for _, line := range numberLines {
		parts := strings.Fields(line)
		last := parts[len(parts)-1]
		if len(last) > lastLen {
			lastLen = len(last)
		}
	}
	columnGroups[len(columnGroups)-1].NumberLen = lastLen

	parsePos := 0

	for _, columnGroup := range columnGroups {
		var total int64
		if columnGroup.Operator == "*" {
			total = 1
		}
		for n := 0; n < columnGroup.NumberLen; n++ {
			numberString := ""
			for j := 0; j < len(numberLines); j++ {
				numberString = numberString + numberLines[j][parsePos+n:parsePos+n+1]
			}
			number, err := strconv.ParseInt(strings.TrimSpace(numberString), 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			if columnGroup.Operator == "*" {
				total *= number
			} else if columnGroup.Operator == "+" {
				total += number
			}
		}
		parsePos = parsePos + columnGroup.NumberLen + 1
		grandTotal += total
	}

	return grandTotal
}

type ColumnGroup struct {
	Operator  string
	NumberLen int
}

func parseOperatorLine(line string) []ColumnGroup {
	var columnGroups []ColumnGroup

	numberLen := 0
	operator := ""

	for _, char := range line {
		numberLen++
		if char != ' ' {
			if operator != "" {
				columnGroups = append(columnGroups, ColumnGroup{operator, numberLen - 2})
				numberLen = 1
			}
			operator = string(char)
		}
	}
	columnGroups = append(columnGroups, ColumnGroup{operator, 0})

	return columnGroups
}
