package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	var result = []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		id, valid := validGame(line)
		if valid {
			result = append(result, id)
		}
	}

	var value = 0
	for _, id := range result {
		value += id
	}

	fmt.Println("Result: ", value)
}
