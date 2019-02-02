package string

import "math"

func reverse(x int) int {
	result := 0
	current := x
	if x < 0 {
		current = current * -1
	}
	for current > 0 {
		digit := current % 10
		result = result*10 + digit
		current = current / 10
	}
	if result > int(math.Exp2(31.0)) || result < int(-math.Exp2(31.0)) {
		return 0
	}
	if x < 0 {
		return result * -1
	}
	return result
}
