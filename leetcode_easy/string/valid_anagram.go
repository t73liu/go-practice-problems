package string

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	countOne := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		val, ok := countOne[s[i]]
		if ok {
			countOne[s[i]] = val + 1
		} else {
			countOne[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		val, ok := countOne[t[i]]
		if ok {
			if val == 1 {
				delete(countOne, t[i])
			} else {
				countOne[t[i]] = val - 1
			}
		} else {
			return false
		}
	}
	if len(countOne) > 0 {
		return false
	}
	return true
}
