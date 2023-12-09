package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Node struct {
	Left  string
	Right string
}

func (n Node) getNextNode(dir string) string {
	if dir == "L" {
		return n.Left
	}

	return n.Right
}

var directions = []string{}
var maps = map[string]Node{}
var startNode = []string{}

func buildMap(line string) {
	var key = strings.TrimSpace(strings.Split(line, "=")[0])
	var values = strings.TrimSpace(strings.Split(line, "=")[1])
	values = values[1 : len(values)-1]

	left := strings.Split(values, ",")[0]
	right := strings.Split(values, ",")[1]

	if key[len(key)-1] == 'Z' {
		startNode = append(startNode, key)
	}

	maps[key] = Node{Left: strings.TrimSpace(left), Right: strings.TrimSpace(right)}
}

func solver(start string) int {
	var instruction = 0
	var mapNode = start
	var count = 0

	for {
		count += 1
		mapNode = maps[mapNode].getNextNode(directions[instruction])

		if mapNode[len(mapNode)-1] == 'Z' {
			break
		} else {
			instruction = (instruction + 1) % len(directions)
		}
	}

	return count
}

func primeFactors(n int) []int {
	var factors []int
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
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

	scanner.Scan()
	directions = strings.Split(scanner.Text(), "")
	scanner.Scan()

	for scanner.Scan() {
		var line = scanner.Text()
		buildMap(line)
	}

	// fmt.Printf("%+v", directions)
	// fmt.Printf("%+v", maps)

	resultMap := map[int]int{}

	for _, start := range startNode {
		steps := solver(start)
		factors := primeFactors(steps)
		for _, factor := range factors {
			resultMap[factor] += 1
		}
	}

	result := 1

	for key := range resultMap {
		result *= key
	}

	println("Steps: ", result)
}
