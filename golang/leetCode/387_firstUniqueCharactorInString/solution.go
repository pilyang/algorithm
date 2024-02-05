// https://leetcode.com/problems/first-unique-character-in-a-string/description/?envType=daily-question&envId=2024-02-05

package uniqueCharacter

func firstUniqChar(s string) int {
	runeCount := make(map[rune]int)
	runeIndex := make(map[rune]int)

	for i, r := range s {
		runeCount[r]++
		if _, exists := runeIndex[r]; !exists {
			runeIndex[r] = i
		}
	}

	mi := len(s)
	for r, c := range runeCount {
		if c == 1 && runeIndex[r] < mi {
			mi = runeIndex[r]
		}
	}

	if mi == len(s) {
		return -1
	}
	return mi
}
