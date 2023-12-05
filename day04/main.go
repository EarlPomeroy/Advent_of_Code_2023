package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	winningNumbers []int
	cardNumbers    []int
	points         int
	matches        int
}

func (c *Card) calculateMatches() {
	var matches = 0

	for _, num := range c.cardNumbers {
		if contains(c.winningNumbers, num) {
			matches += 1
		}
	}

	c.matches = matches
}

func getInstanceTotal(index int) {
	val, ok := instances[index]
	if ok {
		instances[index] = val + 1
	} else {
		instances[index] = 1
	}

	for i := 1; i <= cards[index].matches; i++ {
		getInstanceTotal(index + i)
	}

}

func (c *Card) calculatePoints() {
	var points = 0

	for _, num := range c.cardNumbers {
		if contains(c.winningNumbers, num) {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}

	c.points = points
}

func contains(list []int, num int) bool {
	for _, n := range list {
		if n == num {
			return true
		}
	}

	return false
}

func NewCard(line string) Card {
	var winningNumbers []int
	var cardNumbers []int

	data := strings.Split(line, ":")
	winNumStrs := strings.TrimSpace(strings.Split(data[1], "|")[0])
	cardNumStrs := strings.TrimSpace(strings.Split(data[1], "|")[1])

	for _, num := range strings.Split(winNumStrs, " ") {
		num, err := strconv.Atoi(num)

		if err == nil {
			winningNumbers = append(winningNumbers, num)
		}
	}

	for _, num := range strings.Split(cardNumStrs, " ") {
		num, err := strconv.Atoi(num)

		if err == nil {
			cardNumbers = append(cardNumbers, num)
		}
	}

	card := Card{winningNumbers: winningNumbers, cardNumbers: cardNumbers}

	return card
}

var cards []Card
var instances = map[int]int{}

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
	for scanner.Scan() {
		line := scanner.Text()
		c := NewCard(line)
		c.calculatePoints()
		c.calculateMatches()
		cards = append(cards, c)
	}

	var totalPoints = 0
	var totalInstances = 0

	for i, c := range cards {
		totalPoints += c.points
		getInstanceTotal(i)
	}

	for _, v := range instances {
		totalInstances += v
	}

	fmt.Println("Total points: ", totalPoints)
	fmt.Println("Total instances: ", totalInstances)
}
