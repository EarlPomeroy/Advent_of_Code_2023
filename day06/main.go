package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateWins(time, dist int) int {
	wins := 0

	for i := 0; i < time; i++ {

		mydist := i * (time - i)

		if mydist > dist {
			wins += 1
		}
	}

	return wins
}

func getArray(line string) []int {
	var arr []int

	for _, d := range strings.Split(line, " ") {
		n, err := strconv.Atoi(d)

		if err == nil {
			arr = append(arr, n)
		}
	}

	return arr
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

	// Read time line
	scanner.Scan()
	timeLine := strings.TrimSpace(strings.Split(scanner.Text(), ":")[1])

	// Read distance line
	scanner.Scan()
	distanceLine := strings.TrimSpace(strings.Split(scanner.Text(), ":")[1])

	var timeArr = getArray(timeLine)
	var distanceArr = getArray(distanceLine)

	total := 1

	for i := 0; i < len(timeArr); i++ {
		val := calculateWins(timeArr[i], distanceArr[i])
		total *= val
	}

	fmt.Println("Result: ", total)
}
