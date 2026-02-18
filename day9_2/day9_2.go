package day9_2

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const Up = "Up"
const Right = "Right"
const Down = "Down"
const Left = "Left"

type Point struct {
	X, Y float64
}

type Line struct {
	Start, End Point
}

type Rectangle struct {
	MinX, MaxX, MinY, MaxY float64
}

func NewRectangle(p1, p2 Point) Rectangle {
	return Rectangle{
		MinX: math.Min(p1.X, p2.X),
		MaxX: math.Max(p1.X, p2.X),
		MinY: math.Min(p1.Y, p2.Y),
		MaxY: math.Max(p1.Y, p2.Y),
	}
}

func (r Rectangle) IntersectsLine(l Line) bool {
	return LiangBarsky(r, l)
}

func (r Rectangle) Area() float64 {
	width := r.MaxX - r.MinX + 1
	height := r.MaxY - r.MinY + 1
	return width * height
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
		points = append(points, Point{float64(x), float64(y)})
	}

	boundingLines := getBoundingLines(points)

	var largestArea float64

	for i := 0; i < len(points); i++ {
	OuterLoop:
		for j := i + 1; j < len(points); j++ {
			r := NewRectangle(points[i], points[j])
			area := r.Area()

			if area < largestArea {
				continue
			}

			for _, l := range boundingLines {
				if r.IntersectsLine(l) {
					continue OuterLoop
				}
			}

			largestArea = area
		}
	}

	return int64(largestArea)
}

func getBoundingLines(polygon []Point) []Line {
	n := len(polygon)
	if n < 4 {
		return nil
	}

	var boundingLines []Line

	for i := 0; i < n; i++ {
		prev := polygon[(i+n-1)%n]
		current := polygon[i]
		next := polygon[(i+1)%n]
		next2 := polygon[(i+2)%n]

		prevDir := getDirection(prev, current)
		currentDir := getDirection(current, next)
		nextDir := getDirection(next, next2)

		boundingLineStart := calcBoundingPoint(current, prevDir, currentDir)
		boundingLineEnd := calcBoundingPoint(next, currentDir, nextDir)

		boundingLines = append(boundingLines, Line{boundingLineStart, boundingLineEnd})
	}

	return boundingLines
}

func calcBoundingPoint(p Point, prevDir string, currentDir string) Point {
	var boundingPoint Point

	switch prevDir {
	case Up:
		switch currentDir {
		case Left:
			boundingPoint = Point{
				X: p.X - 0.5,
				Y: p.Y + 0.5,
			}
		case Up:
			boundingPoint = Point{
				X: p.X - 0.5,
				Y: p.Y,
			}
		case Right:
			boundingPoint = Point{
				X: p.X - 0.5,
				Y: p.Y - 0.5,
			}
		}
	case Right:
		switch currentDir {
		case Up:
			boundingPoint = Point{
				X: p.X - 0.5,
				Y: p.Y - 0.5,
			}
		case Right:
			boundingPoint = Point{
				X: p.X,
				Y: p.Y - 0.5,
			}
		case Down:
			boundingPoint = Point{
				X: p.X + 0.5,
				Y: p.Y - 0.5,
			}
		}
	case Down:
		switch currentDir {
		case Right:
			boundingPoint = Point{
				X: p.X + 0.5,
				Y: p.Y - 0.5,
			}
		case Down:
			boundingPoint = Point{
				X: p.X + 0.5,
				Y: p.Y,
			}
		case Left:
			boundingPoint = Point{
				X: p.X + 0.5,
				Y: p.Y + 0.5,
			}
		}
	case Left:
		switch currentDir {
		case Down:
			boundingPoint = Point{
				X: p.X + 0.5,
				Y: p.Y + 0.5,
			}
		case Left:
			boundingPoint = Point{
				X: p.X,
				Y: p.Y + 0.5,
			}
		case Up:
			boundingPoint = Point{
				X: p.X - 0.5,
				Y: p.Y + 0.5,
			}
		}
	}

	return boundingPoint
}

func getDirection(start Point, end Point) string {
	if start.X == end.X {
		if start.Y > end.Y {
			return Up
		}
		return Down
	}
	if start.X < end.X {
		return Right
	}
	return Left
}

// LiangBarsky Liang-Barsky algo to check if line intersects rectangle
func LiangBarsky(rect Rectangle, line Line) bool {
	dx := line.End.X - line.Start.X
	dy := line.End.Y - line.Start.Y

	p := [4]float64{-dx, dx, -dy, dy}
	q := [4]float64{
		line.Start.X - rect.MinX,
		rect.MaxX - line.Start.X,
		line.Start.Y - rect.MinY,
		rect.MaxY - line.Start.Y,
	}

	tEnter := 0.0
	tExit := 1.0

	for i := 0; i < 4; i++ {
		if p[i] == 0 {
			if q[i] < 0 {
				return false
			}
		} else {
			t := q[i] / p[i]
			if p[i] < 0 {
				if t > tEnter {
					tEnter = t
				}
			} else {
				if t < tExit {
					tExit = t
				}
			}
		}
	}

	if tEnter > tExit {
		return false
	}

	return true
}
