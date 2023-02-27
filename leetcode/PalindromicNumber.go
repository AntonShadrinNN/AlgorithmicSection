package leetcode

import "strconv"

func IsPalindrome(x int) bool {
	number := strconv.Itoa(x)
	i, j := 0, len(number)-1
	for i < j {
		if number[i] != number[j] {
			return false
		}
		i++
		j--
	}
	return true
}
