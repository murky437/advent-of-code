package day9

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int64
}

func Run(inputFile string) int64 {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var points []Point

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		points = append(points, Point{x, y})
	}

	largestArea := int64(0)

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			xDistance := points[i].X - points[j].X
			if xDistance < 0 {
				xDistance = -xDistance
			}
			xDistance += 1

			yDistance := points[i].Y - points[j].Y
			if yDistance < 0 {
				yDistance = -yDistance
			}
			yDistance += 1

			area := xDistance * yDistance

			if area > largestArea {
				largestArea = area
			}
		}
	}

	return largestArea
}
