// stack

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner *bufio.Scanner
	reader  *bufio.Reader
	writer  *bufio.Writer
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func scanInt() int {
	n, _ := strconv.Atoi(scanString())
	return n
}

func readString() string {
	str, _ := reader.ReadString('\n')
	return str[:len(str)-1]
}

func main() {
	defer writer.Flush()

	_ = scanInt()

	var skills string
	fmt.Scanf("%s", &skills)
	// skills := readString()

	skillCount := 0

	lStackCount := 0
	sStackCount := 0

loop:
	for _, skill := range skills {
		switch skill {
		case 'L':
			lStackCount++
		case 'R':
			if lStackCount > 0 {
				lStackCount--
				skillCount++
			} else {
				break loop
			}

		case 'S':
			sStackCount++
		case 'K':
			if sStackCount > 0 {
				sStackCount--
				skillCount++
			} else {
				break loop
			}

		default:
			skillCount++
		}
	}

	writer.WriteString(strconv.Itoa(skillCount))
	writer.WriteByte('\n')
}
