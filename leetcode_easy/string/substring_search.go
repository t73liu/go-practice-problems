package string

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	haystackSize := len(haystack)
	needleSize := len(needle)

	for i := 0; i < haystackSize-needleSize+1; i++ {
		isMatch := true
		for j := 0; j < needleSize; j++ {
			if haystack[i+j] != needle[j] {
				isMatch = false
				break
			}
		}
		if isMatch {
			return i
		}
	}

	return -1
}
