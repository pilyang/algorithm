// math, mod, fast-exponential

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

const mod = 1_000_000_007

func fastExp(a, n int) int {
	res := 1
	for n > 0 {
		if n&1 == 1 {
			res = (res * a) % mod
		}
		a = (a * a) % mod
		n >>= 1
	}
	return res
}

func calcModSN(n, s int) int {
	res := fastExp(n, mod-2)
	res = (res * s) % mod
	return res
}

func main() {
	defer writer.Flush()

	M := scanInt()
	ans := 0
	for i := 0; i < M; i++ {
		N, S := scanInt(), scanInt()
		ans += calcModSN(N, S)
		ans %= mod
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
