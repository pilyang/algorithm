// https://www.acmicpc.net/problem/9019
// bfs

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

type Node struct {
	value   int
	command string
}

func readInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func bfs(from, to int) string {
	visited := make([]bool, 10000)
	queue := []Node{{from, ""}}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.value == to {
			return node.command
		}

		commands := []struct {
			value   int
			command string
		}{
			{(node.value * 2) % 10000, "D"},
			{(node.value - 1 + 10000) % 10000, "S"},
			{(node.value%1000)*10 + node.value/1000, "L"},
			{(node.value / 10) + (node.value%10)*1000, "R"},
		}

		for _, cmd := range commands {
			if !visited[cmd.value] {
				visited[cmd.value] = true
				queue = append(queue, Node{cmd.value, node.command + cmd.command})
			}
		}
	}

	return ""
}

func main() {
	defer writer.Flush()

	t := readInt()
	for ; t > 0; t-- {
		from, to := readInt(), readInt()
		writer.WriteString(bfs(from, to))
		writer.WriteByte('\n')
	}
}
