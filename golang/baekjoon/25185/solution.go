// implementation

package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	scanner *bufio.Scanner
	writer  *bufio.Writer
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer = bufio.NewWriter(os.Stdout)
}

type card struct {
	num  int
	suit string
}

type Deck []card

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func scanCard() card {
	scanner.Scan()
	return card{num: int(scanner.Text()[0] - '0'), suit: string(scanner.Text()[1])}
}

func validate(cards Deck) bool {
	// duplicate check
	// duplicate 3 or duplicate 2 2
	set := make(map[card]struct{}, 4)
	for _, c := range cards {
		set[c] = struct{}{}
	}
	// 1 if duplicate 4 cards
	// 2 if duplicate 2-2 or 3 cards
	if len(set) <= 2 {
		return true
	}

	// strait check
	slices.SortFunc(cards, func(i, j card) int {
		if i.suit == j.suit {
			return i.num - j.num
		}
		return strings.Compare(i.suit, j.suit)
	})

	straightCount := 1
	const maxStraightCount = 3
	for i := 0; i < len(cards)-1; i++ {
		if cards[i] == cards[i+1] {
			continue
		}
		if cards[i].num+1 == cards[i+1].num && cards[i].suit == cards[i+1].suit {
			straightCount++
			if straightCount == maxStraightCount {
				return true
			}
		} else {
			straightCount = 1
		}
	}

	return false
}

func main() {
	defer writer.Flush()

	t := scanInt()

	cards := make(Deck, 0, 4)

	for t > 0 {
		t--

		cards = cards[:0]
		for i := 0; i < 4; i++ {
			cards = append(cards, scanCard())
		}
		if validate(cards) {
			writer.WriteString(":)\n")
		} else {
			writer.WriteString(":(\n")
		}
	}
}
