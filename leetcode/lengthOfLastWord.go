package leetcode

func LengthOfLastWord(s string) int {
	runes := []rune(s)
	endOfWord := len(runes) - 1
	for runes[endOfWord] == ' ' {
		endOfWord--
	}

	startOfWord := endOfWord
	for startOfWord >= 0 && runes[startOfWord] != ' ' {
		startOfWord--
	}
	return endOfWord - startOfWord
}
