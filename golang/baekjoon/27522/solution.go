// https://www.acmicpc.net/problem/27522

package main

import (
	"bufio"
	"os"
	"slices"
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

type recordTime struct {
	minute, second, millisecond int
}

type recordInfo struct {
	time recordTime
	team string
}

func scanRecordTime() recordTime {
	scanner.Scan()
	input := scanner.Bytes()

	minute := int(input[0] - '0')
	second := int(input[2]-'0')*10 + int(input[3]-'0')
	millisecond := int(input[5]-'0')*100 + int(input[6]-'0')*10 + int(input[7]-'0')

	return recordTime{minute, second, millisecond}
}

func scanTeam() string {
	scanner.Scan()
	return scanner.Text()
}

func scanRecord() recordInfo {
	return recordInfo{scanRecordTime(), scanTeam()}
}

var scorers = [8]int{10, 8, 6, 5, 4, 3, 2, 1}

func main() {
	defer writer.Flush()

	records := make([]recordInfo, 8)

	for i := 0; i < 8; i++ {
		records[i] = scanRecord()
	}

	slices.SortFunc(records, func(a, b recordInfo) int {
		if a.time.minute != b.time.minute {
			return a.time.minute - b.time.minute
		}
		if a.time.second != b.time.second {
			return a.time.second - b.time.second
		}
		if a.time.millisecond != b.time.millisecond {
			return a.time.millisecond - b.time.millisecond
		}
		return 0
	})

	var redScore, blueScore int
	for i, record := range records {
		if record.team == "R" {
			redScore += scorers[i]
		} else {
			blueScore += scorers[i]
		}
	}

	if redScore > blueScore {
		writer.WriteString("Red\n")
	} else {
		writer.WriteString("Blue\n")
	}
}
