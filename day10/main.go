package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	NorthSouth = iota
	EastWest   = iota
	NorthEast  = iota
	NorthWest  = iota
	SouthWest  = iota
	SouthEast  = iota
	Ground     = iota
	Start      = iota
)

var typeLookupMap = map[string]int{
	"|": NorthSouth, // 0
	"-": EastWest,   // 1
	"L": NorthEast,  // 2
	"J": NorthWest,  // 3
	"7": SouthWest,  // 4
	"F": SouthEast,  // 5
	".": Ground,     // 6
	"S": Start,      // 7
}

type Point struct {
	x        int
	y        int
	nodeType int
	depth    int
	visited  bool
}

func (p *Point) getNextPoints() []*Point {
	nextPoints := []*Point{}
	var upPoint, downPoint, leftPoint, rightPoint *Point

	switch p.nodeType {
	case Start:
		upPoint = findPoint(p.x, p.y-1, []int{NorthSouth, SouthEast, SouthWest})
		downPoint = findPoint(p.x, p.y+1, []int{NorthSouth, NorthEast, NorthWest})
		leftPoint = findPoint(p.x-1, p.y, []int{EastWest, SouthEast, NorthEast})
		rightPoint = findPoint(p.x+1, p.y, []int{EastWest, SouthWest, NorthWest})
	case NorthSouth:
		upPoint = findPoint(p.x, p.y-1, []int{NorthSouth, SouthEast, SouthWest})
		downPoint = findPoint(p.x, p.y+1, []int{NorthSouth, NorthEast, NorthWest})
	case EastWest:
		leftPoint = findPoint(p.x-1, p.y, []int{EastWest, SouthEast, NorthEast})
		rightPoint = findPoint(p.x+1, p.y, []int{EastWest, SouthWest, NorthWest})
	case NorthEast:
		upPoint = findPoint(p.x, p.y-1, []int{NorthSouth, SouthEast, SouthWest})
		rightPoint = findPoint(p.x+1, p.y, []int{EastWest, SouthWest, NorthWest})
	case NorthWest:
		upPoint = findPoint(p.x, p.y-1, []int{NorthSouth, SouthEast, SouthWest})
		leftPoint = findPoint(p.x-1, p.y, []int{EastWest, SouthEast, NorthEast})
	case SouthEast:
		downPoint = findPoint(p.x, p.y+1, []int{NorthSouth, NorthEast, NorthWest})
		rightPoint = findPoint(p.x+1, p.y, []int{EastWest, SouthWest, NorthWest})
	case SouthWest:
		downPoint = findPoint(p.x, p.y+1, []int{NorthSouth, NorthEast, NorthWest})
		leftPoint = findPoint(p.x-1, p.y, []int{EastWest, SouthEast, NorthEast})
	}

	if upPoint != nil {
		upPoint.depth = p.depth + 1
		nextPoints = append(nextPoints, upPoint)
	}

	if downPoint != nil {
		downPoint.depth = p.depth + 1
		nextPoints = append(nextPoints, downPoint)
	}

	if leftPoint != nil {
		leftPoint.depth = p.depth + 1
		nextPoints = append(nextPoints, leftPoint)
	}

	if rightPoint != nil {
		rightPoint.depth = p.depth + 1
		nextPoints = append(nextPoints, rightPoint)
	}

	return nextPoints
}

func print() {
	for _, pl := range points {
		for _, p := range pl {
			if p.nodeType == Ground {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", p.depth)
			}
		}
		fmt.Printf("\n")
	}

	fmt.Println()
}

func recreateMap() {
	for _, pl := range points {
		for _, p := range pl {
			switch p.nodeType {
			case NorthSouth:
				fmt.Printf("|")
			case EastWest:
				fmt.Printf("-")
			case NorthEast:
				fmt.Printf("L")
			case NorthWest:
				fmt.Printf("J")
			case SouthEast:
				fmt.Printf("F")
			case SouthWest:
				fmt.Printf("7")
			case Start:
				fmt.Printf("S")
			default:
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Println()
}

var points [][]*Point
var max = 0

func findPoint(x, y int, expected []int) *Point {
	if x == -1 || y == -1 || y > len(points)-1 || x > len(points[y])-1 {
		return nil
	}

	newPoint := points[y][x]
	if newPoint.visited || newPoint.nodeType == Ground {
		return nil
	}

	for _, nodeType := range expected {
		if newPoint.nodeType == nodeType {
			return newPoint
		}
	}

	return nil
}

func makePoint(ch string, x, y int) Point {
	point := Point{x: x, y: y, nodeType: typeLookupMap[ch]}

	return point
}

func visitor(queue []*Point) {
	if len(queue) == 0 {
		return
	}

	p, q := queue[0], queue[1:]

	if p.depth > max {
		max = p.depth
	}

	p.visited = true
	for _, np := range p.getNextPoints() {
		q = append(q, np)
	}

	visitor(q)
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Missing input file")
		return
	}

	var filename = args[0]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var start *Point

	for scanner.Scan() {
		var line = scanner.Text()
		var row []*Point

		for x, ch := range line {
			point := makePoint(string(ch), x, len(points))
			if point.nodeType == Start {
				start = &point
			}

			row = append(row, &point)
		}

		points = append(points, row)
	}

	// fmt.Printf("%+v\n", start)
	// fmt.Printf("%+v\n", points)

	visitor([]*Point{start})

	print()
	recreateMap()

	fmt.Println("Result: ", max)
}
