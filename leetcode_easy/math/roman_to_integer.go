package math

func romanToInt(s string) int {
	mapping := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	size := len(s)
	sum := 0
	if size == 0 {
		return 0
	}
	prevValue := mapping[string(s[0])]
	if size == 1 {
		return prevValue
	}
	for i := 1; i < size; i++ {
		curr := mapping[string(s[i])]
		if curr > prevValue {
			sum -= prevValue
		} else {
			sum += prevValue
		}
		prevValue = curr
		if i == size-1 {
			sum += curr
		}
	}
	return sum
}
