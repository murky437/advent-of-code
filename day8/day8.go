package day8

import (
	"bufio"
	"cmp"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type JunctionBox struct {
	X, Y, Z      int64
	CircuitIndex int
}

type Pair struct {
	A, B     int
	Distance int64
}

func Run(inputFile string, numOfConnections int) int64 {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var points []*JunctionBox
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, _ := strconv.ParseInt(parts[0], 10, 64)
		y, _ := strconv.ParseInt(parts[1], 10, 64)
		z, _ := strconv.ParseInt(parts[2], 10, 64)
		points = append(points, &JunctionBox{x, y, z, len(points)})
	}

	// Pair up all points and calculate the distance between them
	var pairs []Pair
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dx := points[i].X - points[j].X
			dy := points[i].Y - points[j].Y
			dz := points[i].Z - points[j].Z
			dist := dx*dx + dy*dy + dz*dz // Distance squared, because we only need it for comparison
			pairs = append(pairs, Pair{i, j, dist})
		}
	}

	// Sort all pairs ascending by distance
	slices.SortFunc(pairs, func(p1, p2 Pair) int {
		return cmp.Compare(p1.Distance, p2.Distance)
	})

	// Initialize circuits
	circuits := make([][]*JunctionBox, len(points))
	for i, p := range points {
		circuits[i] = []*JunctionBox{p}
	}

	connections := 0
	for _, pair := range pairs {
		if connections >= numOfConnections {
			break
		}
		a := points[pair.A]
		b := points[pair.B]

		// Merge if in different circuits
		aIndex := a.CircuitIndex
		bIndex := b.CircuitIndex
		if aIndex != bIndex {
			for _, box := range circuits[bIndex] {
				box.CircuitIndex = aIndex
				circuits[aIndex] = append(circuits[aIndex], box)
			}
			circuits[bIndex] = nil
		}

		connections++
	}

	// Sort circuits descending by size
	slices.SortFunc(circuits, func(a, b []*JunctionBox) int {
		return len(b) - len(a)
	})

	return int64(len(circuits[0])) * int64(len(circuits[1])) * int64(len(circuits[2]))
}
