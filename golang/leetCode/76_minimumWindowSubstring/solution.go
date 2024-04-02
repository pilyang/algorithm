// https://leetcode.com/problems/minimum-window-substring/description/?envType=daily-question&envId=2024-02-04

package minimumWindowSubstring

func minWidows(s string, t string) string {
	target := make(map[rune]int)
	for _, r := range t {
		target[r] += 1
	}
	return ""
}

func isMatch(tokens *[]strToken, target map[rune]int) bool {
}

type rInfo struct {
	r rune
	i int
}

type strToken struct {
	first   rInfo
	seccond rInfo
}
