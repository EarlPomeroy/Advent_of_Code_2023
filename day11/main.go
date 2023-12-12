package main

import (
	"bufio"
	"fmt"
	"os"
)

type Galaxy struct {
	x int
	y int
}

func (g *Galaxy) String() string {
	return fmt.Sprintf("{X: %d, Y: %d} ", g.x, g.y)
}

func (g *Galaxy) moveGalaxy() {
	xCount := 0
	yCount := 0
	for _, r := range rows {
		if g.y <= r {
			break
		}

		yCount += 999999
	}

	for _, c := range cols {
		if g.x <= c {
			break
		}

		xCount += 999999
	}

	g.x += xCount
	g.y += yCount
}

var galaxies = []*Galaxy{}
var cols = []int{}
var rows = []int{}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func remove(arr []int, num int) []int {
	result := []int{}
	for _, value := range arr {
		if value != num {
			result = append(result, value)
		}
	}
	return result
}

func calculateDistance() int {
	distances := 0
	for i, g := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			d := abs(galaxies[j].x-g.x) + abs(galaxies[j].y-g.y)
			distances += d
			fmt.Printf("Galaxy %d, Galaxy %d = dist %d\n", i+1, j+1, d)
		}
	}

	return distances

}

func expandUniverse() {
	for _, g := range galaxies {
		g.moveGalaxy()
	}
}

func evaluateUniverse(line string, y int) {
	clean := true
	for i, ch := range line {
		if string(ch) == "#" {
			cols = remove(cols, i)
			g := Galaxy{x: i, y: y}
			galaxies = append(galaxies, &g)
			clean = false
		}
	}

	if clean {
		rows = append(rows, y)
	}
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

	row := 0

	for scanner.Scan() {
		var line = scanner.Text()

		if len(cols) == 0 {
			for i := 0; i < len(line); i++ {
				cols = append(cols, i)
			}
		}
		evaluateUniverse(line, row)
		row++
	}

	//fmt.Printf("rows: %v\n", rows)
	//fmt.Printf("cols: %v\n", cols)
	fmt.Printf("%+v\n", galaxies)

	expandUniverse()

	fmt.Printf("%+v\n", galaxies)

	fmt.Println("Result: ", calculateDistance())
}
