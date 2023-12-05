package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	destination int
	source      int
	length      int
}

type Map struct {
	name   string
	ranges []Range
	next   *Map
}

func (m *Map) getNextSeed(seed int) int {
	for _, r := range m.ranges {
		if seed >= r.source && seed <= r.source+r.length {
			var diff = seed - r.source
			return r.destination + diff
		}
	}

	return seed
}

func (m *Map) walkMaps(seed int) int {
	var nextSeed = m.getNextSeed(seed)

	if m.next != nil {
		return m.next.walkMaps(nextSeed)
	}

	return nextSeed
}

var seedList []int

var firstMap *Map = nil
var prevMap *Map = nil

func readSeeds(line string) {
	seeds := strings.TrimSpace(strings.Split(line, ":")[1])

	for _, n := range strings.Split(seeds, " ") {
		num, _ := strconv.Atoi(n)
		seedList = append(seedList, num)
	}
}

func readSeedPairs(line string) {
	seeds := strings.TrimSpace(strings.Split(line, ":")[1])
	seedPairs := strings.Split(seeds, " ")

	for i := 0; i < len(seedPairs); i += 2 {
		num, _ := strconv.Atoi(seedPairs[i])
		seed := num

		num, _ = strconv.Atoi(seedPairs[i+1])
		count := num

		for j := 0; j < count; j++ {
			seedList = append(seedList, seed+j)
		}
	}
}

func processMap(lines []string) {
	var m = Map{}
	name := lines[0]
	m.name = name

	for _, line := range lines[1:] {
		var r = Range{}

		for i, n := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(n)
			switch i {
			case 0:
				r.destination = num
			case 1:
				r.source = num
			case 2:
				r.length = num
			}
		}

		m.ranges = append(m.ranges, r)
	}

	if firstMap == nil {
		firstMap = &m
	}

	if prevMap != nil {
		prevMap.next = &m
	}

	prevMap = &m
}

func processFile(file io.Reader) {
	seedsRead := false
	var rangeLines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if !seedsRead {
			//readSeeds(line)
			readSeedPairs(line)
			seedsRead = true
		} else {
			if line == "" {
				if len(rangeLines) > 0 {
					processMap(rangeLines)
					rangeLines = rangeLines[:0]
				}
			} else {
				rangeLines = append(rangeLines, line)
			}
		}
	}

	processMap(rangeLines)
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

	processFile(file)

	var location = math.MaxInt

	for _, seed := range seedList {
		val := firstMap.walkMaps(seed)

		if val < location {
			location = val
		}
	}

	fmt.Println("Result: ", location)
}
