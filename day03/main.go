package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Symbol struct {
	row    int
	col    int
	symbol string
}

type PartNumber struct {
	row   int
	col   int
	value string
}

var symbols []Symbol
var partNumbers []PartNumber

func (s Symbol) isAdjacent(pn PartNumber) bool {
	if pn.row == s.row || pn.row+1 == s.row || pn.row-1 == s.row {
		if (pn.col-1) <= s.col && s.col <= (pn.col+len(pn.value)) {
			return true
		}
	}

	return false
}

func (s Symbol) isGear() bool {
	return s.symbol == "*"
}

func (s Symbol) calculateRatio() int {
	var pns []PartNumber

	for _, pn := range partNumbers {
		if s.isAdjacent(pn) {
			pns = append(pns, pn)
		}
	}

	var ratio = 0

	if len(pns) > 1 {
		ratio = 1

		for _, pn := range pns {
			ratio *= pn.getPartNumber()
		}
	}

	return ratio
}

func (pn PartNumber) adjacent() bool {
	for _, s := range symbols {
		if s.isAdjacent(pn) {
			return true
		}
	}

	return false
}

func (pn PartNumber) getPartNumber() int {
	num, err := strconv.Atoi(pn.value)

	if err == nil {
		return num
	}

	fmt.Printf("Could not covert %s to a number", pn.value)
	return 0
}

func isSymbol(s string) bool {
	if s == "*" ||
		s == "/" ||
		s == "-" ||
		s == "+" ||
		s == "&" ||
		s == "=" ||
		s == "@" ||
		s == "$" ||
		s == "%" ||
		s == "#" ||
		s == "." {
		return true
	}

	return false
}

func parseLine(line string, row int) {
	var numStr = ""

	for i, ch := range strings.Split(line, "") {

		if isSymbol(ch) {
			if ch != "." {
				s := Symbol{row: row, col: i, symbol: ch}
				symbols = append(symbols, s)
			}

			if len(numStr) != 0 {
				p := PartNumber{row: row, col: i - len(numStr), value: numStr}
				partNumbers = append(partNumbers, p)
				numStr = ""
			}
		} else {
			numStr = numStr + ch
		}
	}

	if len(numStr) > 0 {
		p := PartNumber{row: row, col: len(line) - 1 - len(numStr), value: numStr}
		partNumbers = append(partNumbers, p)
	}
}

// Part two

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

	scanner := bufio.NewScanner(file)
	row := -1
	for scanner.Scan() {
		row += 1
		line := scanner.Text()
		parseLine(line, row)
	}

	// Part One
	var result = 0
	for _, pn := range partNumbers {
		if pn.adjacent() {
			result += pn.getPartNumber()
		}
	}

	fmt.Println("Result: ", result)

	// Part Two
	var gearRatio = 0

	for _, s := range symbols {
		if s.isGear() {
			gearRatio += s.calculateRatio()
		}
	}

	fmt.Println("Gear Ratio Total: ", gearRatio)
}
