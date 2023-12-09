package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type History struct {
	sequence []int
	base     *History
}

func (h *History) addBase() {
	done := true
	base := History{}

	for i := 0; i < len(h.sequence)-1; i++ {
		first := h.sequence[i]
		second := h.sequence[i+1]

		num := second - first

		if num != 0 {
			done = false
		}

		base.sequence = append(base.sequence, num)
	}

	h.base = &base

	if !done {
		base.addBase()
	}
}

func (h *History) extrapolate() int {
	for {
		if h.base != nil {
			num := h.base.extrapolate()
			return h.sequence[len(h.sequence)-1] + num
		} else {
			return 0
		}
	}
}

func (h *History) reverseExtrapolate() int {
	for {
		if h.base != nil {
			num := h.base.reverseExtrapolate()
			return h.sequence[0] - num
		} else {
			return 0
		}
	}
}

func (h History) print() {
	fmt.Printf("%v\n\t", h.sequence)

	if h.base != nil {
		h.base.print()
	} else {
		fmt.Println("")
	}
}

var HistoryList = []History{}

func makeHistory(line string) {
	h := History{}

	for _, s := range strings.Split(line, " ") {
		num, _ := strconv.Atoi(s)
		h.sequence = append(h.sequence, num)
	}

	h.addBase()

	HistoryList = append(HistoryList, h)
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
		makeHistory(line)
	}

	result := 0

	for _, h := range HistoryList {
		result += h.reverseExtrapolate()
		// h.print()
	}

	fmt.Println("History total: ", result)
}
