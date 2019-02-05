package string

import "strings"

func isAlphanumeric(c byte) bool {
	return c >= 97 && c <= 122 || c >= 48 && c <= 57
}

func isPalindrome(s string) bool {
	lower := strings.ToLower(s)
	start := 0
	end := len(lower) - 1
	for ; start < end; start++ {
		if isAlphanumeric(lower[start]) {
			for ; start < end; end-- {
				if isAlphanumeric(lower[end]) {
					if lower[end] == lower[start] {
						end--
						break
					} else {
						return false
					}
				}
			}
		}
	}
	return true
}
