package string

func firstUniqChar(s string) int {
	seen := make(map[byte]bool)
	unique := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		_, isUnique := unique[s[i]]
		_, isSeen := seen[s[i]]
		if isSeen {
			if isUnique {
				delete(unique, s[i])
			}
		} else {
			seen[s[i]] = true
			unique[s[i]] = i
		}
	}
	if len(unique) == 0 {
		return -1
	}
	first := len(s)
	for _, v := range unique {
		if v < first {
			first = v
		}
	}
	return first
}
