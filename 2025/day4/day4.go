package day4

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

	var data [][]byte

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		row := []byte(line)
		data = append(data, row)
	}

	numOfAccessibleRolls := int64(0)

	for row := 0; row < len(data); row++ {
		for col := 0; col < len(data[row]); col++ {
			if data[row][col] != '@' {
				continue
			}
			numOfAdjacentRolls := 0
			if row > 0 && col > 0 && data[row-1][col-1] == '@' {
				numOfAdjacentRolls++
			}
			if row > 0 && data[row-1][col] == '@' {
				numOfAdjacentRolls++
			}
			if row > 0 && col < len(data[row])-1 && data[row-1][col+1] == '@' {
				numOfAdjacentRolls++
			}
			if col > 0 && data[row][col-1] == '@' {
				numOfAdjacentRolls++
			}
			if col < len(data[row])-1 && data[row][col+1] == '@' {
				numOfAdjacentRolls++
			}
			if row < len(data)-1 && col > 0 && data[row+1][col-1] == '@' {
				numOfAdjacentRolls++
			}
			if row < len(data)-1 && data[row+1][col] == '@' {
				numOfAdjacentRolls++
			}
			if row < len(data)-1 && col < len(data[row])-1 && data[row+1][col+1] == '@' {
				numOfAdjacentRolls++
			}

			if numOfAdjacentRolls < 4 {
				numOfAccessibleRolls += 1
			}
		}
	}

	return numOfAccessibleRolls
}
