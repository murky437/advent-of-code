package day11_2

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Run(inputFile string) int64 {
	devices := parseInputFile(inputFile)
	savedNumOfPaths := make(map[pair]int64)

	if dacToFft := getNumOfPaths(devices, "dac", "fft", savedNumOfPaths); dacToFft != 0 {
		svrToDac := getNumOfPaths(devices, "srv", "dac", savedNumOfPaths)
		fftToOut := getNumOfPaths(devices, "fft", "out", savedNumOfPaths)
		return svrToDac * dacToFft * fftToOut
	}

	fftToDac := getNumOfPaths(devices, "fft", "dac", savedNumOfPaths)
	svrToFft := getNumOfPaths(devices, "svr", "fft", savedNumOfPaths)
	dacToOut := getNumOfPaths(devices, "dac", "out", savedNumOfPaths)

	return svrToFft * fftToDac * dacToOut
}

func parseInputFile(inputFile string) map[string][]string {
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

	return devices
}

type pair struct {
	start, stop string
}

func getNumOfPaths(devices map[string][]string, start string, stop string, savedNumOfPaths map[pair]int64) int64 {
	if v, ok := savedNumOfPaths[pair{start, stop}]; ok {
		return v
	}
	if start == stop {
		return 1
	}
	var numOfPaths int64
	for _, next := range devices[start] {
		numOfPaths += getNumOfPaths(devices, next, stop, savedNumOfPaths)
	}
	savedNumOfPaths[pair{start, stop}] = numOfPaths
	return numOfPaths
}
