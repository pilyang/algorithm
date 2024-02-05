// https://leetcode.com/problems/first-unique-character-in-a-string/description/?envType=daily-question&envId=2024-02-05

package uniqueCharacter

func firstUniqChar(s string) int {
	runeCount := make(map[rune]int)

	for _, r := range s {
		runeCount[r]++
	}

	for i, r := range s {
		if runeCount[r] == 1 {
			return i
		}
	}

	return -1
}
