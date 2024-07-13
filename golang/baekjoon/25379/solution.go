// greedy?

package main

import (
	"bufio"
	"os"
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

func main() {
	defer writer.Flush()

	n := scanInt()

	isEvens := make([]bool, n)
	for i := 0; i < n; i++ {
		isEvens[i] = scanInt()%2 == 0
	}

	var oddToEvenCount, evenToOddCount int64
	var oddCount, evenCount int64

	for _, isEven := range isEvens {
		if isEven {
			evenCount++
			oddToEvenCount += oddCount
		} else {
			oddCount++
			evenToOddCount += evenCount
		}
	}

	count := oddToEvenCount
	if evenToOddCount < count {
		count = evenToOddCount
	}

	writer.WriteString(strconv.FormatInt(count, 10))
	writer.WriteByte('\n')
}
