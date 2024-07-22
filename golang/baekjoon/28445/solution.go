package main

import (
	"bufio"
	"os"
	"slices"
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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer writer.Flush()

	colorSet := make(map[string]struct{}, 4)

	for i := 0; i < 4; i++ {
		colorSet[scanString()] = struct{}{}
	}

	colors := make([]string, 0, len(colorSet))
	for color := range colorSet {
		colors = append(colors, color)
	}

	slices.Sort(colors)

	for _, headColor := range colors {
		for _, tailColor := range colors {
			writer.WriteString(headColor + " " + tailColor + "\n")
		}
	}
}
