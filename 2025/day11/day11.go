package day11

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Run(inputFile string) int64 {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	devices := make(map[string][]string)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		devices[parts[0]] = strings.Fields(parts[1])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var numOfPaths int64

	var followDevice func(name string)
	followDevice = func(name string) {
		for _, output := range devices[name] {
			if output == "you" {
				continue
			}
			if output == "out" {
				numOfPaths++
				continue
			}
			followDevice(output)
		}
	}
	followDevice("you")

	return numOfPaths
}
