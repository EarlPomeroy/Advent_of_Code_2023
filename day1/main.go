package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calibaration struct {
	firstNumber  int
	secondNumber int
}

func (c calibaration) getNumber() int {
	var num = 0
	num += c.firstNumber * 10
	num += c.secondNumber

	return num
}

func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return "", err
	}
	return string(data), nil
}

func makeCalibration(line string) calibaration {
	var c = calibaration{}
	chars := strings.SplitAfter(line, "")

	for _, ch := range chars {
		i, err := strconv.Atoi(ch)

		if err == nil {

			if c.firstNumber == 0 {
				c.firstNumber = i
				c.secondNumber = i
			} else {
				c.secondNumber = i
			}
		}
	}

	return c
}

func main() {
	data, err := readFile("input1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var result = 0

	for _, line := range strings.Split(data, "\n") {
		if len(line) > 1 {
			c := makeCalibration(line)
			result += c.getNumber()
		}
	}

	fmt.Println("Result: ", result)
}
