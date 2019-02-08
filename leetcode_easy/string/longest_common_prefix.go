package string

import "strings"

func longestCommonPrefix(strs []string) string {
	strsCount := len(strs)

	if strsCount == 0 {
		return ""
	} else if strsCount == 1 {
		return strs[0]
	}

	var sb strings.Builder

	first := strs[0]
	isMatch := true
	for i := 0; ; i++ {
		if i >= len(first) {
			break
		}
		c := first[i]
		for j := 1; j < strsCount; j++ {
			str := strs[j]
			if i >= len(str) {
				isMatch = false
				break
			} else if str[i] != c {
				isMatch = false
				break
			}
		}
		if isMatch {
			sb.WriteByte(c)
		} else {
			break
		}
	}
	return sb.String()
}
