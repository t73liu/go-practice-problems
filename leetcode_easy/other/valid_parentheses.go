package other

func isValid(s string) bool {
	size := len(s)
	if size == 0 {
		return true
	} else if size%2 == 1 {
		return false
	}
	mappings := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := make([]rune, 0)
	for i := 0; i < size; i++ {
		char := rune(s[i])
		if char == '{' || char == '[' || char == '(' {
			stack = append(stack, char)
		} else {
			reverse := mappings[char]
			stackSize := len(stack)
			if stackSize > 0 && stack[stackSize-1] == reverse {
				stack = stack[:stackSize-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
