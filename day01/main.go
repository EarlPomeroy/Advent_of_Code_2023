package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Define a map for spelled-out numbers
var spelledNumbers = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

type calibration struct {
	firstNumber  int
	secondNumber int
}

func (c calibration) getNumber() int {
	var num = 0
	num += c.firstNumber * 10
	num += c.secondNumber

	return num
}

func findNumberString(fragment string) (int, error) {
	for key, val := range spelledNumbers {
		if strings.HasPrefix(fragment, key) {
			return val, nil
		}
	}

	return 0, errors.New("number not found")
}

func makeCalibration(line string) calibration {
	var c = calibration{}
	chars := strings.SplitAfter(line, "")

	for pos, ch := range chars {
		num, err := strconv.Atoi(ch)

		if err == nil {
			if c.firstNumber == 0 {
				c.firstNumber = num
				c.secondNumber = num
			} else {
				c.secondNumber = num
			}
		} else {
			num, err = findNumberString(line[pos:])
			if err == nil {
				if c.firstNumber == 0 {
					c.firstNumber = num
					c.secondNumber = num
				} else {
					c.secondNumber = num
				}
			}
		}
	}

	return c
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

	var result = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 1 {
			c := makeCalibration(line)
			result += c.getNumber()
		}
	}

	fmt.Println("Result: ", result)
}
