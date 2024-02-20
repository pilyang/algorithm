// https://www.acmicpc.net/problem/1991
// tree Traversal

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
	scanner.Split(bufio.ScanWords)
}

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

func readInt() int {
	val, _ := strconv.Atoi(readString())
	return val
}

// map[parent] = {leftChild, rightChild}
type Tree map[string][2]string

var tree Tree

func isLeafNode(node string) bool {
	return tree[node][0] == "." && tree[node][1] == "."
}

func visitLeft(current string) (string, bool) {
	if tree[current][0] == "." {
		return "", false
	}
	return tree[current][0], true
}

func visitRight(current string) (string, bool) {
	if tree[current][1] == "." {
		return "", false
	}
	return tree[current][1], true
}

func dfs(current string) {
	preorder = append(preorder, current)

	if left, exist := visitLeft(current); exist {
		dfs(left)
	}

	inorder = append(inorder, current)

	if right, exist := visitRight(current); exist {
		dfs(right)
	}

	postorder = append(postorder, current)
}

//	func preorderDFS(current string) {
//		// visit check
//		writer.WriteString(current)
//
//		if left, exist := visitLeft(current); exist {
//			preorderDFS(left)
//		}
//		if right, exist := visitRight(current); exist {
//			preorderDFS(right)
//		}
//	}
//
//	func inorderDFS(current string) {
//		if left, exist := visitLeft(current); exist {
//			inorderDFS(left)
//		}
//
//		// visit check
//		writer.WriteString(current)
//
//		if right, exist := visitRight(current); exist {
//			inorderDFS(right)
//		}
//	}
//
//	func postorderDFS(current string) {
//		if left, exist := visitLeft(current); exist {
//			postorderDFS(left)
//		}
//
//		if right, exist := visitRight(current); exist {
//			postorderDFS(right)
//		}
//
//		// visit check
//		writer.WriteString(current)
//	}

var preorder, inorder, postorder []string

func main() {
	defer writer.Flush()

	nodeNum := readInt()

	tree = make(Tree)
	for i := 0; i < nodeNum; i++ {
		p, l, r := readString(), readString(), readString()
		tree[p] = [2]string{l, r}
	}

	preorder = make([]string, 0, nodeNum)
	inorder = make([]string, 0, nodeNum)
	postorder = make([]string, 0, nodeNum)

	dfs("A")

	for _, val := range preorder {
		writer.WriteString(val)
	}
	writer.WriteByte('\n')

	for _, val := range inorder {
		writer.WriteString(val)
	}
	writer.WriteByte('\n')

	for _, val := range postorder {
		writer.WriteString(val)
	}
	writer.WriteByte('\n')
}
