// geometry, sort

package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
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

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

type point struct {
	row, column int
}

func distanceSquare(p1, p2 point) int64 {
	rowDiff := int64(p1.row - p2.row)
	columnDiff := int64(p1.column - p2.column)
	return rowDiff*rowDiff + columnDiff*columnDiff
}

func main() {
	defer writer.Flush()

	t := scanInt()
	points := make([]point, 4)
	distances := make([]int64, 0, 6)

	for t > 0 {
		t--
		for i := 0; i < 4; i++ {
			points[i] = point{scanInt(), scanInt()}
		}

		distances = distances[:0]
		for i := 0; i < 3; i++ {
			for j := i + 1; j < 4; j++ {
				distances = append(distances, distanceSquare(points[i], points[j]))
			}
		}
		slices.Sort(distances)

		if distances[0] == distances[1] && distances[1] == distances[2] && distances[2] == distances[3] && distances[4] == distances[5] {
			writer.WriteString("1\n")
		} else {
			writer.WriteString("0\n")
		}

	}
}
