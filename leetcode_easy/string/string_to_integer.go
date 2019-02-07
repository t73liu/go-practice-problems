package string

import (
	"fmt"
	"math"
)

func isInteger(c byte) bool {
	return c >= 48 && c <= 57
}

func isNegative(c byte) bool {
	return c == 45
}

func isPositive(c byte) bool {
	return c == 43
}

func isSpace(c byte) bool {
	return c == 32
}

func myAtoi(str string) int {
	result := 0

	isFirstSeq := false
	isLeadingNegative := false

	for i := 0; i < len(str); i++ {
		if isSpace(str[i]) {
			if isFirstSeq {
				break
			}
		} else {
			if isNegative(str[i]) {
				if isFirstSeq {
					break
				} else {
					isFirstSeq = true
					isLeadingNegative = true
				}
			} else if isInteger(str[i]) {
				isFirstSeq = true
				result = result*10 + int(str[i]) - 48

				if isLeadingNegative && (result*-1) < int(-math.Exp2(31.0)) {
					fmt.Println("test")
					return int(-math.Exp2(31.0))
				} else if result > int(math.Exp2(31.0))-1 {
					fmt.Println("test 2")
					return int(math.Exp2(31.0)) - 1
				}
			} else if isPositive(str[i]) {
				if isFirstSeq {
					break
				} else {
					isFirstSeq = true
				}
			} else {
				break
			}
		}
	}

	if isLeadingNegative {
		result = -1 * result
	}

	return result
}
