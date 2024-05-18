// implementation, string

package main

import (
	"bufio"
	"os"
)

var (
	scanner *bufio.Scanner
	writer  *bufio.Writer
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
}

func scanRunes() []rune {
	scanner.Scan()
	return []rune(scanner.Text())
}

func main() {
	defer writer.Flush()

	origin := scanRunes()
	key := scanRunes()

	for i := 0; i < len(origin); i++ {
		if origin[i] == ' ' {
			writer.WriteByte(' ')
			continue
		}

		keyIdx := i % len(key)
		shift := int(key[keyIdx] - 'a' + 1)
		decrypted := int(origin[i]) - shift
		if decrypted < 'a' {
			decrypted += 26
		}
		writer.WriteByte(byte(decrypted))
	}
	writer.WriteByte('\n')
}
