package string

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n <= 1 {
		return "1"
	}
	val := "1"
	var sb strings.Builder
	for num := 2; num <= n; num++ {
		count := 1
		prev := val[0]
		for i := 1; i < len(val); i++ {
			if val[i] == prev {
				count++
			} else {
				sb.WriteString(strconv.Itoa(count))
				sb.WriteByte(prev)
				prev = val[i]
				count = 1
			}
		}
		sb.WriteString(strconv.Itoa(count))
		sb.WriteByte(prev)
		val = sb.String()
		sb.Reset()
	}
	return val
}
