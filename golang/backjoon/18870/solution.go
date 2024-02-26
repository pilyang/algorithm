// https://www.acmicpc.net/problem/18870
// sort
// binary search
// set

package main

import (
	"bufio"
	"os"
	"slices"
	"sort"
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

func readInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

func main() {
	defer writer.Flush()

	num := readInt()
	arr := make([]int, num)
	distinct := make([]int, 0, num)
	distinctMap := make(map[int]struct{})

	for i := 0; i < num; i++ {
		arr[i] = readInt()
		if _, exist := distinctMap[arr[i]]; !exist {
			distinctMap[arr[i]] = struct{}{}
			distinct = append(distinct, arr[i])
		}
	}

	slices.Sort(distinct)

	for _, val := range arr {
		writer.WriteString(strconv.Itoa(sort.Search(len(distinct), func(i int) bool { return distinct[i] >= val })))
		writer.WriteByte(' ')
	}
	writer.WriteByte('\n')
}
