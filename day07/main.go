package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Define a map for spelled-out numbers
var cardValues = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

const (
	high_card  = iota // 0
	one_pair   = iota // 1
	two_pair   = iota // 2
	three_kind = iota // 3
	full_house = iota // 4
	four_kind  = iota // 5
	five_kind  = iota // 6
)

type Hand struct {
	Cards string
	Bid   int
	Kind  int
}

type HandList []Hand

func (h *Hand) classify() {
	handValue := map[rune]int{}

	for _, c := range h.Cards {
		value, ok := handValue[c]

		if ok {
			handValue[c] = value + 1
		} else {
			handValue[c] = 1
		}
	}

	switch len(handValue) {
	case 5:
		h.Kind = high_card
	case 4:
		h.Kind = one_pair
	case 3:
		for _, value := range handValue {
			if value == 2 {
				h.Kind = two_pair
				break
			}
			if value == 3 {
				h.Kind = three_kind
				break
			}
		}
	case 2:
		for _, value := range handValue {
			if value == 4 || value == 1 {
				h.Kind = four_kind
				break
			} else {
				h.Kind = full_house
				break
			}
		}
	case 1:
		h.Kind = five_kind

	}
}

// Go sort interface Len, Less, and Swap
func (h HandList) Len() int {
	return len(cardList)
}

func (h HandList) Less(i, j int) bool {
	if cardList[i].Kind == cardList[j].Kind {
		for k := 0; k < 5; k++ {
			if cardList[i].Cards[k] == cardList[j].Cards[k] {
				continue
			} else {
				return cardValues[string(cardList[i].Cards[k])] < cardValues[string(cardList[j].Cards[k])]
			}
		}
	}

	return cardList[i].Kind < cardList[j].Kind
}

func (h HandList) Swap(i, j int) {
	tmp := cardList[i]
	cardList[i] = cardList[j]
	cardList[j] = tmp
}

var cardList []Hand

func makeHand(line string) Hand {
	h := Hand{}
	h.Cards = strings.Split(line, " ")[0]
	num, _ := strconv.Atoi(strings.Split(line, " ")[1])
	h.Bid = num

	h.classify()

	return h
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

	for scanner.Scan() {
		var line = scanner.Text()
		hand := makeHand(line)
		cardList = append(cardList, hand)
	}

	sort.Sort(HandList(cardList))
	fmt.Printf("%+v\n", cardList)

	var totalWinnings = 0

	for i, c := range cardList {
		totalWinnings += c.Bid * (i + 1)
	}

	fmt.Println("Total Winnings: ", totalWinnings)
}
