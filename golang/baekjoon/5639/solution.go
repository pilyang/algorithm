// https://www.acmicpc.net/problem/5639
// dfs?, tree

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

func scanInt() (int, bool) {
	if scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, false
		}
		return val, true
	}
	return 0, false
}

// tree[parent] = {leftChild, rightChild}
// 0 for nil
var (
	tree = make(map[int][2]int, 10_000)
	root int
)

func insertNode(current, node int) {
	if node < current {
		if tree[current][0] == 0 {
			tree[current] = [2]int{node, tree[current][1]}
		} else {
			insertNode(tree[current][0], node)
		}
	} else {
		if tree[current][1] == 0 {
			tree[current] = [2]int{tree[current][0], node}
		} else {
			insertNode(tree[current][1], node)
		}
	}
}

func dfs(current int) {
	if tree[current][0] != 0 {
		dfs(tree[current][0])
	}
	if tree[current][1] != 0 {
		dfs(tree[current][1])
	}
	writer.WriteString(strconv.Itoa(current))
	writer.WriteByte('\n')
}

func main() {
	defer writer.Flush()

	root, _ = scanInt()
	tree[root] = [2]int{0, 0}

	// Read the rest of the nodes and insert them into the tree
	for {
		if node, ok := scanInt(); ok {
			insertNode(root, node)
		} else {
			break
		}
	}

	dfs(root)
}
