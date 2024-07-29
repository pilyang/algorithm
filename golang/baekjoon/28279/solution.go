// dequeue

package main

import (
	"bufio"
	"container/list"
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

func writeLn(num int) {
	writer.WriteString(strconv.Itoa(num))
	writer.WriteRune('\n')
}

type Position int

const (
	front Position = iota
	back
)

func getAndRemove(pos Position) int {
	var elem *list.Element
	if pos == front {
		elem = dequeue.Front()
	} else {
		elem = dequeue.Back()
	}
	if elem == nil {
		return -1
	}

	dequeue.Remove(elem)
	return elem.Value.(int)
}

func get(pos Position) int {
	var elem *list.Element
	if pos == front {
		elem = dequeue.Front()
	} else {
		elem = dequeue.Back()
	}
	if elem == nil {
		return -1
	}

	return elem.Value.(int)
}

var dequeue = list.New()

func main() {
	defer writer.Flush()

	n := scanInt()
	for n > 0 {
		n--
		switch scanInt() {
		case 1:
			num := scanInt()
			dequeue.PushFront(num)
		case 2:
			num := scanInt()
			dequeue.PushBack(num)
		case 3:
			num := getAndRemove(front)
			writeLn(num)
		case 4:
			num := getAndRemove(back)
			writeLn(num)
		case 5:
			writeLn(dequeue.Len())
		case 6:
			if dequeue.Len() == 0 {
				writer.WriteString("1\n")
			} else {
				writer.WriteString("0\n")
			}
		case 7:
			num := get(front)
			writeLn(num)
		case 8:
			num := get(back)
			writeLn(num)
		}
	}
}
