package leetcode

func checkPair(a, b rune) bool {
	if (a == '(' && b == ')') || (a == '{' && b == '}') || (a == '[' && b == ']') {
		return true
	}
	return false
}

func IsValid(s string) bool {
	var stack []rune
	for _, bracket := range s {
		if bracket == '(' || bracket == '{' || bracket == '[' {
			stack = append(stack, bracket)
		} else if len(stack) == 0 {
			return false
		} else if bracket == ')' || bracket == '}' || bracket == ']' {
			if !checkPair(stack[len(stack)-1], bracket) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
