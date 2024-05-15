// simulation,implementation

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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func scanInt() int {
	n, _ := strconv.Atoi(scanString())
	return n
}

var (
	stockInfo map[string]int
	groupInfo map[int][]string

	budget  int64
	myStock map[string]int
)

func buyStock(stock string, num int) {
	if budget < int64(stockInfo[stock])*int64(num) {
		return
	}

	budget -= int64(stockInfo[stock]) * int64(num)
	myStock[stock] += num
}

func sellStock(stock string, num int) {
	if myStock[stock] < num {
		initial := myStock[stock]
		myStock[stock] = 0
		budget += int64(stockInfo[stock]) * int64(initial)
		return
	}

	myStock[stock] -= num
	budget += int64(stockInfo[stock]) * int64(num)
}

func changeStockPriceByValue(stoc string, value int) {
	stockInfo[stoc] += value
}

func changeStockPriceByPercentage(stock string, percentage int) {
	// floatPercentage := float64(percentage) / 100.0
	// stockInfo[stock] = int(float64(stockInfo[stock]) * (1.0 + floatPercentage))
	stockInfo[stock] = (stockInfo[stock] * (100 + percentage)) / 100

	stockInfo[stock] = (stockInfo[stock] / 10) * 10
}

func changeGroupByValue(group int, value int) {
	for _, stock := range groupInfo[group] {
		changeStockPriceByValue(stock, value)
	}
}

func changeGroupByPercentage(group int, percentage int) {
	for _, stock := range groupInfo[group] {
		changeStockPriceByPercentage(stock, percentage)
	}
}

func getTotalStockPrice() int64 {
	var total int64
	for stock, num := range myStock {
		total += int64(stockInfo[stock]) * int64(num)
	}

	return total
}

func main() {
	defer writer.Flush()

	numOfStock := scanInt()
	budget = int64(scanInt())
	commandNums := scanInt()

	myStock = make(map[string]int)
	stockInfo = make(map[string]int)
	groupInfo = make(map[int][]string)

	for numOfStock > 0 {
		group, stock, price := scanInt(), scanString(), scanInt()
		stockInfo[stock] = price
		groupInfo[group] = append(groupInfo[group], stock)
		numOfStock--
	}

	for commandNums > 0 {
		switch scanInt() {
		case 1:
			buyStock(scanString(), scanInt())
		case 2:
			sellStock(scanString(), scanInt())
		case 3:
			changeStockPriceByValue(scanString(), scanInt())
		case 4:
			changeGroupByValue(scanInt(), scanInt())
		case 5:
			changeGroupByPercentage(scanInt(), scanInt())
		case 6:
			writer.WriteString(strconv.FormatInt(budget, 10))
			writer.WriteByte('\n')
		case 7:
			writer.WriteString(strconv.FormatInt(getTotalStockPrice()+budget, 10))
			writer.WriteByte('\n')
		}
		commandNums--
	}
}
