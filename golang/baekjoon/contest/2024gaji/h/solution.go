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
	writer = bufio.NewWriter(os.Stdout)
}

func readInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func maze() int {
	writer.WriteString("maze\n")
	writer.Flush()
	return readInt()
}

func gazi(n int) int {
	writer.WriteString("gaji ")
	writer.WriteString(strconv.Itoa(n))
	writer.WriteByte('\n')
	writer.Flush()
	return readInt()
}

func answer(edges [][2]int) {
	writer.WriteString("answer\n")
	for _, e := range edges {
		writer.WriteString(strconv.Itoa(e[0]))
		writer.WriteByte(' ')
		writer.WriteString(strconv.Itoa(e[1]))
		writer.WriteByte('\n')
	}
	writer.Flush()
}

var (
	edges [][2]int
	node  int
)

func dfs(current, parent int) {
	// fmt.Printf("dfs(%d, %d)\n", current, prent)
	nexts := make(map[int]struct{})
	for {
		next := maze()
		// gazi(current)
		if _, ok := nexts[next]; ok || next == parent {
			gazi(current)
			break
		}
		nexts[next] = struct{}{}
		edges = append(edges, [2]int{current, next})
		if len(edges) == node-1 {
			return
		}
		dfs(next, current)
		if len(edges) == node-1 {
			return
		}
		gazi(current)
	}

	// for next := range nexts {
	// 	gazi(next)
	// 	dfs(next, current)
	// 	if len(edges) == node-1 {
	// 		return
	// 	}
	// 	gazi(current)
	// }
}

func main() {
	defer writer.Flush()
	node = readInt()
	edges = make([][2]int, 0, node-1)

	dfs(1, 0)

	answer(edges)
}
