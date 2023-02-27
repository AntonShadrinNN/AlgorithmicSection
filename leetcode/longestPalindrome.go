package leetcode

func max(args ...int) (res int) {
	res = args[0]
	for _, val := range args {
		if val > res {
			res = val
		}
	}
	return
}

func isPalindrome(s *string, j, k int, maxLength int, maxPalindrome *string) int {
	//if j == k {
	//	maxLength--
	//}
	var curLength int
	if j == k {
		curLength = 0
	} else {
		curLength = 1
	}
	for j >= 0 && k < len(*s) && (*s)[j] == (*s)[k] {
		if (*s)[j] == (*s)[k] {
			curLength += 2
			if curLength > maxLength {
				maxLength = curLength
				*maxPalindrome = (*s)[j : k+1]
			}
			j--
			k++
		} else {
			break
		}
	}
	return maxLength
}

func LongestPalindrome(s string) string {
	maxLength := 0
	maxPalindrome := ""
	for i := range s {
		// odd palindromes
		maxLength = max(isPalindrome(&s, i, i, maxLength, &maxPalindrome), maxLength)

		// even palindromes
		maxLength = max(isPalindrome(&s, i, i+1, maxLength, &maxPalindrome), maxLength)
	}
	return maxPalindrome
}
