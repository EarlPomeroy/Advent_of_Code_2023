package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Part One
var limits = map[string]int{"red": 12, "green": 13, "blue": 14}

func checkLimit(color string) bool {
	vals := strings.Split(color, " ")
	num, _ := strconv.Atoi(vals[0])

	return limits[vals[1]] >= num
}

func validGame(line string) (int, bool) {
	game := strings.Split(line, ":")
	gameId, _ := strconv.Atoi(game[0][len("Game "):])
	draws := strings.Split(game[1], ";")

	for _, d := range draws {
		for _, color := range strings.Split(d, ",") {
			if !checkLimit(strings.TrimSpace(color)) {
				return 0, false
			}
		}
	}

	return gameId, true
}

// Part two
type game struct {
	red   int
	green int
	blue  int
}

func (g *game) checkGame(color string) {
	vals := strings.Split(color, " ")
	c := vals[1]
	num, _ := strconv.Atoi(vals[0])

	if c == "red" && num > g.red {
		g.red = num
	} else if c == "green" && num > g.green {
		g.green = num
	} else if c == "blue" && num > g.blue {
		g.blue = num
	}
}

func (g *game) getValue() int {
	return g.blue * g.red * g.green
}

func evaluateGame(line string) game {
	gameStr := strings.Split(line, ":")
	draws := strings.Split(gameStr[1], ";")

	var g = game{}

	for _, d := range draws {
		for _, color := range strings.Split(d, ",") {
			g.checkGame(strings.TrimSpace(color))
		}
	}

	return g
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Missing input file")
		return
	}

	var filename = args[0]

	// Open the calibration document
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var result = []game{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// id, valid := validGame(line)
		// if valid {
		// 	result = append(result, id)
		// }

		result = append(result, evaluateGame(line))
	}

	var value = 0
	for _, game := range result {
		value += game.getValue()
	}

	fmt.Println("Result: ", value)
}
